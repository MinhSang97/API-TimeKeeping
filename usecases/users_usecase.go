package usecases

import (
	"app/dbutil"
	users "app/model/users"
	"app/repo"
	"app/repo/mysql"
	"context"
)

type usersUseCase struct {
	usersRepo repo.UsersRepo
}

func NewUsersUseCase() UsersUsecase {
	db := dbutil.ConnectDB()
	usersRepo := mysql.NewUsersRepository(db)
	return &usersUseCase{
		usersRepo: usersRepo,
	}
}

func (uc *usersUseCase) CreateUsers(ctx context.Context, users *users.Users) error {
	return uc.usersRepo.CreateUsers(ctx, users)
}

//func (uc *adminUseCase) GetAdmin(ctx context.Context, adminreq *admin.ReqSignIn) error {
//	return uc.adminRepo.GetAdmin(ctx, adminreq)
//}
