package view

import (
	"ConfBackend/chat"
	com "ConfBackend/common"
	S "ConfBackend/services"
	"ConfBackend/task"
	"ConfBackend/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"io"
	"path/filepath"
)

type returnBody struct {
	NewMid     string `json:"newMid"`
	MsgType    string `json:"msgType"`
	MsgFileUrl string `json:"msgFileUrl"`
}

func SendMsg(c *gin.Context) {
	//multipart/form-data
	msgType := c.PostForm("msgType")
	msgText := c.PostForm("msgText")
	msgFile, _ := c.FormFile("msgFile")
	isToGroup := c.PostForm("isToGroup")
	toEntityUUID := c.PostForm("toEntityUUID")

	// get fromUserUuid from middleware
	fromUserUuid := c.GetString("uuid")
	if fromUserUuid == "" {
		com.Error(c, "X-User-UUID 不能为空")
		return
	}
	if !checkMsgTypeAllowed(msgType) {
		com.Error(c, "msgType不合法，允许"+allowedMsgTypeToStr())
		return
	}
	if toEntityUUID == "" {
		com.Error(c, "toEntityUUID 收信者uuid不能为空")
		return
	}
	//log.Println(msgType, msgText, msgFile, isToGroup, toEntityUUID)
	if isToGroup != "0" && isToGroup != "1" {
		com.Error(c, "isToGroup只能为0或1")
		return
	}

	// todo 先不开发群聊功能
	if isToGroup == "1" {
		com.Error(c, "群聊功能暂不开发")
		return
	}

	//////// Param Check passed below
	err := error(nil)

	if fromUserUuid == toEntityUUID {
		com.Error(c, "不支持自己给自己发信")
		return
	}

	// 0 is false  1 is true
	isToGroupBool := false
	if isToGroup == "1" {
		isToGroupBool = true
	}

	// msg fromUserUuid init
	msgUuid := ""
	switch msgType {
	case "text":
		{
			if msgText == "" {
				com.Error(c, "类型为text时，msgText不能为空")
				return
			}

			msgUuid, err = chat.IncomingHTTPTextMsg(fromUserUuid, msgText, isToGroupBool, toEntityUUID)
			if err != nil {
				com.Error(c, err.Error())
				return
			}
			ret := returnBody{
				NewMid:     msgUuid,
				MsgType:    msgType,
				MsgFileUrl: "",
			}
			com.OkD(c, ret)
			return

		}
	case "image":
		/*		{
				if msgFile.Size == 0 {
					com.Error(c, "类型为image时，msgFile不能为空")
					return
				}
				// save file
				fileType := filepath.Ext(msgFile.Filename)
				newFileName := uuid.New().String() + fileType
				newFileDir := filepath.Join(S.S.Conf.Chat.SaveStaticFileDirPrefix, newFileName)
				err := c.SaveUploadedFile(msgFile, newFileDir)
				if err != nil {
					com.Error(c, "文件保存失败")
					return
				}

				msgUuid, err = chat.IncomingHTTPFileMsg(fromUserUuid, msgType, isToGroupBool, toEntityUUID, newFileName, newFileDir)
				if err != nil {
					com.Error(c, err.Error())
					return
				}

			}*/
		fallthrough
	case "audio":
		{
			if msgFile.Size == 0 || msgFile == nil {
				com.Error(c, "类型为文件时，msgFile不能为空")
				return
			}
			// save file
			fileType := filepath.Ext(msgFile.Filename)
			newFileName := uuid.New().String() + fileType
			newFileDir := filepath.Join(S.S.Conf.Chat.SaveStaticFileDirPrefix, newFileName)
			err := c.SaveUploadedFile(msgFile, newFileDir)
			if err != nil {
				com.Error(c, "文件保存失败")
				return
			}

			msgUuid, err = chat.IncomingHTTPFileMsg(fromUserUuid, msgType, isToGroupBool, toEntityUUID, newFileName, newFileDir)
			if err != nil {
				com.Error(c, err.Error())
				return
			}

			ret := returnBody{
				NewMid:     msgUuid,
				MsgType:    msgType,
				MsgFileUrl: util.ConcatFullFileUrl(newFileName),
			}

			com.OkD(c, ret)
			return

		}
	}

}

func AllOnline(c *gin.Context) {
	com.OkD(c, chat.GetAllOnlineUsers())
}

func ChatHistory(c *gin.Context) {
	fromUuid := c.GetString("uuid")
	sinceMid := c.PostForm("sinceMid")
	isGroupMsg := c.PostForm("isGroupMsg")
	objectEntityUuid := c.PostForm("objectEntityUUID")

	if fromUuid == "" {
		com.Error(c, "X-User-UUID 禁止为空")
		return
	}
	if isGroupMsg != "0" && isGroupMsg != "1" {
		com.Error(c, "isGroupMsg只能为0或1")
		return
	}
	if objectEntityUuid == "" {
		com.Error(c, "objectEntityUuid不能为空")
		return
	}

	if isGroupMsg == "1" {
		com.Error(c, "群聊功能暂不开发")
		return
	} else {
		// 私聊
		com.OkD(c, chat.GetChatHistory(fromUuid, objectEntityUuid, sinceMid))
	}

}

type batchNicknameRetBody struct {
	Uuid     string      `json:"uuid"`
	Nickname interface{} `json:"nickname"`
}

func GetBatchNicknames(c *gin.Context) {
	// get c req body content
	b := c.Request.Body
	body, err := io.ReadAll(b)
	if err != nil {
		return
	}
	strArr := gjson.Get(string(body), "uuids").Array()
	strs := make([]string, len(strArr))
	// convert strArr to an array of str
	for i, v := range strArr {
		strs[i] = v.String()
	}

	res := task.GetNickNames(strs)

	ret := make([]batchNicknameRetBody, 0)
	for i, v := range res {
		ret = append(ret, batchNicknameRetBody{
			Uuid:     i,
			Nickname: v,
		})
	}

	//batchNicknameRetBody := make([]batchNicknameRetBody, len(res))
	//for i, v := range res {
	//	// uuid is key of map res
	//	batchNicknameRetBody[i] = batchNicknameRetBody{
	//    Uuid:     i,
	//    Nickname: v,
	//  }
	//}
	com.OkD(c, ret)

}
