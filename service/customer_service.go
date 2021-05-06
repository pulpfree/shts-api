package service

import (
	"github.com/pulpfree/shts-api/model"
)

func (svc *Service) CreateCustomer(input *model.CreateCustomer) (*model.Customer, error) {

	return svc.repo.CreateCustomer(input)
}
