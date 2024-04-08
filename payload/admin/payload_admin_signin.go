package admin

import (
	"app/model/admin"
	"encoding/json"
	"log"
)

type GetAdminRequest struct {
	PassWord string `json:"-" db:"password, omitempty" validate:"required"`
	Email    string `json:"email,omitempty" db:"email, omitempty" validate:"required"`
}

func (c *GetAdminRequest) ToModel() *admin.ReqSignIn {
	admin := &admin.ReqSignIn{
		PassWord: c.PassWord,
		Email:    c.Email,
	}

	return admin
}

func (c *GetAdminRequest) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
