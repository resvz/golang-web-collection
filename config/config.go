package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var GlobalConfig Config

type Config struct {
	Port string
}

func init() {
	data, _ := ioutil.ReadFile("config.yml")
	fmt.Println(string(data))
	GlobalConfig = Config{}
	yaml.Unmarshal(data, &GlobalConfig)
	d, _ := yaml.Marshal(&GlobalConfig)
	fmt.Println("配置 :\n", string(d))

}
