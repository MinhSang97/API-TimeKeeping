package mysql

import (
	errors "app/error"
	users "app/model/users"
	"app/repo"
	"context"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func (s usersRepository) CreateUsers(ctx context.Context, users *users.Users) error {

	err := s.db.Table("Users").Create(users).Error
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {

			if driverErr.Number == 1062 {
				return errors.UserConflict
			}
		}
		return errors.SignUpFail
	}
	return nil
}

//func (s adminRepository) GetAdmin(ctx context.Context, admin *admin.ReqSignIn) error {
//	users := admin
//
//	err := s.db.Table("Users").Where("email = ?", users.Email).First(users).Error
//	if err != nil {
//		log.Error(err.Error())
//		if err == gorm.ErrRecordNotFound {
//			return errors.UserNotFound
//		}
//		log.Error(err.Error())
//		return err
//	}
//	return nil
//}

var instancesUsers usersRepository

func NewUsersRepository(db *gorm.DB) repo.UsersRepo {
	if instancesUsers.db == nil {
		instancesUsers.db = db

	}
	return instancesUsers
}
