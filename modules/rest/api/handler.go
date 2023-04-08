package api

import (
	"contacts_api/modules/database/api"
	"go.uber.org/zap"
)

type Handler struct {
	usersRepo  api.UsersRepository
	skillsRepo api.SkillsRepository

	logger *zap.SugaredLogger
}

func New(ur api.UsersRepository, sr api.SkillsRepository, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		usersRepo:  ur,
		skillsRepo: sr,
		logger:     logger,
	}
}
