package integrations

import (
	"encoding/json"
)

type TwilioConfig struct {
	// The SID for your Twilio account
	AccountSid *string `json:"account_sid,omitempty" required:"true" maxLength:"100" minLength:"1"`
	// The API key for Twilio
	ApiKey *string `json:"api_key,omitempty" required:"true" maxLength:"100" minLength:"1"`
	// The API Secret for Twilio
	ApiSecret *string `json:"api_secret,omitempty" required:"true" maxLength:"100" minLength:"1"`
	// The phone number to send from, in E.164 format
	From *string `json:"from,omitempty" required:"true" maxLength:"100" minLength:"1" pattern:"^\+[0-9]{1,14}$"`
	// The region to use for Twilio, defaults to 'us1'
	Region  *TwilioConfigRegion `json:"region,omitempty"`
	touched map[string]bool
}

func (t *TwilioConfig) GetAccountSid() *string {
	if t == nil {
		return nil
	}
	return t.AccountSid
}

func (t *TwilioConfig) SetAccountSid(accountSid string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["AccountSid"] = true
	t.AccountSid = &accountSid
}

func (t *TwilioConfig) SetAccountSidNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["AccountSid"] = true
	t.AccountSid = nil
}

func (t *TwilioConfig) GetApiKey() *string {
	if t == nil {
		return nil
	}
	return t.ApiKey
}

func (t *TwilioConfig) SetApiKey(apiKey string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["ApiKey"] = true
	t.ApiKey = &apiKey
}

func (t *TwilioConfig) SetApiKeyNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["ApiKey"] = true
	t.ApiKey = nil
}

func (t *TwilioConfig) GetApiSecret() *string {
	if t == nil {
		return nil
	}
	return t.ApiSecret
}

func (t *TwilioConfig) SetApiSecret(apiSecret string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["ApiSecret"] = true
	t.ApiSecret = &apiSecret
}

func (t *TwilioConfig) SetApiSecretNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["ApiSecret"] = true
	t.ApiSecret = nil
}

func (t *TwilioConfig) GetFrom() *string {
	if t == nil {
		return nil
	}
	return t.From
}

func (t *TwilioConfig) SetFrom(from string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["From"] = true
	t.From = &from
}

func (t *TwilioConfig) SetFromNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["From"] = true
	t.From = nil
}

func (t *TwilioConfig) GetRegion() *TwilioConfigRegion {
	if t == nil {
		return nil
	}
	return t.Region
}

func (t *TwilioConfig) SetRegion(region TwilioConfigRegion) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Region"] = true
	t.Region = &region
}

func (t *TwilioConfig) SetRegionNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Region"] = true
	t.Region = nil
}
func (t TwilioConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["AccountSid"] && t.AccountSid == nil {
		data["account_sid"] = nil
	} else if t.AccountSid != nil {
		data["account_sid"] = t.AccountSid
	}

	if t.touched["ApiKey"] && t.ApiKey == nil {
		data["api_key"] = nil
	} else if t.ApiKey != nil {
		data["api_key"] = t.ApiKey
	}

	if t.touched["ApiSecret"] && t.ApiSecret == nil {
		data["api_secret"] = nil
	} else if t.ApiSecret != nil {
		data["api_secret"] = t.ApiSecret
	}

	if t.touched["From"] && t.From == nil {
		data["from"] = nil
	} else if t.From != nil {
		data["from"] = t.From
	}

	if t.touched["Region"] && t.Region == nil {
		data["region"] = nil
	} else if t.Region != nil {
		data["region"] = t.Region
	}

	return json.Marshal(data)
}

// The region to use for Twilio, defaults to 'us1'
type TwilioConfigRegion string

const (
	TWILIO_CONFIG_REGION_US1 TwilioConfigRegion = "us1"
	TWILIO_CONFIG_REGION_IE1 TwilioConfigRegion = "ie1"
	TWILIO_CONFIG_REGION_AU1 TwilioConfigRegion = "au1"
)
