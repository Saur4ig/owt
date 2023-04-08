package api

import (
	"net/http"

	"contacts_api/modules/model"
	"contacts_api/modules/rest/response"
	"github.com/gofiber/fiber/v2"
	_ "github.com/swaggo/fiber-swagger"
)

// GetUser returns single user data
// @Summary      Shows user
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  model.User
// @Failure      400
// @Failure      422
// @Router       /v1/users/{id} [get]
func (h *Handler) GetUser(c *fiber.Ctx) error {
	return h.getUserHandler(c)
}

func (h *Handler) getUserHandler(c *fiber.Ctx) error {
	h.logger.Debugf("enter CreateSkill [-]")
	userID, err := h.getUserID(c)
	if err != nil {
		return response.Failed(c, http.StatusUnprocessableEntity, "failed process user_id")
	}

	user, err := h.usersRepo.GetUser(userID)
	if err != nil {
		h.logger.Warnf("failed get user: [%s]", err.Error())
		return response.Failed(c, http.StatusBadRequest, "failed to get user")
	}

	defer h.logger.Debugf("exit CreateUser [%d]", userID)
	return response.Success(c, user)
}

// AllUsers returns all registered users
// @Summary      Shows all users
// @Tags         users
// @Produce      json
// @Success      200  {object}  []model.User
// @Failure      400
// @Router       /v1/users [get]
func (h *Handler) AllUsers(c *fiber.Ctx) error {
	return h.allUsersHandler(c)
}

type UsersResponse struct {
	Users []*model.User `json:"users"`
}

func (h *Handler) allUsersHandler(c *fiber.Ctx) error {
	h.logger.Debugf("enter AllUsers [-]")
	users, err := h.usersRepo.GetUsers()
	if err != nil {
		h.logger.Warnf("failed fetch users: [%s]", err.Error())
		return response.Failed(c, http.StatusBadRequest, "failed fetch users")
	}

	defer h.logger.Debugf("exit AllUsers [%+v]", users)
	return response.Success(c, UsersResponse{Users: users})
}
