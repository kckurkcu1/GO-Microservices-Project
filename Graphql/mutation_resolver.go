package main

import "context"

type MutationResolver struct {
	server *Server
}

func (r *mutationResolver) createAccount(ctx context.Context, in AccountInput) (*Account, error) {

}

func (r *mutationResolver) createProduct(ctx context.Context, in ProductInput) (*Product, error) {

}

func (r *mutationResolver) createOrder(ctx context.Context, in OrderInput) (*Order, error) {

}
