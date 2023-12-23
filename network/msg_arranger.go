package network

import (
	"ConfBackend/util"
	"math"
	"math/rand"
	"time"
)

const (
	PACKET_HEADER_LENGTH = 20
	MAX_CONTENT_LENGTH   = 500
)

func GenerateMessageChunks(messageType byte, allContent []byte, terminalID uint8) []*MsgBody {
	packNum := int(math.Ceil(float64(len(allContent)) / float64(MAX_CONTENT_LENGTH-PACKET_HEADER_LENGTH)))
	nonce := generate3CharNonce()

	msgBodies := make([]*MsgBody, packNum)

	for i := 0; i < packNum; i++ {
		var content []byte
		if i == packNum-1 {
			content = make([]byte, len(allContent)-i*(MAX_CONTENT_LENGTH-PACKET_HEADER_LENGTH))
		} else {
			content = make([]byte, MAX_CONTENT_LENGTH-PACKET_HEADER_LENGTH)
		}
		copy(content, allContent[i*(MAX_CONTENT_LENGTH-PACKET_HEADER_LENGTH):])

		thisPackNumByte := uint16(i + 1)
		sha2564digit := util.FourDigitSHA256(content)
		var sha2564 [4]byte
		copy(sha2564[:], sha2564digit)

		temp := &MsgBody{
			Length:          uint16(len(content)),
			PacketSerialNum: thisPackNumByte,
			TotalPackets:    uint16(packNum),
			MessageType:     messageType,
			TerminalID:      terminalID,
			PayloadSHA256:   sha2564,
			Nonce:           nonce,
			Payload:         content,
		}

		msgBodies[i] = temp
	}

	return msgBodies
}

func generate3CharNonce() [3]byte {
	var nonce [3]byte

	// 使用当前时间的纳秒部分设置随机数种子，以确保每次运行都有不同的种子
	rand.Seed(time.Now().UnixNano())

	// 生成每个字节的随机数，范围是0到127
	for i := 0; i < 3; i++ {
		nonce[i] = byte(rand.Intn(128))
	}

	return nonce
}
