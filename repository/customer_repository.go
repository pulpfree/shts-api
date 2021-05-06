package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/pulpfree/shts-api/model"
	"github.com/pulpfree/shts-api/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var cntx = context.Background()

func (repo *Repository) CreateCustomer(customer *model.CreateCustomer) (*model.Customer, error) {
	collection := repo.db.Client.Database(repo.db.DBName).Collection(mongo.ColCustomer)
	ctx, cancel := context.WithTimeout(cntx, 5*time.Second)
	defer cancel()

	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	res, err := collection.InsertOne(ctx, customer)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	id := res.InsertedID.(primitive.ObjectID)

	// Now fetch new record
	cust := &model.Customer{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	err = collection.FindOne(ctx, filter).Decode(&cust)
	if err != nil {
		log.Fatalf("quote table error: %s", err)
		return nil, errors.New(fmt.Sprintf("customer record for id %s does not exist", id))
	}

	return cust, nil
}
