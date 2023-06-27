package S

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gopkg.in/gcfg.v1"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"path"
	"sync"
)

var S service

// 仅执行一次的控制器，第一个为读取配置文件，第二个为所有服务
var oncerFile sync.Once
var oncerAll sync.Once

// TOtherConf 一些读完conf文件后的配置等结构体声明

// service 声明服务
type service struct {
	Conf AppConfig

	Redis redis.Client
	//  a gorm db
	Mysql *gorm.DB
	// 一个空白context
	Context context.Context

	// 一个logrus的logger
	Logger         *logrus.Logger
	MultipleWriter io.Writer
}

func InitServices() {

	// 首先初始化文件,不要修改oncerFile
	oncerFile.Do(func() {
		S = service{
			Conf: InitConf(),
		}
	})
	oncerAll.Do(func() {
		S = service{
			// 在此初始化所有服务
			Conf:    S.Conf,
			Redis:   initRedis(),
			Mysql:   initMysql(),
			Context: initEmptyContext(),
		}
		S.Logger, S.MultipleWriter = initLogger()

	})
}

func InitConf() AppConfig {
	tempConf := AppConfig{}
	// get environment device_name
	confDir := "./etc/app.conf"

	//read file /etc/tr.txt, and get the content, switch the content
	content, err := os.ReadFile("/etc/tr.txt")
	if err != nil {

	} else {
		switch string(content) {
		case "lab-server":
			confDir = "./etc/app.conf.lab"
		case "box-machine":
			confDir = "./etc/app.conf.box"
		}
	}
	// read file above
	if err := gcfg.ReadFileInto(&tempConf, confDir); err != nil {
		log.Fatalln("读取配置文件错误，", err)
	}

	return tempConf
}

func initRedis() redis.Client {
	log.Println("redis", S.Conf.Redis.Ip+":"+S.Conf.Redis.Port, S.Conf.Redis.Password, S.Conf.Redis.Db)
	tempRedis := *redis.NewClient(&redis.Options{
		Addr:     S.Conf.Redis.Ip + ":" + S.Conf.Redis.Port,
		Password: S.Conf.Redis.Password,
		DB:       S.Conf.Redis.Db,
	})
	pong, err := tempRedis.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("初始化redis失败", err)
	} else {
		log.Println("redis连接成功", pong)
	}
	return tempRedis
}

func initMysql() *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", S.Conf.Database.Username, S.Conf.Database.Password, S.Conf.Database.Url, S.Conf.Database.Port, S.Conf.Database.TableName)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})
	if err != nil {
		log.Fatalln("初始化数据库连接失败", err)
	}
	return db
}

func initEmptyContext() context.Context {
	return context.Background()
}

// initLogger 初始化日志工具
// logrusLogger contains the io.Writer, and the io.Writer returned
// is only for gin to write logs to file
func initLogger() (*logrus.Logger, io.Writer) {
	// init a lumberjack logger
	lumberjackLogger := &lumberjack.Logger{
		Filename:   path.Join(S.Conf.Log.LogFileDirPref, S.Conf.Log.LogFileName), // 日志文件路径
		MaxSize:    1,                                                            // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 1,                                                            // 日志文件最多保存多少个备份
		MaxAge:     7,                                                            // 文件最多保存多少天
		Compress:   false,                                                        // 是否压缩

	}
	// init a logrus logger
	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrusLogger.SetLevel(logrus.InfoLevel)
	//logrusLogger.SetReportCaller(true)
	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)
	logrusLogger.SetOutput(multiWriter)

	return logrusLogger, multiWriter
}
