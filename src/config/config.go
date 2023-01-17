package config

import (
	"cdel/demo/Normal/tool"
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

var C = new()

type Config struct {
	ProjectName    string // 项目名称
	ProjectVersion string // 项目本身的版本信息
	Env            string // 部署环境
	Host           string // 实例部署位置
	Port           string // 对外服务的端口号
	WorkPath       string
	LogPath        string
}

const (
	KEY_ENV      = "env"
	KEY_LOG_ROOT = "log.root"

	VAL_PRODUCT = "product"
)

// 用于反射，将配置文件中的配置项映射到 `Config` 对象的属性上
var fieldMap = map[string]string{
	"prj.name":    "ProjectName",
	"prj.version": "ProjectVersion",
	"port":        "Port",
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
	cEnv := c // 备份环境变量中的配置项

	// 读取缺省配置文件
	c.mergeFile(_fileName + _nameSplitter + _fileDefault)

	// 读取 profile 对应的配置文件
	c.mergeFile(_fileName + _nameSplitter + c.Env)

	// 用环境变量中的配置项进行覆盖
	c.mergeAnother(&cEnv)

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
	c.LogPath = c.LogPath + "/" + c.Host + "-" + c.ProjectName
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
		if fV != "" {
			cV.FieldByName(v).Set(reflect.ValueOf(&fV).Elem())
		}
	}
}

// 读取并合并配置项
func (c *Config) mergeAnother(another *Config) {
	cSelf := reflect.ValueOf(c).Elem()
	cAnother := reflect.ValueOf(another).Elem()

	for _, v := range fieldMap {
		fV := cAnother.FieldByName(v)
		if !fV.IsValid() {
			panic("can't find field: " + v)
		}
		if fV.String() != "" {
			cSelf.FieldByName(v).Set(fV)
		}
	}
}
