package dto

import (
	"app/payload"
	"time"
)

type Admin struct {
	ID        int64     `json:"id"`
	UserId    string    `json:"userid"`
	Name      string    `json:"name"  validate:"required"`
	PassWord  string    `json:"-" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"-"`
}

func (c *Admin) ToPayload() *payload.AddAdminRequest {
	admintPayload := &payload.AddAdminRequest{
		ID:        c.ID,
		UserId:    c.UserId,
		Name:      c.Name,
		PassWord:  c.PassWord,
		Email:     c.Email,
		Role:      c.Role,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Token:     c.Token,
	}

	return admintPayload
}