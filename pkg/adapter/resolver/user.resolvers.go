package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-ent-gqlgen/ent"
	"go-ent-gqlgen/ent/schema/ulid"
	"go-ent-gqlgen/graph/generated"
	"go-ent-gqlgen/pkg/handler"
	"go-ent-gqlgen/pkg/util/datetime"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	u, err := r.controller.User.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input ent.UpdateUserInput) (*ent.User, error) {
	u, err := r.controller.User.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

// CreateUserWithTodo is the resolver for the createUserWithTodo field.
func (r *mutationResolver) CreateUserWithTodo(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	u, err := r.controller.User.CreateWithTodo(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *ulid.ID) (*ent.User, error) {
	u, err := r.controller.User.Get(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	us, err := r.controller.User.List(ctx, after, first, before, last, where)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return us, nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *ent.User) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *ent.User) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
