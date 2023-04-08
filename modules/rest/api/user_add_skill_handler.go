package api

import (
	"encoding/json"
	"net/http"

	"contacts_api/modules/model"
	"contacts_api/modules/rest/response"
	"github.com/gofiber/fiber/v2"
)

// AddUserSkill adds skill to user
// @Summary      Adds skill by name to a user
// @Tags         skills
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        skill  body     model.UserSkill     true  "user skill"
// @Param        token   header      string  true  "auth token"
// @Success      200
// @Failure      400
// @Failure      422
// @Failure      500
// @Router       /v1/users/{id}/skills [post]
func (h *Handler) AddUserSkill(c *fiber.Ctx) error {
	h.logger.Debugf("enter AddUserSkill [-]")

	userID, status, err := h.userValidation(c)
	if err != nil {
		return response.Failed(c, status, err.Error())
	}

	return h.addUserSkillHandler(c, userID)
}

func (h *Handler) addUserSkillHandler(c *fiber.Ctx, userID int) error {
	var skill model.UserSkill
	err := json.Unmarshal(c.Body(), &skill)
	if err != nil {
		h.logger.Warnf("cannot unmarshal user skill: [%s]", err.Error())
		return response.Failed(c, http.StatusBadRequest, "failed process skill")
	}

	// this required for checking: if skill already exists in db
	skills, err := h.skillsRepo.SkillsAsMap()
	if err != nil {
		h.logger.Debugf("failed fetch skills: [%s]", err.Error())
		return response.Failed(c, http.StatusInternalServerError, "failed fetch skills")
	}

	if err = skill.Validate(skills); err != nil {
		h.logger.Debugf("skill validation failed: [%+v]", skill)
		return response.Failed(c, http.StatusBadRequest, err.Error())
	}

	err = h.usersRepo.AddSkill(userID, skill)
	if err != nil {
		h.logger.Debugf("adding skill failed: [%s]", err.Error())
		return response.Failed(c, http.StatusInternalServerError, "adding skill failed")
	}

	defer h.logger.Debugf("exit AddUserSkill [=]")
	return response.Success(c, OkResponse{OK: true})
}
