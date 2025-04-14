package integrations

type ListIntegrationsRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListIntegrationsRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListIntegrationsRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListIntegrationsRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}
