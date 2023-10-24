package server

import (
	"ConfBackend/dto"
	S "ConfBackend/services"
	"ConfBackend/util"
	"ConfBackend/util/coord"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	pkgRetrieveStride int64 = 1
	ctx                     = context.Background()
)

func StartUpdateLocationTask() {

	for {

		updateTask()
		coord.CalcedLinearPTermId.Clear()

		// 休眠一段时间，秒数由配置文件中的update_interval_in_second指定
		time.Sleep(time.Duration(S.S.Conf.Location.UpdateIntervalInSecond) * time.Second)
	}
}

func updateTask() {
	// 获取redis实例
	r := S.S.Redis
	lastUpdateTimestamp := r.Get(S.S.Context, util.GenLatestUpdatePackageTimeKey()).Val()
	if lastUpdateTimestamp == "" {
		r.Set(S.S.Context, util.GenLatestUpdatePackageTimeKey(), "0", 0)
		lastUpdateTimestamp = "0"
	}

	//p := S.S.TaskPool

	queryRedisTimeLogKey := util.GenPacketTimelogPrefix()

	var offset int64 = 0

	// 大for循环用来按时间顺序取包
	for {
		pkgTimes := r.ZRevRangeByScoreWithScores(S.S.Context, queryRedisTimeLogKey, &redis.ZRangeBy{
			Min:    lastUpdateTimestamp,
			Max:    "+inf",
			Offset: offset,
			Count:  pkgRetrieveStride,
		}).Val()

		// len = 0 means no data in redis
		if len(pkgTimes) == 0 {
			break
		}

		hasFoundValid := false

		for _, pkgItem := range pkgTimes {
			pkgNo := pkgItem.Member.(string)
			// float64 (time) UnixMilli()
			pkgTime := pkgItem.Score
			queryCountKey := util.GenDistanceQueryKey(pkgNo)

			// 对于同一个pkg，查看所有的节点上传的距离id
			pkgNodeList := r.Keys(S.S.Context, queryCountKey).Val()

			///////////
			// 新加入一个线性位置判断
			// 如果包的数量大于等于2，可以进行线性的判断

			if len(pkgNodeList) >= 2 {
				// 如果此包的时间等于最新的时间，说明已经计算过了，此包和以前的都不用算了，直接退出
				//fmt.Println("线性 V2 位置计算开始")
				// 开始计算线性位置
				updateLocation(pkgNodeList, pkgTime, true)
			}

			//////////
			if len(pkgNodeList) >= 4 {

				// 如果此包的时间等于最新的时间，说明已经计算过了，此包和以前的都不用算了，直接退出
				if pkgTime == util.StringToFloat64(lastUpdateTimestamp) {
					return
				}

				// 有足够的数据，计算位置
				// 找到了有效的数据，设置为true
				hasFoundValid = true
				r.Set(S.S.Context, util.GenLatestUpdatePackageTimeKey(), pkgTime, 0)
				//todo 计算位置
				S.S.Logger.WithFields(logrus.Fields{
					"task":         "找到可更新的数据，更新操作",
					"pkgNodeCount": len(pkgNodeList),
				}).Infof("最新包时间:%s", pkgTime)
				updateLocation(pkgNodeList, pkgTime, false)
				break
			}

		}

		if hasFoundValid {

			return
		} else {
			offset += pkgRetrieveStride
		}

	}

}

func updateLocation(pkgNodeList []string, timeInUnixMilli float64, isLinearCalc bool) {

	b, done := prepareData(pkgNodeList)

	if done {
		return
	}
	if !isLinearCalc {
		// is not linear calc，不是线性的计算，即普通计算，为第一版的计算方法，也放在这

		// 计算位置
		// b 是计算位置的数据，map[termid]PTermDistanceDTO
		// 每个id对应了一些距离点，如果点数小于某个设定值（如4）则不计算位置
		coord.CalculateCoordinate(b, timeInUnixMilli)
	} else {
		// is linear calc

		coord.CalculateLinearCoordinate(b, timeInUnixMilli)

	}

	// 新增：线性位置计算
	// 直接按比例计算中间点位置
	//go coord.CalculateLinearCoordinate(b, timeInUnixMilli)

}

func prepareData(pkgNodeList []string) (map[string][]dto.PTermDistanceDTO, bool) {
	r := S.S.Redis
	p := r.Pipeline()
	for _, nodeKey := range pkgNodeList {
		p.HGetAll(ctx, nodeKey)
	}
	res, err := p.Exec(S.S.Context)
	if err != nil {
		S.S.Logger.WithFields(logrus.Fields{
			"task":      "更新位置时，从redis中获取数据失败",
			"triedKeys": pkgNodeList,
		}).Error()
		return nil, true
	}

	b := make(map[string][]dto.PTermDistanceDTO)

	for _, cmd := range res {
		nodeNo := util.ParseNodeIdFromPktKey(cmd.(*redis.MapStringStringCmd).Args()[1].(string))
		nodeInfo := cmd.(*redis.MapStringStringCmd).Val()

		for k, v := range nodeInfo {
			// k is termid, v is distance in mm, if k not in b, add it in
			if _, ok := b[k]; !ok {
				b[k] = make([]dto.PTermDistanceDTO, 0)
			}
			b[k] = append(b[k], dto.PTermDistanceDTO{
				NodeId:   nodeNo,
				Distance: util.StringToFloat64(v),
			})

		}

	}
	return b, false
}
