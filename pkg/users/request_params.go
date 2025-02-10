package users

type ListUsersRequestParams struct {
	PageSize   *int64  `explode:"true" serializationStyle:"form" queryParam:"page[size]"`
	PageAfter  *string `explode:"true" serializationStyle:"form" queryParam:"page[after]"`
	PageBefore *string `explode:"true" serializationStyle:"form" queryParam:"page[before]"`
}

func (params *ListUsersRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
func (params *ListUsersRequestParams) SetPageAfter(pageAfter string) {
	params.PageAfter = &pageAfter
}
func (params *ListUsersRequestParams) SetPageBefore(pageBefore string) {
	params.PageBefore = &pageBefore
}
