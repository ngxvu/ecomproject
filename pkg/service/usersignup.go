package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"merakichain.com/golang_ecommerce/pkg/model"
	"merakichain.com/golang_ecommerce/pkg/repo"
	"merakichain.com/golang_ecommerce/pkg/utils"
	"net/mail"
)

type ServiceSignUp struct {
	repoSignUp repo.RepoSignUp
}

func NewServiceSignUp(repoSignUp repo.RepoSignUp) ServiceSignUp {
	return ServiceSignUp{
		repoSignUp: repoSignUp,
	}
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

// ---------------------- CheckFormEmail ----------------------

func (s *ServiceSignUp) ValidFormEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Println("Sai Định Dạng Email.")
		return false
	}
	return true
}

// ---------------------- CheckEmailExist ----------------------

func (s *ServiceSignUp) CheckEmailExistSignUp(email string) bool {
	if s.repoSignUp.CheckEmailExistSignUp(email) {
		return true
	}
	return false
}

// ---------------------- CheckPhoneExist----------------------

func (s *ServiceSignUp) CheckEmailPhoneSignUp(phone string) bool {
	if s.repoSignUp.CheckPhoneExist(phone) {
		return true
	}
	return false
}

// ---------------------- SignUp  ----------------------

func (s *ServiceSignUp) SignUp(email string, password string, phone string) error {

	salt, hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	err = s.repoSignUp.CreateUser(&model.User{
		Email:    email,
		Password: hashedPassword,
		Salt:     salt,
		Phone:    phone,
		Token:    utils.DefaulToken,
	})
	if err != nil {
		return err
	}
	return nil
}

// ---------------------- GetInFo----------------------

func (s *ServiceSignUp) GetInfo(email string) error {
	err := s.repoSignUp.GetInfo(email)
	if err != nil {
		return err
	}
	return nil
}

//// ---------------------- DeHashPassword  ----------------------
//
//func (s *UserService) CheckPassword(password string, salt string) error {
//	// decode Salt
//	deSalt := make([]byte, base64.StdEncoding.DecodedLen(len(salt))
//	n, _ := base64.StdEncoding.Decode(deSalt, []byte(salt))
//	deSalt = deSalt[:n]
//	// decode Password
//	dePass := make([]byte, base64.StdEncoding.DecodedLen(len(password)))
//	m, _ := base64.StdEncoding.Decode(dePass, []byte(password))
//	dePass = dePass[:m] //[]byte
//	saltedPassword := append([]byte(password), deSalt...)
//	hash := sha256.Sum256(saltedPassword)
//	encodedHash := base64.StdEncoding.EncodeToString(hash[:])
//	// so sang voi pass trong database
//	if password != encodedHash {
//		 return fmt.Errorf("Incorrect password. ")
//	}
//		return nil
//	}

//// ---------------------- SignIn  ----------------------
//
//func (s *UserService) SignIn(email string, password string ) error {
//
//	// Retrieve the handleruser record from the database by email
//	err := s.userRepo.LoginUser(email,&model.User{
//		Email: email,
//	})
//	if err != nil {
//		return err
//	}
//
//	// xu ly logic
//	if !validFormEmail(email) {
//		return fmt.Errorf("Sai Định Dạng Email. ")
//	} else if !s.userRepo.CheckEmailExist(email) {
//		return fmt.Errorf(" Email Không Tồn Tại - Mời Bạn Đăng Kí. ")
//	}
//	err = CheckPassword(password,salt)
//	if err != nil {
//		return err
//	}
//	fmt.Println("Login successful")
//
//
//	return nil
//}
