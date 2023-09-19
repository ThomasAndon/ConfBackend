package server

import (
	"ConfBackend/hero"
	S "ConfBackend/services"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"os"
)

func StartListenHeroPort() {
	heroPort := ":" + S.S.Conf.Car.Port
	tcpAddr, err := net.ResolveTCPAddr("tcp", heroPort)
	if err != nil {
		log.Fatalln("resolve err", err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("监听车辆端口错误：", err)
		os.Exit(5)
		log.Fatalln("监听车辆端口错误：", err)
	}
	log.Println("start listen hero port", heroPort)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			S.S.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorf("接受车辆连接错误，端口：%s", heroPort)
		}
		go hero.HandleConnection(conn)
	}

}
