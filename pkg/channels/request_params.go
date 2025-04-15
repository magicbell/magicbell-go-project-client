package channels

type GetDeliveryconfigRequestParams struct {
	Key *string `explode:"true" serializationStyle:"form" queryParam:"key"`
}

func (params *GetDeliveryconfigRequestParams) SetKey(key string) {
	params.Key = &key
}

type GetInAppInboxUserTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetInAppInboxUserTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetInAppInboxUserTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetInAppInboxUserTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type GetMobilePushApnsUserTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetMobilePushApnsUserTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetMobilePushApnsUserTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetMobilePushApnsUserTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type GetMobilePushExpoUserTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetMobilePushExpoUserTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetMobilePushExpoUserTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetMobilePushExpoUserTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type GetMobilePushFcmUserTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetMobilePushFcmUserTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetMobilePushFcmUserTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetMobilePushFcmUserTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type GetSlackUserTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetSlackUserTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetSlackUserTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetSlackUserTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type GetTeamsUserTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetTeamsUserTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetTeamsUserTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetTeamsUserTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type GetWebPushUserTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *GetWebPushUserTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *GetWebPushUserTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *GetWebPushUserTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}
