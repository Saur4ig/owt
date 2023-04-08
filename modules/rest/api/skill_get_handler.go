package api

import (
	"net/http"

	"contacts_api/modules/model"
	"contacts_api/modules/rest/response"
	"github.com/gofiber/fiber/v2"
)

// AllSkills returns all available skills
// @Summary      Shows all skills
// @Tags         skills
// @Produce      json
// @Success      200  {object}  []model.Skill
// @Failure      500
// @Router       /v1/skills [get]
func (h *Handler) AllSkills(c *fiber.Ctx) error {
	return h.allSkillsHandler(c)
}

type SkillsResponse struct {
	Skills []model.Skill `json:"skills"`
}

func (h *Handler) allSkillsHandler(c *fiber.Ctx) error {
	h.logger.Debugf("enter AllSkills [-]")
	skills, err := h.skillsRepo.GetSkills()
	if err != nil {
		// todo: in real world, here, I should analyze error and depending on it use different statuses
		h.logger.Debugf("failed get skills: [%s]", err.Error())
		return response.Failed(c, http.StatusInternalServerError, "failed get skills")
	}

	defer h.logger.Debugf("exit AllSkills [%+v]", skills)
	return response.Success(c, SkillsResponse{Skills: skills})
}
