package dto

type SensorUpdateVO struct {
	NodeId int `json:"node"`
	// PacketId 参数没有用处。
	PacketId   int `json:"range"`
	SensorInfo struct {
		Light1 int `json:"l1"`
		Light2 int `json:"l2"`
		Voice1 int `json:"voice"`
	} `json:"sensor"`
}

type HeroUploadNodeCoordVO struct {
	X                    float64 `json:"x"`
	Y                    float64 `json:"y"`
	Z                    float64 `json:"z"`
	VisualDistInMeter    float64 `json:"visualDistInMeter"`
	DistSinceLastInMeter float64 `json:"distSinceLastInMeter"`
	OverwriteNodeId      int     `json:"overwriteNodeId"`
}
