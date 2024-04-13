package attendance_payload

import (
	"app/model/attendance_model"
	"encoding/json"
	"log"
	"time"
)

type AddAttendaceRequest struct {
	ID                int64      `json:"id" db:"id"`
	UserId            string     `json:"user_id"  db:"user_id, omitempty"`
	Department        string     `json:"department,omitempty" db:"department, omitempty" validate:"required"`
	CheckInTime       *time.Time `json:"check_in_time,omitempty" db:"check_in_time, omitempty" validate:"required"`
	LatitudeCheckIn   float64    `json:"latitude_check_in,omitempty" db:"latitude_check_in, omitempty" validate:"required"`
	LongitudeCheckIn  float64    `json:"longitude_check_in,omitempty" db:"longitude_check_in, omitempty" validate:"required"`
	CheckOutTime      *time.Time `json:"check_out_time,omitempty" db:"check_out_time, omitempty" validate:"required"`
	LatitudeCheckOut  float64    `json:"latitude_check_out,omitempty" db:"latitude_check_out, omitempty" validate:"required"`
	LongitudeCheckOut float64    `json:"longitude_check_out,omitempty" db:"longitude_check_out, omitempty" validate:"required"`
	TotalHours        float64    `json:"total_hours,omitempty" db:"total_hours, omitempty" validate:"required"`
	OverTime          float64    `json:"over_time,omitempty" db:"over_time, omitempty" validate:"required"`
	Bonus             float64    `json:"bonus,omitempty" db:"bonus, omitempty" validate:"required"`
	Salary            float64    `json:"salary,omitempty" db:"salary, omitempty" validate:"required"`
	Location          string     `json:"location,omitempty" db:"location, omitempty" validate:"required"`
}

func (c *AddAttendaceRequest) ToModel() *attendance_model.Attendance {
	attendance := &attendance_model.Attendance{
		ID:                c.ID,
		UserId:            c.UserId,
		Department:        c.Department,
		CheckInTime:       c.CheckInTime,
		LatitudeCheckIn:   c.LatitudeCheckIn,
		LongitudeCheckIn:  c.LongitudeCheckIn,
		CheckOutTime:      c.CheckOutTime,
		LatitudeCheckOut:  c.LatitudeCheckOut,
		LongitudeCheckOut: c.LongitudeCheckOut,
		TotalHours:        c.TotalHours,
		OverTime:          c.OverTime,
		Bonus:             c.Bonus,
		Salary:            c.Salary,
		Location:          c.Location,
	}

	return attendance
}

func (c *AddAttendaceRequest) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
