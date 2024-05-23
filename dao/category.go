package dao

import (
	"TTMS_Web/model"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

// GetCategory 根据id数组返回类型名称
func (dao *CategoryDao) GetCategory(categoryId []uint) (string, error) {
	var err error
	var CategoryString string
	category := model.Category{}

	for _, id := range categoryId {
		err = dao.DB.Model(&model.Category{}).Where("id=?", id).First(&category).Error
		CategoryString += category.CategoryName + " "
	}

	return CategoryString, err
}
