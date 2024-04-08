package admin_dto

import "app/payload/admin_payload"

type ReqSignIn struct {
	PassWord string `json:"-" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

func (c *ReqSignIn) ToPayload() *admin_payload.GetAdminRequest {
	reqSignInPayload := &admin_payload.GetAdminRequest{
		PassWord: c.PassWord,
		Email:    c.Email,
	}

	return reqSignInPayload
}
