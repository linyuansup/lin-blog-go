package response

import "blog/internal/model"

type GetTargetBlogResponse struct {
	Code int        `json:"code"`
	Data model.Blog `json:"data"`
}

type GetBlogListResponse struct {
	Code int          `json:"code"`
	Data []model.Blog `json:"data"`
}
