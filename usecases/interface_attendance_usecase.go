package usecases

import (
	"app/model/attendance_model"
	"context"
)

type AttendanceUsecase interface {
	CheckIn(ctx context.Context, checkin *attendance_model.Attendance) error
	CheckOut(ctx context.Context, id int, checkout *attendance_model.Attendance) error
	History(ctx context.Context, history *attendance_model.Attendance) ([]attendance_model.Attendance, error)
}
