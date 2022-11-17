package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"txapp/graph/generated"
	"txapp/graph/model"
	"txapp/transactions"
)

// FetchTransactions is the resolver for the fetchTransactions field.
func (r *queryResolver) FetchTransactions(ctx context.Context) ([]*model.Transaction, error) {
	return transactions.GetTransactions(), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
