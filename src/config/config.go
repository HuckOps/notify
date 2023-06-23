package config

import (
	"github.com/huckops/auto_config"
	"github.com/huckops/auto_config/source/file"
)

type config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
	JWT    JWT    `yaml:"jwt"`
}

type Server struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type DB struct {
	Mongo Mongo `yaml:"mongo"`
	Redis Redis `yaml:"redis"`
	MySQL MySQL `yaml:"mysql"`
}

type Mongo struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"DB"`
}

type JWT struct {
	Secret string `yaml:"secret"`
	Exp    int    `yaml:"exp"`
	Issuer string `yaml:"issuer"`
}

var Config config
var ConfigInstance *auto_config.Config

func InitConfig(fp string, callback ...func()) {
	var options []auto_config.Option
	options = append(options, auto_config.WithSource(file.NewSource(file.WithPath(fp))))
	options = append(options, auto_config.WithEntity(&Config))
	for _, cb := range callback {
		option := auto_config.WithCallback(cb)
		options = append(options, option)
	}

	c, err := auto_config.NewConfig(options...)
	if err != nil {
		panic(err)
	}
	ConfigInstance = c
	//c.Watcher()
}

func ConfigWatchDog() {
	go ConfigInstance.Watcher()
}
