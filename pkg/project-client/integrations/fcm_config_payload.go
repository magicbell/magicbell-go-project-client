package integrations

import "encoding/json"

type FcmConfigPayload struct {
	AuthProviderX509CertUrl *string `json:"auth_provider_x509_cert_url,omitempty" required:"true"`
	AuthUri                 *string `json:"auth_uri,omitempty" required:"true"`
	ClientEmail             *string `json:"client_email,omitempty" required:"true"`
	ClientId                *string `json:"client_id,omitempty" required:"true"`
	ClientX509CertUrl       *string `json:"client_x509_cert_url,omitempty" required:"true"`
	PrivateKey              *string `json:"private_key,omitempty" required:"true" pattern:"^-+?\s?BEGIN[A-Z ]+-+\n([A-Za-z0-9+/\r\n]+={0,2})\n-+\s?END[A-Z ]+-+\n?$"`
	PrivateKeyId            *string `json:"private_key_id,omitempty" required:"true"`
	ProjectId               *string `json:"project_id,omitempty" required:"true"`
	TokenUri                *string `json:"token_uri,omitempty" required:"true"`
	Type_                   *Type_  `json:"type,omitempty" required:"true"`
	UniverseDomain          *string `json:"universe_domain,omitempty" required:"true"`
}

func (f *FcmConfigPayload) GetAuthProviderX509CertUrl() *string {
	if f == nil {
		return nil
	}
	return f.AuthProviderX509CertUrl
}

func (f *FcmConfigPayload) SetAuthProviderX509CertUrl(authProviderX509CertUrl string) {
	f.AuthProviderX509CertUrl = &authProviderX509CertUrl
}

func (f *FcmConfigPayload) GetAuthUri() *string {
	if f == nil {
		return nil
	}
	return f.AuthUri
}

func (f *FcmConfigPayload) SetAuthUri(authUri string) {
	f.AuthUri = &authUri
}

func (f *FcmConfigPayload) GetClientEmail() *string {
	if f == nil {
		return nil
	}
	return f.ClientEmail
}

func (f *FcmConfigPayload) SetClientEmail(clientEmail string) {
	f.ClientEmail = &clientEmail
}

func (f *FcmConfigPayload) GetClientId() *string {
	if f == nil {
		return nil
	}
	return f.ClientId
}

func (f *FcmConfigPayload) SetClientId(clientId string) {
	f.ClientId = &clientId
}

func (f *FcmConfigPayload) GetClientX509CertUrl() *string {
	if f == nil {
		return nil
	}
	return f.ClientX509CertUrl
}

func (f *FcmConfigPayload) SetClientX509CertUrl(clientX509CertUrl string) {
	f.ClientX509CertUrl = &clientX509CertUrl
}

func (f *FcmConfigPayload) GetPrivateKey() *string {
	if f == nil {
		return nil
	}
	return f.PrivateKey
}

func (f *FcmConfigPayload) SetPrivateKey(privateKey string) {
	f.PrivateKey = &privateKey
}

func (f *FcmConfigPayload) GetPrivateKeyId() *string {
	if f == nil {
		return nil
	}
	return f.PrivateKeyId
}

func (f *FcmConfigPayload) SetPrivateKeyId(privateKeyId string) {
	f.PrivateKeyId = &privateKeyId
}

func (f *FcmConfigPayload) GetProjectId() *string {
	if f == nil {
		return nil
	}
	return f.ProjectId
}

func (f *FcmConfigPayload) SetProjectId(projectId string) {
	f.ProjectId = &projectId
}

func (f *FcmConfigPayload) GetTokenUri() *string {
	if f == nil {
		return nil
	}
	return f.TokenUri
}

func (f *FcmConfigPayload) SetTokenUri(tokenUri string) {
	f.TokenUri = &tokenUri
}

func (f *FcmConfigPayload) GetType_() *Type_ {
	if f == nil {
		return nil
	}
	return f.Type_
}

func (f *FcmConfigPayload) SetType_(type_ Type_) {
	f.Type_ = &type_
}

func (f *FcmConfigPayload) GetUniverseDomain() *string {
	if f == nil {
		return nil
	}
	return f.UniverseDomain
}

func (f *FcmConfigPayload) SetUniverseDomain(universeDomain string) {
	f.UniverseDomain = &universeDomain
}

func (f FcmConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FcmConfigPayload to string"
	}
	return string(jsonData)
}

type Type_ string

const (
	TYPE_SERVICE_ACCOUNT Type_ = "service_account"
)
