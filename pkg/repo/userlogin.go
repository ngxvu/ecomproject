package repo

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"gorm.io/gorm"
	"merakichain.com/golang_ecommerce/pkg/model"
)

type LogIn struct {
	DbConnection *gorm.DB
}

func NewRepoLogIn(Db *gorm.DB) LogIn {

	return LogIn{DbConnection: Db}
}

// ----------------- CheckEmailExist ----------------------

func (r *LogIn) CheckEmailExistLogIn(email string) bool {
	listUser := model.User{}
	err := r.DbConnection.Model(&model.User{}).Where("email = ?", email).First(&listUser).Error
	if err != nil {
		fmt.Println("Email Không Tồn Tại - Mời Bạn Đăng Kí. ")
		return false
	}
	return true
}

// ---------------------- CheckPhoneExist ----------------------

func (r *LogIn) CheckPhoneExistLogIn(phone string) bool {
	listUser := model.User{}
	err := r.DbConnection.Model(&listUser).Where("phone = ?", phone).First(&listUser).Error
	if err != nil {
		// không tìm thấy số điện thoại.
		return false
	}
	fmt.Println("Số Điện Thoại Đã Tồn Tại - Mời Bạn Đăng Kí Số Điện Thoại Khác. ")
	return true
}

// ---------------------- ComparePassword ----------------------

func (r *LogIn) ComparePassword(email string, password string) bool {

	listUser := model.User{}
	r.DbConnection.Table("users").Where("email = ?", email).First(&listUser)
	deSalt := make([]byte, base64.StdEncoding.DecodedLen(len(listUser.Salt)))
	n, _ := base64.StdEncoding.Decode(deSalt, []byte(listUser.Salt))
	deSalt = deSalt[:n] //[]byte
	// decode Password
	dePass := make([]byte, base64.StdEncoding.DecodedLen(len(listUser.Password)))
	m, _ := base64.StdEncoding.Decode(dePass, []byte(listUser.Password))
	dePass = dePass[:m] //[]byte

	saltedPassword := append([]byte(password), deSalt...)
	hash := sha256.Sum256(saltedPassword)
	encodedHash := base64.StdEncoding.EncodeToString(hash[:])
	// so sang voi pass trong database
	if listUser.Password != encodedHash {
		fmt.Println("Incorrect password. ")
		return false
	} else {
		fmt.Println("Login successful")
		fmt.Printf("\n User found: \n- Name: %s \n- Email: %s \n- UserToken: %s \n", listUser.FirstName+" "+listUser.LastName, listUser.Email, listUser.Token)
		return true
	}
}

// ---------------------- Edit Info ----------------------
