package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-ent-gqlgen/ent"
	"go-ent-gqlgen/ent/schema/ulid"
	"go-ent-gqlgen/graph/generated"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id ulid.ID) (ent.Noder, error) {
	n, err := r.client.Noder(ctx, id, ent.WithNodeType(ent.IDToType))
	if err != nil {
		return nil, err
	}
	return n, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
