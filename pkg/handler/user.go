package handler

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
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

	// define
	salt, hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	phone, _ := utils.GetInput(" Mời bạn nhập số điện thoại: ", reader)
	if h.CheckPhoneExist(phone) {
		fmt.Println("Số Điện Thoại Đã Được Đăng Kí")
	}

	if err := h.DbConnection.Create(&model.User{
		Email:    email,
		Password: hashedPassword,
		Phone:    phone,
		Token:    utils.DefaulToken,
		Salt:     salt,
	}).Error; err != nil {
		return err
	}
	return nil
}

// HashPassword

func HashPassword(password string) (string, string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", "", err
	}
	saltedPassword := append([]byte(password), salt...)
	hash := sha256.Sum256(saltedPassword)
	encodeHash := base64.StdEncoding.EncodeToString(hash[:])
	encodeSalt := base64.StdEncoding.EncodeToString(saltedPassword[:])
	return encodeSalt, encodeHash, nil
}

// AuthenticateUser

func (h *UserHandler) AuthenticateUser() error {
	reader := bufio.NewReader(os.Stdin)
	listUser := model.User{}
	// get input email
	email, _ := utils.GetInput("Nhập Email đăng nhập của bạn: ", reader)
	if !validFormEmail(email) {
		return fmt.Errorf("Lỗi Định Dạng Email. ")
	} else if !h.CheckEmailExist(email) {
		return fmt.Errorf(" Email Chưa Được Đăng Kí. Mời Đăng Kí ")
	}

	// Retrieve the user record from the database by email
	rs := h.DbConnection.Table("users").Where("email = ?", email).First(&listUser)
	if rs.Error != nil {
		return rs.Error
	}
	password, _ := utils.GetInput("Nhập Password Của Bạn: ", reader)

	// decode salt
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(listUser.Salt)))
	n, _ := base64.StdEncoding.Decode(dst, []byte(listUser.Salt))
	dst = dst[:n]
	fmt.Printf("%q\n", dst)
	saltedPassword := append([]byte(password), dst...)
	hash := sha256.Sum256(saltedPassword)
	encodedHash := base64.StdEncoding.EncodeToString(hash[:])

	// so sang voi pass trong database
	if listUser.Password == encodedHash {
		fmt.Println("Login successful")
		fmt.Printf("\n User found: \n- Name: %s \n- Email: %s", listUser.FirstName+" "+listUser.LastName, listUser.Email)
		return nil
	}
	return fmt.Errorf("Incorrect password. ")
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
	// In thong tin ban dau
	listUser := model.User{}
	reader := bufio.NewReader(os.Stdin)
	email, _ := utils.GetInput("Nhập Email Của Bạn: ", reader)
	fmt.Println(" Thông Tin Hiện Tại: ")
	h.DbConnection.Table("users").Select("id,first_name,last_name,email,phone").Where(" email = ?", email).First(&listUser)
	fmt.Printf("- First Name: %s\n- Last Name: %s\n- Email: %s\n- Phone: %s\n", listUser.FirstName, listUser.LastName, listUser.Email, listUser.Phone)
	// Thay doi thong tin
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
