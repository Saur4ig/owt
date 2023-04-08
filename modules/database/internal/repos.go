package internal

import (
	"database/sql"

	"contacts_api/modules/database/api"
)

type usersRepo struct {
	db *sql.DB
}

type skillsRepo struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) api.UsersRepository {
	return &usersRepo{
		db: db,
	}
}

func NewSkills(db *sql.DB) api.SkillsRepository {
	return &skillsRepo{
		db: db,
	}
}
