package utils

import (
	"bufio"
	"fmt"
	"merakichain.com/golang_ecommerce/pkg/model"
	"net/mail"
	"os"
	"strings"
)

var listUser []model.User

// ----------------- CheckFormEmail ----------------------
func validFormEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// ----------------- CheckEmailExist ----------------------
func checkEmailExist(email string, listUsers []model.User) bool {
	for _, user := range listUsers {
		if email == user.Email {
			return false
		}
	}
	return true
}

//----------------- Save.Json.File ----------------------
//
//func Save(data []model.User) error {
//	tmp, err := json.Marshal(data)
//	if err != nil {
//		return err
//	}
//
//	err = os.WriteFile("./data.json", tmp, 0644)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//----------------- Save.Json.File ----------------------

// ----------------- GetInput ----------------------

func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func GetInputString() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Đã xảy ra lỗi: ", err)
		return ""
	}
	inputRemovedEnter := strings.Trim(input, "\r\n")
	return inputRemovedEnter
}
