package usecases

import (
	"app/model/users_model"
	"context"
)

type UsersUsecase interface {
	CreateUsers(ctx context.Context, users *users_model.Users) error
	GetUsers(ctx context.Context, users *users_model.ReqUsersSignIn) (*users_model.ReqUsersSignIn, error)
}
