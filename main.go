package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/chat_app/config"
	password "github.com/jcasanella/chat_app/crypto"
	"github.com/jcasanella/chat_app/database"
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

var cf *config.ConfigValues
var db *sql.DB

func init() {
	fmt.Println("Reading config file...")
	cf = config.NewConfigValues()
	database.CreateConnection(cf)
}

func main() {
	fmt.Printf("Postgres %s:%d/%s \n", cf.Host, cf.Port, cf.Database)

	r := gin.Default()

	r.Static("/css", "./assets/css")
	r.Static("/js", "./assets/js")
	r.LoadHTMLFiles("views/index.html")

	r.GET("/", indexHandler)
	r.POST("/login", loginHandler)

	r.Run()
}
