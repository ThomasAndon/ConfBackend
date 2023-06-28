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
		SaveStatidPcdFileDirPrefix       string
	}
	Location struct {
		UpdateIntervalInSecond int
		PktCacheTimeInMinute   int
	}
	Chat struct {
		SaveStaticFileDirPrefix string
		ServerFileUrlPrefix     string
	}
	Log struct {
		LogFileDirPref string
		LogFileName    string
	}
	Pcd struct {
		ServerPCDFileUrlPrefix  string
		SaveStaticFileDirPrefix string
	}
	Node struct {
		NodeInfo []string
	}
	Meta struct {
		ConfDevice string
	}
	Web struct {
		DistFolderDir string
	}
}
