package graph

import "github.com/pulpfree/shts-api/service"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	service service.IService
}

func NewResolver(service service.IService) *Resolver {
	return &Resolver{
		service: service,
	}
}
