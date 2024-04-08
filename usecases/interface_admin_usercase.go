package usecases

import (
	"app/model/admin"
	"context"
)

type AdminUsecase interface {
	CreateAdmin(ctx context.Context, admin *admin.Admin) error
	GetAdmin(ctx context.Context, admin *admin.ReqSignIn) error
}
