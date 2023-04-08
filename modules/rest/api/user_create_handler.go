package api

import (
	"encoding/json"
	"net/http"

	"contacts_api/modules/model"
	"contacts_api/modules/rest/middlewares"
	"contacts_api/modules/rest/response"
	"github.com/golang-jwt/jwt/v5"

	_ "github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
)

// CreateUser creates a user
// @Summary      something like user registration
// @Tags         users
// @Produce      json
// @Param        user  body     model.User     true  "user to create"
// @Success      200
// @Failure      400
// @Failure      422
// @Failure      500
// @Router       /v1/users [post]
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	return h.createUserHandler(c)
}

type UserCreateResponse struct {
	ID    int    `json:"user_id"`
	Token string `json:"token"`
}

func (h *Handler) createUserHandler(c *fiber.Ctx) error {
	h.logger.Debugf("enter CreateUser [-]")

	var user model.User
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		h.logger.Warnf("cannot unmarshal body: [%s]", err.Error())
		return response.Failed(c, http.StatusBadRequest, "failed get body")
	}

	if err = user.Validate(); err != nil {
		h.logger.Debugf("user validation failed: [%+v]", user)
		return response.Failed(c, http.StatusBadRequest, err.Error())
	}

	var userID int
	userID, err = h.usersRepo.CreateUser(user)
	if err != nil {
		h.logger.Debugf("user creation failed: [%s]", err.Error())
		return response.Failed(c, http.StatusInternalServerError, err.Error())
	}

	claims := &middlewares.Claims{
		UserID: userID,
	}
	// token without expiring time, or any other protection
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(middlewares.JwtSecret))
	if err != nil {
		h.logger.Errorf("creating jwt token failed: [%s]", err.Error())
		return response.Failed(c, http.StatusInternalServerError, "Internal Server Error")
	}

	// this is just for simplicity and to avoid adding a lot of logic
	defer h.logger.Debugf("exit CreateUser [%d], token - [%s]", userID, tokenString)
	return response.Success(c, UserCreateResponse{ID: userID, Token: tokenString})
}
