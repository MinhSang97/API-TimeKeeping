package mysql

import (
	errors "app/error"
	"app/log"
	"app/model/admin"
	"app/repo"
	"context"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func (s adminRepository) CreateAdmin(ctx context.Context, admin *admin.Admin) error {
	users := admin

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

func (s adminRepository) GetAdmin(ctx context.Context, admin *admin.ReqSignIn) (*admin.ReqSignIn, error) {
	users := admin

	err := s.db.Table("Users").Where("email = ?", users.Email).First(users).Error
	if err != nil {
		//log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, errors.UserNotFound
		}
		log.Error(err.Error())
		return nil, err
	}
	return users, nil
}

var instances adminRepository

func NewAdminRepository(db *gorm.DB) repo.AdminRepo {
	if instances.db == nil {
		instances.db = db

	}
	return instances
}
