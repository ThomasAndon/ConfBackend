package dto

import "github.com/golang-module/carbon/v2"

type UpdateLocationReq struct {
}

type PTermDistanceDTO struct {
	NodeId   string
	Distance float64
}

type NodeCoordDTO struct {
	NodeId     string          `json:"nodeId"`
	X          float64         `json:"x"`
	Y          float64         `json:"y"`
	Z          float64         `json:"z"`
	UpdateTime carbon.DateTime `json:"updateTime"`
}

type SensorStatsDTO struct {
	NodeId     string          `json:"node"`
	UpdateTime carbon.DateTime `json:"updateTime"`
	Light1     int             `json:"light1"`
	Light2     int             `json:"light2"`
	Voice1     int             `json:"voice1"`
}
