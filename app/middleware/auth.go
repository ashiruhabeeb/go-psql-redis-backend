package middleware

import (
	"strings"

	"github.com/ashiruhabeeb/go-backend/pkg/config"
	"github.com/ashiruhabeeb/go-backend/pkg/response"
	"github.com/ashiruhabeeb/go-backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtService token.JWTService
}

func NewAuthMiddleware(jwtService token.JWTService) gin.HandlerFunc {
	return (&AuthMiddleware{jwtService: jwtService}).Handle
}

func(a *AuthMiddleware) Handle(c *gin.Context){
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		response.Error(c, 401, "[ERROR] missing authorization header")
		c.Abort()
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		response.Error(c, 401, "[ERROR] invalid header format")
		c.Abort()
		return
	}

	if headerParts[0] != "Bearer" {
		response.Error(c, 401, "[ERROR] token must contain Bearer prefix")
		c.Abort()
		return
	}

	user, err := a.jwtService.ValidateToken(headerParts[1], config.Cfg.AccessTokenPublicKey)
	if err != nil {
		response.Error(c, 401, "[ERROR] invalid toekn")
	}

	c.Set("user", user)
	c.Next()
}
