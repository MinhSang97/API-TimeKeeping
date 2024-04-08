package repo

import (
	"app/model/admin"
	"context"
)

type AdminRepo interface {
	CreateAdmin(ctx context.Context, admin *admin.Admin) error
	GetAdmin(ctx context.Context, admin *admin.ReqSignIn) error
}
