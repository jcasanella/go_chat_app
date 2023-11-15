package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/jcasanella/chat_app/model/user"
)

type LoginRoute struct {
	loginUseCase model.UserUsecase
}

func NewLoginRouteController(luc model.UserUsecase) *LoginRoute {
	return &LoginRoute{luc}
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

	// p, err := password.GeneratePassword(login.Password)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, u)
}

func (lrc *LoginRoute) LoginRoute(rg *gin.RouterGroup) {
	rg.POST("/login", lrc.loginHandler)
}
