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
	created := false

	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	res, err := collection.InsertOne(ctx, customer)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	created = true

	// Now fetch new record
	//TODO: use the GetOne method here
	cust := &model.Customer{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	err = collection.FindOne(ctx, filter).Decode(&cust)
	if err != nil {
		log.Fatalf("customer FindOne error: %s", err)
		return nil, errors.New(fmt.Sprintf("customer record for id %s does not exist", id))
	}

	if created {
		select {
		case repo.creationStream <- cust:
		default:
		}
	}

	return cust, nil
}

func (repo *Repository) GetOne(id string) (*model.Customer, error) {
	collection := repo.db.Client.Database(repo.db.DBName).Collection(mongo.ColCustomer)
	ctx, cancel := context.WithTimeout(cntx, 5*time.Second)
	defer cancel()

	mID, _ := primitive.ObjectIDFromHex(id)
	cust := &model.Customer{}
	filter := bson.D{primitive.E{Key: "_id", Value: mID}}
	err := collection.FindOne(ctx, filter).Decode(&cust)
	if err != nil {
		log.Fatalf("customer FindOne error: %s", err)
		return nil, errors.New(fmt.Sprintf("customer record for id %s does not exist", id))
	}
	fmt.Printf("cust: %+v\n", cust.CreatedAt)
	return cust, nil
}

func (repo *Repository) GetAll() ([]*model.Customer, error) {
	collection := repo.db.Client.Database(repo.db.DBName).Collection(mongo.ColCustomer)
	ctx, cancel := context.WithTimeout(cntx, 5*time.Second)
	defer cancel()

	customers := []*model.Customer{}
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatalf("customer Find error: %s", err)
		return nil, errors.New("customer records query failed")
	}
	for cur.Next(ctx) {
		// var result bson.D
		var elem model.Customer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return nil, errors.New("customer records query failed")
		}
		customers = append(customers, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, errors.New("customer records query failed")
	}

	return customers, nil

}

// Returns a channel receiving all Customers that will be created.
func (repo *Repository) GetCreationStream() chan *model.Customer {
	return repo.creationStream
}
