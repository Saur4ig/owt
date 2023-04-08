package rest

import (
	"contacts_api/modules/rest/api"
	"contacts_api/modules/rest/middlewares"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
)

const (
	usersPrefix  = "/users"
	skillsPrefix = "/skills"
)

func routes(router fiber.Router, handler *api.Handler, logger *zap.SugaredLogger) {
	users := router.Group(usersPrefix)
	skills := router.Group(skillsPrefix)

	users.Get("/", handler.AllUsers)
	users.Get("/:id", handler.GetUser)
	users.Post("/", handler.CreateUser)
	users.Put("/:id", middlewares.Auth(logger), handler.UpdateUser)
	users.Delete("/:id", middlewares.Auth(logger), handler.DeleteUser)

	users.Post("/:id/skills", middlewares.Auth(logger), handler.AddUserSkill)
	users.Put("/:id/skills", middlewares.Auth(logger), handler.UpdateUserSkill)
	users.Delete("/:id/skills/:name", middlewares.Auth(logger), handler.DeleteUserSkill)

	skills.Get("/", handler.AllSkills)
	skills.Post("/", handler.CreateSkill)
	skills.Delete("/:name", handler.DeleteSkill)
}
