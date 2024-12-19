package integrations

import (
	"encoding/json"
)

type FcmConfig struct {
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
	touched                 map[string]bool
}

func (f *FcmConfig) GetAuthProviderX509CertUrl() *string {
	if f == nil {
		return nil
	}
	return f.AuthProviderX509CertUrl
}

func (f *FcmConfig) SetAuthProviderX509CertUrl(authProviderX509CertUrl string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["AuthProviderX509CertUrl"] = true
	f.AuthProviderX509CertUrl = &authProviderX509CertUrl
}

func (f *FcmConfig) SetAuthProviderX509CertUrlNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["AuthProviderX509CertUrl"] = true
	f.AuthProviderX509CertUrl = nil
}

func (f *FcmConfig) GetAuthUri() *string {
	if f == nil {
		return nil
	}
	return f.AuthUri
}

func (f *FcmConfig) SetAuthUri(authUri string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["AuthUri"] = true
	f.AuthUri = &authUri
}

func (f *FcmConfig) SetAuthUriNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["AuthUri"] = true
	f.AuthUri = nil
}

func (f *FcmConfig) GetClientEmail() *string {
	if f == nil {
		return nil
	}
	return f.ClientEmail
}

func (f *FcmConfig) SetClientEmail(clientEmail string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["ClientEmail"] = true
	f.ClientEmail = &clientEmail
}

func (f *FcmConfig) SetClientEmailNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["ClientEmail"] = true
	f.ClientEmail = nil
}

func (f *FcmConfig) GetClientId() *string {
	if f == nil {
		return nil
	}
	return f.ClientId
}

func (f *FcmConfig) SetClientId(clientId string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["ClientId"] = true
	f.ClientId = &clientId
}

func (f *FcmConfig) SetClientIdNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["ClientId"] = true
	f.ClientId = nil
}

func (f *FcmConfig) GetClientX509CertUrl() *string {
	if f == nil {
		return nil
	}
	return f.ClientX509CertUrl
}

func (f *FcmConfig) SetClientX509CertUrl(clientX509CertUrl string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["ClientX509CertUrl"] = true
	f.ClientX509CertUrl = &clientX509CertUrl
}

func (f *FcmConfig) SetClientX509CertUrlNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["ClientX509CertUrl"] = true
	f.ClientX509CertUrl = nil
}

func (f *FcmConfig) GetPrivateKey() *string {
	if f == nil {
		return nil
	}
	return f.PrivateKey
}

func (f *FcmConfig) SetPrivateKey(privateKey string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["PrivateKey"] = true
	f.PrivateKey = &privateKey
}

func (f *FcmConfig) SetPrivateKeyNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["PrivateKey"] = true
	f.PrivateKey = nil
}

func (f *FcmConfig) GetPrivateKeyId() *string {
	if f == nil {
		return nil
	}
	return f.PrivateKeyId
}

func (f *FcmConfig) SetPrivateKeyId(privateKeyId string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["PrivateKeyId"] = true
	f.PrivateKeyId = &privateKeyId
}

func (f *FcmConfig) SetPrivateKeyIdNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["PrivateKeyId"] = true
	f.PrivateKeyId = nil
}

func (f *FcmConfig) GetProjectId() *string {
	if f == nil {
		return nil
	}
	return f.ProjectId
}

func (f *FcmConfig) SetProjectId(projectId string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["ProjectId"] = true
	f.ProjectId = &projectId
}

func (f *FcmConfig) SetProjectIdNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["ProjectId"] = true
	f.ProjectId = nil
}

func (f *FcmConfig) GetTokenUri() *string {
	if f == nil {
		return nil
	}
	return f.TokenUri
}

func (f *FcmConfig) SetTokenUri(tokenUri string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["TokenUri"] = true
	f.TokenUri = &tokenUri
}

func (f *FcmConfig) SetTokenUriNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["TokenUri"] = true
	f.TokenUri = nil
}

func (f *FcmConfig) GetType_() *Type_ {
	if f == nil {
		return nil
	}
	return f.Type_
}

func (f *FcmConfig) SetType_(type_ Type_) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Type_"] = true
	f.Type_ = &type_
}

func (f *FcmConfig) SetType_Nil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Type_"] = true
	f.Type_ = nil
}

func (f *FcmConfig) GetUniverseDomain() *string {
	if f == nil {
		return nil
	}
	return f.UniverseDomain
}

func (f *FcmConfig) SetUniverseDomain(universeDomain string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["UniverseDomain"] = true
	f.UniverseDomain = &universeDomain
}

func (f *FcmConfig) SetUniverseDomainNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["UniverseDomain"] = true
	f.UniverseDomain = nil
}
func (f FcmConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if f.touched["AuthProviderX509CertUrl"] && f.AuthProviderX509CertUrl == nil {
		data["auth_provider_x509_cert_url"] = nil
	} else if f.AuthProviderX509CertUrl != nil {
		data["auth_provider_x509_cert_url"] = f.AuthProviderX509CertUrl
	}

	if f.touched["AuthUri"] && f.AuthUri == nil {
		data["auth_uri"] = nil
	} else if f.AuthUri != nil {
		data["auth_uri"] = f.AuthUri
	}

	if f.touched["ClientEmail"] && f.ClientEmail == nil {
		data["client_email"] = nil
	} else if f.ClientEmail != nil {
		data["client_email"] = f.ClientEmail
	}

	if f.touched["ClientId"] && f.ClientId == nil {
		data["client_id"] = nil
	} else if f.ClientId != nil {
		data["client_id"] = f.ClientId
	}

	if f.touched["ClientX509CertUrl"] && f.ClientX509CertUrl == nil {
		data["client_x509_cert_url"] = nil
	} else if f.ClientX509CertUrl != nil {
		data["client_x509_cert_url"] = f.ClientX509CertUrl
	}

	if f.touched["PrivateKey"] && f.PrivateKey == nil {
		data["private_key"] = nil
	} else if f.PrivateKey != nil {
		data["private_key"] = f.PrivateKey
	}

	if f.touched["PrivateKeyId"] && f.PrivateKeyId == nil {
		data["private_key_id"] = nil
	} else if f.PrivateKeyId != nil {
		data["private_key_id"] = f.PrivateKeyId
	}

	if f.touched["ProjectId"] && f.ProjectId == nil {
		data["project_id"] = nil
	} else if f.ProjectId != nil {
		data["project_id"] = f.ProjectId
	}

	if f.touched["TokenUri"] && f.TokenUri == nil {
		data["token_uri"] = nil
	} else if f.TokenUri != nil {
		data["token_uri"] = f.TokenUri
	}

	if f.touched["Type_"] && f.Type_ == nil {
		data["type"] = nil
	} else if f.Type_ != nil {
		data["type"] = f.Type_
	}

	if f.touched["UniverseDomain"] && f.UniverseDomain == nil {
		data["universe_domain"] = nil
	} else if f.UniverseDomain != nil {
		data["universe_domain"] = f.UniverseDomain
	}

	return json.Marshal(data)
}

type Type_ string

const (
	TYPE_SERVICE_ACCOUNT Type_ = "service_account"
)
