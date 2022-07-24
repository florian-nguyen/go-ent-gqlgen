package repository

import (
	"context"
	"go-ent-gqlgen/pkg/entity/model"
)

// User is an interface of repository
type User interface {
	Get(ctx context.Context, id *model.ID) (*model.User, error)
	Create(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput) (*model.UserConnection, error)
	CreateWithTodo(ctx context.Context, input model.CreateUserInput) (*model.User, error)
}
