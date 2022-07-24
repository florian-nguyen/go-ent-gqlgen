package registry

import (
	"go-ent-gqlgen/pkg/adapter/controller"
	"go-ent-gqlgen/pkg/adapter/repository"
	"go-ent-gqlgen/pkg/usecase/usecase"
)

func (r *registry) NewUserController() controller.User {
	repo := repository.NewUserRepository(r.client)
	u := usecase.NewUserUsecase(repo)

	return controller.NewUserController(u)
}
