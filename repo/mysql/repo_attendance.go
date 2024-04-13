package mysql

import (
	errors "app/error"
	"app/log"
	"app/model/attendance_model"
	"app/repo"
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type attendanceRepository struct {
	db *gorm.DB
}

func (s attendanceRepository) CheckIn(ctx context.Context, checkin *attendance_model.Attendance) error {
	attendance := checkin
	//attendance.CheckOutTime = nil // Set check out time to nil

	err := s.db.Table("Attendance").Create(attendance).Error
	//err := s.db.Table("Users").Where("email = ?", users.Email).First(users).Error
	if err != nil {
		//log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return errors.CheckInFail
		}
		log.Error(err.Error())
		return nil
	}
	return nil
}

func (s attendanceRepository) CheckOut(ctx context.Context, id int, checkout *attendance_model.Attendance) error {

	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	// Tạo thời điểm cần so sánh (2024-04-12 17:30:00)
	now := time.Now()
	checkOutTimeLimit := time.Date(now.Year(), now.Month(), now.Day(), 17, 30, 0, 0, loc)

	// Tạo biến để lưu trữ total_hours và over_time
	var totalHours float64
	var overTime float64

	// CheckInTime query from database
	var checkInTime time.Time
	err := s.db.Table("Attendance").Where("id = ?", id).Select("check_in_time").Scan(&checkInTime).Error
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	// Convert checkInTime and checkout.CheckOutTime to UTC
	checkInTime = checkInTime.In(loc)

	checkOutTimeLimit = checkOutTimeLimit.In(loc)

	// Tính total_hours và over_time dựa trên điều kiện của check_out_time
	if checkout.CheckOutTime.After(checkOutTimeLimit) {
		totalHours = checkout.CheckOutTime.Sub(checkInTime).Hours()
		overTime = checkout.CheckOutTime.Sub(checkOutTimeLimit).Hours()
	} else {
		totalHours = checkout.CheckOutTime.Sub(checkInTime).Hours()
		overTime = 0

	}

	attendance := checkout
	attendance.TotalHours = totalHours
	attendance.OverTime = overTime
	err = s.db.Table("Attendance").Where("id = ?", id).Updates(attendance).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.CheckOutFail
		}
		log.Error(err.Error())
		return nil
	}
	return nil
}

func (s attendanceRepository) History(ctx context.Context, history *attendance_model.Attendance) ([]attendance_model.Attendance, error) {

	fmt.Println(history.UserId)
	fmt.Println(history.Department)
	attendances := []attendance_model.Attendance{}
	err := s.db.Table("Attendance").Where("user_id = ? And department = ?", history.UserId, history.Department).
		Find(&attendances).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.HistoryFail
		}
		log.Error(err.Error())
		return nil, nil
	}
	return attendances, nil
}

var instance attendanceRepository

func NewAttendanceRepository(db *gorm.DB) repo.AttendanceRepo {
	if instance.db == nil {
		instance.db = db
	}
	return &instance
}
