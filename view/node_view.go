package view

import (
	com "ConfBackend/common"
	"ConfBackend/dto"
	"ConfBackend/model"
	S "ConfBackend/services"
	_ "ConfBackend/services"
	"ConfBackend/task"
	"ConfBackend/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
)

type locationUpdateDTO struct {
	NodeId   int         `json:"node"`
	PacketId int         `json:"range"`
	Distance locInfoType `json:"distance"`
}

type locInfoType map[string]float64

// UpdateLocation 更新位置
func UpdateLocation(c *gin.Context) {
	// parse req body to locationUpdateDTO
	b := locationUpdateDTO{}
	err := json.NewDecoder(c.Request.Body).Decode(&b)
	if err != nil {
		S.S.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Errorf("parse request body failed")
		return
	}

	// 获取header中X-Node-Id值
	nodeId := util.IntToString(b.NodeId)

	// 获取header中X-Packet-Id值,转换成string
	packetId := util.IntToString(b.PacketId)

	locInfo := b.Distance
	S.S.Logger.WithFields(logrus.Fields{
		"nodeId":   nodeId,
		"packetId": packetId,
		"locInfo":  locInfo,
	}).Infof("update location")
	setToRedis(nodeId, packetId, locInfo)

}

// setToRedis 将距离信息存入redis
func setToRedis(nodeId, packetId string, info locInfoType) {
	r := S.S.Redis

	slice := make([]string, 2*len(info))
	i := 0
	for k, v := range info {
		slice[i] = k
		slice[i+1] = util.Float64ToString(v)
		i += 2
	}

	// 设置距离信息，key编号：dc_:pkt_{包编号}:nd_{节点编号}
	// hash key: 节点编号，hash value: 距离

	p := r.Pipeline()
	p.HSet(S.S.Context, util.GenDistanceCacheKey(packetId, nodeId), slice)
	p.Expire(S.S.Context, util.GenDistanceCacheKey(packetId, nodeId), time.Duration(S.S.Conf.Location.PktCacheTimeInMinute)*time.Minute)

	//r.HSet(c, util.GenDistanceCacheKey(packetId, nodeId), slice)
	// get current timestamp
	timestamp := time.Now().UnixMilli()
	b := redis.Z{
		Score:  float64(timestamp),
		Member: packetId,
	}

	// 在dc_下设置一个有序集合，记录首个该编号上传的时间。key编号：dc_:pkt_tm_
	// zset value: 包编号，score: 时间戳
	//r.ZAddNX(c, util.GenPacketTimelogPrefix(), b)
	p.ZAddNX(S.S.Context, util.GenPacketTimelogPrefix(), b)
	_, err := p.Exec(S.S.Context)
	if err != nil {
		return
	}

}

func SensorStats(c *gin.Context) {
	b := dto.SensorUpdateVO{}
	err := json.NewDecoder(c.Request.Body).Decode(&b)
	if err != nil {
		return
	}

	task.SetSensorStatsToRedis(util.IntToString(b.NodeId), b.SensorInfo.Light1, b.SensorInfo.Light2, b.SensorInfo.Voice1)

}

func GetAllContacts(c *gin.Context) {

	incomeUuid := c.GetHeader("X-User-UUID")
	S.S.Logger.WithFields(logrus.Fields{
		"incomeUuid": incomeUuid,
	}).Infof("GetAllContacts")
	usrs := make([]model.Member, 0)
	// S.S.Mysql find all usrs that are not deleted
	S.S.Mysql.Select("uuid", "nickname").Where("deleted_at is null AND uuid != ?", incomeUuid).Order("id").Find(&usrs)
	com.OkD(c, usrs)
}

func TermData(*gin.Context) {
	//r := S.S.Redis

}
