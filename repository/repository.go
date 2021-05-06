package repository

import (
	"github.com/pulpfree/shts-api/model"
	"github.com/pulpfree/shts-api/mongo"
)

type IRepository interface {
	CreateCustomer(customer *model.CreateCustomer) (*model.Customer, error)
}

type Repository struct {
	db *mongo.MDB
}

func NewRepository(db *mongo.MDB) *Repository {
	return &Repository{
		db: db,
		// creationStream: make(chan *model.Article),
	}
}
