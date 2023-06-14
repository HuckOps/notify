package config

import (
	"github.com/huckops/auto_config"
	"github.com/huckops/auto_config/source/file"
)

type config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
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
