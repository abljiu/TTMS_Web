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
	var categoryString string
	var categories []model.Category

	err = dao.DB.Model(&model.Category{}).Where("id in ?", categoryId).Find(&categories).Error
	for _, category := range categories {
		categoryString += category.CategoryName + " "
	}

	return categoryString, err
}
