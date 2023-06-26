package dto

type SensorUpdateReqModel struct {
	NodeId int `json:"node"`
	// PacketId 参数没有用处。
	PacketId   int `json:"range"`
	SensorInfo struct {
		Light1 int `json:"light1"`
		Light2 int `json:"light2"`
		Voice1 int `json:"voice1"`
	} `json:"sensor"`
}
