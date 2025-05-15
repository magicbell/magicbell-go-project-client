package users

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/project-client/projectclientconfig"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
	"time"
)

type UsersService struct {
	manager *configmanager.ConfigManager
}

func NewUsersService() *UsersService {
	return &UsersService{
		manager: configmanager.NewConfigManager(projectclientconfig.Config{}),
	}
}

func (api *UsersService) WithConfigManager(manager *configmanager.ConfigManager) *UsersService {
	api.manager = manager
	return api
}

func (api *UsersService) getConfig() *projectclientconfig.Config {
	return api.manager.GetUsers()
}

func (api *UsersService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *UsersService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *UsersService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

func (api *UsersService) ListUsers(ctx context.Context, params ListUsersRequestParams) (*shared.ProjectClientResponse[UserCollection], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[UserCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[UserCollection](err)
	}

	return shared.NewProjectClientResponse[UserCollection](resp), nil
}

// Creates a user with the provided details. The user will be associated with the project specified in the request context.
func (api *UsersService) CreateUser(ctx context.Context, user User) (*shared.ProjectClientResponse[User], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/users").
		WithConfig(config).
		WithBody(user).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[User](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[User](err)
	}

	return shared.NewProjectClientResponse[User](resp), nil
}

func (api *UsersService) DeleteUser(ctx context.Context, userId string) (*shared.ProjectClientResponse[any], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[any](err)
	}

	return shared.NewProjectClientResponse[any](resp), nil
}
