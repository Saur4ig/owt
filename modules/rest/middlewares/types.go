package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
)

// just to make it easier
const JwtSecret = "super-secret-key"

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}
