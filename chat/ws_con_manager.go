package chat

import (
	"ConfBackend/chat/unread"
	"ConfBackend/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// _User represents a connected user.
type _User struct {
	UUID     string
	Conn     *websocket.Conn
	Messages chan string
}

// ConnectionManager handles WebSocket connections.
type ConnectionManager struct {
	Users      map[string]*_User
	Register   chan *_User
	Unregister chan *_User
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections
	},
}

// WebSocketHandler is the handler function for the WebSocket route.
func (c *ConnectionManager) WebSocketHandler(ctx *gin.Context) {
	log.Println("incoming ws connection")
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	uuidobj, found := ctx.Get("uuid")

	uuid := ""
	if !found {
		log.Println("No UUID provided.")

	} else {
		uuid = uuidobj.(string)
	}
	// todo check uuid validity

	/*	if uuid == "" {
		log.Println("No UUID provided.")
		conn.Close()
		return
	}*/

	user := &_User{
		UUID:     uuid,
		Conn:     conn,
		Messages: make(chan string),
	}

	c.Register <- user

	// Start goroutine to listen for incoming messages from the user's WebSocket connection
	go c.handleIncomingMessages(user)

	// Start goroutine to send messages to the user's WebSocket connection
	go c.handleOutgoingMessages(user)
}

// handleIncomingMessages listens for incoming messages from a user's WebSocket connection.
func (c *ConnectionManager) handleIncomingMessages(user *_User) {
	defer func() {
		c.Unregister <- user
		user.Conn.Close()
	}()

	for {
		_, msg, err := user.Conn.ReadMessage()
		if err != nil {
			break
		}

		// Process the received message as needed
		fmt.Printf("Received message from user %s: %s\n", user.UUID, string(msg))

	}
}

// handleOutgoingMessages sends messages to a user's WebSocket connection.
func (c *ConnectionManager) handleOutgoingMessages(user *_User) {
	for message := range user.Messages {
		err := user.Conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			break
		}
	}
}

var WsConnectionManager *ConnectionManager

// startUserManagement starts the user management goroutine.
func (c *ConnectionManager) startUserManagement() {

	for {
		select {
		case user := <-c.Register:
			{
				c.Users[user.UUID] = user
				log.Printf("_User %s connected.\n", user.UUID)

				// START 检查是否有未读消息，有的话发送给用户
				{
					// todo 检查是否有未读消息，有的话发送给用户
					unreadMsg := unread.GetUnreadMessage(user.UUID)
					util.PadChatMsgFileUrl(&unreadMsg)
					if len(unreadMsg) > 0 {
						for _, msg := range unreadMsg {
							user.Messages <- util.MarshalString(msg)
						}
					}
				} // END
			}
		case user := <-c.Unregister:
			{
				delete(c.Users, user.UUID)
				close(user.Messages)
				log.Printf("_User %s disconnected.\n", user.UUID)
			}

		}
	}
}

func IsUserOnline(uuid string) bool {
	_, ok := WsConnectionManager.Users[uuid]
	return ok
}

func GetAllOnlineUsers() []string {
	var users []string
	for k := range WsConnectionManager.Users {
		users = append(users, k)
	}
	return users
}
