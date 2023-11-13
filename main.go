package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/chat_app/config"
	"github.com/jcasanella/chat_app/database"
	repository "github.com/jcasanella/chat_app/repository/user"
	usecase "github.com/jcasanella/chat_app/usecase/user"
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

	// Set Up Repository and UseCase
	//timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	timeoutContext := time.Duration(5) * time.Second
	db := database.GetGORM()
	ur := repository.NewDBUserRepository(db)
	uc := usecase.NewUserUsecase(ur, timeoutContext)
	u, err := uc.GetUser(c.Request.Context(), login.Username, login.Password)
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

var cf *config.ConfigValues

func init() {
	fmt.Println("Reading config file...")
	cf = config.NewConfigValues()
	database.CreateConnection(cf)
}

func main() {
	fmt.Printf("Postgres %s:%d/%s \n", cf.Host, cf.Port, cf.Database)

	// Prepare to capture SigInt
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Closing DB Connection...")
		err := database.GetDb().Close()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	// Starting gin
	r := gin.Default()

	r.Static("/css", "./assets/css")
	r.Static("/js", "./assets/js")
	r.LoadHTMLFiles("views/index.html")

	r.GET("/", indexHandler)
	r.POST("/login", loginHandler)

	r.Run()
}
