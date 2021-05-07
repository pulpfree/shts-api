package mongo

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB and Table constants
const (
	ColCustomer = "customers"
)

var ctx = context.Background()

// MDB struct
type MDB struct {
	Client *mongo.Client
	DBName string
	db     *mongo.Database
}

// NewDB sets up new MDB struct
func NewDB(connection string, dbNm string) (*MDB, error) {

	clientOptions := options.Client().ApplyURI(connection)
	err := clientOptions.Validate()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	// defer suite.db.Close()

	return &MDB{
		Client: client,
		DBName: dbNm,
		db:     client.Database(dbNm),
	}, err
}

// Close method
func (db *MDB) Close() {
	err := db.Client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
