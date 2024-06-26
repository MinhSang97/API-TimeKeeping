package mysql

import (
	errors "app/error"
	"app/log"
	"app/model/users_model"
	"app/repo"
	"context"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func (s usersRepository) CreateUsers(ctx context.Context, users *users_model.Users) error {
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

func (s usersRepository) GetUsers(ctx context.Context, users *users_model.ReqUsersSignIn) (*users_model.ReqUsersSignIn, error) {

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

func (s usersRepository) UpdateUsers(ctx context.Context, user_id string, users *users_model.Users) error {
	err := s.db.Table("Users").Where("user_id = ?", user_id).Updates(users).Error
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {

			if driverErr.Number == 1062 {
				return errors.UserNotUpdated
			}
		}
		return errors.SignUpFail
	}
	return nil
}

func (s usersRepository) DeleteUsers(ctx context.Context, user_id string) error {
	var user users_model.Users

	// Check if user exists
	if err := s.db.Table("Users").Where("user_id = ?", user_id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.UserNotFound
		}
		return err
	}

	// If user exists, delete the user
	if err := s.db.Table("Users").Where("user_id = ?", user_id).Delete(&user).Error; err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 {
				return errors.UserNotDeleted
			}
		}
		return err
	}

	return nil
}

var instancesUsers usersRepository

func NewUsersRepository(db *gorm.DB) repo.UsersRepo {
	if instancesUsers.db == nil {
		instancesUsers.db = db

	}
	return instancesUsers
}
