package service

import (
	"github.com/pulpfree/shts-api/model"
	"github.com/pulpfree/shts-api/repository"
)

type IService interface {
	CreateCustomer(*model.CreateCustomer) (*model.Customer, error)
}

type Service struct {
	repo repository.IRepository
	// articleCreationObservers []*ArticleServiceObserver
	// mutex                    sync.Mutex
}

func NewService(repo repository.IRepository) *Service {
	service := &Service{repo: repo}
	// go service.articleCreationStreamMultiplexer()
	return service
}
