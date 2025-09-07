package outbound

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Base struct {
	ID         uint `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	CreateTime time.Time
	UpdateTime time.Time
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	if b.CreateTime.IsZero() {
		b.CreateTime = now
	}
	if b.UpdateTime.IsZero() {
		b.UpdateTime = now
	}
	return
}

func (b *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdateTime = time.Now()
	return
}

type User struct {
	Name string
	Age  int
	Base
}

var db *gorm.DB

func init() {
	print("执行 init\n")
	// 1. 连接 SQLite

	home, err := os.UserHomeDir()
	if err != nil {
		panic("get home path error")
	}

	db, err := gorm.Open(sqlite.Open(filepath.Join(home, ".finance-go", "test.db")), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect database, %v", err)
	}

	// 2. 自动迁移表结构
	db.AutoMigrate(&User{})
}

// demo
// // 3. 插入数据
// db.Create(&User{Name: "Alice", Age: 25})

// // 4. 查询数据
// var user User
// db.First(&user, "name = ?", "Alice")
// fmt.Println("查询结果:", user)

// // 5. 更新数据
// db.Model(&user).Update("Age", 26)

// // 6. 删除数据
// db.Delete(&user)
