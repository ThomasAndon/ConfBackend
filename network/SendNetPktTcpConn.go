package network

import (
	S "ConfBackend/services"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"os"
	"time"
)

// 规定的tcp发信时间间隔
var SendTCPNetPktIntervalInMilliSec = 500

// / 下面不修改
var lastSend time.Time

var NetPktChan = make(chan []byte, 200)

func StartListenNetPktPort() {
	netPktPort := ":8052"
	tcpAddr, err := net.ResolveTCPAddr("tcp", netPktPort)
	if err != nil {
		log.Fatalln("resolve err", err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("监听UWB TCP端口错误：", err)
		os.Exit(5)
		log.Fatalln("监听UWB TCP端口错误：", err)
	}
	log.Println("start listen net pkt port", netPktPort)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			S.S.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorf("接受UWB网络包连接错误，端口：%s", netPktPort)
		}
		S.S.Logger.Infof("已经建立8052的发包TCP连接")

		go handleNetPktCon(conn)
	}

}

func handleNetPktCon(conn net.Conn) {
	defer func() {
		conn.Close()
		S.S.Logger.Infof("UWB Network trans disconnected")
	}()

	for {
		select {
		case pkt := <-NetPktChan:
			// send to conn
			// if not long enough, wait
			if time.Now().Sub(lastSend) < time.Duration(SendTCPNetPktIntervalInMilliSec)*time.Millisecond {
				time.Sleep(time.Duration(SendTCPNetPktIntervalInMilliSec)*time.Millisecond - time.Now().Sub(lastSend))
			}
			_, err := conn.Write(pkt)
			lastSend = time.Now()
			if err != nil {
				S.S.Logger.WithFields(logrus.Fields{
					"err": err,
				}).Errorf("发送UWB网络包到板错误")
			}
		}
	}
}
