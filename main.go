package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jcasanella/chat_app/config"
	"github.com/jcasanella/chat_app/db"
	repository "github.com/jcasanella/chat_app/repository/user"
	routes "github.com/jcasanella/chat_app/routes"
	"github.com/jcasanella/chat_app/security"
	usecase "github.com/jcasanella/chat_app/usecase/user"
)

var routeLogin *routes.LoginRoute
var routeIndex *routes.IndexRoute

var routeChat *routes.ChatRoute

func init() {
	fmt.Println("Reading config file...")
	cf := config.NewConfigValues()
	db.CreateConnection(cf)

	timeoutContext := time.Duration(5) * time.Second
	db := db.GetGORM()

	// Index
	routeIndex = routes.NewIndexRoute()

	// Login
	u := repository.NewDBUserRepository(db)
	s := usecase.NewUserService(u, timeoutContext)
	jwt := security.NewJWTService()
	routeLogin = routes.NewLoginRouteController(s, jwt)

	// Chat & jwt validation
	wu := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	routeChat = routes.NewChatRouteController(wu, jwt)
}

func main() {
	// Prepare to capture SigInt
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Closing DB Connection...")
		err := db.GetDb().Close()
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
	r.Static("/img", "./assets/img")
	r.LoadHTMLFiles("views/index.html")

	// Index
	routeIndex.IndexRoute(r.Group("/"))

	// Login
	api := r.Group("/api")
	routeLogin.LoginRoute(api)

	// Chat App
	routeChat.ChatRoute(api)

	r.Run()
}
