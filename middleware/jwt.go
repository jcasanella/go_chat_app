package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jcasanella/chat_app/security"
)

// Middleware
func AuthorizeJWT(jwtService security.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.Query("token")
		token, err := jwtService.ValidateToken(jwtToken)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			c.Next()
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
