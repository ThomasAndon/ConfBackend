package view

import (
	com "ConfBackend/common"
	"ConfBackend/dto"
	"ConfBackend/hero"
	"ConfBackend/model"
	S "ConfBackend/services"
	"ConfBackend/task"
	"ConfBackend/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
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
	S.S.Logger.Infof("当前小车控制已经清空")
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
	S.S.Logger.Infof("小车控制请求 IN")

	// 如果当前已经有人在控制了，那么就不允许再有人控制了
	if !IsControlAvailable() {
		com.Error(ctx, "当前小车正在被他人控制")
		S.S.Logger.Infof("小车当前正在被他人控制")
		return
	}
	// handler the connection to websocket
	handler, err := upgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header())
	log.Println("接入ws车辆控制")
	S.S.Logger.Infof("接入ws车辆控制")
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

type loginBodyType struct {
	Token string `json:"token"`
}

func CCLogin(c *gin.Context) {
	// json has two fields: loginId and pw
	body, _ := io.ReadAll(c.Request.Body)
	loginId := gjson.Get(string(body), "loginId").String()
	pw := gjson.Get(string(body), "pw").String()
	// get ip addr of the client
	S.S.Logger.WithFields(logrus.Fields{
		"loginId": loginId,
		"pw":      pw,
	}).Info("login")

	members := make([]model.Member, 0)

	if loginId == "" || pw == "" {
		com.Error(c, "用户名或密码不能为空")
		return
	}

	S.S.Mysql.Where("login_id = ?", loginId).Find(&members)

	if len(members) == 0 {
		com.Error(c, "用户名不存在")
		S.S.Logger.Infof("用户名不存在")
		return
	}
	member := members[0]
	if member.Password != pw {
		com.Error(c, "密码错误")
		S.S.Logger.Infof("密码错误")
		return
	}
	retBody := loginBodyType{
		Token: member.UUID,
	}
	S.S.Logger.Infof("登录成功，返回uuid: %s", member.UUID)
	com.OkD(c, retBody)

	log.Println("member", members)

}

func LatestPcdLink(c *gin.Context) {
	// First get from the db the latest file
	latestUploadRecord := model.HeroPcdUoload{}
	res := S.S.Mysql.Order("id desc").Where("pcd_file_type = ?", "3d").First(&latestUploadRecord)
	if res.RowsAffected == 0 {
		S.S.Logger.Infof("没上传过PCD文件")
		com.Error(c, "现在还没有上传过pcd文件")
		return
	}
	fullLink := util.PadUrlLinkToPcdFile(latestUploadRecord.SavedFilename)
	S.S.Logger.Infof("返回最新的PCD文件链接: %s", fullLink)
	com.OkD(c, fullLink)

}

func LatestPcdLink2D(c *gin.Context) {
	// First get from the db the latest file
	latestUploadRecord := model.HeroPcdUoload{}
	res := S.S.Mysql.Order("id desc").Where("pcd_file_type = ?", "2d").First(&latestUploadRecord)
	if res.RowsAffected == 0 {
		S.S.Logger.Infof("没上传过PCD文件")
		com.Error(c, "现在还没有上传过pcd文件")
		return
	}
	fullLink := util.PadUrlLinkToPcdFile(latestUploadRecord.SavedFilename)
	S.S.Logger.Infof("返回最新的PCD文件链接: %s", fullLink)
	com.OkD(c, fullLink)
}

func MemLocation(c *gin.Context) {
	r := S.S.Redis
	b := make([]dto.PTermCalcedCoordDTO, 0)
	r.HGetAll(S.S.Context, util.GenCalcedPTermCoordKey()).Val()
	for _, v := range r.HGetAll(S.S.Context, util.GenCalcedPTermCoordKey()).Val() {
		temp := dto.PTermCalcedCoordDTO{}
		json.Unmarshal([]byte(v), &temp)
		b = append(b, temp)
	}
	com.OkD(c, b)

}

func NodeSensorStats(ctx *gin.Context) {
	b := make([]dto.SensorStatsDTO, 0)
	r := S.S.Redis

	for _, v := range r.HGetAll(S.S.Context, util.GenNodeStatsKey()).Val() {
		temp := dto.SensorStatsDTO{}
		json.Unmarshal([]byte(v), &temp)
		b = append(b, temp)
	}

	com.OkD(ctx, b)
}

func NodeCoords(ctx *gin.Context) {
	res := task.GetNodeCoord()
	ret := make([]dto.NodeCoordDTO, 0)
	for _, v := range res {
		ret = append(ret, v)
	}
	com.OkD(ctx, ret)
}

func MemLocationLinear(c *gin.Context) {
	r := S.S.Redis
	b := make([]dto.PTermCalcedCoordDTO, 0)
	r.HGetAll(S.S.Context, util.GenCalcedPTermCoordKey()).Val()
	for _, v := range r.HGetAll(S.S.Context, util.GenLinearCalcedPTermKey()).Val() {
		temp := dto.PTermCalcedCoordDTO{}
		json.Unmarshal([]byte(v), &temp)
		b = append(b, temp)
	}
	com.OkD(c, b)

}

func GetVibs(ctx *gin.Context) {
	m := S.S.Mysql
	res := make([]model.VibrationInfo, 0)
	/*	record := ctx.Query("time_in_min")
		var timeMinute int
		if record != "" {
			// record must be an int
			var err error
			timeMinute, err = strconv.Atoi(record)
			if err != nil {
				com.Error(ctx, "record 参数必须是整数")
				return
			}
			goto withRecord
		}*/

	// first group by created_at, then find all the records with this created_at
	m.Raw("select * from t_vibration_info " +
		"where detected_time " +
		"= (select detected_time " +
		"from t_vibration_info " +
		"group by detected_time " +
		"order by detected_time " +
		"desc limit 1) ").Scan(&res)

	/*withRecord:
	  // time of timeMinute minutes ago
	  timeMinuteAgo := time.Now().Add(-time.Minute * time.Duration(timeMinute))*/

	//m.Raw("select * from t_vibration_info " +
	//	"where detected_time >=").Scan(&res)
	com.OkD(ctx, res)
}
