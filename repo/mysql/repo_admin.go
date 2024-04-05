package mysql

import (
	"app/model"
	"app/repo"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

// func (s adminRepository) CreateAdmin(ctx context.Context, admin *model.Admin) error {
//
//		users := admin
//		if err := s.db.Create(users).Error; err != nil {
//			return fmt.Errorf("create student error: %w", err)
//		}
//		return nil
//	}
func (s adminRepository) CreateAdmin(ctx context.Context, admin *model.Admin) error {
	users := admin
	if err := s.db.Table("Users").Create(users).Error; err != nil {
		return fmt.Errorf("create user error: %w", err)
	}
	return nil
}

var instances adminRepository

func NewAdminRepository(db *gorm.DB) repo.AdminRepo {
	if instances.db == nil {
		instances.db = db

	}
	return instances
}
