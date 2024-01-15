package network

import (
	"ConfBackend/chat"
	S "ConfBackend/services"
	"ConfBackend/util"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type MsgBody struct {
	Length          uint16  // 0、1前两个字节，表示小包长度（仅载荷部分长度，0-65535）
	PacketSerialNum uint16  // 第2、3字节，表示包序号——1-65535）
	TotalPackets    uint16  // 第4、5字节，表示消息总共拆成了多少包（1-65535）
	MessageType     byte    // 第6个字节，消息类型（'a'为文本，'b'为语音）
	TerminalID      uint8   // 第7个字节，终端编号
	PayloadSHA256   [4]byte // 第8-11个字节，载荷部分的SHA256前4位
	Nonce           [3]byte // 第12-14个字节，随机三个字符
	Reserved        [5]byte // 第15-19个字节，保留
	Payload         []byte  // 第20位起到X-1位，载荷部分
}

func (a MsgBody) ToByteArray() []byte {
	res := make([]byte, 0)
	res = append(res, util.Uint16ToTwoBytes(a.Length)...)
	res = append(res, util.Uint16ToTwoBytes(a.PacketSerialNum)...)
	res = append(res, util.Uint16ToTwoBytes(a.TotalPackets)...)
	res = append(res, a.MessageType)
	res = append(res, a.TerminalID)
	res = append(res, a.PayloadSHA256[:]...)
	res = append(res, a.Nonce[:]...)
	res = append(res, a.Reserved[:]...)
	res = append(res, a.Payload[:]...)
	return res
}

func ParseMsgBody(data []byte) (*MsgBody, error) {
	if len(data) < 20 {
		return nil, fmt.Errorf("字节流长度不足20字节")
	}

	msg := &MsgBody{}

	msg.Length = binary.BigEndian.Uint16(data[0:2])
	msg.PacketSerialNum = binary.BigEndian.Uint16(data[2:4])
	msg.TotalPackets = binary.BigEndian.Uint16(data[4:6])
	msg.MessageType = data[6]
	msg.TerminalID = data[7]
	copy(msg.PayloadSHA256[:], data[8:12])
	copy(msg.Nonce[:], data[12:15])
	copy(msg.Reserved[:], data[15:20])

	payloadStartIndex := 20
	if len(data) < payloadStartIndex+int(msg.Length)-20 {
		return nil, fmt.Errorf("字节流长度不足 %d 字节", msg.Length)
	}

	msg.Payload = data[payloadStartIndex:]

	return msg, nil
}

func HandleSinglePacket(input []byte) {
	S.S.Logger.Infof("接收到一个请求上传，字节流长度为%d", len(input))
	b, _ := ParseMsgBody(input)
	r := S.S.Redis

	if b.TotalPackets == 1 {
		// todo handle single packet
		handleFullData(&b.Payload, b)
	} else {
		// todo handle more than one
		/*		if b.PacketSerialNum < b.TotalPackets {
					// handle not the last one
					// save to redis
					key := util.GenNetworkPktKey(int(b.TerminalID), b.Nonce, int(b.PacketSerialNum))
					r.Set(context.Background(), key, b.Payload, time.Duration(time.Minute*10))
					fmt.Println(key)

				} else {
					// todo handle the last one
				}*/
		isFull, keys := isFillingTheLast(int(b.PacketSerialNum), int(b.TotalPackets), int(b.TerminalID), b.Nonce)

		if !isFull {
			key := util.GenNetworkPktKey(int(b.TerminalID), b.Nonce, int(b.PacketSerialNum))
			r.Set(context.Background(), key, b.Payload, time.Duration(time.Minute*3))
		} else {
			// get previous saved packets and concat them with this last one
			// get all the packets
			// concat them
			mapp := getSavedByteDict(keys)
			// concat them
			fullData := concatAll(mapp, b)
			handleFullData(fullData, b)

		}

	}

}

func handleFullData(fullData *[]byte, b *MsgBody) {

	if b.MessageType == 'a' {
		S.S.Logger.Infof("开始处理%d号终端发来的文字消息，byte数%d", b.TerminalID, len(*fullData))
		chat.IncomingHTTPTextMsg(strconv.Itoa(int(b.TerminalID)), string(*fullData), false, "0")
	}
	if b.MessageType == 'b' {
		// todo handle voice
		S.S.Logger.Infof("开始处理%d号终端发来的语音消息", b.TerminalID)
		fileType := ".aac"
		newFileName := uuid.New().String() + fileType
		newFileDir := filepath.Join(S.S.Conf.Chat.SaveStaticFileDirPrefix, newFileName)
		saveByteFile(*fullData, newFileDir)
		chat.IncomingHTTPFileMsg(strconv.Itoa(int(b.TerminalID)), "audio", false, "0", newFileName, newFileDir)

	}
}

func saveByteFile(data []byte, path string) {
	os.Mkdir(filepath.Dir(path), os.ModePerm)
	// write byte stream to file
	os.WriteFile(path, data, os.ModePerm)
	S.S.Logger.Infof("尝试写入完成，保存路径 %s，byte数为 %d", path, len(data))

}

func concatAll(prevSaved map[int][]byte, cur *MsgBody) *[]byte {
	mapp := prevSaved
	mapp[int(cur.PacketSerialNum)] = cur.Payload
	ret := []byte{}
	for i := 0; i < len(mapp); i++ {
		ret = append(ret, mapp[i+1]...)
	}
	return &ret
}

// 获取一个终端一次上传任务中已经存放的字节流（总数减1）
// 返回形式为map[完整key名]字节流
func getSavedByteDict(keys []string) map[int][]byte {
	r := S.S.Redis
	pipe := r.Pipeline()
	for _, v := range keys {
		pipe.Get(context.Background(), v)

	}
	res, err := pipe.Exec(context.Background())
	if err != nil {
		return map[int][]byte{}
	}
	ret := map[int][]byte{}
	for _, v := range res {
		//ret[keyname] = v.Byte()
		// split and get last
		splited := strings.Split(v.(*redis.StringCmd).Args()[1].(string), ":")
		// splited[len(splited)-1] to int
		index := util.StringToInt(splited[len(splited)-1])
		ret[index], _ = v.(*redis.StringCmd).Bytes()

	}
	fmt.Println(ret)
	return ret
}

func fetchRedisNetPacket(key string) []byte {
	r := S.S.Redis
	bytes, err := r.Get(context.Background(), key).Bytes()
	if err != nil {
		return []byte{}
	}
	return bytes
}

func isFillingTheLast(thisPacketSerialNum int, totalPacketsNum int, termId int, nonce [3]byte) (bool, []string) {
	// first get all keys
	keys := getAllPacketKeyFromSingleTask(termId, nonce)
	// split each key from ":",check get the last part, which is the serial number
	serialNums := splitAndGetLast(keys)
	// convert to int
	serialNumsInt := []int{}
	for _, v := range serialNums {
		serialNumsInt = append(serialNumsInt, util.StringToInt(v))
	}
	isSeq := isSequence(serialNumsInt, totalPacketsNum, thisPacketSerialNum)
	return isSeq, keys

}

func splitAndGetLast(inp []string) []string {
	ret := []string{}
	for _, v := range inp {
		splited := strings.Split(v, ":")
		ret = append(ret, splited[len(splited)-1])
	}
	return ret
}

func isSequence(arr []int, totalNum, currentNum int) bool {
	// 将currentNum添加到数组中
	arr = append(arr, currentNum)

	// 对数组进行排序
	sort.Ints(arr)
	if len(arr) != totalNum {
		return false
	}

	// 检查是否构成从1到totalNum的数组
	for i, num := range arr {
		if i+1 != num {
			return false
		}
	}

	return true
}

// return all the keys in redis that sets this packet
func getAllPacketKeyFromSingleTask(termId int, nonce [3]byte) []string {
	r := S.S.Redis
	queryKey := util.GenNetworkPktQueryKey(termId, nonce)
	keys, err := r.Keys(context.Background(), queryKey).Result()
	if err != nil {
		return []string{}
	}
	return keys
}
