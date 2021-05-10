package model

import (
	"errors"
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
	Province   string `bson:"province" json:"province"`
	Street1    string `bson:"street1" json:"street1"`
	Street2    string `bson:"street2" json:"street2"`
}

type CreateAddress struct {
	City       string  `bson:"city" json:"city"`
	PostalCode string  `bson:"postalCode" json:"postalCode"`
	Province   *string `bson:"province" json:"province"`
	Street1    string  `bson:"street1" json:"street1"`
	Street2    *string `bson:"street2" json:"street2"`
}

type CreateCustomer struct {
	Address   *CreateAddress `bson:"address" json:"address"`
	Email     *string        `bson:"email" json:"email"`
	Name      *CreateName    `bson:"name" json:"name"`
	Phone     *string        `bson:"phone" json:"phone"`
	CreatedAt time.Time      `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time      `bson:"updatedAt" json:"updatedAt"`
}

type CreateName struct {
	First  string      `bson:"first" json:"first"`
	Last   string      `bson:"last" json:"last"`
	Prefix *NamePrefix `bson:"prefix" json:"prefix"`
}

type NamePrefix string

// Update =============================================================================================================

type UpdateCustomer struct {
	Address *UpdateAddress `json:"address"`
	Email   *string        `json:"email"`
	Name    *CreateName    `json:"name"`
	Phone   *string        `json:"phone"`
}

type UpdateAddress struct {
	City       *string `json:"city"`
	PostalCode *string `json:"postalCode"`
	Province   *string `json:"province"`
	Street1    *string `json:"street1"`
	Street2    *string `json:"street2"`
}

type UpdateName struct {
	First  *string     `json:"first"`
	Last   *string     `json:"last"`
	Prefix *NamePrefix `json:"prefix"`
}

// Returns if the UpdateCustomer-Object is valid
// Could be improved by returning a list of validation errors instead of just a boolean.
func (update *UpdateCustomer) IsValid() bool {
	return true // Could add some validation logic here
}

// Merge changes from the UpdateCustomer-Request into an existing Customer.
// Returns the modified Customer or an error, if the UpdateCustomer-Object is not valid.
func (update *UpdateCustomer) MergeChanges(customer *Customer) (*Customer, error) {
	if !update.IsValid() {
		return nil, errors.New("customer object is not valid")
	}
	customer.Address.City = setStringIfNotNil(customer.Address.City, update.Address.City)
	customer.Address.PostalCode = setStringIfNotNil(customer.Address.PostalCode, update.Address.PostalCode)
	customer.Address.Province = setStringIfNotNil(customer.Address.Province, update.Address.Province)
	customer.Address.Street1 = setStringIfNotNil(customer.Address.Street1, update.Address.Street1)
	customer.Address.Street2 = setStringIfNotNil(customer.Address.Street2, update.Address.Street2)
	customer.Email = setStringIfNotNil(customer.Email, update.Email)
	customer.Name.First = setStringIfNotNil(customer.Name.First, &update.Name.First)
	customer.Name.Last = setStringIfNotNil(customer.Name.Last, &update.Name.Last)
	// customer.Name.Prefix = setStringIfNotNil(customer.Name.Prefix, &update.Name.Prefix)
	customer.Phone = setStringIfNotNil(customer.Phone, update.Phone)
	return customer, nil
}

func setStringIfNotNil(oldValue string, newValue *string) string {
	if newValue != nil {
		return *newValue
	}
	return oldValue
}
