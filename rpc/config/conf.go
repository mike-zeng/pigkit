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
	ServiceName string `yaml:"service_name"`
	NodeId string `yaml:"node_id"`
	TimeOut int      `yaml:"timeout"`
	Etcd    EtcdConf `yaml:"etcd"` // etcd连接信息
	JaegerConf JaegerConf `yaml:"jaeger"`
}

type EtcdConf struct{
	Hosts []string `yaml:"hosts,flow"`
	TimeOut int `yaml:"timeout"`
}
type JaegerConf struct {
	CollectorEndpoint string `yaml:"collector_endpoint"`
}
var confInit sync.Once
var conf *config

func InitConf(path string){

}

func GetConfig()*config {
	confInit.Do(func() {
		conf = &config{}
		yamlFile, err := ioutil.ReadFile("./pig.yaml")
		if err != nil {
			log.Fatalln(err)
		}
		err = yaml.Unmarshal(yamlFile, conf)
		if err != nil {
			log.Fatalln(err)
		}
	})
	if conf == nil {
		log.Fatalln(errors.New("conf not read"))
	}
	return conf
}