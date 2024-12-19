package jwt

import (
	"context"
	restClient "github.com/magicbell/magicbell-go-project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go-project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go-project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
	"time"
)

type JwtService struct {
	manager *configmanager.ConfigManager
}

func NewJwtService(manager *configmanager.ConfigManager) *JwtService {
	return &JwtService{
		manager: manager,
	}
}

func (api *JwtService) getConfig() *magicbellprojectclientconfig.Config {
	return api.manager.GetJwt()
}

func (api *JwtService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *JwtService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *JwtService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Retrieves a list of all active project-level JWT tokens. Returns a paginated list showing token metadata including creation date, last used date, and expiration time. For security reasons, the actual token values are not included in the response.
func (api *JwtService) FetchProjectTokens(ctx context.Context, params FetchProjectTokensRequestParams) (*shared.MagicbellProjectClientResponse[ArrayOfFetchTokensResponseTokens], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/jwt/project").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfFetchTokensResponseTokens](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfFetchTokensResponseTokens](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfFetchTokensResponseTokens](resp), nil
}

// Creates a new project-level JWT token. These tokens provide project-wide access and should be carefully managed. Only administrators can create project tokens. The returned token should be securely stored as it cannot be retrieved again after creation.
func (api *JwtService) CreateProjectJwt(ctx context.Context, createProjectTokenRequest CreateProjectTokenRequest) (*shared.MagicbellProjectClientResponse[AccessToken], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/jwt/project").
		WithConfig(config).
		WithBody(createProjectTokenRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[AccessToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[AccessToken](err)
	}

	return shared.NewMagicbellProjectClientResponse[AccessToken](resp), nil
}

// Immediately revokes a project-level JWT token. Once revoked, any requests using this token will be rejected. This action is immediate and cannot be undone. Active sessions using this token will be terminated.
func (api *JwtService) DiscardProjectJwt(ctx context.Context, tokenId string) (*shared.MagicbellProjectClientResponse[DiscardTokenResponse], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/jwt/project/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardTokenResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[DiscardTokenResponse](err)
	}

	return shared.NewMagicbellProjectClientResponse[DiscardTokenResponse](resp), nil
}

// Issues a new user-specific JWT token. These tokens are scoped to individual user permissions and access levels. Only administrators can create user tokens. The token is returned only once at creation time and cannot be retrieved later.
func (api *JwtService) CreateUserJwt(ctx context.Context, createUserTokenRequest CreateUserTokenRequest) (*shared.MagicbellProjectClientResponse[AccessToken], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/jwt/user").
		WithConfig(config).
		WithBody(createUserTokenRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[AccessToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[AccessToken](err)
	}

	return shared.NewMagicbellProjectClientResponse[AccessToken](resp), nil
}

// Revokes a specific user's JWT token. This immediately invalidates the token and terminates any active sessions using it. This action cannot be undone. Administrators should use this to revoke access when needed for security purposes.
func (api *JwtService) DiscardUserJwt(ctx context.Context, tokenId string) (*shared.MagicbellProjectClientResponse[DiscardTokenResponse], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/jwt/user/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardTokenResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[DiscardTokenResponse](err)
	}

	return shared.NewMagicbellProjectClientResponse[DiscardTokenResponse](resp), nil
}

// Lists all JWT tokens associated with a specific user. Returns token metadata including creation time, last access time, and expiration date. Administrators can use this to audit user token usage and manage active sessions. Token values are not included in the response for security reasons.
func (api *JwtService) FetchUserTokens(ctx context.Context, userId string, params FetchUserTokensRequestParams) (*shared.MagicbellProjectClientResponse[ArrayOfFetchTokensResponseTokens], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/jwt/user/{user_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfFetchTokensResponseTokens](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfFetchTokensResponseTokens](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfFetchTokensResponseTokens](resp), nil
}