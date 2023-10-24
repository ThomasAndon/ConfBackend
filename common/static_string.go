package com

// 通用固定字符串
var (
	ProjectPref = "tr_"

	DistanceCachePrefix = "dc_"
	PacketsPrefix       = "pkts_"
	PacketPrefix        = "pkt_"
	NodePrefix          = "nd_"

	NodePositionPrefix      = "np_"
	PacketTimelogPrefix     = "pkt_tm_"
	LatestPackageTimePrefix = "latest_package_time_"
)

// IMB部分的一些固定字符串

var (
	IMStr = "im_"

	InboxStr = "inbox_"

	UserStr = "user_"

	MsgStr = "allMsgLst_"

	// ImInboxWholePrefix 是预拼接好的key，最后没有冒号，用的时候需要后面拼useruuid
	ImInboxWholePrefix = ProjectPref + ":" + IMStr + ":" + InboxStr + ":" + UserStr

	// MsgListStaticKey 是静态key，不用拼接什么东西，最后没有冒号
	MsgListStaticKey = ProjectPref + ":" + IMStr + ":" + MsgStr
)

var (
	MemberStr = "members_"

	MemberWholeKey = ProjectPref + ":" + MemberStr
)

var (
	// NicknameKey hash key
	NicknameKey = "nickname_"
)

var (
	// CoordPref 节点坐标部分
	CoordPref      = "coord_"
	NodeCoordsPref = "node-coords_"
)

var (
	PtermCalcedCoordPrefix = "pterm-calced-coord_"
)

var (
	LinearCalcedPTermPrefix = "linear-calced-pterm_"
)

// /// node sensor stats
var (
	SensorPrefix          = "sensor_"
	NodeSensorStatsPrefix = "node-sensor-stats_"
)
