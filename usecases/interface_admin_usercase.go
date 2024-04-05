package usecases

import (
	"app/model"
	"context"
)

type AdminUsecase interface {
	CreateAdmin(ctx context.Context, admin *model.Admin) error
}
