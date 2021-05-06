package graph

import "github.com/pulpfree/shts-api/service"

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	service service.IService
}

func NewResolver(service service.IService) *Resolver {
	return &Resolver{
		service: service,
	}
}
