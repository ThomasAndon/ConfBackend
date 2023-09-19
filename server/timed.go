package server

import (
	S "ConfBackend/services"
	"ConfBackend/util"
	"time"
)

func StartTimed() {

	// 定时更新位置
	go StartUpdateLocationTask()

	go CleanPktTimeLog()

}

// CleanPktTimeLog 清理redis中的包时间记录
func CleanPktTimeLog() {
	for {

		lastUpdateTimestamp := S.S.Redis.Get(S.S.Context, util.GenLatestUpdatePackageTimeKey()).Val()

		key := util.GenPacketTimelogPrefix()

		if lastUpdateTimestamp == "" || lastUpdateTimestamp == "0" {
			goto SLEEP
		}

		//根据Score删除zset，删除范围是所有算过的包都可以删除
		S.S.Redis.ZRemRangeByScore(S.S.Context, key, "-inf", util.Float64ToString(util.StringToFloat64(lastUpdateTimestamp)-float64(1)))
	SLEEP:
		time.Sleep(10 * time.Second)
	}

}
