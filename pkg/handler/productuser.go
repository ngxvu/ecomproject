package handler

import (
	"bufio"
	"fmt"
	"gorm.io/gorm"
	"merakichain.com/golang_ecommerce/pkg/model"
	"merakichain.com/golang_ecommerce/pkg/utils"
	"os"
)

type ProductUserHandler struct {
	DbConnection *gorm.DB
}

func NewProProductUserHandler(db *gorm.DB) ProductUserHandler {
	return ProductUserHandler{
		DbConnection: db,
	}
}

// AddProductToUserProduct

func (h *ProductUserHandler) AddProductToFavorite() error {

	reader := bufio.NewReader(os.Stdin)
	for {
		opt, _ := utils.GetInput("Chọn Option Của Bạn: \n 1 - Nhập Sản Phẩm Bạn Muốn Thêm Vào Watchlist.\n 2 - Quay Lại Menu.", reader)
		switch opt {
		case "1":
			var products []model.Product
			input, _ := utils.GetInput("Nhập Tên Sản Phẩm: ", reader)
			err := h.DbConnection.Model(&products).Where("product_name = ?", input).First(&products).Error
			if err != nil {
				return err
			}
			var listproducts []model.Product
			result := h.DbConnection.Table("products").Find(&listproducts)
			if result.Error != nil {
				return nil
			}
			var listproductsUser []model.ProductUser
			for _, product := range products {
				productsUser := model.ProductUser{
					ProductUserID: product.ProductID,
					ProductName:   product.ProductName,
					Price:         product.Price,
					Rating:        product.Rating,
					Image:         product.Image,
				}
				listproductsUser = append(listproductsUser, productsUser)
				fmt.Sprintf("Đã Add Sản Phẩm %s Vào Danh Mục Yêu Thích", product.ProductName)
			}
			result = h.DbConnection.Table("product_users").Create(&listproductsUser)
			if result.Error != nil {
				fmt.Println("Đã xảy ra lỗi. ")
				return nil
			}
		case "2":
			return nil
		case "6":
			return nil
		default:
			fmt.Println("Lựa Chọn đó không có - Hãy Chọn Lại.")
		}
	}
}
