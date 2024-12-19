package events

type GetEventsRequestParams struct {
	PageSize   *int64  `queryParam:"page[size]"`
	PageAfter  *string `queryParam:"page[after]"`
	PageBefore *string `queryParam:"page[before]"`
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
