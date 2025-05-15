package channels

type GetInAppInboxTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetInAppInboxTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetInAppInboxTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetInAppInboxTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type GetMobilePushApnsTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetMobilePushApnsTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetMobilePushApnsTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetMobilePushApnsTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type GetMobilePushExpoTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetMobilePushExpoTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetMobilePushExpoTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetMobilePushExpoTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type GetMobilePushFcmTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetMobilePushFcmTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetMobilePushFcmTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetMobilePushFcmTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type GetSlackTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetSlackTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetSlackTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetSlackTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type GetTeamsTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetTeamsTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetTeamsTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetTeamsTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type GetWebPushTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetWebPushTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetWebPushTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetWebPushTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}
