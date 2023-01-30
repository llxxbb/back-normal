package config

import (
	"cdel/demo/Normal/tool"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/spf13/viper"
)

var C = new()

type Config struct {
	ProjectName    string // 项目名称
	ProjectVersion string // 项目本身的版本信息
	Env            string // 部署环境
	Port           string // 对外服务的端口号
	GinRelease     bool   // gin 是否以 release 模式工作
	LogRoot        string // 保存日志的根路径

	// 注意：下面的配置项目在运行时自动设置，无需配置。
	Host     string // 实例部署位置
	WorkPath string // 项目启动所在的目录
	LogPath  string // 日志输出的位置，由 log.root 主机IP 项目名 等组成
}

const (
	KEY_ENV     = "env"
	VAL_PRODUCT = "product"
)

// 用于反射，将配置文件中的配置项映射到 `Config` 对象的属性上
var fieldMap = map[string]string{
	"prj.name":    "ProjectName",
	"prj.version": "ProjectVersion",
	"port":        "Port",
	"gin.release": "GinRelease",
	"log.root":    "LogRoot",
}

const (
	_fileDefault  = "default"
	_fileName     = "config"
	_fileType     = "yaml"
	_nameSplitter = "_"
)

func new() (c Config) {

	// 可以从环境变量中取值
	viper.SetDefault(KEY_ENV, VAL_PRODUCT)
	viper.AutomaticEnv()
	c = Config{}
	c.Env = viper.GetString(KEY_ENV)
	c.mergeFile("")

	// 读取缺省配置文件
	c.mergeFile(_fileName + _nameSplitter + _fileDefault)

	// 读取 profile 对应的配置文件
	c.mergeFile(_fileName + _nameSplitter + c.Env)

	// 设置工作目录、日志目录等
	c.setWdAndLogPath()

	return c
}

func (c *Config) setWdAndLogPath() {
	var e error
	c.WorkPath, e = os.Getwd()
	if e != nil {
		panic(e)
	}
	c.Host, e = tool.GetOutBoundIP()
	if e != nil {
		panic(e)
	}
	c.LogPath = c.LogRoot + "/" + c.Host + "-" + c.ProjectName
}

// 读取并合并配置项
// `path` 为 "" 则从环境变量中读取
func (c *Config) mergeFile(path string) {
	if path != "" {
		viper.SetConfigName(path)      // name of config file (without extension)
		viper.SetConfigType(_fileType) // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath(".")       // optionally look for config in the working directory
		err := viper.ReadInConfig()    // Find and read the config file
		if err != nil {                // Handle errors reading the config file
			panic(fmt.Errorf("fatal error: read config file: %s, %w", path, err))
		}
	}

	// 利用反射进行赋值
	cV := reflect.ValueOf(c).Elem()
	for k, v := range fieldMap {
		fV := viper.GetString(k)
		if fV == "" {
			continue
		}
		var rV reflect.Value
		var err error
		switch cV.FieldByName(v).Type().Name() {
		case "string":
			rV = reflect.ValueOf(fV)
		case "bool":
			rtn, e := strconv.ParseBool(fV)
			if e == nil {
				rV = reflect.ValueOf(rtn)
			} else {
				err = e
			}
		case "int":
			rtn, e := strconv.ParseInt(fV, 0, 32)
			if e == nil {
				rV = reflect.ValueOf(rtn)
			} else {
				err = e
			}
		default:
			panic(fmt.Sprintf("config item: %s, unhandled type.", k))
		}
		if err != nil {
			panic(fmt.Sprintf("config item: %s, value type error. %v", k, err))
		}
		cV.FieldByName(v).Set(rV)
	}
}
