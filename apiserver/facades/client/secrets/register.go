// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package secrets

import (
	"context"
	"reflect"

	"github.com/juju/errors"

	apiservererrors "github.com/juju/juju/apiserver/errors"
	"github.com/juju/juju/apiserver/facade"
	secretservice "github.com/juju/juju/domain/secret/service"
	secretbackendservice "github.com/juju/juju/domain/secretbackend/service"
)

// Register is called to expose a package of facades onto a given registry.
func Register(registry facade.FacadeRegistry) {
	registry.MustRegister("Secrets", 1, func(stdCtx context.Context, ctx facade.ModelContext) (facade.Facade, error) {
		return newSecretsAPIV1(stdCtx, ctx)
	}, reflect.TypeOf((*SecretsAPI)(nil)))
	registry.MustRegister("Secrets", 2, func(stdCtx context.Context, ctx facade.ModelContext) (facade.Facade, error) {
		return newSecretsAPI(stdCtx, ctx)
	}, reflect.TypeOf((*SecretsAPI)(nil)))
}

func newSecretsAPIV1(stdCtx context.Context, context facade.ModelContext) (*SecretsAPIV1, error) {
	api, err := newSecretsAPI(stdCtx, context)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &SecretsAPIV1{SecretsAPI: api}, nil
}

// newSecretsAPI creates a SecretsAPI.
func newSecretsAPI(stdCtx context.Context, ctx facade.ModelContext) (*SecretsAPI, error) {
	if !ctx.Auth().AuthClient() {
		return nil, apiservererrors.ErrPerm
	}

	domainServices := ctx.DomainServices()
	backendService := domainServices.SecretBackend()

	modelInfoService := domainServices.ModelInfo()
	modelInfo, err := modelInfoService.GetModelInfo(stdCtx)
	if err != nil {
		return nil, errors.Trace(err)
	}

	secretService := domainServices.Secret(
		secretservice.SecretServiceParams{
			BackendAdminConfigGetter: secretbackendservice.AdminBackendConfigGetterFunc(
				backendService, ctx.ModelUUID(),
			),
			BackendUserSecretConfigGetter: secretbackendservice.UserSecretBackendConfigGetterFunc(
				backendService, ctx.ModelUUID(),
			),
		},
	)

	return &SecretsAPI{
		authorizer:           ctx.Auth(),
		authTag:              ctx.Auth().GetAuthTag(),
		controllerUUID:       ctx.ControllerUUID(),
		modelUUID:            ctx.State().ModelUUID(),
		modelName:            modelInfo.Name,
		secretService:        secretService,
		secretBackendService: backendService,
	}, nil
}
