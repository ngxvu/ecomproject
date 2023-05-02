package handler

import (
	"bufio"
	"fmt"
	"gorm.io/gorm"
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
	return err == nil
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

// ----------------- CheckSoDienThoai ----------------------

func (h *UserHandler) CheckPhoneExist(phone string) bool {
	listUser := model.User{}
	err := h.DbConnection.Model(&listUser).Where("phone = ?", phone).First(&listUser).Error
	if err != nil {
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
	phone, _ := utils.GetInput(" Mời bạn nhập số điện thoại: ", reader)
	if h.CheckPhoneExist(phone) {
		fmt.Println("Số Điện Thoại Đã Được Đăng Kí")
	}

	if err := h.DbConnection.Create(&model.User{
		Email:    email,
		Password: password,
		Phone:    phone,
	}).Error; err != nil {
		return err
	}

	return nil
}

// Login ()

func (h *UserHandler) LogIn() error {
	listUser := model.User{}
	reader := bufio.NewReader(os.Stdin)
	email, _ := utils.GetInput("Nhập Email đăng nhập của bạn: ", reader)
	if !validFormEmail(email) {
		return fmt.Errorf("Lỗi định dạng email. ")
	} else if !h.CheckEmailExist(email) {
		return fmt.Errorf("Email chưa được đăng kí, mời đăng kí. ")
	}
	h.DbConnection.First(&listUser, "email = ?", email)
	h.VerifyPassword()
	return nil
}

// HashPassword

// VerifyPassword

func (h *UserHandler) VerifyPassword() bool {

	listUser := model.User{}
	reader := bufio.NewReader(os.Stdin)
	password, _ := utils.GetInput("Nhập Password Của Bạn: ", reader)
	h.DbConnection.First(&listUser, "password = ?", password)
	if password != listUser.Password {
		fmt.Println("Incorrect password")
		return false
	}
	fmt.Println("Login successful")
	fmt.Printf("\n User found: \n- Name: %s \n- Email: %s", listUser.First_Name+" "+listUser.Last_Name, listUser.Email)
	return true
}

// SearchProduct

// SearchProductByQuery
