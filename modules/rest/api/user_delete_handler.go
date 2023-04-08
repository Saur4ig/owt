package api

import (
	"net/http"
	"strconv"

	"contacts_api/modules/rest/response"
	"github.com/gofiber/fiber/v2"
)

// DeleteUser deletes user
// @Summary      deletes user with all his skills
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "user id"
// @Param        token   header      string  true  "auth token"
// @Success      200
// @Failure      500
// @Router       /v1/users/{id} [delete]
func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	h.logger.Debugf("enter DeleteUser [-]")

	userID, status, err := h.userValidation(c)
	if err != nil {
		return response.Failed(c, status, err.Error())
	}

	return h.deleteUserHandler(c, userID)
}

type OkResponse struct {
	OK bool `json:"ok"`
}

func (h *Handler) getUserID(c *fiber.Ctx) (int, error) {
	idRaw := c.Params("id")

	userID, err := strconv.Atoi(idRaw)
	if err != nil {
		h.logger.Warnf("cannot parce user_id: [%s]", err.Error())
		return 0, err
	}

	return userID, nil
}

func (h *Handler) deleteUserHandler(c *fiber.Ctx, userID int) error {
	err := h.usersRepo.DeleteUser(userID)
	if err != nil {
		h.logger.Debugf("user deletion failed: [%s]", err.Error())
		return response.Failed(c, http.StatusBadRequest, "failed delete user")
	}

	defer h.logger.Debugf("exit DeleteUser [=]")
	return response.Success(c, OkResponse{OK: true})
}
