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
	); err != nil {
		fmt.Println("An error occurred: ", err)
		panic("Failed to connect database")
	}

	fmt.Println("Migrated models to DB successfully")

	// khởi tạo user handler
	userHandler := handler.NewUserHandler(db)
	// khởi tạo menu
	fmt.Println("Lựa Chọn Chức Năng.")
	reader := bufio.NewReader(os.Stdin)
	for {
		opt, _ := utils.GetInput("\n   1 - View Product\n   2 - Search Product.\n   3 - Sign Up.\n   4 - Login.\n   5 - Update & Exit.", reader)
		switch opt {
		case "1":
			fmt.Println("View Product")
		case "2":
			fmt.Println("Search Product")
		case "3":
			if err := userHandler.SignUp(); err != nil {
				fmt.Println("Lỗi tạo user: ", err)
			} else {
				fmt.Println("Tạo user thành công")
			}
		case "4":
			fmt.Println("SignIn")
			
		case "5":
			return
		default:
			fmt.Println("Lựa chọn đó không có - Hãy Chọn Lại.")
		}
	}

}
