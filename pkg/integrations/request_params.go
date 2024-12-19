package integrations

type ListIntegrationsRequestParams struct {
	PageSize   *int64  `queryParam:"page[size]"`
	PageAfter  *string `queryParam:"page[after]"`
	PageBefore *string `queryParam:"page[before]"`
}

func (params *ListIntegrationsRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
func (params *ListIntegrationsRequestParams) SetPageAfter(pageAfter string) {
	params.PageAfter = &pageAfter
}
func (params *ListIntegrationsRequestParams) SetPageBefore(pageBefore string) {
	params.PageBefore = &pageBefore
}
