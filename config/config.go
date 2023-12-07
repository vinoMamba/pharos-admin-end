package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type JwtConfig struct {
	Secret    string
	ExpiresIn int64
}

type MysqlConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type config struct {
	Jwt   JwtConfig
	Mysql MysqlConfig
	Salt  string
}

var _config config

func LoadConfig(path string) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)     // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := viper.Unmarshal(&_config); err != nil {
		panic(err)
	}
}

func GetJwtSecret() string {
	return _config.Jwt.Secret
}

func GetJwtExpiresIn() int64 {
	return _config.Jwt.ExpiresIn
}

func GetSalt() string {
	return _config.Salt
}

func GetMysqlConfig() MysqlConfig {
	return _config.Mysql
}
