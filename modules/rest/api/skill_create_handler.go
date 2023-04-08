package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"contacts_api/modules/model"
	"contacts_api/modules/rest/response"
	"github.com/gofiber/fiber/v2"
)

// CreateSkill creates skills
// @Summary      creates all skills from body
// @Tags         skills
// @Accept       json
// @Produce      json
// @Param        skills  body     []model.Skill     true  "skills to create"
// @Success      200
// @Failure      400
// @Failure      422
// @Failure      500
// @Router       /v1/skills [post]
func (h *Handler) CreateSkill(c *fiber.Ctx) error {
	return h.createSkillHandler(c)
}

func (h *Handler) createSkillHandler(c *fiber.Ctx) error {
	h.logger.Debugf("enter CreateSkill [-]")

	var newSkills []model.Skill
	err := json.Unmarshal(c.Body(), &newSkills)
	if err != nil {
		h.logger.Warnf("cannot unmarshal skills: [%s]", err.Error())
		return response.Failed(c, http.StatusBadRequest, "failed get skills")
	}

	// get all skills as map, to make faster checks
	skills, err := h.skillsRepo.SkillsAsMap()
	if err != nil {
		h.logger.Debugf("failed fetch skills: [%s]", err.Error())
		return response.Failed(c, http.StatusInternalServerError, "failed fetch skills")
	}

	// check all requirements for new skills
	ok, skillsExists := validateNewSkills(newSkills, skills)
	if !ok {
		return response.Failed(c, http.StatusBadRequest, fmt.Sprintf("already exists: %s", sliceToString(skillsExists)))
	}

	err = h.skillsRepo.CreateSkills(newSkills)
	if err != nil {
		h.logger.Debugf("failed create skills: [%s]", err.Error())
		return response.Failed(c, http.StatusBadRequest, "failed create skills")
	}

	defer h.logger.Debugf("exit CreateSkill [=]")
	return response.Success(c, OkResponse{OK: true})
}

// validates every skill
func validateNewSkills(newSkills []model.Skill, skills map[string]struct{}) (bool, []string) {
	exists := make([]string, 0)

	for _, skill := range newSkills {
		if _, ok := skills[skill.Name]; ok {
			exists = append(exists, skill.Name)
		}
	}

	return len(exists) == 0, exists
}

// creates string of all skills, that fas failed: [skill1, skill22]
func sliceToString(strs []string) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for _, val := range strs {
		buffer.WriteString(val)
		buffer.WriteString(", ")
	}

	buffer.WriteString("]")
	return buffer.String()
}
