package broadcasts

type ListBroadcastsRequestParams struct {
	PageSize   *int64  `explode:"true" serializationStyle:"form" queryParam:"page[size]"`
	PageAfter  *string `explode:"true" serializationStyle:"form" queryParam:"page[after]"`
	PageBefore *string `explode:"true" serializationStyle:"form" queryParam:"page[before]"`
}

func (params *ListBroadcastsRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
func (params *ListBroadcastsRequestParams) SetPageAfter(pageAfter string) {
	params.PageAfter = &pageAfter
}
func (params *ListBroadcastsRequestParams) SetPageBefore(pageBefore string) {
	params.PageBefore = &pageBefore
}
