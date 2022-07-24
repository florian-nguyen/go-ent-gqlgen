package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-ent-gqlgen/ent"
	"go-ent-gqlgen/ent/schema/ulid"
	"go-ent-gqlgen/ent/todo"
	"go-ent-gqlgen/graph/generated"
	"go-ent-gqlgen/pkg/util/datetime"
)

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id *ulid.ID) (*ent.Todo, error) {
	t, err := r.client.Todo.Query().Where(todo.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
	t, err := r.client.Todo.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *todoResolver) CreatedAt(ctx context.Context, obj *ent.Todo) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *todoResolver) UpdatedAt(ctx context.Context, obj *ent.Todo) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
