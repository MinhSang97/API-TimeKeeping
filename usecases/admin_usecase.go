package usecases

import (
	"app/dbutil"
	"app/model/admin_model"
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

func (uc *adminUseCase) CreateAdmin(ctx context.Context, admin *admin_model.Admin) error {
	return uc.adminRepo.CreateAdmin(ctx, admin)
}

func (uc *adminUseCase) GetAdmin(ctx context.Context, adminreq *admin_model.ReqSignIn) (*admin_model.ReqSignIn, error) {
	return uc.adminRepo.GetAdmin(ctx, adminreq)
}
