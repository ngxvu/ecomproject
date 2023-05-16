package handleruser

import (
	"bufio"
	"merakichain.com/golang_ecommerce/pkg/service"
	"merakichain.com/golang_ecommerce/pkg/utils"
	"os"
)

type HandlerEdit struct {
	serviceEdit service.ServiceEdit
}

func NewHandlerEdit(serviceEdit service.ServiceEdit) HandlerEdit {
	return HandlerEdit{
		serviceEdit: serviceEdit,
	}
}

func (h *HandlerEdit) EditInformation() error {
	reader := bufio.NewReader(os.Stdin)
	email, _ := utils.GetInput(" Mời Bạn Nhập Email: ", reader)
	err := h.serviceEdit.EditInformation(email)
	if err != nil {
		return err
	}
	return nil
}
