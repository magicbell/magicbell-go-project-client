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

func (api *UsersService) ListUsers(ctx context.Context, params ListUsersRequestParams) (*shared.MagicbellProjectClientResponse[ArrayOfUsers], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfUsers](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfUsers](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfUsers](resp), nil
}
