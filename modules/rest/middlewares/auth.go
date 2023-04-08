package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

// just the basic example, for a challange, almost nothing similar with real production auth
func Auth(logger *zap.SugaredLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// token should be in headers
		token := c.GetReqHeaders()["Token"]
		if token == "" {
			logger.Debug("token not found")
			return c.SendStatus(http.StatusForbidden)
		}

		// parse provided token
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtSecret), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				logger.Debug("not valid token provided")
				return c.SendStatus(http.StatusUnauthorized)
			}
			logger.Debugf("failed process token [%s]", err.Error())
			return c.SendStatus(http.StatusBadRequest)
		}

		if !tkn.Valid {
			logger.Debug("invalid token")
			return c.SendStatus(http.StatusUnauthorized)
		}

		// set current user_id to context, to have possibility to check it in handlers
		c.Locals("user_id", claims.UserID)
		logger.Debugf("user [%d], reuests [%s]", claims.UserID, c.BaseURL())
		return c.Next()
	}
}
