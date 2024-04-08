package repo

import (
	"app/model/admin_model"
	"context"
)

type AdminRepo interface {
	CreateAdmin(ctx context.Context, admin *admin_model.Admin) error
	GetAdmin(ctx context.Context, admin *admin_model.ReqSignIn) (*admin_model.ReqSignIn, error)
}
