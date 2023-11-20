package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/chat_app/middleware"
	"github.com/jcasanella/chat_app/security"
)

type AuthRoute struct {
	jwtService security.JWTService
}

func NewAuthRouteController(jwt security.JWTService) *AuthRoute {
	return &AuthRoute{jwt}
}

func (a *AuthRoute) authHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Auth": "ok"})
}

func (a *AuthRoute) AuthRoute(rg *gin.RouterGroup) {
	rg.GET("/auth", middleware.AuthorizeJWT(a.jwtService), a.authHandler)
	// rg.GET("/auth", a.authHandler)
}
