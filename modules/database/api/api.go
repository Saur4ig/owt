package api

import (
	"contacts_api/modules/model"
)

type UsersRepository interface {
	CreateUser(user model.User) (int, error)
	UpdateUser(user model.User) error
	DeleteUser(userID int) error
	GetUsers() ([]*model.User, error)
	GetUser(userID int) (*model.User, error)

	AddSkill(userID int, skill model.UserSkill) error
	UpdateSkill(userID int, skill model.UserSkill) error
	DeleteSkill(userID int, skillName string) error
}

type SkillsRepository interface {
	GetSkills() ([]model.Skill, error)
	SkillsAsMap() (map[string]struct{}, error)
	CreateSkills([]model.Skill) error
	DeleteSkill(name string) error
}
