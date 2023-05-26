package S

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gcfg.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var S service

// 仅执行一次的控制器，第一个为读取配置文件，第二个为所有服务
var oncerFile sync.Once
var oncerAll sync.Once

// service 声明服务
type service struct {
	Conf     AppConfig
	Redis    redis.Client
	TaskPool *ants.Pool
	//  a gorm db
	Mysql *gorm.DB
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
			Conf:     S.Conf,
			Redis:    initRedis(),
			TaskPool: initTaskPool(),
			Mysql:    initMysql(),
		}
	})
}

func InitConf() AppConfig {
	tempConf := AppConfig{}
	if err := gcfg.ReadFileInto(&tempConf, "./etc/app.conf"); err != nil {
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
	return tempRedis
}

func initTaskPool() *ants.Pool {
	pool, err := ants.NewPool(30)
	if err != nil {
		log.Fatalln("初始化任务池失败", err)
	}
	return pool
}

func initMysql() *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", S.Conf.Database.Username, S.Conf.Database.Password, S.Conf.Database.Url, S.Conf.Database.Port, S.Conf.Database.TableName)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalln("初始化数据库连接失败", err)
	}
	return db
}
