package config

import (
	"cdel/demo/Normal/tool"
	"os"

	"github.com/spf13/viper"
)

var C = new()

type Config struct {
	ProjectName    string // 项目名称
	ProjectVersion string // 项目本身的版本信息
	Env            string // 部署环境
	Host           string // 实例部署位置
	WorkPath       string
	LogPath        string
}

const (
	KEY_PROJECT_NAME    = "prj.name"
	KEY_PROJECT_VERSION = "prj.version"
	KEY_ENV             = "env"
	KEY_LOG_ROOT        = "log.root"

	VAL_PRODUCT = "product"
)

// 设置配置的缺省值
func setDefault() {
	viper.SetDefault(KEY_PROJECT_NAME, "back-normal")
	viper.SetDefault(KEY_PROJECT_VERSION, "v0.0.1")
	viper.SetDefault(KEY_LOG_ROOT, "/web/logs")
	viper.SetDefault(KEY_ENV, VAL_PRODUCT)
}

func new() (c Config) {
	setDefault()
	viper.AutomaticEnv()

	c = Config{}
	var e error

	c.ProjectName = viper.GetString(KEY_PROJECT_NAME)
	c.ProjectVersion = viper.GetString(KEY_PROJECT_VERSION)
	c.Env = viper.GetString(KEY_ENV)
	c.WorkPath, e = os.Getwd()
	if e != nil {
		panic(e)
	}
	c.Host, e = tool.GetOutBoundIP()
	if e != nil {
		panic(e)
	}
	c.LogPath = viper.GetString(KEY_LOG_ROOT) + "/" + c.Host + "-" + c.ProjectName

	// viper.SetConfigName("config")   // name of config file (without extension)
	// viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	// viper.AddConfigPath("./config") // path to look for the config file in
	// viper.AddConfigPath(".")        // optionally look for config in the working directory
	// err := viper.ReadInConfig()     // Find and read the config file
	// if err != nil {                 // Handle errors reading the config file
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }
	return
}
