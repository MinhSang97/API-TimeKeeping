package req

type ReqUpdateUser struct {
	FullName string `json:"full_name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required"`
}
