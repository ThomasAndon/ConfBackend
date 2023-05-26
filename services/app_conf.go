package S

// AppConfig 项目的配置文件的结构体声明
type AppConfig struct {
	App struct {
		Port string
	}
	Database struct {
		Url       string
		Port      string
		Username  string
		Password  string
		TableName string
	}
	Redis struct {
		Ip       string
		Port     string
		Password string
		Db       int
	}
	Car struct {
		Port                             string
		SendCommandIntervalInMillisecond int64
	}
	Location struct {
		UpdateIntervalInSecond int
	}
}
