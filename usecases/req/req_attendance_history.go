package req

type ReqAttendaceHistory struct {
	UserId     string `json:"user_id,omitempty" validate:"required"`
	Department string `json:"department,omitempty" validate:"required"`
}
