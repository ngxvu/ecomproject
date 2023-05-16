package repo

import (
	"fmt"
	"gorm.io/gorm"
	"merakichain.com/golang_ecommerce/pkg/model"
)

type RepoSignUp struct {
	DbConnection *gorm.DB
}

func NewRepoSignUp(db *gorm.DB) RepoSignUp {
	return RepoSignUp{
		DbConnection: db,
	}
}

// ----------------- CheckEmailExist ----------------------

func (r *RepoSignUp) CheckEmailExistSignUp(email string) bool {
	listUser := model.User{}
	err := r.DbConnection.Model(&model.User{}).Where("email = ?", email).First(&listUser).Error
	if err != nil {
		// không tìm thấy email
		return false
	}
	fmt.Println("Email Đã Tồn Tại - Mời Bạn Đăng Kí Email Khác. ")
	return true
}

// ---------------------- CheckEmailExist ----------------------

func (r *RepoSignUp) CheckPhoneExist(phone string) bool {
	listUser := model.User{}
	err := r.DbConnection.Model(&listUser).Where("phone = ?", phone).First(&listUser).Error
	if err != nil {
		// không tìm thấy số điện thoại.
		return false
	}
	fmt.Println("Số Điện Thoại Đã Tồn Tại - Mời Bạn Đăng Kí Số Điện Thoại Khác. ")
	return true
}

// ---------------------- GetToken ----------------------

func (r *RepoSignUp) GetToken() error {
	listUser := model.User{}
	r.DbConnection.First(&listUser)
	fmt.Printf("Token Của Bạn Là: %s", listUser.Token)
	return nil
}

// ---------------------- Create User ----------------------

func (r *RepoSignUp) CreateUser(newUser *model.User) error {
	// update database
	err := r.DbConnection.Create(newUser).Error
	if err != nil {
		return err
	}
	return nil
}

// ---------------------- Login User ----------------------

func (r *RepoSignUp) LoginUser(email string, user *model.User) error {
	listUser := user
	err := r.DbConnection.Table("users").Where("email = ?", email).First(&listUser).Error
	if err != nil {
		return err
	}
	return nil
}

// ---------------------- GetInfo ----------------------

func (r *RepoSignUp) GetInfo(email string) error {
	listUser := model.User{}
	r.DbConnection.Table("users").Where("email = ?", email).First(&listUser)
	fmt.Printf("\n User founded: \n- Name: %s \n- Email: %s", listUser.FirstName+" "+listUser.LastName, listUser.Email)
	return nil
}
