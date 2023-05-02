package handler

import (
	"bufio"
	"fmt"
	"gorm.io/gorm"
	"merakichain.com/golang_ecommerce/pkg/model"
	"merakichain.com/golang_ecommerce/pkg/utils"
	"os"
)

type ProductHandler struct {
	DbConnection *gorm.DB
}

// Method ProductHandler

func NewProductHandler(db *gorm.DB) ProductHandler {
	return ProductHandler{
		DbConnection: db,
	}
}

// Search Product

func (h *ProductHandler) SearchProduct() error {

	reader := bufio.NewReader(os.Stdin)
	productname, _ := utils.GetInput(" Nhập Tên Sản Phẩm Cần Tìm: ", reader)
	listProduct := model.Product{}
	h.DbConnection.First(&listProduct, "product_name = ?", productname)
	fmt.Printf("User found: Product Name=%s Product Price=%f", listProduct.Product_Name, listProduct.Price)
	return nil
}

// View Product

func (h *ProductHandler) ViewProduct() error {
	listProduct := model.Product{}
	result := h.DbConnection.Find(&listProduct)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(listProduct.Product_Name)
	return nil
}
