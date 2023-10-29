package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	password "github.com/jcasanella/chat_app/crypto"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(c *gin.Context) {
	var login Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := password.GeneratePassword(login.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}

func main() {
	r := gin.Default()

	r.Static("/css", "./assets/css")
	r.Static("/js", "./assets/js")
	r.LoadHTMLFiles("views/index.html")

	r.GET("/", indexHandler)
	r.POST("/login", loginHandler)

	r.Run()
}
