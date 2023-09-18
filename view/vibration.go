package view

import (
	com "ConfBackend/common"
	"ConfBackend/dto"
	"ConfBackend/model"
	S "ConfBackend/services"
	"ConfBackend/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func VibUpload(c *gin.Context) {
	up := new(dto.VibrationUploadVO)
	c.BindJSON(up)
	// a time.Time parsed with format 2023-08-23 12:34:45 with GMT+8
	detectedTime, _ := time.ParseInLocation("2006-01-02 15:04:05", up.RecordTime, time.Local)
	// rows is the number of the times that "|" appears in the string up.PointsData
	rows := strings.Count(up.PointsData, "|")

	stringInfo := strings.Split(up.PointsData, ",")

	if len(stringInfo) != rows {
		com.Error(c, "错误：逗号数+1 不等于 竖线数！")
		S.S.Logger.WithFields(logrus.Fields{
			"stringInfo": stringInfo,
			"rows":       rows,
		}).Error()
	}

	data := make([]model.VibrationInfo, rows)

	for i := 0; i < rows; i++ {
		sp := strings.Split(stringInfo[i], "|")
		data[i].DetectedTime = detectedTime
		data[i].CreatedAt = time.Now()
		data[i].Value = util.StringToFloat64(sp[0])
		data[i].LocationInMeter = util.StringToFloat64(sp[1])
	}

	// save data
	S.S.Mysql.Create(&data)

}
