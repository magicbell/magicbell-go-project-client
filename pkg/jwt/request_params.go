package jwt

type FetchProjectTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *FetchProjectTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *FetchProjectTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *FetchProjectTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type FetchUserTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *FetchUserTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *FetchUserTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *FetchUserTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}
