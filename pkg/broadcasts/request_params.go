package broadcasts

type ListBroadcastsRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListBroadcastsRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListBroadcastsRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListBroadcastsRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}
