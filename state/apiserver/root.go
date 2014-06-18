// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package apiserver

import (
	"reflect"
	"time"

	"github.com/juju/errors"

	"github.com/juju/juju/rpc"
	"github.com/juju/juju/rpc/rpcreflect"
	"github.com/juju/juju/state"
	"github.com/juju/juju/state/apiserver/common"
)

type taggedAuthenticator interface {
	state.Entity
	state.Authenticator
}

var (
	// maxClientPingInterval defines the timeframe until the ping timeout
	// closes the monitored connection. TODO(mue): Idea by Roger:
	// Move to API (e.g. params) so that the pinging there may
	// depend on the interval.
	maxClientPingInterval = 3 * time.Minute

	// mongoPingInterval defines the interval at which an API server
	// will ping the mongo session to make sure that it's still
	// alive. When the ping returns an error, the server will be
	// terminated.
	mongoPingInterval = 10 * time.Second
)

type objectKey struct {
	name    string
	version int
	objId   string
}

// srvRoot represents a single client's connection to the state
// after it has logged in.
type srvRoot struct {
	state       *state.State
	rpcConn     *rpc.Conn
	resources   *common.Resources
	entity      taggedAuthenticator
	objectCache map[objectKey]interface{}
}

// newSrvRoot creates the client's connection representation
// and starts a ping timeout for the monitoring of this
// connection.
func newSrvRoot(root *initialRoot, entity taggedAuthenticator) *srvRoot {
	r := &srvRoot{
		state:       root.srv.state,
		rpcConn:     root.rpcConn,
		resources:   common.NewResources(),
		entity:      entity,
		objectCache: make(map[objectKey]interface{}),
	}
	r.resources.RegisterNamed("dataDir", common.StringResource(root.srv.dataDir))
	return r
}

// Kill implements rpc.Killer.  It cleans up any resources that need
// cleaning up to ensure that all outstanding requests return.
func (r *srvRoot) Kill() {
	r.resources.StopAll()
}

// srvCaller is our implementation of the rpcreflect.MethodCaller interface.
// It lives just long enough to encapsulate the methods that should be
// available for an RPC call and allow the RPC code to instantiate an object
// and place a call on its method.
type srvCaller struct {
	objMethod rpcreflect.ObjMethod
	creator   func(id string) (interface{}, error)
}

// ParamsType defines the parameters that should be supplied to this function.
// See rpcreflect.MethodCaller for more detail.
func (s *srvCaller) ParamsType() reflect.Type {
	return s.objMethod.Params
}

// ReturnType defines the object that is returned from the function.`
// See rpcreflect.MethodCaller for more detail.
func (s *srvCaller) ResultType() reflect.Type {
	return s.objMethod.Result
}

// Call takes the object Id and an instance of ParamsType to create an object and place
// a call on its method. It then returns an instance of ResultType.
func (s *srvCaller) Call(objId string, arg reflect.Value) (reflect.Value, error) {
	obj, err := s.creator(objId)
	if err != nil {
		return reflect.Value{}, err
	}
	return s.objMethod.Call(reflect.ValueOf(obj), arg)
}

// FindMethod looks up the given rootName and version in our facade registry
// and returns a MethodCaller that will be used by the RPC code to place calls on
// that facade.
// FindMethod uses the global registry state/apiserver/common.Facades.
// For more information about how FindMethod should work, see rpc/server.go and
// rpc/rpcreflect/value.go
func (r *srvRoot) FindMethod(rootName string, version int, methodName string) (rpcreflect.MethodCaller, error) {
	goType, err := common.Facades.GetType(rootName, version)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, &rpcreflect.CallNotImplementedError{
				RootMethod: rootName,
				Version:    version,
			}
		}
		return nil, err
	}
	rpcType := rpcreflect.ObjTypeOf(goType)
	objMethod, err := rpcType.Method(methodName)
	if err != nil {
		if err == rpcreflect.ErrMethodNotFound {
			return nil, &rpcreflect.CallNotImplementedError{
				RootMethod: rootName,
				Version:    version,
				Method:     methodName,
			}
		}
		return nil, err
	}
	creator := func(id string) (interface{}, error) {
		objKey := objectKey{name: rootName, version: version, objId: id}
		if obj, ok := r.objectCache[objKey]; ok {
			return obj, nil
		}
		factory, err := common.Facades.GetFactory(rootName, version)
		if err != nil {
			// We don't check for IsNotFound here, because it
			// should have already been handled in the GetType
			// check.
			return nil, err
		}
		obj, err := factory(r.state, r.resources, r, id)
		if err != nil {
			return nil, err
		}
		if !reflect.TypeOf(obj).AssignableTo(goType) {
			return nil, errors.Errorf(
				"internal error, %s(%d) claimed to return %s but returned %T",
				rootName, version, goType, obj)
		}
		r.objectCache[objKey] = obj
		return obj, nil
	}
	return &srvCaller{
		creator:   creator,
		objMethod: objMethod,
	}, nil
}

// AuthMachineAgent returns whether the current client is a machine agent.
func (r *srvRoot) AuthMachineAgent() bool {
	_, ok := r.entity.(*state.Machine)
	return ok
}

// AuthUnitAgent returns whether the current client is a unit agent.
func (r *srvRoot) AuthUnitAgent() bool {
	_, ok := r.entity.(*state.Unit)
	return ok
}

// AuthOwner returns whether the authenticated user's tag matches the
// given entity tag.
func (r *srvRoot) AuthOwner(tag string) bool {
	return r.entity.Tag() == tag
}

// AuthEnvironManager returns whether the authenticated user is a
// machine with running the ManageEnviron job.
func (r *srvRoot) AuthEnvironManager() bool {
	return isMachineWithJob(r.entity, state.JobManageEnviron)
}

// AuthClient returns whether the authenticated entity is a client
// user.
func (r *srvRoot) AuthClient() bool {
	_, isUser := r.entity.(*state.User)
	return isUser
}

// GetAuthTag returns the tag of the authenticated entity.
func (r *srvRoot) GetAuthTag() string {
	return r.entity.Tag()
}

// GetAuthEntity returns the authenticated entity.
func (r *srvRoot) GetAuthEntity() state.Entity {
	return r.entity
}
