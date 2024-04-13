package repo

import (
	"app/model/attendance_model"
	"context"
)

type AttendanceRepo interface {
	CheckIn(ctx context.Context, checkin *attendance_model.Attendance) error
	CheckOut(ctx context.Context, id int, checkout *attendance_model.Attendance) error
	History(ctx context.Context, history *attendance_model.Attendance) ([]attendance_model.Attendance, error)
	//GetAdmin(ctx context.Context, admin *admin_model.ReqSignIn) (*admin_model.ReqSignIn, error)
	//UpdateAdmin(ctx context.Context, user_id string, admin *admin_model.Admin) error
	//DeleteAdmin(ctx context.Context, user_id string) error
}
