package util

import com "ConfBackend/common"

func GenDistanceCacheKey(packetId, nodeId string) string {
	return com.ProjectPref + ":" + com.DistanceCachePrefix + ":" + com.PacketsPrefix + ":" + com.PacketPrefix + packetId + ":" + com.NodePrefix + nodeId
}

func GenDistanceQueryKey(packetId string) string {
	return com.ProjectPref + ":" + com.DistanceCachePrefix + ":" + com.PacketsPrefix + ":" + com.PacketPrefix + packetId + ":" + com.NodePrefix + "*"
}

// GenPacketTimelogPrefix 生成节点位置的key,即 tr_:dc_:pkt_tm_
// score float64 unixMilli
func GenPacketTimelogPrefix() string {
	return com.ProjectPref + ":" + com.DistanceCachePrefix + ":" + com.PacketTimelogPrefix
}

// GenLatestUpdatePackageTimeKey 生成最新更新的包的时间的key
// 为固定的"dc_:latest_package_time_"，值是一个时间戳
func GenLatestUpdatePackageTimeKey() string {
	return com.ProjectPref + ":" + com.DistanceCachePrefix + ":" + com.LatestPackageTimePrefix
}

// GenUserInboxKey 用于获取一个用户的inbox的key，输入user uuid，返回key
func GenUserInboxKey(useruuid string) string {
	return com.ImInboxWholePrefix + ":" + useruuid
}

// GenMsgListStaticKey 返回一个静态的消息列表list，即 tr_:im_:msg_
func GenMsgListStaticKey() string {
	return com.MsgListStaticKey
}

func GenMemberInfoNicknameKey(uuid string) string {
	return com.ProjectPref + ":" + com.MemberStr + ":" + com.NicknameKey + ":" + uuid
}

///// 节点坐标部分

// GenNodeCoordKey 返回tr_:dc_:coord_:node-coord_
// 是一个hash的key，用来存放所有已知节点的坐标
func GenNodeCoordKey() string {
	return com.ProjectPref + ":" + com.DistanceCachePrefix + ":" + com.CoordPref + ":" + com.NodeCoordsPref
}

// GenCalcedPTermCoordKey 返回tr_:dc_:pterm-calced-coord_
// 是一个hash，key是termid，value是一个json，是PTermCalcedCoordDTO结构体 json
func GenCalcedPTermCoordKey() string {
	return com.ProjectPref + ":" + com.DistanceCachePrefix + ":" + com.PtermCalcedCoordPrefix
}

func GenLinearCalcedPTermKey() string {
	return com.ProjectPref + ":" + com.DistanceCachePrefix + ":" + com.LinearCalcedPTermPrefix
}

///// node stats

// GenNodeStatsKey 返回tr_:sensor_:node-sensor-stats_
// 是一个hash，key是nodeid，value是一个json，是SensorStatsDTO结构体 json
func GenNodeStatsKey() string {
	return com.ProjectPref + ":" + com.SensorPrefix + ":" + com.NodeSensorStatsPrefix
}
