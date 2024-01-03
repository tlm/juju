// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package model_test

import (
	"bytes"
	"time"

	"github.com/juju/clock/testclock"
	"github.com/juju/cmd/v3"
	"github.com/juju/cmd/v3/cmdtesting"
	"github.com/juju/errors"
	"github.com/juju/names/v5"
	jutesting "github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/api/base"
	apiservererrors "github.com/juju/juju/apiserver/errors"
	"github.com/juju/juju/cmd/juju/model"
	"github.com/juju/juju/cmd/modelcmd"
	coremodel "github.com/juju/juju/core/model"
	"github.com/juju/juju/jujuclient"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/testing"
)

type DestroySuite struct {
	testing.FakeJujuXDGDataHomeSuite
	api             *fakeAPI
	stub            *jutesting.Stub
	budgetAPIClient *mockBudgetAPIClient
	store           *jujuclient.MemStore

	clock testclock.AdvanceableClock
}

var _ = gc.Suite(&DestroySuite{})

// fakeDestroyAPI mocks out the client API
type fakeAPI struct {
	*jutesting.Stub
	statusCallCount    int
	modelInfoErr       []*params.Error
	modelStatusPayload []base.ModelStatus
}

func (f *fakeAPI) Close() error { return nil }

func (f *fakeAPI) DestroyModel(tag names.ModelTag, destroyStorage *bool, force *bool, maxWait *time.Duration, timeout *time.Duration) error {
	f.MethodCall(f, "DestroyModel", tag, destroyStorage, force, maxWait, timeout)
	return f.NextErr()
}

func (f *fakeAPI) ModelStatus(models ...names.ModelTag) ([]base.ModelStatus, error) {
	var err error
	if f.statusCallCount < len(f.modelInfoErr) {
		modelInfoErr := f.modelInfoErr[f.statusCallCount]
		if modelInfoErr != nil {
			err = modelInfoErr
		}
	} else {
		err = &params.Error{Code: params.CodeNotFound}
	}
	f.statusCallCount++

	if f.modelStatusPayload == nil {
		f.modelStatusPayload = []base.ModelStatus{{
			Volumes: []base.Volume{
				{Detachable: true},
				{Detachable: true},
			},
			Filesystems: []base.Filesystem{{Detachable: true}},
		}}
	}
	return f.modelStatusPayload, err
}

func (s *DestroySuite) SetUpTest(c *gc.C) {
	s.FakeJujuXDGDataHomeSuite.SetUpTest(c)
	s.stub = &jutesting.Stub{}
	s.api = &fakeAPI{
		Stub: s.stub,
	}
	s.clock = testclock.NewDilatedWallClock(100 * time.Millisecond)

	s.store = jujuclient.NewMemStore()
	s.store.CurrentControllerName = "test1"
	s.store.Controllers["test1"] = jujuclient.ControllerDetails{ControllerUUID: "test1-uuid"}
	s.store.Models["test1"] = &jujuclient.ControllerModels{
		Models: map[string]jujuclient.ModelDetails{
			"admin/test1": {ModelUUID: "test1-uuid", ModelType: coremodel.IAAS},
			"admin/test2": {ModelUUID: "test2-uuid", ModelType: coremodel.IAAS},
		},
	}
	s.store.Accounts["test1"] = jujuclient.AccountDetails{
		User: "admin",
	}

	s.budgetAPIClient = &mockBudgetAPIClient{Stub: s.stub}
}

func (s *DestroySuite) runDestroyCommand(c *gc.C, args ...string) (*cmd.Context, error) {
	command := model.NewDestroyCommandForTest(s.api, s.clock, noOpRefresh, s.store)
	return cmdtesting.RunCommand(c, command, args...)
}

func (s *DestroySuite) NewDestroyCommand() cmd.Command {
	return model.NewDestroyCommandForTest(s.api, s.clock, noOpRefresh, s.store)
}

func checkModelExistsInStore(c *gc.C, name string, store jujuclient.ClientStore) {
	controller, amodel := modelcmd.SplitModelName(name)
	_, err := store.ModelByName(controller, amodel)
	c.Assert(err, jc.ErrorIsNil)
}

func checkModelRemovedFromStore(c *gc.C, name string, store jujuclient.ClientStore) {
	controller, amodel := modelcmd.SplitModelName(name)
	_, err := store.ModelByName(controller, amodel)
	c.Assert(err, jc.ErrorIs, errors.NotFound)
}

func (s *DestroySuite) TestDestroyNoModelNameError(c *gc.C) {
	_, err := s.runDestroyCommand(c)
	c.Assert(err, gc.ErrorMatches, "no model specified")
}

func (s *DestroySuite) TestDestroyBadFlags(c *gc.C) {
	_, err := s.runDestroyCommand(c, "-n")
	c.Assert(err, gc.ErrorMatches, "option provided but not defined: -n")
}

func (s *DestroySuite) TestDestroyUnknownArgument(c *gc.C) {
	_, err := s.runDestroyCommand(c, "model", "whoops")
	c.Assert(err, gc.ErrorMatches, `unrecognized args: \["whoops"\]`)
}

func (s *DestroySuite) TestDestroyMaxWaitWithoutForce(c *gc.C) {
	_, err := s.runDestroyCommand(c, "model", "--no-wait")
	c.Assert(err, gc.ErrorMatches, `--no-wait without --force not valid`)
}

func (s *DestroySuite) TestDestroyUnknownModelCallsRefresh(c *gc.C) {
	called := false
	refresh := func(jujuclient.ClientStore, string) error {
		called = true
		return nil
	}

	command := model.NewDestroyCommandForTest(s.api, s.clock, refresh, s.store)
	_, err := cmdtesting.RunCommand(c, command, "foo")
	c.Check(called, jc.IsTrue)
	c.Check(err, gc.ErrorMatches, `model test1:admin/foo not found`)
}

func (s *DestroySuite) TestDestroyCannotConnectToAPI(c *gc.C) {
	s.stub.SetErrors(errors.New("connection refused"))
	_, err := s.runDestroyCommand(c, "test2", "--no-prompt")
	c.Assert(err, gc.ErrorMatches, "cannot destroy model: connection refused")
	c.Check(c.GetTestLog(), jc.Contains, "failed to destroy model \"test2\"")
	checkModelExistsInStore(c, "test1:admin/test2", s.store)
}

func (s *DestroySuite) TestSystemDestroyFails(c *gc.C) {
	_, err := s.runDestroyCommand(c, "test1", "--no-prompt")
	c.Assert(err, gc.ErrorMatches, `"test1" is a controller; use 'juju destroy-controller' to destroy it`)
	checkModelExistsInStore(c, "test1:admin/test1", s.store)
}

func (s *DestroySuite) TestDestroy(c *gc.C) {
	checkModelExistsInStore(c, "test1:admin/test2", s.store)
	_, err := s.runDestroyCommand(c, "test2", "--no-prompt")
	c.Assert(err, jc.ErrorIsNil)
	checkModelRemovedFromStore(c, "test1:admin/test2", s.store)
	s.stub.CheckCalls(c, []jutesting.StubCall{
		{"DestroyModel",
			[]interface{}{names.NewModelTag("test2-uuid"), (*bool)(nil), (*bool)(nil), (*time.Duration)(nil), (*time.Duration)(nil)}},
	})
}

func (s *DestroySuite) TestDestroyWithPartModelUUID(c *gc.C) {
	checkModelExistsInStore(c, "test1:admin/test2", s.store)
	_, err := s.runDestroyCommand(c, "test2-uu", "--no-prompt")
	c.Assert(err, jc.ErrorIsNil)
	checkModelRemovedFromStore(c, "test1:admin/test2", s.store)
	s.stub.CheckCalls(c, []jutesting.StubCall{
		{"DestroyModel",
			[]interface{}{names.NewModelTag("test2-uuid"), (*bool)(nil), (*bool)(nil), (*time.Duration)(nil), (*time.Duration)(nil)}},
	})
}

func (s *DestroySuite) TestDestroyWithForce(c *gc.C) {
	checkModelExistsInStore(c, "test1:admin/test2", s.store)
	_, err := s.runDestroyCommand(c, "test2", "--no-prompt", "--force")
	c.Assert(err, jc.ErrorIsNil)
	checkModelRemovedFromStore(c, "test1:admin/test2", s.store)
	force := true
	s.stub.CheckCalls(c, []jutesting.StubCall{
		{"DestroyModel", []interface{}{names.NewModelTag("test2-uuid"), (*bool)(nil), &force, (*time.Duration)(nil), (*time.Duration)(nil)}},
	})
}

func (s *DestroySuite) TestDestroyWithForceTimeout(c *gc.C) {
	checkModelExistsInStore(c, "test1:admin/test2", s.store)
	_, err := s.runDestroyCommand(c, "test2", "--no-prompt", "--force", "--timeout", "30m")
	c.Assert(err, jc.ErrorIsNil)
	checkModelRemovedFromStore(c, "test1:admin/test2", s.store)
	force := true
	timeout := 30 * time.Minute
	s.stub.CheckCalls(c, []jutesting.StubCall{
		{"DestroyModel", []interface{}{names.NewModelTag("test2-uuid"), (*bool)(nil), &force, (*time.Duration)(nil), &timeout}},
	})
}

func (s *DestroySuite) TestDestroyWithTimeoutNoForce(c *gc.C) {
	checkModelExistsInStore(c, "test1:admin/test2", s.store)
	_, err := s.runDestroyCommand(c, "test2", "--no-prompt", "--timeout", "30m")
	c.Assert(err, gc.ErrorMatches, `--timeout can only be used with --force \(dangerous\)`)
}

func (s *DestroySuite) TestDestroyWithForceNoWait(c *gc.C) {
	checkModelExistsInStore(c, "test1:admin/test2", s.store)
	_, err := s.runDestroyCommand(c, "test2", "--no-prompt", "--force", "--no-wait")
	c.Assert(err, jc.ErrorIsNil)
	checkModelRemovedFromStore(c, "test1:admin/test2", s.store)
	force := true
	maxWait := 0 * time.Second
	s.stub.CheckCalls(c, []jutesting.StubCall{
		{"DestroyModel", []interface{}{names.NewModelTag("test2-uuid"), (*bool)(nil), &force, &maxWait, (*time.Duration)(nil)}},
	})
}

func (s *DestroySuite) TestDestroyBlocks(c *gc.C) {
	checkModelExistsInStore(c, "test1:admin/test2", s.store)
	s.api.modelInfoErr = []*params.Error{{}, {Code: params.CodeNotFound}}
	_, err := s.runDestroyCommand(c, "test2", "--no-prompt")
	c.Assert(err, jc.ErrorIsNil)
	checkModelRemovedFromStore(c, "test1:admin/test2", s.store)
	c.Assert(s.api.statusCallCount, gc.Equals, 1)
}

func (s *DestroySuite) TestFailedDestroyModel(c *gc.C) {
	s.stub.SetErrors(errors.New("permission denied"))
	_, err := s.runDestroyCommand(c, "test1:test2", "--no-prompt")
	c.Assert(err, gc.ErrorMatches, "cannot destroy model: permission denied")
	checkModelExistsInStore(c, "test1:admin/test2", s.store)
}

func (s *DestroySuite) TestDestroyDestroyStorage(c *gc.C) {
	_, err := s.runDestroyCommand(c, "test2", "--no-prompt", "--destroy-storage")
	c.Assert(err, jc.ErrorIsNil)
	destroyStorage := true
	s.stub.CheckCalls(c, []jutesting.StubCall{
		{"DestroyModel", []interface{}{names.NewModelTag("test2-uuid"), &destroyStorage, (*bool)(nil), (*time.Duration)(nil), (*time.Duration)(nil)}},
	})
}

func (s *DestroySuite) TestDestroyReleaseStorage(c *gc.C) {
	_, err := s.runDestroyCommand(c, "test2", "--no-prompt", "--release-storage")
	c.Assert(err, jc.ErrorIsNil)
	destroyStorage := false
	s.stub.CheckCalls(c, []jutesting.StubCall{
		{"DestroyModel", []interface{}{names.NewModelTag("test2-uuid"), &destroyStorage, (*bool)(nil), (*time.Duration)(nil), (*time.Duration)(nil)}},
	})
}

func (s *DestroySuite) TestDestroyDestroyReleaseStorageFlagsMutuallyExclusive(c *gc.C) {
	_, err := s.runDestroyCommand(c, "test2", "--no-prompt", "--destroy-storage", "--release-storage")
	c.Assert(err, gc.ErrorMatches, "--destroy-storage and --release-storage cannot both be specified")
}

func (s *DestroySuite) TestDestroyDestroyStorageFlagUnspecified(c *gc.C) {
	s.stub.SetErrors(&params.Error{Code: params.CodeHasPersistentStorage})
	s.api.modelInfoErr = []*params.Error{nil}
	_, err := s.runDestroyCommand(c, "test2", "--no-prompt")
	c.Assert(err, gc.ErrorMatches, `cannot destroy model "test2"

The model has persistent storage remaining:
	2 volumes and 1 filesystem

To destroy the storage, run the destroy-model
command again with the "--destroy-storage" option.

To release the storage from Juju's management
without destroying it, use the "--release-storage"
option instead. The storage can then be imported
into another Juju model.

`)
}

func (s *DestroySuite) resetModel(c *gc.C) {
	s.store.Models["test1"] = &jujuclient.ControllerModels{
		Models: map[string]jujuclient.ModelDetails{
			"admin/test1": {ModelUUID: "test1-uuid", ModelType: coremodel.IAAS},
			"admin/test2": {ModelUUID: "test2-uuid", ModelType: coremodel.IAAS},
		},
	}
}

func (s *DestroySuite) TestDestroyCommandConfirmation(c *gc.C) {
	var stdin, stdout, stderr bytes.Buffer
	ctx, err := cmd.DefaultContext()
	c.Assert(err, jc.ErrorIsNil)
	ctx.Stdout = &stdout
	ctx.Stdin = &stdin
	ctx.Stderr = &stderr
	s.api.modelInfoErr = []*params.Error{nil, nil, nil}

	// Ensure confirmation is requested if "--no-prompt" is not specified.
	stdin.WriteString("n")
	errc := cmdtesting.RunCommandWithContext(ctx, s.NewDestroyCommand(), "test2")
	select {
	case err := <-errc:
		c.Check(err, gc.ErrorMatches, "model destruction: aborted")
	case <-time.After(testing.LongWait):
		c.Fatalf("command took too long")
	}
	testLog := c.GetTestLog()
	c.Check(testLog, gc.Matches, "(.|\n)*WARNING.*test2(.|\n)*")
	checkModelExistsInStore(c, "test1:admin/test1", s.store)

	// EOF on stdin: equivalent to answering no.
	stdin.Reset()
	stdout.Reset()
	stderr.Reset()
	errc = cmdtesting.RunCommandWithContext(ctx, s.NewDestroyCommand(), "test2")
	select {
	case err := <-errc:
		c.Check(err, gc.ErrorMatches, "model destruction: aborted")
	case <-time.After(testing.LongWait):
		c.Fatalf("command took too long")
	}
	testLog = c.GetTestLog()
	c.Check(testLog, gc.Matches, "(.|\n)*WARNING.*test2(.|\n)*")
	checkModelExistsInStore(c, "test1:admin/test2", s.store)

	answer := "test2"
	stdin.Reset()
	stdout.Reset()
	stderr.Reset()
	stdin.WriteString(answer)
	errc = cmdtesting.RunCommandWithContext(ctx, s.NewDestroyCommand(), "test2")
	select {
	case err := <-errc:
		c.Check(err, jc.ErrorIsNil)
	case <-time.After(testing.LongWait):
		c.Fatalf("command took too long")
	}
	checkModelRemovedFromStore(c, "test1:admin/test2", s.store)

	// Add the test2 model back into the store for the next test
	s.resetModel(c)

}

func (s *DestroySuite) TestBlockedDestroy(c *gc.C) {
	s.stub.SetErrors(apiservererrors.OperationBlockedError("TestBlockedDestroy"))
	_, err := s.runDestroyCommand(c, "test2", "--no-prompt")
	testing.AssertOperationWasBlocked(c, err, ".*TestBlockedDestroy.*")
}

// mockBudgetAPIClient implements the budgetAPIClient interface.
type mockBudgetAPIClient struct {
	*jutesting.Stub
}

func (c *mockBudgetAPIClient) DeleteBudget(model string) (string, error) {
	c.MethodCall(c, "DeleteBudget", model)
	return "Budget removed.", c.NextErr()
}
