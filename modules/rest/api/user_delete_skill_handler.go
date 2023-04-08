package api

import (
	"net/http"

	"contacts_api/modules/rest/response"
	"github.com/gofiber/fiber/v2"
)

// DeleteUserSkill deletes user skill
// @Summary      deletes just a user skill
// @Tags         skills
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        name   path      string  true  "skill name"
// @Param        token   header      string  true  "auth token"
// @Success      200
// @Failure      422
// @Failure      500
// @Router       /v1/users/{id}/skills/{name} [delete]
func (h *Handler) DeleteUserSkill(c *fiber.Ctx) error {
	h.logger.Debugf("enter DeleteUserSkill [-]")

	userID, status, err := h.userValidation(c)
	if err != nil {
		return response.Failed(c, status, err.Error())
	}

	return h.deleteUserSkillHandler(c, userID)
}

func (h *Handler) deleteUserSkillHandler(c *fiber.Ctx, userID int) error {
	skillName := c.Params("name")

	err := h.usersRepo.DeleteSkill(userID, skillName)
	if err != nil {
		h.logger.Debugf("skill deletion failed: [%s]", err.Error())
		return response.Failed(c, http.StatusInternalServerError, err.Error())
	}
	defer h.logger.Debugf("exit DeleteUserSkill [=]")
	return response.Success(c, OkResponse{OK: true})
}
