package task

import (
	"ConfBackend/dto"
	S "ConfBackend/services"
	"ConfBackend/util"
	"encoding/json"
	"github.com/golang-module/carbon/v2"
	"github.com/sirupsen/logrus"
	"strings"
)

// SetNodeCoordV1 弃用版本。使用SetNodeCoordV2
func SetNodeCoordV1(inferredColor string, x, y, z float64) {
	nodeInfo := S.S.Conf.Node.NodeInfo

	color2NodeId := map[string]string{}
	for _, node := range nodeInfo {
		res := strings.Split(node, "|")
		color2NodeId[strings.ToLower(res[0])] = res[1]
	}

	if !util.ContainKey(inferredColor, color2NodeId, false) {
		S.S.Logger.WithFields(logrus.Fields{
			"inferredColor": inferredColor,
			"x":             x,
			"y":             y,
			"z":             z,
		}).Errorf("操作失败：conf中未定义此颜色节点")
		return
	}

	targetNodeId := color2NodeId[strings.ToLower(inferredColor)]
	S.S.Logger.WithFields(logrus.Fields{
		"inferredColor": inferredColor,
		"x":             x,
		"y":             y,
		"z":             z,
	}).Infof("准备设置节点坐标")

	setNodeCoordToRedis(targetNodeId, x, y, z)

}

// SetNodeCoordV2 设置节点坐标。若Redis已存在第N个节点坐标（已设置坐标节点ID的最大值），则此时调用该函数就设置第N+1（最大ID+1）个节点坐标。
// 除非overwriteNodeId 大于0 （等于0或者小于0则不覆写），则直接设置第n个节点坐标，已存在就修改，不存在则创建。供调试方便。
// 同时，此阶段不管vi...和dist...两个参数，忽视它们。他们的作用是未来增强鲁棒性。
func SetNodeCoordV2(x, y, z, visualDistInMeter, distSinceLastInMeter float64, overwriteNodeId int) {
	targetNodeId := 0
	if overwriteNodeId > 0 {
		targetNodeId = overwriteNodeId
	} else {
		res := GetNodeCoord()
		// map[nodeId(string)]NodeCoordDTO find the largest nodeId, targetNodeId = largestNodeId + 1
		temp := make([]string, 0)
		for k, _ := range res {
			temp = append(temp, k)
		}
		targetNodeId = util.FindLargestNumberOfNodeId(temp) + 1
	}
	S.S.Logger.Infof("调用设置节点函数V2")

	setNodeCoordToRedis(util.IntToString(targetNodeId), x, y, z)

}

func setNodeCoordToRedis(nodeId string, x, y, z float64) {
	r := S.S.Redis

	// check if node coord already exists in redis
	{
		e := r.HExists(S.S.Context, util.GenNodeCoordKey(), nodeId).Val()
		if e {
			S.S.Logger.WithFields(logrus.Fields{
				"nodeId": nodeId,
				"x":      x,
				"y":      y,
				"z":      z,
			}).Warn("警告：尝试修改坐标的节点，节点坐标已存在。操作仍会继续。")

		}
	}
	rkey := util.GenNodeCoordKey()
	nodeBody := dto.NodeCoordDTO{
		NodeId:     nodeId,
		X:          x,
		Y:          y,
		Z:          z,
		UpdateTime: carbon.DateTime{carbon.Now()},
	}
	nbstr, err := json.Marshal(nodeBody)
	if err != nil {
		return
	}

	// set to redis
	r.HSet(S.S.Context, rkey, nodeId, nbstr)
	S.S.Logger.WithFields(logrus.Fields{
		"nodeId": nodeId,
		"x":      x,
		"y":      y,
		"z":      z,
	}).Infof("节点坐标设置成功")
}

// GetNodeCoord get node coord from redis
// return map[nodeId]nodeCoord
func GetNodeCoord() map[string]dto.NodeCoordDTO {
	r := S.S.Redis
	rkey := util.GenNodeCoordKey()

	res := r.HGetAll(S.S.Context, rkey).Val()

	nodeCoord := map[string]dto.NodeCoordDTO{}

	for _, v := range res {
		var nodeBody dto.NodeCoordDTO
		err := json.Unmarshal([]byte(v), &nodeBody)
		if err != nil {
			continue
		}
		nodeCoord[nodeBody.NodeId] = nodeBody
	}

	return nodeCoord
}
