package repo

import (
	"app/model"
	"context"
)

type AdminRepo interface {
	CreateAdmin(ctx context.Context, admin *model.Admin) error
}
