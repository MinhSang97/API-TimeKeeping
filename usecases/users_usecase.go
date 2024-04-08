package usecases

import (
	"app/dbutil"
	"app/model/users_model"
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

func (uc *usersUseCase) CreateUsers(ctx context.Context, users *users_model.Users) error {
	return uc.usersRepo.CreateUsers(ctx, users)
}
func (uc *usersUseCase) GetUsers(ctx context.Context, users *users_model.ReqUsersSignIn) (*users_model.ReqUsersSignIn, error) {
	return uc.usersRepo.GetUsers(ctx, users)
}
