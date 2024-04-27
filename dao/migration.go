package dao

import "fmt"

// 自动迁移数据
func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate()
	if err != nil {
		fmt.Println("err:", err)
	}
	return
}
