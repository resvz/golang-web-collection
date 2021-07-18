package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var globalConfig Config

type Config struct {
	Port   string
	Upload struct {
		BasePath string
	}
}

func init() {
	data, _ := ioutil.ReadFile("config.yml")
	fmt.Println(string(data))
	globalConfig = Config{}
	yaml.Unmarshal(data, &globalConfig)
	d, _ := yaml.Marshal(&globalConfig)
	fmt.Println("配置 :\n", string(d))
}

func ServerPort() string {
	return globalConfig.Port
}

func UploadBasePath() string {
	return globalConfig.Upload.BasePath
}
