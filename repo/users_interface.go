package repo

import (
	users "app/model/users"
	"context"
)

type UsersRepo interface {
	CreateUsers(ctx context.Context, users *users.Users) error
	//GetAdmin(ctx context.Context, admin *admin.ReqSignIn) error
}
