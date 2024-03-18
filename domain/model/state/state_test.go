// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"
	"database/sql"

	"github.com/juju/collections/set"
	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/cloud"
	coremodel "github.com/juju/juju/core/model"
	modeltesting "github.com/juju/juju/core/model/testing"
	"github.com/juju/juju/core/user"
	dbcloud "github.com/juju/juju/domain/cloud/state"
	"github.com/juju/juju/domain/credential"
	credentialstate "github.com/juju/juju/domain/credential/state"
	"github.com/juju/juju/domain/model"
	modelerrors "github.com/juju/juju/domain/model/errors"
	schematesting "github.com/juju/juju/domain/schema/testing"
	usererrors "github.com/juju/juju/domain/user/errors"
	userstate "github.com/juju/juju/domain/user/state"
	"github.com/juju/juju/version"
)

type modelSuite struct {
	schematesting.ControllerSuite
	uuid     coremodel.UUID
	userUUID user.UUID
	userName string
}

var _ = gc.Suite(&modelSuite{})

func (m *modelSuite) SetUpTest(c *gc.C) {
	m.ControllerSuite.SetUpTest(c)

	// We need to generate a user in the database so that we can set the model
	// owner.
	userUUID, err := user.NewUUID()
	m.userUUID = userUUID
	m.userName = "test-user"
	c.Assert(err, jc.ErrorIsNil)
	userState := userstate.NewState(m.TxnRunnerFactory())
	err = userState.AddUser(
		context.Background(),
		m.userUUID,
		m.userName,
		m.userName,
		m.userUUID,
	)
	c.Assert(err, jc.ErrorIsNil)

	cloudSt := dbcloud.NewState(m.TxnRunnerFactory())
	err = cloudSt.UpsertCloud(context.Background(), cloud.Cloud{
		Name:      "my-cloud",
		Type:      "ec2",
		AuthTypes: cloud.AuthTypes{cloud.AccessKeyAuthType, cloud.UserPassAuthType},
		Regions: []cloud.Region{
			{
				Name: "my-region",
			},
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	cred := credential.CloudCredentialInfo{
		Label:    "foobar",
		AuthType: string(cloud.AccessKeyAuthType),
		Attributes: map[string]string{
			"foo": "foo val",
			"bar": "bar val",
		},
	}

	credSt := credentialstate.NewState(m.TxnRunnerFactory())
	_, err = credSt.UpsertCloudCredential(
		context.Background(), credential.ID{
			Cloud: "my-cloud",
			Owner: "test-user",
			Name:  "foobar",
		},
		cred,
	)
	c.Assert(err, jc.ErrorIsNil)

	m.uuid = modeltesting.GenModelUUID(c)
	modelSt := NewState(m.TxnRunnerFactory())
	err = modelSt.Create(
		context.Background(),
		m.uuid,
		coremodel.IAAS,
		model.ModelCreationArgs{
			AgentVersion: version.Current,
			Cloud:        "my-cloud",
			CloudRegion:  "my-region",
			Credential: credential.ID{
				Cloud: "my-cloud",
				Owner: "test-user",
				Name:  "foobar",
			},
			Name:  "my-test-model",
			Owner: m.userUUID,
		},
	)
	c.Assert(err, jc.ErrorIsNil)

	err = modelSt.Finalise(context.Background(), m.uuid)
	c.Assert(err, jc.ErrorIsNil)
}

// TestModelCloudNameAndCredential tests the happy path for getting a models
// cloud name and credential.
func (m *modelSuite) TestModelCloudNameAndCredential(c *gc.C) {
	st := NewState(m.TxnRunnerFactory())
	// We are relying on the model setup as part of this suite.
	cloudName, credentialID, err := st.ModelCloudNameAndCredential(context.Background(), "my-test-model", "test-user")
	c.Assert(err, jc.ErrorIsNil)
	c.Check(cloudName, gc.Equals, "my-cloud")
	c.Check(credentialID, gc.Equals, credential.ID{
		Cloud: "my-cloud",
		Owner: string(m.userUUID),
		Name:  "foobar",
	})
}

// TestModelCloudNameAndEmptyCredential tests that if a model does not have a
// credential set the zero value for the [credential.Id] is returned.
func (m *modelSuite) TestModelCloudNameAndEmptyCredential(c *gc.C) {
	st := NewState(m.TxnRunnerFactory())

	modelUUID := modeltesting.GenModelUUID(c)
	// We need to first inject a model that does not have a cloud credential set
	err := st.Create(
		context.Background(),
		modelUUID,
		coremodel.IAAS,
		model.ModelCreationArgs{
			AgentVersion: version.Current,
			Cloud:        "my-cloud",
			Name:         "no-credential-model",
			Owner:        m.userUUID,
		},
	)
	c.Assert(err, jc.ErrorIsNil)

	err = st.Finalise(context.Background(), modelUUID)
	c.Assert(err, jc.ErrorIsNil)

	cloudName, credentialID, err := st.ModelCloudNameAndCredential(context.Background(), "no-credential-model", "test-user")
	c.Assert(err, jc.ErrorIsNil)
	c.Check(cloudName, gc.Equals, "my-cloud")
	c.Check(credentialID.IsZero(), jc.IsTrue)
}

// TestModelCloudNameAndCredentialController is testing the cloud name and
// credential id is returned for the controller model and owner.
func (m *modelSuite) TestModelCloudNameAndCredentialController(c *gc.C) {
	userUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	userState := userstate.NewState(m.TxnRunnerFactory())
	err = userState.AddUser(
		context.Background(),
		userUUID,
		coremodel.ControllerModelOwnerUsername,
		coremodel.ControllerModelOwnerUsername,
		userUUID,
	)
	c.Assert(err, jc.ErrorIsNil)

	st := NewState(m.TxnRunnerFactory())
	modelUUID := modeltesting.GenModelUUID(c)
	// We need to first inject a model that does not have a cloud credential set
	err = st.Create(
		context.Background(),
		modelUUID,
		coremodel.IAAS,
		model.ModelCreationArgs{
			AgentVersion: version.Current,
			Cloud:        "my-cloud",
			Credential: credential.ID{
				Cloud: "my-cloud",
				Owner: "test-user",
				Name:  "foobar",
			},
			Name:  coremodel.ControllerModelName,
			Owner: userUUID,
		},
	)
	c.Assert(err, jc.ErrorIsNil)

	err = st.Finalise(context.Background(), modelUUID)

	cloudName, credentialID, err := st.ModelCloudNameAndCredential(
		context.Background(),
		coremodel.ControllerModelName,
		coremodel.ControllerModelOwnerUsername,
	)

	c.Assert(err, jc.ErrorIsNil)
	c.Check(cloudName, gc.Equals, "my-cloud")
	c.Check(credentialID, gc.Equals, credential.ID{
		Cloud: "my-cloud",
		Owner: string(m.userUUID),
		Name:  "foobar",
	})
}

// TestModelCloudNameAndCredentialNotFound is testing that if we pass a model
// that doesn't exist we get back a [modelerrors.NotFound] error.
func (m *modelSuite) TestModelCloudNameAndCredentialNotFound(c *gc.C) {
	st := NewState(m.TxnRunnerFactory())
	// We are relying on the model setup as part of this suite.
	cloudName, credentialID, err := st.ModelCloudNameAndCredential(context.Background(), "does-not-exist", "test-user")
	c.Assert(err, jc.ErrorIs, modelerrors.NotFound)
	c.Check(cloudName, gc.Equals, "")
	c.Check(credentialID.IsZero(), jc.IsTrue)
}

// TestCreateModelAgentWithNoModel is asserting that if we attempt to make a
// model agent record where no model already exists that we get back a
// [modelerrors.NotFound] error.
func (m *modelSuite) TestCreateModelAgentWithNoModel(c *gc.C) {
	runner, err := m.TxnRunnerFactory()()
	c.Assert(err, jc.ErrorIsNil)

	testUUID := modeltesting.GenModelUUID(c)
	err = runner.StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		return createModelAgent(context.Background(), tx, testUUID, version.Current)
	})

	c.Assert(err, jc.ErrorIs, modelerrors.NotFound)
}

// TestCreateModelAgentAlreadyExists is asserting that if we attempt to make a
// model agent record when one already exists we get a
// [modelerrors.AlreadyExists] back.
func (m *modelSuite) TestCreateModelAgentAlreadyExists(c *gc.C) {
	runner, err := m.TxnRunnerFactory()()
	c.Assert(err, jc.ErrorIsNil)

	err = runner.StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		return createModelAgent(context.Background(), tx, m.uuid, version.Current)
	})

	c.Assert(err, jc.ErrorIs, modelerrors.AlreadyExists)
}

func (m *modelSuite) TestCreateModelWithExisting(c *gc.C) {
	runner, err := m.TxnRunnerFactory()()
	c.Assert(err, jc.ErrorIsNil)

	err = runner.StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		return createModel(
			ctx,
			tx,
			m.uuid,
			coremodel.IAAS,
			model.ModelCreationArgs{
				Cloud:       "my-cloud",
				CloudRegion: "my-region",
				Name:        "fantasticmodel",
				Owner:       m.userUUID,
			},
		)
	})
	c.Assert(err, jc.ErrorIs, modelerrors.AlreadyExists)
}

func (m *modelSuite) TestCreateModelWithSameNameAndOwner(c *gc.C) {
	modelSt := NewState(m.TxnRunnerFactory())
	testUUID := modeltesting.GenModelUUID(c)
	err := modelSt.Create(
		context.Background(),
		testUUID,
		coremodel.IAAS,
		model.ModelCreationArgs{
			Cloud:       "my-cloud",
			CloudRegion: "my-region",
			Name:        "my-test-model",
			Owner:       m.userUUID,
		},
	)
	c.Assert(err, jc.ErrorIs, modelerrors.AlreadyExists)
}

func (m *modelSuite) TestCreateModelWithInvalidCloudRegion(c *gc.C) {
	modelSt := NewState(m.TxnRunnerFactory())
	testUUID := modeltesting.GenModelUUID(c)
	err := modelSt.Create(
		context.Background(),
		testUUID,
		coremodel.IAAS,
		model.ModelCreationArgs{
			Cloud:       "my-cloud",
			CloudRegion: "noexist",
			Name:        "noregion",
			Owner:       m.userUUID,
		},
	)
	c.Assert(err, jc.ErrorIs, errors.NotFound)
}

// TestCreateModelWithNonExistentOwner is here to assert that if we try and make
// a model with a user/owner that does not exist a [usererrors.NotFound] error
// is returned.
func (m *modelSuite) TestCreateModelWithNonExistentOwner(c *gc.C) {
	modelSt := NewState(m.TxnRunnerFactory())
	testUUID := modeltesting.GenModelUUID(c)
	err := modelSt.Create(
		context.Background(),
		testUUID,
		coremodel.IAAS,
		model.ModelCreationArgs{
			Cloud:       "my-cloud",
			CloudRegion: "noexist",
			Name:        "noregion",
			Owner:       user.UUID("noexist"), // does not exist
		},
	)
	c.Assert(err, jc.ErrorIs, usererrors.NotFound)
}

// TestCreateModelWithRemovedOwner is here to test that if we try and create a
// new model with an owner that has been removed from the Juju user base that
// the operation fails with a [usererrors.NotFound] error.
func (m *modelSuite) TestCreateModelWithRemovedOwner(c *gc.C) {
	userState := userstate.NewState(m.TxnRunnerFactory())
	err := userState.RemoveUser(context.Background(), m.userName)
	c.Assert(err, jc.ErrorIsNil)

	modelSt := NewState(m.TxnRunnerFactory())
	testUUID := modeltesting.GenModelUUID(c)
	err = modelSt.Create(
		context.Background(),
		testUUID,
		coremodel.IAAS,
		model.ModelCreationArgs{
			Cloud:       "my-cloud",
			CloudRegion: "noexist",
			Name:        "noregion",
			Owner:       m.userUUID,
		},
	)
	c.Assert(err, jc.ErrorIs, usererrors.NotFound)
}

func (m *modelSuite) TestCreateModelWithInvalidCloud(c *gc.C) {
	modelSt := NewState(m.TxnRunnerFactory())
	testUUID := modeltesting.GenModelUUID(c)
	err := modelSt.Create(
		context.Background(),
		testUUID,
		coremodel.IAAS,
		model.ModelCreationArgs{
			Cloud:       "noexist",
			CloudRegion: "my-region",
			Name:        "noregion",
			Owner:       m.userUUID,
		},
	)
	c.Assert(err, jc.ErrorIs, errors.NotFound)
}

func (m *modelSuite) TestUpdateCredentialForDifferentCloud(c *gc.C) {
	cloudSt := dbcloud.NewState(m.TxnRunnerFactory())
	err := cloudSt.UpsertCloud(context.Background(), cloud.Cloud{
		Name:      "my-cloud2",
		Type:      "ec2",
		AuthTypes: cloud.AuthTypes{cloud.AccessKeyAuthType, cloud.UserPassAuthType},
		Regions: []cloud.Region{
			{
				Name: "my-region",
			},
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	cred := credential.CloudCredentialInfo{
		Label:    "foobar",
		AuthType: string(cloud.AccessKeyAuthType),
		Attributes: map[string]string{
			"foo": "foo val",
			"bar": "bar val",
		},
	}

	credSt := credentialstate.NewState(m.TxnRunnerFactory())
	_, err = credSt.UpsertCloudCredential(
		context.Background(), credential.ID{
			Cloud: "my-cloud2",
			Owner: "test-user",
			Name:  "foobar1",
		},
		cred,
	)
	c.Assert(err, jc.ErrorIsNil)

	modelSt := NewState(m.TxnRunnerFactory())
	err = modelSt.UpdateCredential(
		context.Background(),
		m.uuid,
		credential.ID{
			Cloud: "my-cloud2",
			Owner: "test-user",
			Name:  "foobar1",
		},
	)
	c.Assert(err, jc.ErrorIs, errors.NotValid)
}

// We are trying to test here that we can set a cloud credential on the model
// for the same cloud as the model when no cloud region has been set. This is a
// regression test discovered during DQlite development where we messed up the
// logic assuming that a cloud region was always set for a model.
func (m *modelSuite) TestSetModelCloudCredentialWithoutRegion(c *gc.C) {
	cloudSt := dbcloud.NewState(m.TxnRunnerFactory())
	err := cloudSt.UpsertCloud(context.Background(), cloud.Cloud{
		Name:      "minikube",
		Type:      "kubernetes",
		AuthTypes: cloud.AuthTypes{cloud.UserPassAuthType},
		Regions:   []cloud.Region{},
	})
	c.Assert(err, jc.ErrorIsNil)

	cred := credential.CloudCredentialInfo{
		Label:    "foobar",
		AuthType: string(cloud.UserPassAuthType),
		Attributes: map[string]string{
			"username": "myuser",
			"password": "secret",
		},
	}

	credSt := credentialstate.NewState(m.TxnRunnerFactory())
	_, err = credSt.UpsertCloudCredential(
		context.Background(), credential.ID{
			Cloud: "minikube",
			Owner: "test-user",
			Name:  "foobar",
		},
		cred,
	)
	c.Assert(err, jc.ErrorIsNil)

	m.uuid = modeltesting.GenModelUUID(c)
	modelSt := NewState(m.TxnRunnerFactory())
	err = modelSt.Create(
		context.Background(),
		m.uuid,
		coremodel.CAAS,
		model.ModelCreationArgs{
			Cloud: "minikube",
			Credential: credential.ID{
				Cloud: "minikube",
				Owner: "test-user",
				Name:  "foobar",
			},
			Name:  "controller",
			Owner: m.userUUID,
		},
	)
	c.Assert(err, jc.ErrorIsNil)

	err = modelSt.Finalise(context.Background(), m.uuid)
	c.Assert(err, jc.ErrorIsNil)
}

// TestDeleteModel tests that we can delete a model that is already created in
// the system. We also confirm that list models returns no models after the
// deletion.
func (m *modelSuite) TestDeleteModel(c *gc.C) {
	modelSt := NewState(m.TxnRunnerFactory())
	err := modelSt.Delete(
		context.Background(),
		m.uuid,
	)
	c.Assert(err, jc.ErrorIsNil)

	runner, err := m.TxnRunnerFactory()()
	c.Assert(err, jc.ErrorIsNil)

	err = runner.StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		row := tx.QueryRowContext(
			context.Background(),
			"SELECT uuid FROM model WHERE uuid = ?",
			m.uuid,
		)
		var val string
		err := row.Scan(&val)
		c.Assert(err, jc.ErrorIs, sql.ErrNoRows)
		return nil
	})
	c.Assert(err, jc.ErrorIsNil)

	modelUUIDS, err := modelSt.List(context.Background())
	c.Check(err, jc.ErrorIsNil)
	c.Check(modelUUIDS, gc.HasLen, 0)
}

func (m *modelSuite) TestDeleteModelNotFound(c *gc.C) {
	uuid := modeltesting.GenModelUUID(c)
	modelSt := NewState(m.TxnRunnerFactory())
	err := modelSt.Delete(context.Background(), uuid)
	c.Assert(err, jc.ErrorIs, modelerrors.NotFound)
}

// TestListModels is testing that once we have created several models calling
// list returns all of the models created.
func (m *modelSuite) TestListModels(c *gc.C) {
	uuid1 := modeltesting.GenModelUUID(c)
	modelSt := NewState(m.TxnRunnerFactory())
	err := modelSt.Create(
		context.Background(),
		uuid1,
		coremodel.IAAS,
		model.ModelCreationArgs{
			AgentVersion: version.Current,
			Cloud:        "my-cloud",
			CloudRegion:  "my-region",
			Credential: credential.ID{
				Cloud: "my-cloud",
				Owner: "test-user",
				Name:  "foobar",
			},
			Name:  "listtest1",
			Owner: m.userUUID,
		},
	)
	c.Assert(err, jc.ErrorIsNil)

	uuid2 := modeltesting.GenModelUUID(c)
	err = modelSt.Create(
		context.Background(),
		uuid2,
		coremodel.IAAS,
		model.ModelCreationArgs{
			AgentVersion: version.Current,
			Cloud:        "my-cloud",
			CloudRegion:  "my-region",
			Credential: credential.ID{
				Cloud: "my-cloud",
				Owner: "test-user",
				Name:  "foobar",
			},
			Name:  "listtest2",
			Owner: m.userUUID,
		},
	)
	c.Assert(err, jc.ErrorIsNil)

	uuids, err := modelSt.List(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Check(uuids, gc.HasLen, 3)

	uuidsSet := set.NewStrings()
	for _, uuid := range uuids {
		uuidsSet.Add(uuid.String())
	}

	c.Check(uuidsSet.Contains(uuid1.String()), jc.IsTrue)
	c.Check(uuidsSet.Contains(uuid2.String()), jc.IsTrue)
	c.Check(uuidsSet.Contains(m.uuid.String()), jc.IsTrue)
}
