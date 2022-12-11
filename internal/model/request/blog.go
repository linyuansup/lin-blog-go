package request

type GetTargetBlogRequest struct {
	ID string
}

type GetBlogListRequest struct {
	CategoryName string
	PageNum      int
	PageSize     int
}
