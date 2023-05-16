package service

import "merakichain.com/golang_ecommerce/pkg/repo"

type ServiceEdit struct {
	repoEdit repo.Edit
}

func NewServiceEdit(repoEdit repo.Edit) ServiceEdit {
	return ServiceEdit{repoEdit: repoEdit}
}

func (s *ServiceEdit) EditInformation(email string) error {
	err := s.repoEdit.EditInformation(email)
	if err != nil {
		return err
	}
	return nil
}
