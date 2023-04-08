package api

import (
	"encoding/json"
	"net/http"

	"contacts_api/modules/model"
	"contacts_api/modules/rest/response"
	"github.com/gofiber/fiber/v2"
)

// UpdateUserSkill updates skill
// @Summary      updates level of the skill for user
// @Tags         skills
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        name   path      string  true  "skill name"
// @Param        skill  body     model.UserSkill     true  "user skill"
// @Param        token   header      string  true  "auth token"
// @Success      200
// @Failure      400
// @Failure      422
// @Router       /v1/users/{id}/skills [put]
func (h *Handler) UpdateUserSkill(c *fiber.Ctx) error {
	h.logger.Debugf("enter UpdateUserSkill [-]")

	userID, status, err := h.userValidation(c)
	if err != nil {
		return response.Failed(c, status, err.Error())
	}

	return h.updateUserSkillHandler(c, userID)
}

func (h *Handler) updateUserSkillHandler(c *fiber.Ctx, userID int) error {
	var skill model.UserSkill
	err := json.Unmarshal(c.Body(), &skill)
	if err != nil {
		h.logger.Warnf("cannot unmarshal user skill: [%s]", err.Error())
		return response.Failed(c, http.StatusBadRequest, "failed process skill")
	}

	err = h.usersRepo.UpdateSkill(userID, skill)
	if err != nil {
		h.logger.Debugf("failed update user skill: [%s]", err.Error())
		return response.Failed(c, http.StatusInternalServerError, "failed update skill")
	}

	defer h.logger.Debugf("exit UpdateUserSkill [=]")
	return response.Success(c, OkResponse{OK: true})
}
