package broadcasts

type ListBroadcastsRequestParams struct {
	PageSize   *int64  `queryParam:"page[size]"`
	PageAfter  *string `queryParam:"page[after]"`
	PageBefore *string `queryParam:"page[before]"`
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
