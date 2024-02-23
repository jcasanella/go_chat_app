package routes

import (
	"fmt"
	"time"

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
		fmt.Printf("%s received: %s\n", conn.RemoteAddr(), msg)

		// Test send conversation from Client
		now := time.Now()
		msgResponse := fmt.Sprintf(`<div class="outgoing-chats">
			<div class="outgoing-chats-img">
				<img src="/img/avatar.png">
			</div>
			<div class="outgoing-msg">
				<div class="outgoing-chats-msg">
					<p class="multi-msg">Hi this is the Response from Server.</p>
					<span class="time">%s</span>
				</div>
			</div>
		</div>`, now.Format("15:04 | Jan 2 Mon, 2006"))
		// Print msg to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), msgResponse)

		// Write msg back to the browser
		if err = conn.WriteMessage(msgType, []byte(msgResponse)); err != nil {
			return
		}
	}
}

func (cr *ChatRoute) ChatRoute(rg *gin.RouterGroup) {
	rg.GET("/chat", cr.chatHandler)
}
