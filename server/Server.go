package server

import (
	"ConfBackend/chat"
	com "ConfBackend/common"
	S "ConfBackend/services"
	"ConfBackend/task"
	"ConfBackend/view"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func StartApi() {
	s := gin.Default()
	// set log writer

	s.Use(cors())
	//s.Use(printRequest)

	// preflight request if method is OPTIONS
	s.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// pad PTerm Services
	{
		// Human's terminal url group
		pt := s.Group("/pterm")
		pt.GET("/ping", view.PTerm)
		pt.POST("/file", view.FileReceived)
	}

	{
		// node api group
		node := s.Group("/node")
		node.POST("/update_location", view.UpdateLocation)
		node.POST("/sensor_stats", view.SensorStats)
		node.POST("/echo", func(context *gin.Context) {
			// echo whatever is in the request body
			content, _ := io.ReadAll(context.Request.Body)
			S.S.Logger.Infof("echo to node: %s", string(content))
			context.String(200, string(content))

		})
	}

	{
		// The car's api group
		car := s.Group("/hero")
		car.POST("/upload", view.HeroUpload)
		car.GET("/ping", func(context *gin.Context) {
			com.OkD(context, "Hello!!!!!-lab-server")
		})
	}

	// pad Com url
	{
		// The internal_model center's api
		cc := s.Group("/cc")
		cc.GET("/hero_control", view.HeroControl)
		cc.POST("/login", view.CCLogin)
		cc.GET("/latest_pcd_link", view.LatestPcdLink)
	}

	{
		// instant messaging im common api
		im := s.Group("/im")
		im.Use(MustHasUserUUID())
		im.POST("/sendmsg", view.SendMsg)
		im.GET("/ws", chat.WsConnectionManager.WebSocketHandler)
		im.GET("/all_online", view.AllOnline)
		im.POST("/chat_history", view.ChatHistory)
		im.POST("/get_batch_nicknames", view.GetBatchNicknames)
		im.GET("/get_all_contacts", view.GetAllContacts)

		// file system for /static/file
		s.Static("im/static/file", S.S.Conf.Chat.SaveStaticFileDirPrefix)
	}

	{
		test := s.Group("/test")
		test.POST("/db", view.TestDb)
		test.GET("/hasid", func(c *gin.Context) {
			id := c.Query("id")
			task.HaveValidUser(id)
		})

	}
	// set release mode
	err := s.Run(":" + S.S.Conf.App.Port)
	if err != nil {
		log.Fatalln(err)
	}

}
