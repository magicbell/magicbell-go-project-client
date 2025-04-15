package users

import (
	"context"
	restClient "github.com/magicbell/magicbell-go-project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go-project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go-project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
	"time"
)

type UsersService struct {
	manager *configmanager.ConfigManager
}

func NewUsersService() *UsersService {
	return &UsersService{
		manager: configmanager.NewConfigManager(magicbellprojectclientconfig.Config{}),
	}
}

func (api *UsersService) WithConfigManager(manager *configmanager.ConfigManager) *UsersService {
	api.manager = manager
	return api
}

func (api *UsersService) getConfig() *magicbellprojectclientconfig.Config {
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

func (api *UsersService) ListUsers(ctx context.Context, params ListUsersRequestParams) (*shared.MagicbellProjectClientResponse[UserCollection], *shared.MagicbellProjectClientError) {
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
		return nil, shared.NewMagicbellProjectClientError[UserCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[UserCollection](resp), nil
}

// Creates a user with the provided details. The user will be associated with the project specified in the request context.
func (api *UsersService) CreateUser(ctx context.Context, user User) (*shared.MagicbellProjectClientResponse[User], *shared.MagicbellProjectClientError) {
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
		return nil, shared.NewMagicbellProjectClientError[User](err)
	}

	return shared.NewMagicbellProjectClientResponse[User](resp), nil
}

func (api *UsersService) DeleteUser(ctx context.Context, userId string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
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
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}
