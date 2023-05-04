package main

import (
	"bufio"
	"fmt"
	"merakichain.com/golang_ecommerce/pkg/database"
	"merakichain.com/golang_ecommerce/pkg/handler"
	"merakichain.com/golang_ecommerce/pkg/model"
	"merakichain.com/golang_ecommerce/pkg/utils"
	"os"
)

func main() {

	// Ket noi voi database
	db, err := database.DbConnection(model.ConfigDB{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "547436",
		Dbname:   "postgres",
	})

	if err != nil {
		fmt.Println("An error occurred: ", err)

	}
	if err := db.AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.ProductUser{},
		&model.Order{},
		&model.CartItem{},
		&model.Product{},
	); err != nil {
		fmt.Println("An error occurred: ", err)
		panic("Failed to connect database")
	}

	fmt.Println("Migrated models to DB successfully")

	// khởi tạo user handler
	userHandler := handler.NewUserHandler(db)
	// khởi tạo product handler
	productHandler := handler.NewProductHandler(db)
	// khởi tạo product handler
	//cartHandler := handler.NewCartHandler()
	// khởi tạo productuserhandler
	productUserHandler := handler.NewProProductUserHandler(db)

	// khởi tạo menu
	fmt.Println("Lựa Chọn Chức Năng.")
	reader := bufio.NewReader(os.Stdin)
	for {
		opt, _ := utils.GetInput("\n   1 - View Product\n   2 - Search Product.\n   3 - Sign Up.\n   4 - Login.\n   5 - Exit.", reader)
		switch opt {
		case "1":
			err := productHandler.ViewProduct()
			if err != nil {
				return
			}
		case "2":
			err := productHandler.SearchProduct()
			if err != nil {
				return
			}
		case "3":
			if err := userHandler.SignUp(); err != nil {
				fmt.Println("Lỗi tạo user:", err)
			} else {
				fmt.Println("Tạo user thành công")
			}
		case "4":
			if err := userHandler.LogIn(); err != nil {
				fmt.Println(" Lỗi Đăng Nhập:", err)
			} else {
				for {
					opt, _ := utils.GetInput("\n   1 - Add Product.\n   2 - Add To Cart.\n   3 - Remove Product Prom Cart.\n   4 - Check Out.\n   5 - Edit Your Information.\n   6 - Exit.", reader)
					switch opt {
					case "1":
						fmt.Println("Add Product.")
						err := productUserHandler.AddProductToUserProduct()
						if err != nil {
							return
						}
					}
				}
			}
		case "5":
			return
		default:
			fmt.Println("Lựa Chọn đó không có - Hãy Chọn Lại.")
		}
	}
}
