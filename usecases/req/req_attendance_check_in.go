package req

import "time"

// ReqAttendaceCheckIn is the request structure for attendance check in
type ReqAttendaceCheckIn struct {
	UserId           string    `json:"user_id,omitempty" validate:"required"`
	Department       string    `json:"department,omitempty" validate:"required"`
	CheckInTime      time.Time `json:"check_in_time,omitempty" validate:"required"`
	LatitudeCheckIn  float64   `json:"latitude_check_in,omitempty" db:"latitude_check_in, omitempty" validate:"required"`
	LongitudeCheckIn float64   `json:"longitude_check_in,omitempty" db:"longitude_check_in, omitempty" validate:"required"`
}
