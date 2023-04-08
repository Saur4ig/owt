package api

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) userValidation(c *fiber.Ctx) (int, int, error) {
	userID, err := h.getUserID(c)
	if err != nil {
		return 0, http.StatusUnprocessableEntity, errors.New("failed process user_id")
	}

	// user_id should be placed to locals by middleware
	if c.Locals("user_id") != userID {
		h.logger.Errorf("User [%v] trying to update user [%d]", c.Locals("user_id"), userID)
		return 0, http.StatusForbidden, errors.New("can't update user")
	}
	return userID, 0, nil
}
