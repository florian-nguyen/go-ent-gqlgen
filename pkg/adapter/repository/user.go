package repository

import (
	"context"
	"go-ent-gqlgen/ent"
	"go-ent-gqlgen/ent/user"
	"go-ent-gqlgen/pkg/entity/model"
	usecaseRepository "go-ent-gqlgen/pkg/usecase/repository"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) usecaseRepository.User {
	return &userRepository{client: client}
}

func (r *userRepository) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	u, err := r.client.User.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (r *userRepository) Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	u, err := r.client.User.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (r *userRepository) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	u, err := r.client.User.Query().Where(user.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (r *userRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput) (*model.UserConnection, error) {
	us, err := r.client.User.Query().Paginate(ctx, after, first, before, last, ent.WithUserFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return us, nil
}

func (r *userRepository) CreateWithTodo(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	client := WithTransactionalMutation(ctx)

	todo, err := client.
		Todo.
		Create().
		Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	u, err := client.User.
		Create().
		SetInput(input).
		AddTodos(todo).
		Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return u, nil
}
