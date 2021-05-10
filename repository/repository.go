package repository

import (
	"github.com/pulpfree/shts-api/model"
	"github.com/pulpfree/shts-api/mongo"
)

type IRepository interface {
	CreateCustomer(customer *model.CreateCustomer) (*model.Customer, error)
	SaveCustomer(customer *model.Customer) (*model.Customer, error)
	GetOne(string) (*model.Customer, error)
	GetAll() ([]*model.Customer, error)
	GetCreationStream() chan *model.Customer
}

type Repository struct {
	db             *mongo.MDB
	creationStream chan *model.Customer
}

func NewRepository(db *mongo.MDB) *Repository {
	return &Repository{
		db:             db,
		creationStream: make(chan *model.Customer),
	}
}
