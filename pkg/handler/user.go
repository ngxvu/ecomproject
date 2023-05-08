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
		Token:    utils.DefaulToken,
	}).Error; err != nil {
		return err
	}

	return nil
}

// Login ()

func (h *UserHandler) LogIn() error {
	reader := bufio.NewReader(os.Stdin)
	email, _ := utils.GetInput("Nhập Email đăng nhập của bạn: ", reader)
	if !validFormEmail(email) {
		return fmt.Errorf("Lỗi định dạng email. ")
	} else if !h.CheckEmailExist(email) {
		return fmt.Errorf("Email chưa được đăng kí, mời đăng kí. ")
	}
	h.DbConnection.First(&model.User{}, "email = ?", email)
	err := h.VerifyPassword()
	if err != nil {
		return err
	}
	return nil
}

// VerifyPassword

func (h *UserHandler) VerifyPassword() error {

	listUser := model.User{}
	reader := bufio.NewReader(os.Stdin)
	password, _ := utils.GetInput("Nhập Password Của Bạn: ", reader)
	h.DbConnection.First(&listUser, "password = ?", password)
	if password != listUser.Password {
		return fmt.Errorf("Incorrect password. ")
	}
	fmt.Println("Login successful")
	fmt.Printf("\n User found: \n- Name: %s \n- Email: %s", listUser.FirstName+" "+listUser.LastName, listUser.Email)
	return nil
}

// GetToken

func (h *UserHandler) GetToken() error {
	listUser := model.User{}
	h.DbConnection.First(&listUser)
	fmt.Printf("Token Của Bạn Là: %s", listUser.Token)
	return nil
}

// PrintUserInformation

func (h *UserHandler) EditInfo() error {
	listUser := model.User{}
	reader := bufio.NewReader(os.Stdin)
	email, _ := utils.GetInput("Nhập Email Của Bạn: ", reader)
	fmt.Println(" Thông Tin Hiện Tại: ")
	h.DbConnection.Table("users").Select("id,first_name,last_name,email,phone").Where(" email = ?", email).First(&listUser)
	fmt.Printf("- First Name: %s\n- Last Name: %s\n- Email: %s\n- Phone: %s\n", listUser.FirstName, listUser.LastName, listUser.Email, listUser.Phone)
	id := listUser.ID
	for {
		opt, _ := utils.GetInput("Chọn Thông Tin Bạn Muốn Thay Đổi: \n 1 - Name \n 2 - Email \n 3 - Phone \n 4 - Return", reader)
		switch opt {
		case "1":
			newfname, _ := utils.GetInput("Nhập First Name: ", reader)
			newlname, _ := utils.GetInput("Nhập Last Name: ", reader)
			h.DbConnection.Table("users").Where("id = ?", id).Update("first_name", newfname)
			h.DbConnection.Table("users").Where("id = ?", id).Update("last_name", newlname)
		case "2":
			newemail, _ := utils.GetInput("Nhập Email: ", reader)
			if !h.CheckEmailExist(newemail) {
				h.DbConnection.Table("users").Where("id = ?", id).Update("email", newemail)
			} else {
				return fmt.Errorf("Email Này Đã Được Đăng Kí, Vui Lòng Chọn Email Khác. ")
			}
		case "3":
			newphone, _ := utils.GetInput("Nhập Số Điện Thoại", reader)
			if !h.CheckPhoneExist(newphone) {
				h.DbConnection.Table("user").Where("id = ?", id).Update("phone", newphone)
			} else {
				return fmt.Errorf("Số Điện Thoại Này Đã Được Đăng Kí, Vui Lòng Chọn Số Điện Thoại Khác. ")
			}
		case "4":
			return nil
		default:
			fmt.Println("Lựa Chọn Đó Không Có")
		}
		fmt.Printf("- First Name: %s\n- Last Name: %s\n- Email: %s\n- Phone: %s\n", listUser.FirstName, listUser.LastName, listUser.Email, listUser.Phone)
	}
}

// SearchProductByQuery
