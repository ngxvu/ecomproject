package service

import (
	"fmt"
	"merakichain.com/golang_ecommerce/pkg/repo"
	"net/mail"
)

type LogIn struct {
	repoLogIn repo.LogIn
}

func NewServiceLogIn(repoLogIn repo.LogIn) LogIn {
	return LogIn{
		repoLogIn: repoLogIn,
	}
}

// ---------------------- CheckFormEmail ----------------------

func (s *LogIn) ValidFormEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Println("Sai Định Dạng Email.")
		return false
	}
	return true
}

// ---------------------- CheckEmailExist ----------------------

func (s *LogIn) CheckEmailExistLogIn(email string) bool {
	if s.repoLogIn.CheckEmailExistLogIn(email) {
		return true
	}
	return false
}

// ---------------------- CheckPhoneExist----------------------

func (s *LogIn) CheckPhoneExistLogIn(phone string) bool {
	if s.repoLogIn.CheckPhoneExistLogIn(phone) {
		return true
	}
	return false
}

// ---------------------- Login----------------------

func (s *LogIn) Login(email string, password string) bool {
	s.repoLogIn.ComparePassword(email, password)
	return false
}

// ---------------------- Edit Info ----------------------
