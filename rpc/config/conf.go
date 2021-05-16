package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

type config struct {
	configFilePath string
	PigServer      PigServerConfig `yaml:"server"`
}

type PigServerConfig struct {
	Port    int      `yaml:"port"` // 端口号
	TimeOut int      `yaml:"timeout"`
	Etcd    EtcdConf `yaml:"etcd"` // etcd连接信息
}

type EtcdConf struct{
	Hosts []string `yaml:"hosts,flow"`
	TimeOut int `yaml:"timeout"`
}

var once sync.Once
var conf *config

func InitConf(path string){
	once.Do(func() {
		conf = &config{}
		yamlFile, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalln(err)
		}
		err = yaml.Unmarshal(yamlFile, conf)
		if err != nil {
			log.Fatalln(err)
		}
	})
}

func GetConfig()*config {
	if conf == nil {
		log.Fatalln(errors.New("conf not read"))
	}
	return conf
}