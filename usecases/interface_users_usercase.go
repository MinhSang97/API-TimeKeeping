package usecases

import (
	users "app/model/users"
	"context"
)

type UsersUsecase interface {
	//CreateUsers(ctx context.Context, users *users.Users) error
	CreateUsers(ctx context.Context, users *users.Users) error
	//GetAdmin(ctx context.Context, admin *admin.ReqSignIn) error
}
