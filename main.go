package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pulpfree/shts-api/config"
	"github.com/pulpfree/shts-api/graph"
	"github.com/pulpfree/shts-api/graph/generated"
	"github.com/pulpfree/shts-api/mongo"
	"github.com/pulpfree/shts-api/repository"
	"github.com/pulpfree/shts-api/service"
)

var cfg *config.Config

func init() {
	cfg = &config.Config{}
	err := cfg.Load()
	if err != nil {
		log.Fatal(err)
	}
}

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := mongo.NewDB(cfg.GetMongoConnectURL(), cfg.DBName)
	if err != nil {
		log.Errorf("Failed to connect to db with error %s", err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	res := graph.NewResolver(serv)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: res,
		},
	))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
