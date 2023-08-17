package handler

import (
	"net/http"
	"os"

	"seacrust-backend/src/jwt"
	"seacrust-backend/src/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, code int64, message string, data interface{}) {
	c.JSON(int(code), HTTPResponse{
		Message:   message,
		IsSuccess: false,
		Data:      data,
	})
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorization := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			ErrorResponse(c, http.StatusUnauthorized, "UnauthorizedJWT1", nil)
			c.Abort()
			return
		}

		tokenJwt := authorization[7:]
		claims := models.UserClaims{}
		jwtKey := os.Getenv("SECRET_KEY")

		if err := jwt.DecodeToken(tokenJwt, &claims, jwtKey); err != nil {
			ErrorResponse(c, http.StatusUnauthorized, "UnauthorizedJWT2", nil)
			c.Abort()
			return
		}

		c.Set("user", claims)
	}
}

func JwtMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			ErrorResponse(c, http.StatusUnauthorized, "UnauthorizedJWT1Admin", nil)
			c.Abort()
			return
		}

		tokenJwt := authorization[7:]
		claims := models.AdminClaims{}
		jwtKey := os.Getenv("SECRET_KEY")

		if err := jwt.DecodeToken(tokenJwt, &claims, jwtKey); err != nil {
			ErrorResponse(c, http.StatusUnauthorized, "Unauthorized2", nil)
			c.Abort()
			return
		}

		c.Set("admin", claims)
	}
}
