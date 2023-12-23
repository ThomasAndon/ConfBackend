package main

import (
	"ConfBackend/chat"
	"ConfBackend/network"
	"ConfBackend/server"
	S "ConfBackend/services"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	S.InitServices()
	log.Println("init services")
}

// main 入口函数
func main() {
	// 单独的协程监听车的socket端口
	go server.StartListenHeroPort()

	// 单独的携程监听UWB的tcp发包端口
	go network.StartListenNetPktPort()

	// 启动聊天部分功能
	log.Println("init chat services")
	chat.InitChatServices()

	// 协程，启动所有其他的定时/周期或其它线程任务
	go server.StartTimed()

	// 设置gin的运行模式 调试/生产
	//gin.SetMode(gin.ReleaseMode)
	// web 服务器
	gin.DefaultWriter = S.S.MultipleWriter
	gin.SetMode(gin.ReleaseMode)
	server.StartApi()

}
