package view

import (
	com "ConfBackend/common"
	"ConfBackend/model"
	S "ConfBackend/services"
	"ConfBackend/task"
	"ConfBackend/util"
	"github.com/gin-gonic/gin"
	"os"
)

type d struct {
	TotalPage int64 `json:"totalPage"`
	Result    any   `json:"result"`
}

func TestDb(c *gin.Context) {
	memberMgr := model.MemberMgr(S.S.Mysql)
	memberMgr.Omit("nickName")
	//q := c.Query("q")
	cur := c.Query("pg")
	size := c.Query("psize")
	page := model.Page{}
	// cur to int
	curNo := util.StringToInt64(cur)
	page.SetCurrent(curNo)
	page.SetSize(util.StringToInt64(size))
	//order := model.BuildDesc("id")
	//page.AddOrderItem(order)
	// ignore nickname
	Member, err := memberMgr.SelectPage(&page, memberMgr.WithNickname("a"))

	res := d{TotalPage: page.GetPages(), Result: Member.GetRecords()}

	if err != nil {
		return
	}
	com.OkD(c, res)
}

func TestAddNode(c *gin.Context) {
	// color from form
	color := c.PostForm("color")
	x := c.PostForm("x")
	y := c.PostForm("y")
	z := c.PostForm("z")
	// if any is "", return 400
	if color == "" || x == "" || y == "" || z == "" {
		com.Error(c, "must provide color, x, y, z")
	}

	xd := util.StringToFloat64(x)
	yd := util.StringToFloat64(y)
	zd := util.StringToFloat64(z)
	task.SetNodeCoord(color, xd, yd, zd)
}

func ConfView(c *gin.Context) {
	type ret struct {
		Conf S.AppConfig
		Env  string
	}
	con, err := os.ReadFile("/etc/tr.txt")
	if err != nil {
		com.ErrorD(c, "err", err)
	}
	com.OkD(c, ret{
		Conf: S.S.Conf,
		Env:  string(con),
	})
}
