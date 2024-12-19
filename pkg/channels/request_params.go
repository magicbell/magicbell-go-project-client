package channels

type GetMobilePushApnsUserTokensRequestParams struct {
	PageSize   *int64  `queryParam:"page[size]"`
	PageAfter  *string `queryParam:"page[after]"`
	PageBefore *string `queryParam:"page[before]"`
}

func (params *GetMobilePushApnsUserTokensRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
func (params *GetMobilePushApnsUserTokensRequestParams) SetPageAfter(pageAfter string) {
	params.PageAfter = &pageAfter
}
func (params *GetMobilePushApnsUserTokensRequestParams) SetPageBefore(pageBefore string) {
	params.PageBefore = &pageBefore
}

type GetMobilePushExpoUserTokensRequestParams struct {
	PageSize   *int64  `queryParam:"page[size]"`
	PageAfter  *string `queryParam:"page[after]"`
	PageBefore *string `queryParam:"page[before]"`
}

func (params *GetMobilePushExpoUserTokensRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
func (params *GetMobilePushExpoUserTokensRequestParams) SetPageAfter(pageAfter string) {
	params.PageAfter = &pageAfter
}
func (params *GetMobilePushExpoUserTokensRequestParams) SetPageBefore(pageBefore string) {
	params.PageBefore = &pageBefore
}

type GetMobilePushFcmUserTokensRequestParams struct {
	PageSize   *int64  `queryParam:"page[size]"`
	PageAfter  *string `queryParam:"page[after]"`
	PageBefore *string `queryParam:"page[before]"`
}

func (params *GetMobilePushFcmUserTokensRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
func (params *GetMobilePushFcmUserTokensRequestParams) SetPageAfter(pageAfter string) {
	params.PageAfter = &pageAfter
}
func (params *GetMobilePushFcmUserTokensRequestParams) SetPageBefore(pageBefore string) {
	params.PageBefore = &pageBefore
}

type GetSlackUserTokensRequestParams struct {
	PageSize   *int64  `queryParam:"page[size]"`
	PageAfter  *string `queryParam:"page[after]"`
	PageBefore *string `queryParam:"page[before]"`
}

func (params *GetSlackUserTokensRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
func (params *GetSlackUserTokensRequestParams) SetPageAfter(pageAfter string) {
	params.PageAfter = &pageAfter
}
func (params *GetSlackUserTokensRequestParams) SetPageBefore(pageBefore string) {
	params.PageBefore = &pageBefore
}

type GetTeamsUserTokensRequestParams struct {
	PageSize   *int64  `queryParam:"page[size]"`
	PageAfter  *string `queryParam:"page[after]"`
	PageBefore *string `queryParam:"page[before]"`
}

func (params *GetTeamsUserTokensRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
func (params *GetTeamsUserTokensRequestParams) SetPageAfter(pageAfter string) {
	params.PageAfter = &pageAfter
}
func (params *GetTeamsUserTokensRequestParams) SetPageBefore(pageBefore string) {
	params.PageBefore = &pageBefore
}

type GetWebPushUserTokensRequestParams struct {
	PageSize   *int64  `queryParam:"page[size]"`
	PageAfter  *string `queryParam:"page[after]"`
	PageBefore *string `queryParam:"page[before]"`
}

func (params *GetWebPushUserTokensRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
func (params *GetWebPushUserTokensRequestParams) SetPageAfter(pageAfter string) {
	params.PageAfter = &pageAfter
}
func (params *GetWebPushUserTokensRequestParams) SetPageBefore(pageBefore string) {
	params.PageBefore = &pageBefore
}
