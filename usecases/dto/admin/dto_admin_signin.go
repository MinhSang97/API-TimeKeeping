package admin

import (
	"app/payload/admin"
)

type ReqSignIn struct {
	PassWord string `json:"-" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

func (c *ReqSignIn) ToPayload() *admin.GetAdminRequest {
	reqSignInPayload := &admin.GetAdminRequest{
		PassWord: c.PassWord,
		Email:    c.Email,
	}

	return reqSignInPayload
}
