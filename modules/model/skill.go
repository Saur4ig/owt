package model

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type UserSkill struct {
	Name  string `json:"name" validate:"required"`
	Level int8   `json:"level" validate:"required,gte=0,lte=10"`
}

type Skill struct {
	Name string `json:"name" validate:"required"`
}

func (us *UserSkill) Validate(skills map[string]struct{}) error {
	validate := validator.New()
	err := validate.Struct(us)
	if err != nil {
		return err
	}

	if _, ok := skills[us.Name]; !ok {
		return errors.New("skill not found")
	}
	return nil
}

func (s *Skill) Validate(skills map[string]struct{}) error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return err
	}

	if _, ok := skills[s.Name]; ok {
		return errors.New("skill already exists")
	}
	return nil
}

func (s *Skill) String() string {
	return fmt.Sprintf(`'%s'`, s.Name)
}
