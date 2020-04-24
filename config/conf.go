package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf *Config

type Config struct {
	RabbitMQ RabbitMQ
}

type RabbitMQ struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

func init() {
	GetConf()
}

func GetConf() {
	yamlFile, err := ioutil.ReadFile("config/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		fmt.Println(err.Error())
	}

}
