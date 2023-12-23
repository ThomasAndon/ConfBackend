package network

import (
	"ConfBackend/model"
	S "ConfBackend/services"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
)

func SendUWBMsg(msg model.ImMessage) {
	toUserIdInt, err := strconv.Atoi(msg.ToEntityUUID)
	if err != nil {
		S.S.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("UWB发送消息时，转换toUserId到int失败，发送停止。toUserId: %s", msg.ToEntityUUID)
		return
	}
	var msgType byte
	var rawContent []byte

	if msg.MsgType == "text" {
		// todo 调用UWB发送
		msgType = 'a'
		// raw content is utf-8 encoded msg.TextTypeText
		rawContent = []byte(msg.TextTypeText)

	}

	if msg.MsgType == "audio" {
		// todo 调用UWB发送
		msgType = 'b'
		fl := filepath.Join(S.S.Conf.Chat.SaveStaticFileDirPrefix, msg.FileTypeURI)
		fullInBytes, _ := readFileGetBytes(fl)
		rawContent = fullInBytes

	}

	msgBodies := GenerateMessageChunks(msgType, rawContent, uint8(toUserIdInt))

	S.S.Logger.Infof("调用发送UWB消息，msgType: %c, toUserId: %d, raw Content length(byte): %d", msgType, toUserIdInt, len(rawContent))
	for _, msgBody := range msgBodies {
		NetPktChan <- msgBody.ToByteArray()
	}

}

func readFileGetBytes(fp string) ([]byte, error) {
	content, err := os.ReadFile(fp)
	if err != nil {
		S.S.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("UWB发TCP给板过程中打开文件失败，文件路径：%s", fp)
		return nil, err
	}
	return content, nil
}
