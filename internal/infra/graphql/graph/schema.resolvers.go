package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	"github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/infra/graphql/graph/model"
	"github.com/rafaelsouzaribeiro/clean-architecture-golang/internal/usecase"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.OrderInput) (*model.Order, error) {
	dto := usecase.InputOrderDto{
		Price: float64(input.Price),
		Tax:   float64(input.Tax),
	}

	ouput, err := r.CreateOrderUseCase.Execute(dto)

	if err != nil {
		return nil, err
	}

	return &model.Order{
		ID:         ouput.ID,
		Price:      ouput.Price,
		Tax:        ouput.Tax,
		FinalPrice: ouput.FinalPrice,
	}, nil
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	orders, err := r.ListOrderUseCase.Execute()

	if err != nil {
		return nil, err
	}

	listOrder := []*model.Order{}

	for _, v := range orders {
		listOrder = append(listOrder, &model.Order{
			ID:         v.ID,
			Price:      v.Price,
			Tax:        v.Tax,
			FinalPrice: v.FinalPrice,
		})
	}

	return listOrder, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!