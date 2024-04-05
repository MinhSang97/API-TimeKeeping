package payload

import (
	"app/model"
	"encoding/json"
	"log"
	"time"
)

type AddStudentRequest struct {
	FirstName    string    `json:"first_name" validate:"required"`
	LastName     string    `json:"last_name" validate:"required"`
	Age          int       `json:"age" validate:"required,gt=0"`
	Grade        float32   `json:"grade" validate:"gte=0,lte=10"`
	ClassName    string    `json:"class_name"`
	EntranceDate time.Time `json:"entrance_date" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	OldDueDate   time.Time
	NewDueDate   *time.Time
}

func (c *AddStudentRequest) ToModel() *model.Student {
	student := &model.Student{
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		Age:          c.Age,
		Grade:        c.Grade,
		ClassName:    c.ClassName,
		EntranceDate: c.EntranceDate,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}

	return student
}

func (c *AddStudentRequest) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}

type AddAdminRequest struct {
	ID        int64     `json:"id" db:"id" `
	UserId    string    `json:"userid"  db:"user_id, omitempty"`
	Name      string    `json:"name,omitempty" db:"name, omitempty"`
	PassWord  string    `json:"-" db:"password, omitempty"`
	Email     string    `json:"email,omitempty" db:"email, omitempty"`
	Role      string    `json:"role" db:"role, omitempty"`
	CreatedAt time.Time `json:"created_at" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at, omitempty"`
	Token     string    `json:"-" db:"token"`
}

func (c *AddAdminRequest) ToModel() *model.Admin {
	admin := &model.Admin{
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

	return admin
}

func (c *AddAdminRequest) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
