// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/juju/errors"
	"github.com/juju/version/v2"

	"github.com/juju/juju/core/database"
	coremodel "github.com/juju/juju/core/model"
	"github.com/juju/juju/domain"
	clouderrors "github.com/juju/juju/domain/cloud/errors"
	"github.com/juju/juju/domain/credential"
	"github.com/juju/juju/domain/model"
	modelerrors "github.com/juju/juju/domain/model/errors"
	usererrors "github.com/juju/juju/domain/user/errors"
	jujudb "github.com/juju/juju/internal/database"
)

// State represents a type for interacting with the underlying model state.
type State struct {
	*domain.StateBase
}

// NewState returns a new State for interacting with the underlying model state.
func NewState(
	factory database.TxnRunnerFactory,
) *State {
	return &State{
		StateBase: domain.NewStateBase(factory),
	}
}

// ModelCloudNameAndCredential returns the cloud name and credential id for a
// model identified by the model name and the owner. If no model exists for the
// provided name and user a [modelerrors.NotFound] error is returned.
func (s *State) ModelCloudNameAndCredential(
	ctx context.Context,
	modelName string,
	modelOwnerName string,
) (string, credential.ID, error) {
	db, err := s.DB()
	if err != nil {
		return "", credential.ID{}, errors.Trace(err)
	}

	stmt := `
SELECT cloud_name, cloud_credential_name, cloud_credential_owner_uuid
FROM v_model
WHERE name = ?
AND owner_name = ?
`

	var (
		cloudName       string
		credentialName  sql.NullString
		credentialOwner sql.NullString
	)
	err = db.StdTxn(ctx, func(ctx context.Context, tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, stmt, modelName, modelOwnerName)
		if err := row.Scan(&cloudName, &credentialName, &credentialOwner); err != nil {
			return err
		}
		return nil
	})

	if errors.Is(err, sql.ErrNoRows) {
		return "", credential.ID{}, fmt.Errorf("%w for name %q and owner %q",
			modelerrors.NotFound, modelName, modelOwnerName,
		)
	} else if err != nil {
		return "", credential.ID{}, fmt.Errorf(
			"getting cloud name and credential for model %q with owner %q: %w",
			modelName, modelOwnerName, domain.CoerceError(err),
		)
	}

	if !credentialName.Valid {
		return cloudName, credential.ID{}, nil
	}

	credId := credential.ID{
		Cloud: cloudName,
		Name:  credentialName.String,
		Owner: credentialOwner.String,
	}

	return cloudName, credId, nil
}

// CloudType is responsible for reporting the type for a given cloud name. If no
// cloud exists for the provided name then an error of [clouderrors.NotFound]
// will be returned.
func (s *State) CloudType(
	ctx context.Context,
	name string,
) (string, error) {
	db, err := s.DB()
	if err != nil {
		return "", errors.Trace(err)
	}

	ctFunc := CloudType()

	var cloudType string
	return cloudType, db.StdTxn(ctx, func(ctx context.Context, tx *sql.Tx) error {
		var err error
		cloudType, err = ctFunc(ctx, tx, name)
		return err
	})
}

// CloudType is responsible for reporting the type for a given cloud name. If no
// cloud exists for the provided name then an error of [clouderrors.NotFound]
// will be returned.
func CloudType() func(context.Context, *sql.Tx, string) (string, error) {
	stmt := `
SELECT ct.type
FROM cloud_type ct
INNER JOIN cloud c
ON c.cloud_type_id = ct.id
WHERE c.name = ?
`

	return func(ctx context.Context, tx *sql.Tx, name string) (string, error) {
		var cloudType string
		err := tx.QueryRowContext(ctx, stmt, name).Scan(&cloudType)
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%w for name %q", clouderrors.NotFound, name)
		} else if err != nil {
			return "", fmt.Errorf("determing type for cloud %q: %w", name, domain.CoerceError(err))
		}
		return cloudType, nil
	}
}

// Create is responsible for creating a new moddel from start to finish. It will
// register the model existence and associate all of the model metadata.
// If a model already exists with the same name and owner then an error
// satisfying modelerrors.AlreadyExists will be returned.
// If the model type is not found then an error satisfying errors.NotSupported
// will be returned.
func (s *State) Create(
	ctx context.Context,
	uuid coremodel.UUID,
	modelType coremodel.ModelType,
	input model.ModelCreationArgs,
) error {
	db, err := s.DB()
	if err != nil {
		return errors.Trace(err)
	}

	return db.StdTxn(ctx, func(ctx context.Context, tx *sql.Tx) error {
		return Create(ctx, tx, uuid, modelType, input)
	})
}

// Create is responsible for creating a new model from start to finish. It will
// register the model existence and associate all of the model metadata.
// If a model already exists with the same name and owner then an error
// satisfying modelerrors.AlreadyExists will be returned.
// If the model type is not found then an error satisfying errors.NotSupported
// will be returned.
func Create(
	ctx context.Context,
	tx *sql.Tx,
	uuid coremodel.UUID,
	modelType coremodel.ModelType,
	input model.ModelCreationArgs,
) error {
	if err := createModel(ctx, tx, uuid, modelType, input); err != nil {
		return err
	}

	if err := createModelAgent(ctx, tx, uuid, input.AgentVersion); err != nil {
		return err
	}

	return nil
}

// createModelAgent is responsible for create a new model's agent record for the
// given model UUID. If a model agent record already exists for the given model
// uuid then an error satisfying [modelerrors.AlreadyExists] is returned. If no
// model exists for the provided UUID then a [modelerrors.NotFound] is returned.
func createModelAgent(
	ctx context.Context,
	tx *sql.Tx,
	modelUUID coremodel.UUID,
	agentVersion version.Number,
) error {
	stmt := `
INSERT INTO model_agent (model_uuid, previous_version, target_version)
    VALUES (?, ?, ?)
`

	res, err := tx.ExecContext(ctx, stmt, modelUUID, agentVersion.String(), agentVersion.String())
	if jujudb.IsErrConstraintPrimaryKey(err) {
		return fmt.Errorf(
			"%w for uuid %q while setting model agent version",
			modelerrors.AlreadyExists, modelUUID,
		)
	} else if jujudb.IsErrConstraintForeignKey(err) {
		return fmt.Errorf(
			"%w for uuid %q while setting model agent version",
			modelerrors.NotFound,
			modelUUID,
		)
	} else if err != nil {
		return fmt.Errorf("creating model %q agent information: %w", modelUUID, err)
	}

	if num, err := res.RowsAffected(); err != nil {
		return errors.Trace(err)
	} else if num != 1 {
		return fmt.Errorf("creating model agent record, expected 1 row to be inserted got %d", num)
	}

	return nil
}

// createModel is responsible for creating a new model record
// for the given model UUID. If a model record already exists for the
// given model uuid then an error satisfying modelerrors.AlreadyExists is
// returned. Conversely should the owner already have a model that exists with
// the provided name then a modelerrors.AlreadyExists error will be returned. If
// the model type supplied is not found then a errors.NotSupported error is
// returned.
//
// Should the provided cloud and region not be found an error matching
// errors.NotFound will be returned.
// If the ModelCreationArgs contains a non zero value cloud credential this func
// will also attempt to set the model cloud credential using updateCredential. In
// this  scenario the errors from updateCredential are also possible.
// If the model owner does not exist an error satisfying [usererrors.NotFound]
// will be returned.
func createModel(
	ctx context.Context,
	tx *sql.Tx,
	uuid coremodel.UUID,
	modelType coremodel.ModelType,
	input model.ModelCreationArgs,
) error {
	cloudStmt := `
SELECT uuid
FROM cloud
WHERE name = ?
`
	var cloudUUID string
	err := tx.QueryRowContext(ctx, cloudStmt, input.Cloud).Scan(&cloudUUID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w cloud %q", errors.NotFound, input.Cloud)
	} else if err != nil {
		return fmt.Errorf("getting cloud %q uuid: %w", input.Cloud, err)
	}

	userStmt := `
SELECT uuid
FROM user
WHERE uuid = ?
AND removed = false
`
	var userUUID string
	err = tx.QueryRowContext(ctx, userStmt, input.Owner).Scan(&userUUID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w for model owner %q", usererrors.NotFound, input.Owner)
	} else if err != nil {
		return fmt.Errorf("getting user uuid for setting model %q owner: %w", input.Name, err)
	}

	stmt := `
INSERT INTO model (uuid,
            cloud_uuid,
            model_type_id,
            name,
            owner_uuid)
SELECT ?, ?, model_type.id, ?, ?
FROM model_type
WHERE model_type.type = ?
`

	res, err := tx.ExecContext(ctx, stmt,
		uuid, cloudUUID, input.Name, input.Owner, modelType,
	)
	if jujudb.IsErrConstraintPrimaryKey(err) {
		return fmt.Errorf("%w for uuid %q", modelerrors.AlreadyExists, uuid)
	} else if jujudb.IsErrConstraintForeignKey(err) {
		return fmt.Errorf("%w for uuid %q", modelerrors.NotFound, uuid)
	} else if jujudb.IsErrConstraintUnique(err) {
		return fmt.Errorf("%w for name %q and owner %q", modelerrors.AlreadyExists, input.Name, input.Owner)
	} else if err != nil {
		return fmt.Errorf("setting model %q metadata: %w", uuid, err)
	}

	if num, err := res.RowsAffected(); err != nil {
		return errors.Trace(err)
	} else if num != 1 {
		return fmt.Errorf("creating model metadata, expected 1 row to be inserted, got %d", num)
	}

	if input.CloudRegion != "" {
		err := setCloudRegion(ctx, uuid, input.CloudRegion, tx)
		if err != nil {
			return err
		}
	}

	if !input.Credential.IsZero() {
		err := updateCredential(ctx, tx, uuid, input.Credential)
		if err != nil {
			return err
		}
	}
	return nil
}

// Delete will remove all data associated with the provided model uuid removing
// the models existence from Juju. If the model does not exist then a error
// satisfying modelerrors.NotFound will be returned.
func (s *State) Delete(
	ctx context.Context,
	uuid coremodel.UUID,
) error {
	db, err := s.DB()
	if err != nil {
		return errors.Trace(err)
	}

	deleteModelAgent := "DELETE FROM model_agent WHERE model_uuid = ?"
	deleteModel := "DELETE FROM model WHERE uuid = ?"
	return db.StdTxn(ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, deleteModelAgent, uuid)
		if err != nil {
			return fmt.Errorf("delete model %q agent: %w", uuid, err)
		}

		res, err := tx.ExecContext(ctx, deleteModel, uuid)
		if err != nil {
			return fmt.Errorf("deleting model %q metadata: %w", uuid, err)
		}

		if num, err := res.RowsAffected(); err != nil {
			return errors.Trace(err)
		} else if num != 1 {
			return fmt.Errorf("%w %q", modelerrors.NotFound, uuid)
		}
		return nil
	})
}

// Finalise is responsible for setting a model as fully constructed and
// indicates the final system state for the model is ready for use. This is used
// because the model creation process involves several transactions with which
// anyone could fail at a given time.
//
// If no model exists for the provided id then a [modelerrors.NotFound] will be
// returned. If the model as previously been finalised a
// [modelerrors.AlreadyFinalised] error will be returned.
func (s *State) Finalise(ctx context.Context, uuid coremodel.UUID) error {
	db, err := s.DB()
	if err != nil {
		return errors.Trace(err)
	}

	finalise := Finalise()

	return db.StdTxn(ctx, func(ctx context.Context, tx *sql.Tx) error {
		return finalise(ctx, tx, uuid)
	})
}

// FinaliseFunc is responsible for setting a model as fully constructed and
// indicates the final system state for the model is ready for use. This is used
// because the model creation process involves several transactions with which
// anyone could fail at a given time.
//
// If no model exists for the provided id then a [modelerrors.NotFound] will be
// returned. If the model as previously been finalised a
// [modelerrors.AlreadyFinalised] error will be returned.
type FinaliseFunc func(context.Context, *sql.Tx, coremodel.UUID) error

// Finalise constructs a [FinaliseFunc] that can safely be used over several
// transaction retry's.
func Finalise() FinaliseFunc {
	existsStmt := `
SELECT finalised FROM model WHERE uuid = ?
`

	stmt := `
UPDATE model 
SET finalised = TRUE
WHERE uuid = ?
`

	return func(ctx context.Context, tx *sql.Tx, uuid coremodel.UUID) error {
		var finalised bool
		err := tx.QueryRowContext(ctx, existsStmt, uuid).Scan(&finalised)
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%w for id %q", modelerrors.NotFound, uuid)
		} else if err != nil {
			return fmt.Errorf("determining finalised status for model with id %q: %w", uuid, err)
		}

		if finalised {
			return fmt.Errorf("%w for id %q", modelerrors.AlreadyFinalised, uuid)
		}

		if _, err := tx.ExecContext(ctx, stmt, uuid); err != nil {
			return fmt.Errorf("finalising model with id %q: %w", uuid, err)
		}
		return nil
	}
}

// GetModelTypes returns the slice of model.Type's supported by state.
func (s *State) GetModelTypes(ctx context.Context) ([]coremodel.ModelType, error) {
	rval := []coremodel.ModelType{}

	db, err := s.DB()
	if err != nil {
		return rval, errors.Trace(err)
	}

	stmt := `
SELECT type FROM model_type;
	`

	return rval, db.StdTxn(ctx, func(ctx context.Context, tx *sql.Tx) error {
		rows, err := tx.QueryContext(ctx, stmt)
		if err != nil {
			return fmt.Errorf("getting supported model types: %w", err)
		}
		defer rows.Close()

		var t coremodel.ModelType
		for rows.Next() {
			if err := rows.Scan(&t); err != nil {
				return fmt.Errorf("building model type: %w", err)
			}
			rval = append(rval, t)
		}
		return nil
	})
}

// List returns a list of all model UUIDs in the system that have not been
// deleted.
func (s *State) List(ctx context.Context) ([]coremodel.UUID, error) {
	db, err := s.DB()
	if err != nil {
		return nil, errors.Trace(err)
	}

	var models []coremodel.UUID
	err = db.StdTxn(ctx, func(ctx context.Context, tx *sql.Tx) error {
		stmt := `SELECT uuid FROM model_list;`
		rows, err := tx.QueryContext(ctx, stmt)
		if err != nil {
			return errors.Trace(err)
		}
		defer rows.Close()

		for rows.Next() {
			var model coremodel.UUID
			if err := rows.Scan(&model); err != nil {
				return errors.Trace(err)
			}
			if err := rows.Err(); err != nil {
				return errors.Trace(err)
			}
			models = append(models, model)
		}
		return nil
	})
	return models, errors.Trace(err)
}

// setCloudRegion is responsible for setting a model's cloud region. This
// operation can only be performed once and will fail with an error that
// satisfies errors.AlreadyExists on subsequent tries.
// If no cloud region is found for the model's cloud then an error that satisfies
// errors.NotFound will be returned.
func setCloudRegion(
	ctx context.Context,
	uuid coremodel.UUID,
	region string,
	tx *sql.Tx,
) error {
	cloudRegionStmt := `
SELECT cr.uuid
FROM cloud_region cr
INNER JOIN cloud c
ON c.uuid = cr.cloud_uuid
INNER JOIN model m
ON m.cloud_uuid = c.uuid
WHERE m.uuid = ?
AND cr.name = ?
`

	var cloudRegionUUID string
	err := tx.QueryRowContext(ctx, cloudRegionStmt, uuid, region).
		Scan(&cloudRegionUUID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf(
			"%w cloud region %q for model uuid %q",
			errors.NotFound,
			region,
			uuid,
		)
	} else if err != nil {
		return fmt.Errorf(
			"getting cloud region %q uuid for model %q: %w",
			region,
			uuid,
			err,
		)
	}

	modelMetadataStmt := `
UPDATE model
SET cloud_region_uuid = ?
WHERE uuid = ?
AND cloud_region_uuid IS NULL
`

	res, err := tx.ExecContext(ctx, modelMetadataStmt, cloudRegionUUID, uuid)
	if err != nil {
		return fmt.Errorf(
			"setting cloud region uuid %q for model uuid %q: %w",
			cloudRegionUUID,
			uuid,
			err,
		)
	}
	if num, err := res.RowsAffected(); err != nil {
		return errors.Trace(err)
	} else if num != 1 {
		return fmt.Errorf(
			"model %q already has a cloud region set%w",
			uuid,
			errors.Hide(errors.AlreadyExists),
		)
	}
	return nil
}

// UpdateCredential is responsible for updating the cloud credential in use
// by model. If the cloud credential is not found an error that satisfies
// errors.NotFound is returned.
// If the credential being updated to is not of the same cloud that is currently
// set for the model then an error that satisfies errors.NotValid is returned.
func (s *State) UpdateCredential(
	ctx context.Context,
	uuid coremodel.UUID,
	id credential.ID,
) error {
	db, err := s.DB()
	if err != nil {
		return errors.Trace(err)
	}

	return db.StdTxn(ctx, func(ctx context.Context, tx *sql.Tx) error {
		return updateCredential(ctx, tx, uuid, id)
	})
}

//func (s *State) UpsertCloudCredential(
//	ctx context.Context,
//	cloud string,
//	owner user.UUID,
//	name string,
//	credential.CloudCredentialInfo,
//) error {
//	db, err := s.DB()
//	if err != nil {
//		return errors.Trace(err)
//	}
//
//
//}

// updateCredential is responsible for updating the cloud credential in use
// by model. If the cloud credential is not found an error that satisfies
// errors.NotFound is returned.
// If the credential being updated to is not of the same cloud that is currently
// set for the model then an error that satisfies errors.NotValid is returned.
func updateCredential(
	ctx context.Context,
	tx *sql.Tx,
	uuid coremodel.UUID,
	id credential.ID,
) error {
	cloudCredUUIDStmt := `
SELECT cc.uuid,
       c.uuid
FROM cloud_credential cc
INNER JOIN cloud c
ON c.uuid = cc.cloud_uuid
INNER JOIN user u
ON cc.owner_uuid = u.uuid
WHERE c.name = ?
AND u.name = ?
AND u.removed = false
AND cc.name = ?
`

	stmt := `
UPDATE model
SET cloud_credential_uuid = ?
WHERE uuid= ?
AND cloud_uuid = ?
`

	var cloudCredUUID, cloudUUID string
	err := tx.QueryRowContext(ctx, cloudCredUUIDStmt, id.Cloud, id.Owner, id.Name).
		Scan(&cloudCredUUID, &cloudUUID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf(
			"%w cloud credential %q%w",
			errors.NotFound, id, errors.Hide(err),
		)
	} else if err != nil {
		return fmt.Errorf(
			"getting cloud credential uuid for %q: %w",
			id, err,
		)
	}

	res, err := tx.ExecContext(ctx, stmt, cloudCredUUID, uuid, cloudUUID)
	if err != nil {
		return fmt.Errorf(
			"setting cloud credential %q for model %q: %w",
			id, uuid, err)
	}

	if num, err := res.RowsAffected(); err != nil {
		return errors.Trace(err)
	} else if num != 1 {
		return fmt.Errorf(
			"%w model %q has different cloud to credential %q",
			errors.NotValid, uuid, id)
	}
	return nil
}
