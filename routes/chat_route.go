package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatRoute struct {
	upgrader websocket.Upgrader
}

func NewChatRouteController(u websocket.Upgrader) *ChatRoute {
	return &ChatRoute{u}
}

func (cr *ChatRoute) chatHandler(c *gin.Context) {
	// Error ignored just for simplicity
	conn, err := cr.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print msg to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write msg back to the browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func (cr *ChatRoute) ChatRoute(rg *gin.RouterGroup) {
	rg.GET("/chat", cr.chatHandler)
}
