package repo

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"gorm.io/gorm"
	"merakichain.com/golang_ecommerce/pkg/model"
	"merakichain.com/golang_ecommerce/pkg/utils"
	"os"
)

type Edit struct {
	DbConnection *gorm.DB
}

func NewRepoEdit(Db *gorm.DB) Edit {
	return Edit{DbConnection: Db}
}

// ----------------- CheckEmailExist ----------------------

func (r *Edit) CheckEmailExist(email string) bool {
	listUser := model.User{}
	err := r.DbConnection.Model(&model.User{}).Where("email = ?", email).First(&listUser).Error
	if err != nil {
		fmt.Println("Email Không Tồn Tại - Mời Bạn Đăng Kí. ")
		return false
	}
	return true
}

// ---------------------- CheckPhoneExist ----------------------

func (r *Edit) CheckPhoneExist(phone string) bool {
	listUser := model.User{}
	err := r.DbConnection.Model(&listUser).Where("phone = ?", phone).First(&listUser).Error
	if err != nil {
		// không tìm thấy số điện thoại.
		return false
	}
	fmt.Println("Số Điện Thoại Đã Tồn Tại - Mời Bạn Đăng Kí Số Điện Thoại Khác. ")
	return true
} // ---------------------- GetToken ----------------------

func (r *Edit) GetToken() error {
	listUser := model.User{}
	r.DbConnection.First(&listUser)
	fmt.Printf("Token Của Bạn Là: %s", listUser.Token)
	return nil
}

// ---------------------- HashPassword  ----------------------

func HashPassword(password string) (string, string, error) {
	salt := make([]byte, 16) // salt = 123
	_, err := rand.Read(salt)
	if err != nil {
		return "", "", err
	}
	saltedPassword := append([]byte(password), salt...)
	hash := sha256.Sum256(saltedPassword)
	encodeSalt := base64.StdEncoding.EncodeToString(salt[:])
	encodeHash := base64.StdEncoding.EncodeToString(hash[:])
	return encodeSalt, encodeHash, nil
}

// ---------------------- EditInformation ----------------------

func (r *Edit) EditInformation(email string) error {
	reader := bufio.NewReader(os.Stdin)
	listUser := model.User{}
	fmt.Println(" Thông Tin Hiện Tại: ")
	r.DbConnection.Table("users").Select("id,first_name,last_name,email,phone,password").Where(" email = ?", email).First(&listUser)
	fmt.Printf("- First Name: %s\n- Last Name: %s\n- Email: %s\n- Phone: %s\n", listUser.FirstName, listUser.LastName, listUser.Email, listUser.Phone)
	id := listUser.ID
	for {
		opt, _ := utils.GetInput("Chọn Thông Tin Bạn Muốn Thay Đổi: \n 1 - Name \n 2 - Email \n 3 - Phone \n 4 - Password \n 5 - Return", reader)
		switch opt {
		case "1":
			newfname, _ := utils.GetInput("Nhập First Name: ", reader)
			newlname, _ := utils.GetInput("Nhập Last Name: ", reader)
			r.DbConnection.Table("users").Where("id = ?", id).Update("first_name", newfname)
			r.DbConnection.Table("users").Where("id = ?", id).Update("last_name", newlname)
		case "2":
			newemail, _ := utils.GetInput("Nhập Email: ", reader)
			if !r.CheckEmailExist(newemail) {
				r.DbConnection.Table("users").Where("id = ?", id).Update("email", newemail)
				fmt.Println("Email Đã Được Thay Đổi.")
			} else {
				return fmt.Errorf("Email Này Đã Được Đăng Kí, Vui Lòng Chọn Email Khác. ")
			}
		case "3":
			newphone, _ := utils.GetInput("Nhập Số Điện Thoại. ", reader)
			if !r.CheckPhoneExist(newphone) {
				r.DbConnection.Table("users").Where("id = ?", id).Update("phone", newphone)
				fmt.Println("Số Điện Thoại Đã Được Thay Đổi.")
			} else {
				return fmt.Errorf("Số Điện Thoại Này Đã Được Đăng Kí, Vui Lòng Chọn Số Điện Thoại Khác. ")
			}
		case "4":
			newPassword, _ := utils.GetInput("Nhập Password Mới Của Bạn. ", reader)
			newsalt, newhashedPassword, _ := HashPassword(newPassword)
			r.DbConnection.Table("users").Where("id = ?", id).Update("password", newhashedPassword)
			r.DbConnection.Table("users").Where("id = ?", id).Update("salt", newsalt)
			fmt.Println("Password Đã Được Thay Đổi.")
		case "5":
			return nil
		default:
			fmt.Println("Lựa Chọn Đó Không Có")
		}
		r.DbConnection.Table("users").Select("id,first_name,last_name,email,phone,password").Where(" email = ?", email).First(&listUser)
		fmt.Printf("- First Name: %s\n- Last Name: %s\n- Email: %s\n- Phone: %s\n", listUser.FirstName, listUser.LastName, listUser.Email, listUser.Phone)
	}
	return nil
}
