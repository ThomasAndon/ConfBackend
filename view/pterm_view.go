package view

import (
	com "ConfBackend/common"
	"github.com/gin-gonic/gin"
)

// PTerm 废弃的接口
func PTerm(c *gin.Context) {
	com.OkM(c, "pong")
}

// FileReceived 废弃的接口
func FileReceived(c *gin.Context) {
	f, err := c.MultipartForm()
	if err != nil {
		com.Error(c, err.Error())
	}
	print(len(f.Value))
	com.Ok(c)

}
