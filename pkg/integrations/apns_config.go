package integrations

import (
	"encoding/json"
)

type ApnsConfig struct {
	// The default bundle identifier of the application that is configured with this project. It can be overriden on a per token basis, when registering device tokens.
	AppId *string `json:"app_id,omitempty" required:"true" pattern:"^[a-zA-Z0-9]+(.[a-zA-Z0-9]+)*$"`
	Badge *Badge  `json:"badge,omitempty" required:"true"`
	// The APNs certificate in P8 format. Generate it at [developer.apple.com](https://developer.apple.com/account/resources/authkeys/add) with the 'Apple Push Notification service (APNs)' option selected.
	Certificate    *string         `json:"certificate,omitempty" required:"true" pattern:"^-+?\s?BEGIN PRIVATE KEY-+\n([A-Za-z0-9+/\r\n]+={0,2})\n-+\s?END PRIVATE KEY+-+\n?$"`
	KeyId          *string         `json:"key_id,omitempty" required:"true" maxLength:"10" minLength:"10"`
	PayloadVersion *PayloadVersion `json:"payload_version,omitempty"`
	TeamId         *string         `json:"team_id,omitempty" required:"true" maxLength:"10" minLength:"10"`
	touched        map[string]bool
}

func (a *ApnsConfig) GetAppId() *string {
	if a == nil {
		return nil
	}
	return a.AppId
}

func (a *ApnsConfig) SetAppId(appId string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["AppId"] = true
	a.AppId = &appId
}

func (a *ApnsConfig) SetAppIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["AppId"] = true
	a.AppId = nil
}

func (a *ApnsConfig) GetBadge() *Badge {
	if a == nil {
		return nil
	}
	return a.Badge
}

func (a *ApnsConfig) SetBadge(badge Badge) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Badge"] = true
	a.Badge = &badge
}

func (a *ApnsConfig) SetBadgeNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Badge"] = true
	a.Badge = nil
}

func (a *ApnsConfig) GetCertificate() *string {
	if a == nil {
		return nil
	}
	return a.Certificate
}

func (a *ApnsConfig) SetCertificate(certificate string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Certificate"] = true
	a.Certificate = &certificate
}

func (a *ApnsConfig) SetCertificateNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Certificate"] = true
	a.Certificate = nil
}

func (a *ApnsConfig) GetKeyId() *string {
	if a == nil {
		return nil
	}
	return a.KeyId
}

func (a *ApnsConfig) SetKeyId(keyId string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["KeyId"] = true
	a.KeyId = &keyId
}

func (a *ApnsConfig) SetKeyIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["KeyId"] = true
	a.KeyId = nil
}

func (a *ApnsConfig) GetPayloadVersion() *PayloadVersion {
	if a == nil {
		return nil
	}
	return a.PayloadVersion
}

func (a *ApnsConfig) SetPayloadVersion(payloadVersion PayloadVersion) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["PayloadVersion"] = true
	a.PayloadVersion = &payloadVersion
}

func (a *ApnsConfig) SetPayloadVersionNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["PayloadVersion"] = true
	a.PayloadVersion = nil
}

func (a *ApnsConfig) GetTeamId() *string {
	if a == nil {
		return nil
	}
	return a.TeamId
}

func (a *ApnsConfig) SetTeamId(teamId string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["TeamId"] = true
	a.TeamId = &teamId
}

func (a *ApnsConfig) SetTeamIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["TeamId"] = true
	a.TeamId = nil
}
func (a ApnsConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["AppId"] && a.AppId == nil {
		data["app_id"] = nil
	} else if a.AppId != nil {
		data["app_id"] = a.AppId
	}

	if a.touched["Badge"] && a.Badge == nil {
		data["badge"] = nil
	} else if a.Badge != nil {
		data["badge"] = a.Badge
	}

	if a.touched["Certificate"] && a.Certificate == nil {
		data["certificate"] = nil
	} else if a.Certificate != nil {
		data["certificate"] = a.Certificate
	}

	if a.touched["KeyId"] && a.KeyId == nil {
		data["key_id"] = nil
	} else if a.KeyId != nil {
		data["key_id"] = a.KeyId
	}

	if a.touched["PayloadVersion"] && a.PayloadVersion == nil {
		data["payload_version"] = nil
	} else if a.PayloadVersion != nil {
		data["payload_version"] = a.PayloadVersion
	}

	if a.touched["TeamId"] && a.TeamId == nil {
		data["team_id"] = nil
	} else if a.TeamId != nil {
		data["team_id"] = a.TeamId
	}

	return json.Marshal(data)
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
