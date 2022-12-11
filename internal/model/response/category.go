package response

import "blog/internal/model"

type GetAllCategoryResponse struct {
	Code int              `json:"code"`
	Data []model.Category `json:"data"`
}
