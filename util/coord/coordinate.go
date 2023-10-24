package coord

import (
	"ConfBackend/dto"
	S "ConfBackend/services"
	"ConfBackend/task"
	"ConfBackend/util"
	"encoding/json"
	"errors"
	"github.com/golang-module/carbon/v2"
	"github.com/sirupsen/logrus"
	"math"
)

// CalculateCoordinate 计算节点坐标
// b是一个map，key是termID，value是一个数组，数组中的元素是PTermDistanceDTO
// {NodeId: str, Distance: float64(mm)}
// 要求：conf定义的color｜nodeId和updateLocation api中的nodeId严格一致
func CalculateCoordinate(b map[string][]dto.PTermDistanceDTO, timeInUnixMilli float64) {
	// 获取已知节点的坐标
	nodeCoords := task.GetNodeCoord()

	// 检查已知节点个数是否符合要求
	if len(nodeCoords) < 4 {
		S.S.Logger.WithFields(logrus.Fields{
			"task":      "计算节点坐标时，已知节点个数不足。需等待小车返回足够的已知节点坐标。",
			"nodeCount": len(nodeCoords),
		}).Info()

		return
	}

	// 获取已知节点的ID slice
	knownNodeIds := []string{}
	for k, _ := range nodeCoords {
		knownNodeIds = append(knownNodeIds, k)
	}

	toSaveToRedis := make([]dto.PTermCalcedCoordDTO, 0)

	// 逐个处理每个人/pterm 的数据
	// ptermId, distanceDTOS 就是一个人的数据
	for ptermId, distanceDTOS := range b {
		// ptermId是人/pterm的ID, v是一个数组，数组中的元素是PTermDistanceDTO
		if len(distanceDTOS) < 4 {
			// 跳过不足4个距节点距离的pterm数据
			S.S.Logger.WithFields(logrus.Fields{
				"task":      "计算节点坐标时，pterm数据中有不足4个距节点距离的数据",
				"distCount": len(distanceDTOS),
				"ptermId":   ptermId,
			}).Info()
			continue
		}

		///// 检查已知的节点的坐标是否都已经已知（小车已回传）
		knownDistNodeIds := []string{}
		for _, dist := range distanceDTOS {
			knownDistNodeIds = append(knownDistNodeIds, dist.NodeId)
		}

		// nodeIdsInters 是已知坐标node和入参中已知距离pterm nodeId交集
		nodeIdsInters := util.Intersection(knownNodeIds, knownDistNodeIds)

		if len(nodeIdsInters) < 4 {
			// 跳过不足4个已知节点的pterm数据
			S.S.Logger.WithFields(logrus.Fields{
				"task":               "计算节点坐标时，pterm距离数据和已知节点数据的交集不足4个",
				"knownNodeCount":     len(knownNodeIds),
				"knownDistNodeCount": len(knownDistNodeIds),
			}).Info()
			continue
		}
		// todo 修改下面，需要根据intersection的结果获取已知，进行拼接
		// todo for _, dist := range nodeIdsInters { 之类的

		// 计算坐标，此处距离/1000，单位转 换为m

		nodeCoordArray := [][]float64{}
		distArrayInMeter := [][]float64{}
		for _, eachNodeIdInInters := range nodeIdsInters {
			for _, distDTO := range distanceDTOS {
				if distDTO.NodeId == eachNodeIdInInters {
					nodeCoordArray = append(nodeCoordArray,
						[]float64{nodeCoords[eachNodeIdInInters].X,
							nodeCoords[eachNodeIdInInters].Y,
							nodeCoords[eachNodeIdInInters].Z})
					distArrayInMeter = append(distArrayInMeter, []float64{distDTO.Distance / 1000})
				}
			}
		}

		// Slices padded. Then calculate the coords.
		coord, err := doCalcCoord(ptermId, nodeCoordArray, distArrayInMeter, timeInUnixMilli)
		if err != nil {

		} else {
			toSaveToRedis = append(toSaveToRedis, coord)
		}

	}

	// save toSaveToRedis to Redis
	{
		r := S.S.Redis
		pipe := r.Pipeline()
		rkey := util.GenCalcedPTermCoordKey()
		for _, coordDTO := range toSaveToRedis {
			hkey := coordDTO.PTermId
			hvalue, _ := json.Marshal(coordDTO)
			pipe.HSet(S.S.Context, rkey, hkey, string(hvalue))
		}
		_, err := pipe.Exec(S.S.Context)
		if err != nil {
			S.S.Logger.WithFields(logrus.Fields{
				"task": "计算节点坐标时，保存计算结果到Redis时出错",
			}).Error(err)
		}
	}

}

// doCalcCoord 计算坐标
// 传进来的坐标和距离必须是匹配好的，下标对应.
// 距离单位必须已经转成单位 米
func doCalcCoord(termId string, nodeCoordArray, distArrayInMeter [][]float64, timeInUnixMilli float64) (dto.PTermCalcedCoordDTO, error) {
	x, y, z, err := doCalcCoordMath(nodeCoordArray, distArrayInMeter)
	if err != nil {
		S.S.Logger.WithFields(logrus.Fields{
			"task":   "计算节点坐标时，计算出错-doCalcCoord-doCalcCoordMath",
			"termId": termId,
		}).Error(err)
		return dto.PTermCalcedCoordDTO{}, err
	}

	body := dto.PTermCalcedCoordDTO{
		PTermId:    termId,
		UpdateTime: carbon.DateTimeMilli{Carbon: carbon.CreateFromTimestampMilli(int64(timeInUnixMilli))},
		X:          x,
		Y:          y,
		Z:          z,
	}
	return body, nil

}

func CalculateLinearCoordinate(b map[string][]dto.PTermDistanceDTO, timeInUnixMilli float64) {
	// todo 线性上的距离
	// 只保证至少有两个已知节点
	// 直接计算沿着线的距离

	nodeCoords := task.GetNodeCoord()

	// 检查已知节点个数是否符合要求
	if len(nodeCoords) < 2 {
		S.S.Logger.WithFields(logrus.Fields{
			"task":      "计算节点坐标时，已知节点个数不足(<2个)。需等待小车返回足够的已知节点坐标。",
			"nodeCount": len(nodeCoords),
		}).Info()

		return
	}

	// 获取已知节点的ID slice
	knownNodeIds := []string{}
	for k, _ := range nodeCoords {
		knownNodeIds = append(knownNodeIds, k)
	}

	toSaveToRedis := make([]dto.PTermCalcedCoordDTO, 0)

	for ptermId, distanceDTOS := range b {
		// 如果此次已经算过此人则不算
		if CalcedLinearPTermId.Contains(ptermId) {
			continue
		}

		// ptermId 是一个人的id，distanceDTOs是一个数组，数组里面是每个人距离各个节点距离值

		///// 检查已知的节点的坐标是否都已经已知（小车已回传）
		knownDistNodeIds := []string{}
		for _, dist := range distanceDTOS {
			knownDistNodeIds = append(knownDistNodeIds, dist.NodeId)
		}

		// nodeIdsInters 是已知坐标node和入参中已知距离pterm nodeId交集
		nodeIdsInters := util.Intersection(knownNodeIds, knownDistNodeIds)

		if len(nodeIdsInters) < 2 {
			// 跳过不足4个已知节点的pterm数据
			S.S.Logger.WithFields(logrus.Fields{
				"task":               "计算线性中点节点坐标时，pterm距离数据和已知节点数据的交集不足2个",
				"knownNodeCount":     len(knownNodeIds),
				"knownDistNodeCount": len(knownDistNodeIds),
			}).Info()
			continue
		}

		// 找到dto中数值最小的两个节点
		min1 := math.Inf(1)
		min2 := math.Inf(1)

		var minCoord1, minCoord2 dto.NodeCoordDTO

		for _, intersNodeId := range nodeIdsInters {
			var dist dto.PTermDistanceDTO
			for _, distDTO := range distanceDTOS {
				if distDTO.NodeId == intersNodeId {
					dist = distDTO
				}
			}
			//if dist.Distance < min1 {
			//	min1 = dist.Distance
			//	minCoord1 = nodeCoords[dist.NodeId]
			//} else if dist.Distance < min2 {
			//	min2 = dist.Distance
			//	minCoord2 = nodeCoords[dist.NodeId]
			//}
			// 找最小的两个距离
			if dist.Distance < min1 {
				min2 = min1
				min1 = dist.Distance
				minCoord2 = minCoord1
				minCoord1 = nodeCoords[dist.NodeId]
			} else if dist.Distance < min2 {
				min2 = dist.Distance
				minCoord2 = nodeCoords[dist.NodeId]
			}

		}

		//fmt.Println(ptermId)

		// todo 使用min1， min2， minCoord1， minCoord2 计算中点坐标
		res, err := doLinearCalcCoord(ptermId, [][]float64{{minCoord1.X, minCoord1.Y, minCoord1.Z}, {minCoord2.X, minCoord2.Y, minCoord2.Z}}, [][]float64{{min1 / 1000}, {min2 / 1000}}, timeInUnixMilli)

		if err == nil {
			toSaveToRedis = append(toSaveToRedis, res)
		}

	}
	{
		r := S.S.Redis
		pipe := r.Pipeline()
		rkey := util.GenLinearCalcedPTermKey()
		for _, coordDTO := range toSaveToRedis {
			hkey := coordDTO.PTermId
			hvalue, _ := json.Marshal(coordDTO)
			pipe.HSet(S.S.Context, rkey, hkey, string(hvalue))
		}
		_, err := pipe.Exec(S.S.Context)
		if err != nil {
			S.S.Logger.WithFields(logrus.Fields{
				"task": "计算节点坐标时，保存计算结果到Redis时出错",
			}).Error(err)
		}
	}

}

func doLinearCalcCoord(termId string, nodeCoordArray, distArrayInMeter [][]float64, timeInUnixMilli float64) (dto.PTermCalcedCoordDTO, error) {

	if len(nodeCoordArray) != 2 {
		return dto.PTermCalcedCoordDTO{}, errors.New("nodeCoordArray length must be 2")
	}

	if len(distArrayInMeter) != 2 {
		return dto.PTermCalcedCoordDTO{}, errors.New("distArrayInMeter length must be 2")
	}

	// 计算两点之间的距离

	x, y, z, _ := CalcLinearCoord(nodeCoordArray, distArrayInMeter)
	body := dto.PTermCalcedCoordDTO{
		PTermId:    termId,
		UpdateTime: carbon.DateTimeMilli{Carbon: carbon.CreateFromTimestampMilli(int64(timeInUnixMilli))},
		X:          x,
		Y:          y,
		Z:          z,
	}

	return body, nil
}
