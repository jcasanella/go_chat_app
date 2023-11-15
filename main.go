package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/chat_app/config"
	"github.com/jcasanella/chat_app/database"
	repository "github.com/jcasanella/chat_app/repository/user"
	routes "github.com/jcasanella/chat_app/routes"
	usecase "github.com/jcasanella/chat_app/usecase/user"
)

var routeLogin *routes.LoginRoute
var routeIndex *routes.IndexRoute

func init() {
	fmt.Println("Reading config file...")
	cf := config.NewConfigValues()
	database.CreateConnection(cf)

	timeoutContext := time.Duration(5) * time.Second
	db := database.GetGORM()

	// Index
	routeIndex = routes.NewIndexRoute()

	// Login
	ur := repository.NewDBUserRepository(db)
	luc := usecase.NewUserUsecase(ur, timeoutContext)
	routeLogin = routes.NewLoginRouteController(luc)
}

func main() {
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

	// Static resources
	r.Static("/css", "./assets/css")
	r.Static("/js", "./assets/js")
	r.LoadHTMLFiles("views/index.html")

	// Index
	routeIndex.IndexRoute(r.Group("/"))

	// Login
	api := r.Group("/api")
	routeLogin.LoginRoute(api)

	r.Run()
}
