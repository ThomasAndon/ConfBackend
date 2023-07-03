package dto

type SensorUpdateReqModel struct {
	NodeId int `json:"node"`
	// PacketId 参数没有用处。
	PacketId   int `json:"range"`
	SensorInfo struct {
		Light1 int `json:"l1"`
		Light2 int `json:"l2"`
		Voice1 int `json:"voice"`
	} `json:"sensor"`
}
