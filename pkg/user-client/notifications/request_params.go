package notifications

type ListNotificationsRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
	Status        *string `explode:"true" serializationStyle:"form" queryParam:"status"`
	Topic         *string `explode:"true" serializationStyle:"form" queryParam:"topic"`
}

func (params *ListNotificationsRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListNotificationsRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListNotificationsRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}
func (params *ListNotificationsRequestParams) SetStatus(status string) {
	params.Status = &status
}
func (params *ListNotificationsRequestParams) SetTopic(topic string) {
	params.Topic = &topic
}
