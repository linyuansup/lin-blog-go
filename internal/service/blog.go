package service

import (
	"blog/internal/dao"
	"blog/internal/model"
	"blog/internal/model/request"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var blogDao *dao.BlogDao

func init() {
	blogDao = dao.NewBlogDao()
	categoryDao = dao.NewCategoryDao()
}

func GetBlogList(param *request.GetBlogListRequest) ([]model.Blog, error) {
	var err error
	var blogs []model.Blog
	var cat model.Category
	if param.CategoryName != "" {
		cat, err = categoryDao.FindOne(map[string]interface{}{"title": param.CategoryName})
		if err != nil {
			return nil, err
		}
		blogs, err = blogDao.FindMany(map[string]interface{}{"category": cat.ID}, int64(param.PageSize), int64(param.PageNum))
	} else {
		blogs, err = blogDao.FindMany(map[string]interface{}{}, int64(param.PageSize), int64(param.PageNum))
	}
	if err != nil {
		return nil, err
	}
	for _, b := range blogs {
		b.Content = ""
	}
	return blogs, nil
}

func GetTargetBlog(param *request.GetTargetBlogRequest) (model.Blog, error) {
	var blog model.Blog
	oid, err := primitive.ObjectIDFromHex(param.ID)
	if err != nil {
		return blog, err
	}
	return blogDao.FindOne(map[string]interface{}{"_id": oid})
}
