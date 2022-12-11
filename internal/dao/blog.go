package dao

import (
	"blog/global"
	"blog/internal/model"
	"context"
)

type BlogDao struct {
	*Dao
}

func NewBlogDao() *BlogDao {
	return &BlogDao{
		Dao: &Dao{
			collection: global.BlogCollection,
		},
	}
}

func (d *BlogDao) FindOne(filter map[string]interface{}) (model.Blog, error) {
	res := d.findOne(filter)
	blog := model.Blog{}
	if res.Err() != nil {
		return blog, res.Err()
	}
	err := res.Decode(&blog)
	return blog, err
}

func (d *BlogDao) FindMany(filter map[string]interface{}, limit, page int64) ([]model.Blog, error) {
	cur, err := d.findMany(filter, limit, page)
	if err != nil {
		return nil, err
	}
	var blogs []model.Blog
	err = cur.All(context.Background(), &blogs)
	return blogs, err
}
