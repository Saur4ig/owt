package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
)

type User struct {
	ID      int         `json:"id"`
	Name    string      `json:"name" validate:"required"`
	Surname string      `json:"surname" validate:"required"`
	Address string      `json:"address" validate:"required"`
	Email   string      `json:"email" validate:"required,email"`
	Phone   string      `json:"phone" validate:"required,e164"`
	Skills  []UserSkill `json:"skills"`
}

func (u *User) String() string {
	return fmt.Sprintf(`'%s', '%s', '%s', '%s', '%s'`, u.Name, u.Surname, u.Address, u.Email, u.Phone)
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
