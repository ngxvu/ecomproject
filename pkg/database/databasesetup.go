package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"merakichain.com/golang_ecommerce/pkg/model"
)

func DbConnection(config model.ConfigDB) (*gorm.DB, error) {

	// Tạo connection string với các thông số đã được định nghĩa ở trên
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Host, config.Username, config.Password, config.Dbname, config.Port)
	// Kết nối tới database dựa vào connection string
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db, err
}
