package hero

import (
	S "ConfBackend/services"
	"log"
	"net"
	"sync"
	"time"
)

var Info = heroInfoType{
	Connected: false,
}

func IsCarConnected() bool {
	return Info.Connected
}

type heroInfoType struct {
	IP        string
	Port      string
	Connected bool

	// 当前连接控制者
	CurController string
}

// 只运行一次的初始化车辆控制参数的函数
var oncer sync.Once

var CommandStringChan = make(chan string)

// SendCommandInterval 配置文件中读取发送命令间隔
var SendCommandInterval time.Duration

type heroCommand struct {
	Forward    bool
	Backward   bool
	Left       bool
	Right      bool
	SpeedLevel string
}

func (hc *heroCommand) ValidateDirection() {
	if hc.Forward && hc.Backward {
		hc.Forward = false
		hc.Backward = false
	}
	if hc.Left && hc.Right {
		hc.Left = false
		hc.Right = false
	}

}

// ToCommandString 将HeroCommand转换为字符串.
// 格式是"--------"前4位分别代表前后左右，后四位保留。如果前4位中某一位是1，则表示激活，为"-"则不激活。
func (hc *heroCommand) ToCommandString() string {
	var s string
	if hc.Forward {
		s += "1"
	} else {
		s += "-"
	}
	if hc.Backward {
		s += "1"
	} else {
		s += "-"
	}
	if hc.Left {
		s += "1"
	} else {
		s += "-"
	}
	if hc.Right {
		s += "1"
	} else {
		s += "-"
	}
	s += hc.SpeedLevel
	s += "---"
	return s
}

// StringToSafeHeroCommand 将字符串解析为HeroCommand，没有检查合法性。
// 格式是"--------"前4位分别代表前后左右，后四位保留。如果前4位中某一位是1，则表示激活。
func StringToSafeHeroCommand(s string) (hc heroCommand) {
	if s[0] == '1' {
		hc.Forward = true
	}
	if s[1] == '1' {
		hc.Backward = true
	}
	if s[2] == '1' {
		hc.Left = true
	}
	if s[3] == '1' {
		hc.Right = true
	}

	// speed level is the 5th char as string
	hc.SpeedLevel = string(s[4])

	hc.ValidateDirection()
	return

}

func initParams() {
	SendCommandInterval = time.Duration(S.S.Conf.Car.SendCommandIntervalInMillisecond) * time.Millisecond
}

// StartListenHeroPort 监听小车端口，向小车发送指令。
// 注意，该函数读取 HeroCommandStringChan 的内容，该channel没有缓冲，即入即出；
// 如果输入该chan频率过大，只会定期取每 SendCommandInterval 最新的指令。如果这 SendCommandInterval 中间有指令，
// 则会被丢弃。

// 每 SendCommandInterval 读取一次HeroCommandStringChan，如果有指令，则发送给小车。
// 如果 SendCommandInterval 内有多余指令，则只发送最新的指令，其他指令被丢弃。
// 这样做的目的是为了防止指令发送过快，导致小车无法处理，节省带宽。

func HandleConnection(conn net.Conn) {
	oncer.Do(initParams)

	defer func() {
		conn.Close()
		S.S.Logger.Infof("Car disconnected")
		Info.Connected = false
	}()

	Info.Connected = true
	S.S.Logger.Infof("Car connected")

	// 记录一个最新发送命令的时间戳
	var lastSendCommandTime time.Time

	for {
		select {
		case commandString := <-CommandStringChan:
			{
				log.Println("commandString:", commandString)
				// 如果指令不合法，则不发送
				if !IsStringCommandValid(commandString) {
					continue
				}
				// 如果距离上次发送命令的时间小于 SendCommandInterval ，则不发送
				if time.Now().Sub(lastSendCommandTime) < SendCommandInterval {
					continue
				}

				//turn it into safe command and turn it into string and then send it
				command := StringToSafeHeroCommand(commandString)
				if !shouldSendToHero(command) {
					continue
				}
				_, err := conn.Write([]byte(command.ToCommandString()))
				println("发送给车：", command.ToCommandString())
				// 记录最新发送命令的时间戳
				lastSendCommandTime = time.Now()

				if err != nil {
					log.Println("Send Hero Command Error", err)
				}

			}

		}
	}

}

// IsStringCommandValid 判断字符串是否是合法的命令字符串
// 此处只判断长度是否为8，前4位是否为"1"或"-"，第五位是否为合法速度档位，
// 不判断方向是否合法
func IsStringCommandValid(s string) bool {
	if len(s) != 8 {
		return false
	}

	// if the first 4 char is "----", then it is not valid
	if s[0] == '-' && s[1] == '-' && s[2] == '-' && s[3] == '-' {
		return false
	}

	// the first 4 char must be either "1" or "-"
	for i := 0; i < 4; i++ {
		if s[i] != '1' && s[i] != '-' {
			return false
		}
	}

	// the 5th char can only be from 1-4
	if s[4] < '1' || s[4] > '4' {
		return false
	}

	// last three must be "-"
	for i := 5; i < 8; i++ {
		if s[i] != '-' {
			return false
		}
	}

	return true
}

// shouldSendToHero Whether the command should be sent to hero.
// e.g., if the command contains no direction, then it should not be sent to hero.
func shouldSendToHero(command heroCommand) bool {
	if command.Forward || command.Backward || command.Left || command.Right {
		return true
	}
	return false
}
