package chat

import (
	"ConfBackend/model"
	"ConfBackend/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func SendOnlineMsg(msg model.ImMessage) error {
	if !IsUserOnline(msg.ToEntityUUID) {
		return fmt.Errorf("user %s is not online", msg.ToEntityUUID)
	}
	//put in that user's message queue
	msgStr, _ := json.Marshal(msg)
	WsConnectionManager.Users[msg.ToEntityUUID].Messages <- string(msgStr)
	return nil

}

// DispatchToSingleOnlineUser Dispatch 尝试将消息下发至一个用户，用户不在线则返回error，同时返回
// 所生成的msg对象
func DispatchToSingleOnlineUser(msg model.ImMessage) error {
	// todo 调用UWB发送

	// 下面的原有的WS发送逻辑不用管了
	util.PadSingleChatMsgFileUrl(&msg)
	err := SendOnlineMsg(msg)
	if err != nil {
		return errors.New("send online msg failed")

	}
	return nil
}

// GenNewMsg 生成新的消息对象, 传入的参数都是必须的
// 文件类型需要先处理再传入uri
func GenNewMsg(fromUUID string, msgType string, msgText string, fileTypeUri string, toEntityUUID string) model.ImMessage {
	newUuid := uuid.New()
	msg := model.ImMessage{
		UUID:         "mid-" + newUuid.String(),
		MsgType:      msgType,
		TextTypeText: msgText,
		FileTypeURI:  fileTypeUri,
		FromUserUUID: fromUUID,
		ToEntityUUID: toEntityUUID,
		CreatedAt:    time.Now(),
	}
	return msg
}
