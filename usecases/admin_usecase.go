package usecases

import (
	"app/dbutil"
	"app/model"
	"app/repo"
	"app/repo/mysql"
	"context"
)

type adminUseCase struct {
	adminRepo repo.AdminRepo
}

func NewAdminUseCase() AdminUsecase {
	db := dbutil.ConnectDB()
	adminRepo := mysql.NewAdminRepository(db)
	return &adminUseCase{
		adminRepo: adminRepo,
	}
}

func (uc *adminUseCase) CreateAdmin(ctx context.Context, admin *model.Admin) error {
	return uc.adminRepo.CreateAdmin(ctx, admin)
}
