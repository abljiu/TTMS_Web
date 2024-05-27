package dao

import (
	"TTMS_Web/model"
	"fmt"
)

// 自动迁移数据
func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.Category{},
			&model.Notice{},
			&model.Order{},
			&model.Movie{},
			&model.User{},
			&model.Admin{},
			&model.BasePage{},
			&model.Carousel{})
	if err != nil {
		fmt.Println("err:", err)
	}
	return
}
