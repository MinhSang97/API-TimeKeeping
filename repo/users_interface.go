//package repo
//
//import (
// users "app/model/users_model"
// "context"
//)
//
//type UsersRepo interface {
// CreateUsers(ctx context.Context, users *users.Users) error
// GetUsers(ctx context.Context, users *users.Users) (*users.Users, error)
//}

package repo

import (
	"app/model/users_model"
	"context"
)

type UsersRepo interface {
	CreateUsers(ctx context.Context, users *users_model.Users) error
	GetUsers(ctx context.Context, users *users_model.ReqUsersSignIn) (*users_model.ReqUsersSignIn, error)
}
