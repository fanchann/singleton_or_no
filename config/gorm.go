package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormConnectDB() *gorm.DB {
	fmt.Println("Database connection opened")
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/articles"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(0)    // ← tidak menyimpan koneksi idle
	sqlDB.SetMaxOpenConns(10)   // ← bisa membuka banyak koneksi
	sqlDB.SetConnMaxLifetime(0) // ← koneksi tidak kedaluwarsa

	return db
}
