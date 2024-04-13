package usecases

import (
	"app/dbutil"
	"app/model/attendance_model"
	"app/repo"
	"app/repo/mysql"
	"context"
)

type attendanceUseCase struct {
	attendanceRepo repo.AttendanceRepo
}

func NewAttendanceUseCase() AttendanceUsecase {
	db := dbutil.ConnectDB()
	attendanceRepo := mysql.NewAttendanceRepository(db)
	return &attendanceUseCase{
		attendanceRepo: attendanceRepo,
	}
}

func (uc *attendanceUseCase) CheckIn(ctx context.Context, checkin *attendance_model.Attendance) error {
	return uc.attendanceRepo.CheckIn(ctx, checkin)
}

func (uc *attendanceUseCase) CheckOut(ctx context.Context, id int, checkout *attendance_model.Attendance) error {
	return uc.attendanceRepo.CheckOut(ctx, id, checkout)
}

func (uc *attendanceUseCase) History(ctx context.Context, history *attendance_model.Attendance) ([]attendance_model.Attendance, error) {
	return uc.attendanceRepo.History(ctx, history)
}
