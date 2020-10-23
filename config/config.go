package config

import (
	"blog.zozoo.net/common"
	"encoding/json"
	"io/ioutil"
	"os"
)

type (

	//数据库配置
	Mysql struct {
		Username string
		Password string
		Database string
		Host string
	}

	//总配置
	Config struct {
		Mysql Mysql
	}
)

//读取配置文件
func LoadConfig()*Config {
	config := &Config{}
	file, err := os.Open("config/config.json")
	common.HandleErr(err,"打开配置文件出错")

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	common.HandleErr(err,"读取配置文件出错")

	err = json.Unmarshal(bytes, config)
	common.HandleErr(err,"解析配置文件出错")
	return config
}