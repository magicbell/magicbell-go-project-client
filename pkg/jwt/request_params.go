package jwt

type FetchProjectTokensRequestParams struct {
	PageSize   *int64  `queryParam:"page[size]"`
	PageAfter  *string `queryParam:"page[after]"`
	PageBefore *string `queryParam:"page[before]"`
}

func (params *FetchProjectTokensRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
func (params *FetchProjectTokensRequestParams) SetPageAfter(pageAfter string) {
	params.PageAfter = &pageAfter
}
func (params *FetchProjectTokensRequestParams) SetPageBefore(pageBefore string) {
	params.PageBefore = &pageBefore
}

type FetchUserTokensRequestParams struct {
	PageSize   *int64  `queryParam:"page[size]"`
	PageAfter  *string `queryParam:"page[after]"`
	PageBefore *string `queryParam:"page[before]"`
}

func (params *FetchUserTokensRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
func (params *FetchUserTokensRequestParams) SetPageAfter(pageAfter string) {
	params.PageAfter = &pageAfter
}
func (params *FetchUserTokensRequestParams) SetPageBefore(pageBefore string) {
	params.PageBefore = &pageBefore
}
