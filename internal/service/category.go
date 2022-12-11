package service

import (
	"blog/internal/dao"
	"blog/internal/model"
	"blog/internal/model/request"
)

var categoryDao *dao.CategoryDao

func init() {
	categoryDao = dao.NewCategoryDao()
}

func GetAllCategory(param *request.GetAllCategoryRequest) ([]model.Category, error) {
	return categoryDao.FindMany(make(map[string]interface{}), int64(param.PageSize), int64(param.PageNum))
}
