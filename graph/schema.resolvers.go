package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/pulpfree/shts-api/graph/generated"
	"github.com/pulpfree/shts-api/model"
)

func (r *customerResolver) ID(ctx context.Context, obj *model.Customer) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CreateCustomer) (*model.Customer, error) {
	return r.service.CreateCustomer(&input)
}

func (r *mutationResolver) UpdateCustomer(ctx context.Context, id string, update model.UpdateCustomer) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	return r.service.ListCustomers()
}

func (r *queryResolver) Customer(ctx context.Context, id string) (*model.Customer, error) {
	return r.service.GetCustomer(id)
}

func (r *subscriptionResolver) CustomerCreated(ctx context.Context) (<-chan *model.Customer, error) {
	subscription := r.service.SubscribeCustomerCreation()
	go func() {
		<-ctx.Done()
		r.service.UnsubscribeCustomerCreation(subscription)
	}()
	return subscription.CreationStream, nil
}

// Customer returns generated.CustomerResolver implementation.
func (r *Resolver) Customer() generated.CustomerResolver { return &customerResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type customerResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
