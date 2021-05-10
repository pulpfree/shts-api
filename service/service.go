package service

import (
	"sync"

	"github.com/pulpfree/shts-api/model"
	"github.com/pulpfree/shts-api/repository"
)

type IService interface {
	CreateCustomer(*model.CreateCustomer) (*model.Customer, error)
	SaveCustomer(*model.Customer) (*model.Customer, error)
	UpdateCustomer(string, *model.UpdateCustomer) (*model.Customer, error)
	GetCustomer(string) (*model.Customer, error)
	ListCustomers() ([]*model.Customer, error)
	SubscribeCustomerCreation() *CustomerServiceObserver
	UnsubscribeCustomerCreation(*CustomerServiceObserver)
}

type Service struct {
	repo                      repository.IRepository
	customerCreationObservers []*CustomerServiceObserver
	mutex                     sync.Mutex
}

type CustomerServiceObserver struct {
	CreationStream chan *model.Customer
}

func NewService(repo repository.IRepository) *Service {
	service := &Service{repo: repo}
	go service.customerCreationStreamMultiplexer()
	return service
}
