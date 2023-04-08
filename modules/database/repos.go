package database

import (
	"database/sql"

	"contacts_api/modules/database/api"
	"contacts_api/modules/database/internal"
)

func NewUsers(db *sql.DB) api.UsersRepository {
	return internal.NewUsers(db)
}

func NewSkills(db *sql.DB) api.SkillsRepository {
	return internal.NewSkills(db)
}
