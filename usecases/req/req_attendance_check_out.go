package req

import "time"

type ReqAttendaceCheckOut struct {
	//ID                int64     `json:"user_id,omitempty" validate:"required"`
	CheckOutTime      time.Time `json:"check_out_time,omitempty" validate:"required"`
	LatitudeCheckOut  float64   `json:"latitude_check_out,omitempty" db:"latitude_check_out, omitempty" validate:"required"`
	LongitudeCheckOut float64   `json:"longitude_check_out,omitempty" db:"longitude_check_out, omitempty" validate:"required"`
	//TotalHours        float64 `json:"total_hours,omitempty" db:"total_hours, omitempty" validate:"required"`
	//OverTime          float64 `json:"over_time,omitempty" db:"over_time, omitempty" validate:"required"`
	//Bonus             float64 `json:"bonus,omitempty" db:"bonus, omitempty" validate:"required"`
	//Salary            float64 `json:"salary,omitempty" db:"salary, omitempty" validate:"required"`
	//Location          string     `json:"location,omitempty" db:"location, omitempty" validate:"required"`
}
