package integrations

import "encoding/json"

type TwilioConfigPayload struct {
	// The SID for your Twilio account
	AccountSid *string `json:"account_sid,omitempty" required:"true" maxLength:"100" minLength:"1"`
	// A US1 API key for Twilio-  - https://www.twilio.com/docs/iam/api-keys
	ApiKey *string `json:"api_key,omitempty" required:"true" maxLength:"100" minLength:"1"`
	// The API Secret for Twilio
	ApiSecret *string `json:"api_secret,omitempty" required:"true" maxLength:"100" minLength:"1"`
	// The phone number to send from, in E.164 format
	From *string `json:"from,omitempty" required:"true" maxLength:"100" minLength:"1" pattern:"^\+[0-9]{1,14}$"`
}

func (t *TwilioConfigPayload) GetAccountSid() *string {
	if t == nil {
		return nil
	}
	return t.AccountSid
}

func (t *TwilioConfigPayload) SetAccountSid(accountSid string) {
	t.AccountSid = &accountSid
}

func (t *TwilioConfigPayload) GetApiKey() *string {
	if t == nil {
		return nil
	}
	return t.ApiKey
}

func (t *TwilioConfigPayload) SetApiKey(apiKey string) {
	t.ApiKey = &apiKey
}

func (t *TwilioConfigPayload) GetApiSecret() *string {
	if t == nil {
		return nil
	}
	return t.ApiSecret
}

func (t *TwilioConfigPayload) SetApiSecret(apiSecret string) {
	t.ApiSecret = &apiSecret
}

func (t *TwilioConfigPayload) GetFrom() *string {
	if t == nil {
		return nil
	}
	return t.From
}

func (t *TwilioConfigPayload) SetFrom(from string) {
	t.From = &from
}

func (t TwilioConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TwilioConfigPayload to string"
	}
	return string(jsonData)
}
