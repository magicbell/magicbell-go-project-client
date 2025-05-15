package events

type ListEventsRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListEventsRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListEventsRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListEventsRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}
