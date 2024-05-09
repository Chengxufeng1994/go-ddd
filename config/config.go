package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Transport      Transport      `mapstructure:"transport"`
	Infrastructure Infrastructure `mapstructure:"infrastructure"`
}

type Transport struct {
	HTTP HTTP `mapstructure:"http"`
}

type HTTP struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Infrastructure struct {
	Persistence Persistence `mapstructure:"persistence"`
	Cache       Cache       `mapstructure:"cache"`
}

type Persistence struct {
	Type     string   `mapstructure:"type"`
	Postgres Postgres `mapstructure:"postgres"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

type Cache struct{}

func Load() (cfg *Configuration, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return
}
