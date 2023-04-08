package api

import (
	"net/http"

	"contacts_api/modules/rest/response"
	"github.com/gofiber/fiber/v2"
)

// DeleteSkill deletes skill from set
// @Summary      deletes skill from pool and removes it for every user
// @Tags         skills
// @Produce      json
// @Param        id   path      string  true  "skill"
// @Success      200
// @Failure      500
// @Router       /v1/skills/{name} [delete]
func (h *Handler) DeleteSkill(c *fiber.Ctx) error {
	return h.deleteSkillHandler(c)
}

func (h *Handler) deleteSkillHandler(c *fiber.Ctx) error {
	h.logger.Debugf("enter DeleteSkill [-]")
	name := c.Params("name")

	err := h.skillsRepo.DeleteSkill(name)
	if err != nil {
		h.logger.Debugf("failed remvoe skill: [%s]", err.Error())
		return response.Failed(c, http.StatusInternalServerError, "failed delete skill")
	}
	defer h.logger.Debugf("exit DeleteSkill [=]")
	return response.Success(c, OkResponse{OK: true})
}
