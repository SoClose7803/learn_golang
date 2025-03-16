package config

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Biến DB toàn cục
var DB *gorm.DB

// Định nghĩa model User
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

// Kết nối và tạo bảng
func ConnectDatabase() {
	dsn := "sqlserver://sa:123456123456@localhost:1433?database=user_db"
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Không thể kết nối đến SQL Server:", err)
	}

	fmt.Println("✅ Kết nối SQL Server thành công!")

	// GORM AutoMigrate: Tạo bảng nếu chưa có
	err = database.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("❌ Lỗi khi tạo bảng:", err)
	}

	fmt.Println("✅ Tạo bảng thành công!")

	DB = database
}
