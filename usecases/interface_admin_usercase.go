package usecases

import (
	"app/model/admin_model"
	"context"
)

type AdminUsecase interface {
	CreateAdmin(ctx context.Context, admin *admin_model.Admin) error
	GetAdmin(ctx context.Context, admin *admin_model.ReqSignIn) (*admin_model.ReqSignIn, error)
	UpdateAdmin(ctx context.Context, user_id string, admin *admin_model.Admin) error
	DeleteAdmin(ctx context.Context, user_id string) error
}
