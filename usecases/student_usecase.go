package usecases

import (
	"app/dbutil"
	"app/model/admin"
	"app/repo"
	"app/repo/mysql"
	"context"
)

type studentUseCase struct {
	studentRepo repo.StudentRepo
}

func NewStudentUseCase() StudentUsecase {
	db := dbutil.ConnectDB()
	studentRepo := mysql.NewStudentRepository(db)
	return &studentUseCase{
		studentRepo: studentRepo,
	}
}

// Tested
func (uc *studentUseCase) GetOneByID(ctx context.Context, id int) (admin.Student, error) {
	return uc.studentRepo.GetOneByID(ctx, id)
}

// Tested
func (uc *studentUseCase) GetAll(ctx context.Context) ([]admin.Student, error) {
	return uc.studentRepo.GetAll(ctx)
}

// Tested
func (uc *studentUseCase) UpdateOne(ctx context.Context, id int, student *admin.Student) error {
	return uc.studentRepo.UpdateOne(ctx, id, student)
}

// Tested
func (uc *studentUseCase) DeleteOne(ctx context.Context, id int) error {
	return uc.studentRepo.DeleteOne(ctx, id)
}

func (uc *studentUseCase) Search(ctx context.Context, Value string) ([]admin.Student, error) {
	return uc.studentRepo.Search(ctx, Value)
}

// Tested
func (uc *studentUseCase) CreateStudent(ctx context.Context, student *admin.Student) error {
	return uc.studentRepo.CreateStudent(ctx, student)
}
