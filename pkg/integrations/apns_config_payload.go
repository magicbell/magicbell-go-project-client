package integrations

import "encoding/json"

type ApnsConfigPayload struct {
	// The default bundle identifier of the application that is configured with this project. It can be overriden on a per token basis, when registering device tokens.
	AppId *string `json:"app_id,omitempty" required:"true" pattern:"^[a-zA-Z0-9]+(.[a-zA-Z0-9]+)*$"`
	Badge *Badge  `json:"badge,omitempty" required:"true"`
	// The APNs certificate in P8 format. Generate it at [developer.apple.com](https://developer.apple.com/account/resources/authkeys/add) with the 'Apple Push Notification service (APNs)' option selected.
	Certificate    *string         `json:"certificate,omitempty" required:"true" pattern:"^-+?\s?BEGIN PRIVATE KEY-+\n([A-Za-z0-9+/\r\n]+={0,2})\n-+\s?END PRIVATE KEY+-+\n?$"`
	KeyId          *string         `json:"key_id,omitempty" required:"true" maxLength:"10" minLength:"10"`
	PayloadVersion *PayloadVersion `json:"payload_version,omitempty"`
	TeamId         *string         `json:"team_id,omitempty" required:"true" maxLength:"10" minLength:"10"`
}

func (a *ApnsConfigPayload) GetAppId() *string {
	if a == nil {
		return nil
	}
	return a.AppId
}

func (a *ApnsConfigPayload) SetAppId(appId string) {
	a.AppId = &appId
}

func (a *ApnsConfigPayload) GetBadge() *Badge {
	if a == nil {
		return nil
	}
	return a.Badge
}

func (a *ApnsConfigPayload) SetBadge(badge Badge) {
	a.Badge = &badge
}

func (a *ApnsConfigPayload) GetCertificate() *string {
	if a == nil {
		return nil
	}
	return a.Certificate
}

func (a *ApnsConfigPayload) SetCertificate(certificate string) {
	a.Certificate = &certificate
}

func (a *ApnsConfigPayload) GetKeyId() *string {
	if a == nil {
		return nil
	}
	return a.KeyId
}

func (a *ApnsConfigPayload) SetKeyId(keyId string) {
	a.KeyId = &keyId
}

func (a *ApnsConfigPayload) GetPayloadVersion() *PayloadVersion {
	if a == nil {
		return nil
	}
	return a.PayloadVersion
}

func (a *ApnsConfigPayload) SetPayloadVersion(payloadVersion PayloadVersion) {
	a.PayloadVersion = &payloadVersion
}

func (a *ApnsConfigPayload) GetTeamId() *string {
	if a == nil {
		return nil
	}
	return a.TeamId
}

func (a *ApnsConfigPayload) SetTeamId(teamId string) {
	a.TeamId = &teamId
}

func (a ApnsConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApnsConfigPayload to string"
	}
	return string(jsonData)
}

type Badge string

const (
	BADGE_UNREAD Badge = "unread"
	BADGE_UNSEEN Badge = "unseen"
)

type PayloadVersion string

const (
	PAYLOAD_VERSION_1 PayloadVersion = "1"
	PAYLOAD_VERSION_2 PayloadVersion = "2"
)
