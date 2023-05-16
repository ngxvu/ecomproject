package handleruser

import (
	"bufio"
	"fmt"
	"merakichain.com/golang_ecommerce/pkg/service"
	"merakichain.com/golang_ecommerce/pkg/utils"
	"os"
)

type HandlerSignUp struct {
	userService service.ServiceSignUp
}

func NewHandlerSignUp(userService service.ServiceSignUp) HandlerSignUp {
	return HandlerSignUp{
		userService: userService,
	}
}

// ---------------------- Register  ----------------------

func (h *HandlerSignUp) Register() error {
	reader := bufio.NewReader(os.Stdin)
	email, _ := utils.GetInput(" Mời Bạn Nhập Email: ", reader)
	// check form email
	if !h.userService.ValidFormEmail(email) {
		return nil
		// check trùng email
	} else if h.userService.CheckEmailExistSignUp(email) {
		return nil
	}
	password, _ := utils.GetInput(" Mời Bạn Nhập Password: ", reader)
	// check trùng phone
	phone, _ := utils.GetInput(" Mời Bạn Nhập Số Điện Thoại: ", reader)
	if h.userService.CheckEmailPhoneSignUp(phone) {
		return nil
	}
	err := h.userService.SignUp(email, password, phone)
	if err != nil {
		fmt.Println("Lỗi Tạo User. ", err)
		return err
	}
	fmt.Println("Tạo User Thành Công. ")
	err = h.userService.GetInfo(email)
	if err != nil {
		return err
	}
	return nil
}
