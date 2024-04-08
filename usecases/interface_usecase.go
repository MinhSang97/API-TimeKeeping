package usecases

import (
	"app/model/admin"
	"context"
)

type StudentUsecase interface {
	GetOneByID(ctx context.Context, id int) (admin.Student, error)
	GetAll(ctx context.Context) ([]admin.Student, error)
	UpdateOne(ctx context.Context, id int, student *admin.Student) error
	DeleteOne(ctx context.Context, id int) error
	Search(ctx context.Context, Value string) ([]admin.Student, error)
	CreateStudent(ctx context.Context, student *admin.Student) error
}
