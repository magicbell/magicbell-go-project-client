package integrations

import "encoding/json"

type SlackInstallation struct {
	AccessToken         *string          `json:"access_token,omitempty" required:"true"`
	AppId               *string          `json:"app_id,omitempty" required:"true"`
	AuthedUser          *AuthedUser      `json:"authed_user,omitempty" required:"true"`
	BotUserId           *string          `json:"bot_user_id,omitempty"`
	Enterprise          *Enterprise      `json:"enterprise,omitempty"`
	ExpiresIn           *int64           `json:"expires_in,omitempty"`
	Id                  *string          `json:"id,omitempty" pattern:"^[A-Z0-9]+-.*$"`
	IncomingWebhook     *IncomingWebhook `json:"incoming_webhook,omitempty"`
	IsEnterpriseInstall *bool            `json:"is_enterprise_install,omitempty"`
	RefreshToken        *string          `json:"refresh_token,omitempty"`
	Scope               *string          `json:"scope,omitempty"`
	Team                *Team            `json:"team,omitempty" required:"true"`
	TokenType           *string          `json:"token_type,omitempty"`
}

func (s *SlackInstallation) GetAccessToken() *string {
	if s == nil {
		return nil
	}
	return s.AccessToken
}

func (s *SlackInstallation) SetAccessToken(accessToken string) {
	s.AccessToken = &accessToken
}

func (s *SlackInstallation) GetAppId() *string {
	if s == nil {
		return nil
	}
	return s.AppId
}

func (s *SlackInstallation) SetAppId(appId string) {
	s.AppId = &appId
}

func (s *SlackInstallation) GetAuthedUser() *AuthedUser {
	if s == nil {
		return nil
	}
	return s.AuthedUser
}

func (s *SlackInstallation) SetAuthedUser(authedUser AuthedUser) {
	s.AuthedUser = &authedUser
}

func (s *SlackInstallation) GetBotUserId() *string {
	if s == nil {
		return nil
	}
	return s.BotUserId
}

func (s *SlackInstallation) SetBotUserId(botUserId string) {
	s.BotUserId = &botUserId
}

func (s *SlackInstallation) GetEnterprise() *Enterprise {
	if s == nil {
		return nil
	}
	return s.Enterprise
}

func (s *SlackInstallation) SetEnterprise(enterprise Enterprise) {
	s.Enterprise = &enterprise
}

func (s *SlackInstallation) GetExpiresIn() *int64 {
	if s == nil {
		return nil
	}
	return s.ExpiresIn
}

func (s *SlackInstallation) SetExpiresIn(expiresIn int64) {
	s.ExpiresIn = &expiresIn
}

func (s *SlackInstallation) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *SlackInstallation) SetId(id string) {
	s.Id = &id
}

func (s *SlackInstallation) GetIncomingWebhook() *IncomingWebhook {
	if s == nil {
		return nil
	}
	return s.IncomingWebhook
}

func (s *SlackInstallation) SetIncomingWebhook(incomingWebhook IncomingWebhook) {
	s.IncomingWebhook = &incomingWebhook
}

func (s *SlackInstallation) GetIsEnterpriseInstall() *bool {
	if s == nil {
		return nil
	}
	return s.IsEnterpriseInstall
}

func (s *SlackInstallation) SetIsEnterpriseInstall(isEnterpriseInstall bool) {
	s.IsEnterpriseInstall = &isEnterpriseInstall
}

func (s *SlackInstallation) GetRefreshToken() *string {
	if s == nil {
		return nil
	}
	return s.RefreshToken
}

func (s *SlackInstallation) SetRefreshToken(refreshToken string) {
	s.RefreshToken = &refreshToken
}

func (s *SlackInstallation) GetScope() *string {
	if s == nil {
		return nil
	}
	return s.Scope
}

func (s *SlackInstallation) SetScope(scope string) {
	s.Scope = &scope
}

func (s *SlackInstallation) GetTeam() *Team {
	if s == nil {
		return nil
	}
	return s.Team
}

func (s *SlackInstallation) SetTeam(team Team) {
	s.Team = &team
}

func (s *SlackInstallation) GetTokenType() *string {
	if s == nil {
		return nil
	}
	return s.TokenType
}

func (s *SlackInstallation) SetTokenType(tokenType string) {
	s.TokenType = &tokenType
}

func (s SlackInstallation) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackInstallation to string"
	}
	return string(jsonData)
}

type AuthedUser struct {
	AccessToken  *string `json:"access_token,omitempty"`
	ExpiresIn    *int64  `json:"expires_in,omitempty"`
	Id           *string `json:"id,omitempty" required:"true"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	Scope        *string `json:"scope,omitempty"`
	TokenType    *string `json:"token_type,omitempty"`
}

func (a *AuthedUser) GetAccessToken() *string {
	if a == nil {
		return nil
	}
	return a.AccessToken
}

func (a *AuthedUser) SetAccessToken(accessToken string) {
	a.AccessToken = &accessToken
}

func (a *AuthedUser) GetExpiresIn() *int64 {
	if a == nil {
		return nil
	}
	return a.ExpiresIn
}

func (a *AuthedUser) SetExpiresIn(expiresIn int64) {
	a.ExpiresIn = &expiresIn
}

func (a *AuthedUser) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *AuthedUser) SetId(id string) {
	a.Id = &id
}

func (a *AuthedUser) GetRefreshToken() *string {
	if a == nil {
		return nil
	}
	return a.RefreshToken
}

func (a *AuthedUser) SetRefreshToken(refreshToken string) {
	a.RefreshToken = &refreshToken
}

func (a *AuthedUser) GetScope() *string {
	if a == nil {
		return nil
	}
	return a.Scope
}

func (a *AuthedUser) SetScope(scope string) {
	a.Scope = &scope
}

func (a *AuthedUser) GetTokenType() *string {
	if a == nil {
		return nil
	}
	return a.TokenType
}

func (a *AuthedUser) SetTokenType(tokenType string) {
	a.TokenType = &tokenType
}

func (a AuthedUser) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AuthedUser to string"
	}
	return string(jsonData)
}

type Enterprise struct {
	Id   *string `json:"id,omitempty" required:"true"`
	Name *string `json:"name,omitempty" required:"true"`
}

func (e *Enterprise) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *Enterprise) SetId(id string) {
	e.Id = &id
}

func (e *Enterprise) GetName() *string {
	if e == nil {
		return nil
	}
	return e.Name
}

func (e *Enterprise) SetName(name string) {
	e.Name = &name
}

func (e Enterprise) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: Enterprise to string"
	}
	return string(jsonData)
}

type IncomingWebhook struct {
	Channel          *string `json:"channel,omitempty" required:"true"`
	ConfigurationUrl *string `json:"configuration_url,omitempty" required:"true"`
	Url              *string `json:"url,omitempty" required:"true"`
}

func (i *IncomingWebhook) GetChannel() *string {
	if i == nil {
		return nil
	}
	return i.Channel
}

func (i *IncomingWebhook) SetChannel(channel string) {
	i.Channel = &channel
}

func (i *IncomingWebhook) GetConfigurationUrl() *string {
	if i == nil {
		return nil
	}
	return i.ConfigurationUrl
}

func (i *IncomingWebhook) SetConfigurationUrl(configurationUrl string) {
	i.ConfigurationUrl = &configurationUrl
}

func (i *IncomingWebhook) GetUrl() *string {
	if i == nil {
		return nil
	}
	return i.Url
}

func (i *IncomingWebhook) SetUrl(url string) {
	i.Url = &url
}

func (i IncomingWebhook) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: IncomingWebhook to string"
	}
	return string(jsonData)
}

type Team struct {
	Id   *string `json:"id,omitempty" required:"true"`
	Name *string `json:"name,omitempty"`
}

func (t *Team) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *Team) SetId(id string) {
	t.Id = &id
}

func (t *Team) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *Team) SetName(name string) {
	t.Name = &name
}

func (t Team) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: Team to string"
	}
	return string(jsonData)
}
