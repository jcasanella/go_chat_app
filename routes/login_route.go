package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/chat_app/security"
	usecase "github.com/jcasanella/chat_app/usecase/user"
)

type LoginRoute struct {
	loginUseCase usecase.UserHandler
	jwtService   security.JWTService
}

func NewLoginRouteController(luc usecase.UserHandler, jwt security.JWTService) *LoginRoute {
	return &LoginRoute{luc, jwt}
}

func (lrc *LoginRoute) loginHandler(c *gin.Context) {
	var login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := lrc.loginUseCase.GetUser(c.Request.Context(), login.Username, login.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Generating token!!!")
	var token string
	token, err = lrc.jwtService.GenerateToken(u.Username, u.Password)
	if token != "" {
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, err.Error())
	}

	// p, err := password.GeneratePassword(login.Password)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
}

func (lrc *LoginRoute) LoginRoute(rg *gin.RouterGroup) {
	rg.POST("/login", lrc.loginHandler)
}
