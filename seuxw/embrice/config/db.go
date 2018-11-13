package config

import (
	"gopkg.in/ini.v1"
	"seuxw/embrice/entity"
)

var configPath = "/data/config/seuxw.cfg"

func ReadDBConfig() (*entity.Database, error) {
	config, err := ReadConfig()
	return config.Database, err
}

func ReadConfig() (config entity.Config, err error) {

	if err = ini.MapTo(&config, configPath); err != nil {
		return
	}
	return
}
