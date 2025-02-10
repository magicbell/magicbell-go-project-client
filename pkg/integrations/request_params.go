package integrations

type ListIntegrationsRequestParams struct {
	PageSize   *int64  `explode:"true" serializationStyle:"form" queryParam:"page[size]"`
	PageAfter  *string `explode:"true" serializationStyle:"form" queryParam:"page[after]"`
	PageBefore *string `explode:"true" serializationStyle:"form" queryParam:"page[before]"`
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
