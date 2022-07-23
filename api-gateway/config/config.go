package config

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
)

func InitConfig() Config {
	var cfg Config
	content, err := ioutil.ReadFile("./config/config.toml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = toml.Unmarshal(content, &cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return cfg
}
