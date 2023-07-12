package view

import (
	"ConfBackend/dto"
	"ConfBackend/model"
	S "ConfBackend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"time"
)

func HeroUpload(c *gin.Context) {

	upFile, _ := c.FormFile("file")
	if upFile == nil || upFile.Size == 0 {
		c.String(400, "文件file字段不能为空")
		S.S.Logger.Info("上传pcd文件, file is nil or size = 0")
		return
	}
	S.S.Logger.Info("HeroUpload, size = ", upFile.Size)

	// 尝试保存PCD文件
	fileType := filepath.Ext(upFile.Filename)
	t := time.Now()
	newFileName := t.Format("2006-01-02T15_04_05") + fileType
	newFileDir := filepath.Join(S.S.Conf.Car.SaveStatidPcdFileDirPrefix, newFileName)

	// 保存有误
	err := c.SaveUploadedFile(upFile, newFileDir)
	if err != nil {
		c.String(500, "文件保存失败")
		return
	}

	// archive the record
	go func() {
		rec := model.HeroPcdUoload{
			CreatedAt:                t,
			OriginalUploadedFilename: upFile.Filename,
			SavedFilename:            newFileName,
			FileSize:                 upFile.Size,
			SaveDuration:             int(time.Since(t).Seconds()),
			FileUUID:                 "pcd3d-" + uuid.New().String(),
			PcdFileType:              "3d",
		}
		S.S.Mysql.Create(&rec)

	}()

	c.String(200, "文件保存成功")

}

func HeroUpload2D(c *gin.Context) {

	upFile, _ := c.FormFile("file")
	if upFile == nil || upFile.Size == 0 {
		c.String(400, "文件file字段不能为空")
		S.S.Logger.Info("上传pcd文件, file is nil or size = 0")
		return
	}
	S.S.Logger.Info("HeroUpload, size = ", upFile.Size)

	// 尝试保存PCD文件
	fileType := filepath.Ext(upFile.Filename)
	t := time.Now()
	newFileName := t.Format("2006-01-02T15_04_05") + fileType
	newFileDir := filepath.Join(S.S.Conf.Car.SaveStatidPcdFileDirPrefix, newFileName)

	// 保存有误
	err := c.SaveUploadedFile(upFile, newFileDir)
	if err != nil {
		c.String(500, "文件保存失败")
		return
	}

	// archive the record
	go func() {
		rec := model.HeroPcdUoload{
			CreatedAt:                t,
			OriginalUploadedFilename: upFile.Filename,
			SavedFilename:            newFileName,
			FileSize:                 upFile.Size,
			SaveDuration:             int(time.Since(t).Seconds()),
			FileUUID:                 "pcd2d-" + uuid.New().String(),
			PcdFileType:              "2d",
		}
		S.S.Mysql.Create(&rec)

	}()

	c.String(200, "文件保存成功")

}

func SingleNodeCoord(c *gin.Context) {
	b := dto.HeroUploadNodeCoordVO{}

	// parse body to b
	err := c.ShouldBindJSON(&b)
	if err != nil {
		S.S.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("json解析错误")
	}

}
