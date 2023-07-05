package server

import (
	"ConfBackend/chat"
	com "ConfBackend/common"
	S "ConfBackend/services"
	"ConfBackend/task"
	"ConfBackend/view"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"strings"
)

func StartApi() {
	//ss := gin.Default()
	// set log writer

	s := gin.New()
	// a conditional logger
	logger := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			if strings.HasPrefix(c.Request.URL.Path, "/trweb") {
				return
			} else {
				gin.Logger()(c)
			}
		}
	}()
	s.Use(logger, gin.Recovery())
	s.Use(gzip.Gzip(gzip.DefaultCompression))
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
		node.POST("/sensors", view.SensorStats)
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
		car.POST("upload_2d", view.HeroUpload2D)
		car.GET("/ping", func(context *gin.Context) {
			com.OkD(context, "test new cicd in systemd ")
		})
	}

	// pad Com url
	{
		// The internal_model center's api
		cc := s.Group("/cc")
		cc.GET("/hero_control", view.HeroControl)
		cc.POST("/login", view.CCLogin)
		cc.GET("/latest_pcd_link", view.LatestPcdLink)
		cc.GET("/latest_pcd_link_2d", view.LatestPcdLink2D)
		cc.GET("/mem_location", view.MemLocation)
		cc.GET("/node_stats", view.NodeSensorStats)
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

	//join dir of current and static
	s.Static("pcd/static/", S.S.Conf.Pcd.SaveStaticFileDirPrefix)

	{
		test := s.Group("/test")
		test.POST("/db", view.TestDb)
		test.GET("/hasid", func(c *gin.Context) {
			id := c.Query("id")
			task.HaveValidUser(id)
		})
		test.POST("/add_node_coord", view.TestAddNode)
		test.GET("/check_config", view.ConfView)
		test.GET("/version", func(c *gin.Context) {
			c.String(200, "0704:1755")
		})

	}
	// frontend web
	s.Static("/trweb/", S.S.Conf.Web.DistFolderDir)
	// set release mode
	err := s.Run(":" + S.S.Conf.App.Port)
	if err != nil {
		log.Fatalln(err)
	}

}
