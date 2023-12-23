package chat

import (
	"ConfBackend/chat/archive"
	"ConfBackend/model"
	S "ConfBackend/services"
	"ConfBackend/task"
	"ConfBackend/util"
	"errors"
	"os"
)

// IncomingHTTPTextMsg 接受通过HTTP方式发来的文本消息
func IncomingHTTPTextMsg(fromUUID string, msgText string, isToGroup bool, toEntityUUID string) (string, error) {

	if !isToGroup {
		//检查私聊合法性
		isToValidUser := task.HaveValidUser(toEntityUUID)
		if !isToValidUser {
			return "", errors.New("找不到你填写的收件用户：" + toEntityUUID)
		}

	} else {
		//检查群合法性

	}

	// 首先存档聊天
	msg := GenNewMsg(fromUUID, "text", msgText, "", toEntityUUID)

	go archive.ArchiveSingleMsg(msg)

	go func() {
		// 私聊，非群发
		if !isToGroup {
			// 尝试将消息下发至一个用户，用户不在线则返回error
			err := DispatchToSingleOnlineUser(msg)
			if err != nil {
				//todo 不在线，存入未读消息
				// 先不存redis了，不在线则不管了
				//go unread.SaveSingleUnreadMsg(msg)
			}
		} else {
			// 群发
			// 先不开发

		}
	}()

	return msg.UUID, nil
}

func IncomingHTTPFileMsg(fromUUID string, msgType string, isToGroup bool, toEntityUUID string, newFileName, newFileDir string) (string, error) {

	if !isToGroup {
		//检查私聊合法性
		isToValidUser := task.HaveValidUser(toEntityUUID)
		if !isToValidUser {
			// 删除文件
			os.Remove(newFileDir)
			return "", errors.New("找不到你填写的收件用户：" + toEntityUUID)
		}

	} else {
		//检查群合法性

	}

	msg := GenNewMsg(fromUUID, msgType, "", newFileName, toEntityUUID)

	go archive.ArchiveSingleMsg(msg)

	go func() {
		// 私聊，非群发
		if !isToGroup {
			err := DispatchToSingleOnlineUser(msg)
			if err != nil {
				//todo 不在线，存入未读消息
				// 先不存redis了，不在线则不管了
				//go unread.SaveSingleUnreadMsg(msg)
			}
		} else {
			// 群发
			// 先不开发
		}
	}()

	return msg.UUID, nil
}

func IncomingWebSocketTextMsg() {

}

// InitChatServices 聊天部分入口初始化函数
func InitChatServices() {

	// 启动聊天WS管理
	{
		WsConnectionManager = &ConnectionManager{
			Users:      make(map[string]*_User),
			Register:   make(chan *_User),
			Unregister: make(chan *_User),
		}

		go WsConnectionManager.startUserManagement()

	}

}

func GetChatHistory(fromUuid, objectEntityUuid, sinceMid string) []model.ImMessage {
	msg := make([]model.ImMessage, 0)

	if sinceMid != "" {
		S.S.Mysql.Raw(util.ChatHistorySqlBeforeMsg, fromUuid, objectEntityUuid, objectEntityUuid, fromUuid, sinceMid, 10).Scan(&msg)
	} else {
		S.S.Mysql.Raw(util.ChatHistorySqlLatest, fromUuid, objectEntityUuid, objectEntityUuid, fromUuid, 10).Scan(&msg)
	}

	util.PadChatMsgFileUrl(&msg)

	return msg
}
