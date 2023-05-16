package handleruser

import (
	"bufio"
	"merakichain.com/golang_ecommerce/pkg/service"
	"merakichain.com/golang_ecommerce/pkg/utils"
	"os"
)

type HandlerLogin struct {
	serviceLogIn service.LogIn
}

func NewHandlerLogin(serviceLogIn service.LogIn) HandlerLogin {
	return HandlerLogin{
		serviceLogIn: serviceLogIn,
	}
}

func (h *HandlerLogin) LogIn() error {
	reader := bufio.NewReader(os.Stdin)
	email, _ := utils.GetInput(" Mời Bạn Nhập Email: ", reader)
	if !h.serviceLogIn.ValidFormEmail(email) {
		return nil
		// check trùng email
	} else if !h.serviceLogIn.CheckEmailExistLogIn(email) {
		return nil
	}
	password, _ := utils.GetInput(" Mời Bạn Password: ", reader)
	h.serviceLogIn.Login(email, password)
	return nil
}
