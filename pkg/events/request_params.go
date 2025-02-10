package events

type GetEventsRequestParams struct {
	PageSize   *int64  `explode:"true" serializationStyle:"form" queryParam:"page[size]"`
	PageAfter  *string `explode:"true" serializationStyle:"form" queryParam:"page[after]"`
	PageBefore *string `explode:"true" serializationStyle:"form" queryParam:"page[before]"`
}

func (params *GetEventsRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
func (params *GetEventsRequestParams) SetPageAfter(pageAfter string) {
	params.PageAfter = &pageAfter
}
func (params *GetEventsRequestParams) SetPageBefore(pageBefore string) {
	params.PageBefore = &pageBefore
}
