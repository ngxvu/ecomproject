package handler

import (
	"bufio"
	"fmt"
	"gorm.io/gorm"
	"log"
	"merakichain.com/golang_ecommerce/pkg/model"
	"merakichain.com/golang_ecommerce/pkg/utils"
	"net/mail"
	"os"
)

type UserHandler struct {
	DbConnection *gorm.DB
}

// Method UserHandler

func NewUserHandler(db *gorm.DB) UserHandler {
	return UserHandler{
		DbConnection: db,
	}
}

// ----------------- CheckFormEmail ----------------------
func validFormEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

// ----------------- CheckEmailExist ----------------------

func (h *UserHandler) CheckEmailExist(email string) bool {
	listUser := model.User{}
	err := h.DbConnection.Model(&model.User{}).Where("email = ?", email).First(&listUser).Error
	if err != nil {
		// email chưa tôn tại, nên không tìm tim thấy
		return false
	}
	return true
}

// Signup ()

func (h *UserHandler) SignUp() error {
	reader := bufio.NewReader(os.Stdin)
	email, _ := utils.GetInput(" Mời bạn nhập email: ", reader)
	if !validFormEmail(email) {
		return fmt.Errorf(" Lỗi định dạng email. ")
	} else if h.CheckEmailExist(email) {
		return fmt.Errorf(" Email Đã Tồn Tại. ")
	}
	password, _ := utils.GetInput(" Mời bạn nhập password: ", reader)
	if err := h.DbConnection.Create(&model.User{
		Email:    email,
		Password: password,
	}).Error; err != nil {
		return err
	}
	return nil
}

// Login ()
//
//func ( h *UserHandler) SignIn() error{
//	reader := bufio.NewReader(os.Stdin)
//	email,_ := utils.GetInput("")
//}

// HashPassword

// VerifyPassword

// SearchProduct

// SearchProductByQuery
