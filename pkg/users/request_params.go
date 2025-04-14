package users

type ListUsersRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListUsersRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListUsersRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListUsersRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}
