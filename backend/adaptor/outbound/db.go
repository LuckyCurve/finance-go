package outbound

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Age  int
}

var db *gorm.DB

func init() {

	print("执行 init\n")
	// 1. 连接 SQLite

	db, err := gorm.Open(sqlite.Open("../data/test.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect database, %v", err)
	}

	// 2. 自动迁移表结构
	db.AutoMigrate(&User{})

	// 3. 插入数据
	db.Create(&User{Name: "Alice", Age: 25})

	// 4. 查询数据
	var user User
	db.First(&user, "name = ?", "Alice")
	fmt.Println("查询结果:", user)

	// 5. 更新数据
	db.Model(&user).Update("Age", 26)

	// 6. 删除数据
	db.Delete(&user)
}
