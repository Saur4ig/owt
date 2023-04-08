package api

import (
	"encoding/json"
	"net/http"

	"contacts_api/modules/model"
	"contacts_api/modules/rest/response"
	"github.com/gofiber/fiber/v2"
)

// UpdateUser updates user
// @Summary      updates basic info about the user
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        token   header      string  true  "auth token"
// @Success      200
// @Failure      400
// @Failure      422
// @Router       /v1/users/{id} [put]
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	h.logger.Debugf("enter UpdateUser [-]")

	userID, status, err := h.userValidation(c)
	if err != nil {
		return response.Failed(c, status, err.Error())
	}

	return h.updateUserHandler(c, userID)
}

func (h *Handler) updateUserHandler(c *fiber.Ctx, userID int) error {
	var user model.User
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		h.logger.Warnf("cannot unmarshal user: [%s]", err.Error())
		return response.Failed(c, http.StatusBadRequest, "failed get user")
	}

	user.ID = userID

	err = h.usersRepo.UpdateUser(user)
	if err != nil {
		h.logger.Debugf("failed update user: [%s]", err.Error())
		return response.Failed(c, http.StatusInternalServerError, "update failed")
	}

	defer h.logger.Debugf("exit UpdateUser [%d]", userID)
	return response.Success(c, UserCreateResponse{ID: userID})
}
