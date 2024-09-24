// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package httpserverargs_test

import (
	"context"
	"time"

	"github.com/juju/clock"
	"github.com/juju/clock/testclock"
	"github.com/juju/errors"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v4"
	"github.com/juju/worker/v4/dependency"
	dt "github.com/juju/worker/v4/dependency/testing"
	"github.com/juju/worker/v4/workertest"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/apiserverhttp"
	"github.com/juju/juju/apiserver/authentication"
	"github.com/juju/juju/apiserver/authentication/macaroon"
	accessservice "github.com/juju/juju/domain/access/service"
	controllerconfigservice "github.com/juju/juju/domain/controllerconfig/service"
	macaroonservice "github.com/juju/juju/domain/macaroon/service"
	"github.com/juju/juju/internal/services"
	"github.com/juju/juju/internal/worker/httpserverargs"
	"github.com/juju/juju/state"
	statetesting "github.com/juju/juju/state/testing"
)

type ManifoldSuite struct {
	statetesting.StateSuite

	config         httpserverargs.ManifoldConfig
	manifold       dependency.Manifold
	getter         dependency.Getter
	clock          *testclock.Clock
	stateTracker   stubStateTracker
	authenticator  mockLocalMacaroonAuthenticator
	domainServices stubDomainServices

	stub testing.Stub
}

var _ = gc.Suite(&ManifoldSuite{})

func (s *ManifoldSuite) SetUpTest(c *gc.C) {
	s.StateSuite.SetUpTest(c)

	s.clock = testclock.NewClock(time.Time{})
	s.stateTracker = stubStateTracker{
		pool: s.StatePool,
	}
	s.domainServices = stubDomainServices{}
	s.stub.ResetCalls()

	s.getter = s.newGetter(nil)
	s.config = httpserverargs.ManifoldConfig{
		ClockName:             "clock",
		StateName:             "state",
		DomainServicesName:    "domain-services",
		NewStateAuthenticator: s.newStateAuthenticator,
	}
	s.manifold = httpserverargs.Manifold(s.config)
}

func (s *ManifoldSuite) newGetter(overlay map[string]any) dependency.Getter {
	resources := map[string]any{
		"clock":           s.clock,
		"state":           &s.stateTracker,
		"domain-services": &s.domainServices,
	}
	for k, v := range overlay {
		resources[k] = v
	}
	return dt.StubGetter(resources)
}

func (s *ManifoldSuite) newStateAuthenticator(
	ctx context.Context,
	statePool *state.StatePool,
	modelUUID string,
	controllerConfig httpserverargs.ControllerConfigService,
	accessService httpserverargs.AccessService,
	macaroonService httpserverargs.MacaroonService,
	mux *apiserverhttp.Mux,
	clock clock.Clock,
	abort <-chan struct{},
) (macaroon.LocalMacaroonAuthenticator, error) {
	s.stub.MethodCall(s, "NewStateAuthenticator", ctx, statePool, controllerConfig, accessService, mux, clock, abort)
	if err := s.stub.NextErr(); err != nil {
		return nil, err
	}
	return &s.authenticator, nil
}

var expectedInputs = []string{"state", "clock", "domain-services"}

func (s *ManifoldSuite) TestInputs(c *gc.C) {
	c.Assert(s.manifold.Inputs, jc.SameContents, expectedInputs)
}

func (s *ManifoldSuite) TestMissingInputs(c *gc.C) {
	for _, input := range expectedInputs {
		getter := s.newGetter(map[string]any{
			input: dependency.ErrMissing,
		})
		_, err := s.manifold.Start(context.Background(), getter)
		c.Assert(errors.Cause(err), gc.Equals, dependency.ErrMissing)
	}
}

func (s *ManifoldSuite) TestMuxOutput(c *gc.C) {
	w := s.startWorkerClean(c)
	defer workertest.CleanKill(c, w)

	var mux *apiserverhttp.Mux
	err := s.manifold.Output(w, &mux)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(mux, gc.NotNil)
}

func (s *ManifoldSuite) TestAuthenticatorOutput(c *gc.C) {
	w := s.startWorkerClean(c)
	defer workertest.CleanKill(c, w)

	var auth1 authentication.RequestAuthenticator
	var auth2 macaroon.LocalMacaroonAuthenticator
	for _, out := range []any{&auth1, &auth2} {
		err := s.manifold.Output(w, out)
		c.Assert(err, jc.ErrorIsNil)
	}
	c.Assert(auth1, gc.NotNil)
	c.Assert(auth1, gc.Equals, auth2)
}

func (s *ManifoldSuite) startWorkerClean(c *gc.C) worker.Worker {
	w, err := s.manifold.Start(context.Background(), s.getter)
	c.Assert(err, jc.ErrorIsNil)
	workertest.CheckAlive(c, w)
	return w
}

func (s *ManifoldSuite) TestStopWorkerClosesState(c *gc.C) {
	w := s.startWorkerClean(c)
	defer workertest.CleanKill(c, w)

	s.stateTracker.CheckCallNames(c, "Use")

	workertest.CleanKill(c, w)
	s.stateTracker.CheckCallNames(c, "Use", "Done")
}

func (s *ManifoldSuite) TestStoppingWorkerClosesAuthenticator(c *gc.C) {
	w := s.startWorkerClean(c)
	s.stub.CheckCallNames(c, "NewStateAuthenticator")
	authArgs := s.stub.Calls()[0].Args
	c.Assert(authArgs, gc.HasLen, 7)
	abort := authArgs[6].(<-chan struct{})

	// abort should still be open at this point.
	select {
	case <-abort:
		c.Fatalf("abort closed while worker still running")
	default:
	}

	workertest.CleanKill(c, w)
	select {
	case <-abort:
	default:
		c.Fatalf("authenticator abort channel not closed")
	}
}

func (s *ManifoldSuite) TestValidate(c *gc.C) {
	type test struct {
		f      func(*httpserverargs.ManifoldConfig)
		expect string
	}
	tests := []test{{
		f:      func(cfg *httpserverargs.ManifoldConfig) { cfg.StateName = "" },
		expect: "empty StateName not valid",
	}, {
		f:      func(cfg *httpserverargs.ManifoldConfig) { cfg.ClockName = "" },
		expect: "empty ClockName not valid",
	}, {
		f:      func(cfg *httpserverargs.ManifoldConfig) { cfg.DomainServicesName = "" },
		expect: "empty DomainServicesName not valid",
	}, {
		f:      func(cfg *httpserverargs.ManifoldConfig) { cfg.NewStateAuthenticator = nil },
		expect: "nil NewStateAuthenticator not valid",
	}}

	for i, test := range tests {
		c.Logf("test #%d (%s)", i, test.expect)

		config := s.config
		test.f(&config)

		manifold := httpserverargs.Manifold(config)
		w, err := manifold.Start(context.Background(), s.getter)
		workertest.CheckNilOrKill(c, w)
		c.Check(err, gc.ErrorMatches, test.expect)
	}
}

type mockLocalMacaroonAuthenticator struct {
	macaroon.LocalMacaroonAuthenticator
}

type stubStateTracker struct {
	testing.Stub
	pool  *state.StatePool
	state *state.State
}

func (s *stubStateTracker) Use() (*state.StatePool, *state.State, error) {
	s.MethodCall(s, "Use")
	return s.pool, s.state, s.NextErr()
}

func (s *stubStateTracker) Done() error {
	s.MethodCall(s, "Done")
	return s.NextErr()
}

func (s *stubStateTracker) Report() map[string]any {
	s.MethodCall(s, "Report")
	return nil
}

type stubDomainServices struct {
	testing.Stub
	services.ControllerDomainServices
}

func (s *stubDomainServices) ControllerConfig() *controllerconfigservice.WatchableService {
	s.MethodCall(s, "ControllerConfig")
	return nil
}

func (s *stubDomainServices) Access() *accessservice.Service {
	s.MethodCall(s, "Access")
	return nil
}

func (s *stubDomainServices) Macaroon() *macaroonservice.Service {
	s.MethodCall(s, "Macaroon")
	return nil
}
