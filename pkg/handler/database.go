package handler

import (
	"fmt"
	"gorm.io/gorm"
	"merakichain.com/golang_ecommerce/pkg/model"
)

type DbHandler struct {
	DbConnection *gorm.DB
}

func NewDbHandler(db *gorm.DB) DbHandler {
	return DbHandler{DbConnection: db}
}

func (h *DbHandler) Migrate(db *gorm.DB) error {
	_ = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	err := db.AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.Product{},
		&model.ProductUser{})
	if err != nil {
		fmt.Println("Fail to Connect Database. ")
	}
	fmt.Println("Migrated Models To DB Successfully. ")
	return nil

}
