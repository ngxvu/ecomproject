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
	if !h.ProductExist(productname) {
		fmt.Println("Lỗi: Không Tìm Thấy Tên Sản Phẩm.")
		return nil
	}
	fmt.Printf("\nThông Tin Sản Phẩm: \nProduct Name=%s \nProduct Price=%d\n", listProduct.ProductName, listProduct.Price)
	return nil
}

// checkProductExist

func (h *ProductHandler) ProductExist(product_name string) bool {
	var listProduct []model.Product
	err := h.DbConnection.Model(&listProduct).Where("product_name = ?", product_name).First(&listProduct).Error
	if err != nil {
		return false
	}
	return true
}

// View Product

func (h *ProductHandler) ViewProduct() error {
	var listProduct []model.Product
	fmt.Println(" *** Danh Sách Sản Phẩm ***")
	result := h.DbConnection.Find(&listProduct)
	if result.Error != nil {
		panic(result.Error)
	}
	for _, product := range listProduct {

		fmt.Println("\tLoại Sản Phẩm: \n\t- Tên Sản phẩm: ", product.ProductName, "\n\t- Giá Sản Phẩm: ", product.Price)
	}
	return nil
}
