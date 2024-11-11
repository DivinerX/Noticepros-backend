package config

import (
	"noticepros/config/app_config"
	"noticepros/config/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
}
