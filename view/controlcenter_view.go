package view

import (
	com "ConfBackend/common"
	"ConfBackend/hero"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 同一时间只能有一个控制器，也就是下面这个实例, current controller 当前控制者
var curController *websocket.Conn

// SetCurController setter for curController
func SetCurController(conn *websocket.Conn) {
	curController = conn
}

// ClearCurController clear the current Controller
func ClearCurController() {
	err := curController.Close()
	if err != nil {
		log.Println("close the current controller error: ", err)
	}
	curController = nil
}

// IsControlAvailable 查看控制位置是否可用，如果不可用说明当前已经有人在控制
func IsControlAvailable() bool {
	return curController == nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HeroControl(ctx *gin.Context) {

	// 如果当前已经有人在控制了，那么就不允许再有人控制了
	if !IsControlAvailable() {
		com.Error(ctx, "当前小车正在被他人控制")
		return
	}
	// handler the connection to websocket
	handler, err := upgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header())
	curController = handler
	if err != nil {
		log.Println("handler error:", err)
	}
	defer func(handler *websocket.Conn) {
		handler.Close()
	}(handler)
	// read the message from the client
	/*	for {
		_, p, err := handler.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}
		log.Printf("recv: %s", p)
		// write the message back to the client
		server.HeroCommandStringChan <- string(p)
	}*/
	processControl()
	ClearCurController()

}

func processControl() {
	for {
		_, p, err := curController.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}
		// write the message back to the client
		//server.HeroCommandStringChan <- string(p)
		// send to channel
		if hero.IsCarConnected() {
			hero.CommandStringChan <- string(p)
		}

	}
}