package dao

import (
	"blog/global"
	"blog/internal/model"
	"context"
)

type CategoryDao struct {
	*Dao
}

func NewCategoryDao() *CategoryDao {
	return &CategoryDao{
		Dao: &Dao{
			collection: global.CategoryCollection,
		},
	}
}

func (d *CategoryDao) FindOne(filter map[string]interface{}) (model.Category, error) {
	res := d.findOne(filter)
	category := model.Category{}
	if res.Err() != nil {
		return category, res.Err()
	}
	err := res.Decode(&category)
	return category, err
}

func (d *CategoryDao) FindMany(filter map[string]interface{}, limit, page int64) ([]model.Category, error) {
	cur, err := d.findMany(filter, limit, page)
	if err != nil {
		return nil, err
	}
	var categorys []model.Category
	err = cur.All(context.Background(), &categorys)
	return categorys, err
}
