package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Customer struct
type Customer struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Address   *Address
	Email     string `bson:"email" json:"email"`
	Name      *Name
	Phone     string    `bson:"phone" json:"phone"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type Name struct {
	First  string      `bson:"first" json:"first"`
	Last   string      `bson:"last" json:"last"`
	Prefix *NamePrefix `bson:"prefix" json:"prefix,omitempty"`
}

// Address struct
type Address struct {
	City       string `bson:"city" json:"city"`
	PostalCode string `bson:"postalCode" json:"postalCode"`
	Province   string `bson:"provinceCode" json:"province"`
	Street1    string `bson:"street1" json:"street1"`
	Street2    string `bson:"street2" json:"street2"`
}

type CreateAddress struct {
	City       string  `json:"city"`
	PostalCode string  `json:"postalCode"`
	Province   *string `json:"province"`
	Street1    string  `json:"street1"`
	Street2    *string `json:"street2"`
}

type CreateCustomer struct {
	Address   *CreateAddress `json:"address"`
	Email     *string        `json:"email"`
	Name      *CreateName    `json:"name"`
	Phone     *string        `json:"phone"`
	CreatedAt time.Time      `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time      `bson:"updatedAt" json:"updatedAt"`
}

type CreateName struct {
	First  string      `json:"first"`
	Last   string      `json:"last"`
	Prefix *NamePrefix `json:"prefix"`
}

type NamePrefix string

type UpdateAddress struct {
	City       *string `json:"city"`
	PostalCode *string `json:"postalCode"`
	Province   *string `json:"province"`
	Street1    *string `json:"street1"`
	Street2    *string `json:"street2"`
}

type UpdateCustomer struct {
	Address *UpdateAddress `json:"address"`
	Email   *string        `json:"email"`
	Name    *CreateName    `json:"name"`
	Phone   *string        `json:"phone"`
}

type UpdateName struct {
	First  *string     `json:"first"`
	Last   *string     `json:"last"`
	Prefix *NamePrefix `json:"prefix"`
}
