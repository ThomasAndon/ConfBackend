package view

import (
	com "ConfBackend/common"
	"ConfBackend/model"
	"ConfBackend/network"
	S "ConfBackend/services"
	"ConfBackend/task"
	"ConfBackend/util"
	"fmt"
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
	task.SetNodeCoordV1(color, xd, yd, zd)
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

func SetByte(c *gin.Context) {
	b1 := []byte{0, 5}
	b2 := []byte{0, 1}
	b3 := []byte{0, 2}
	b4 := byte('a')
	b5 := byte(1)
	b6 := []byte{1, 2, 3, 4}
	b7 := []byte{'a', 'f', 'e'}
	b8 := []byte{0, 0, 0, 0, 0}
	// b9 contains 260 bytes
	b9 := []byte{1, 3, 5, 7, 9}
	//for i := 0; i < 260; i++ {
	//	b9 = append(b9, 1)
	//}

	// concat them all
	b := append(b1, b2...)
	b = append(b, b3...)
	b = append(b, b4)
	b = append(b, b5)
	b = append(b, b6...)
	b = append(b, b7...)
	b = append(b, b8...)
	b = append(b, b9...)

	network.HandleSinglePacket(b)
	// 可以成功设置字节流
}

func GetByte(c *gin.Context) {
	r := S.S.Redis
	bytes, _ := r.Get(c, "tr_:net_:1:97102101:1").Bytes()
	fmt.Println(bytes)
	// 可以成功获取字节流

}

func Set1(c *gin.Context) {
	b1 := []byte{0, 5}
	b2 := []byte{0, 1}
	b3 := []byte{0, 3}
	b4 := byte('a')
	b5 := byte(1)
	b6 := []byte{1, 2, 3, 4}
	b7 := []byte{'a', 'f', 'e'}
	b8 := []byte{0, 0, 0, 0, 0}
	// b9 contains 260 bytes
	b9 := []byte{1, 3, 5, 7, 9}
	//for i := 0; i < 260; i++ {
	//	b9 = append(b9, 1)
	//}

	// concat them all
	b := append(b1, b2...)
	b = append(b, b3...)
	b = append(b, b4)
	b = append(b, b5)
	b = append(b, b6...)
	b = append(b, b7...)
	b = append(b, b8...)
	b = append(b, b9...)

	network.HandleSinglePacket(b)
	// 可以成功设置字节流
}

func Set2(c *gin.Context) {
	b1 := []byte{0, 5}
	b2 := []byte{0, 3}
	b3 := []byte{0, 3}
	b4 := byte('a')
	b5 := byte(1)
	b6 := []byte{1, 2, 3, 4}
	b7 := []byte{'a', 'f', 'e'}
	b8 := []byte{0, 0, 0, 0, 0}
	// b9 contains 260 bytes
	b9 := []byte{2, 4, 6, 8, 10}
	//for i := 0; i < 260; i++ {
	//	b9 = append(b9, 1)
	//}

	// concat them all
	b := append(b1, b2...)
	b = append(b, b3...)
	b = append(b, b4)
	b = append(b, b5)
	b = append(b, b6...)
	b = append(b, b7...)
	b = append(b, b8...)
	b = append(b, b9...)

	network.HandleSinglePacket(b)
	// 可以成功设置字节流
}

func Set3(c *gin.Context) {
	b1 := []byte{0, 5}
	b2 := []byte{0, 2}
	b3 := []byte{0, 3}
	b4 := byte('a')
	b5 := byte(1)
	b6 := []byte{1, 2, 3, 4}
	b7 := []byte{'a', 'f', 'e'}
	b8 := []byte{0, 0, 0, 0, 0}
	// b9 contains 260 bytes
	b9 := []byte{111, 112, 113, 114, 115}
	//for i := 0; i < 260; i++ {
	//	b9 = append(b9, 1)
	//}

	// concat them all
	b := append(b1, b2...)
	b = append(b, b3...)
	b = append(b, b4)
	b = append(b, b5)
	b = append(b, b6...)
	b = append(b, b7...)
	b = append(b, b8...)
	b = append(b, b9...)

	network.HandleSinglePacket(b)
	// 可以成功设置字节流
}
