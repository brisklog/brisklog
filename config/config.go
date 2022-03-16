package config

/**
 * @author: xaohuihui
 * @Path: brisklog/config/config.go
 * @Description:
 * @datetime: 2022/3/16 17:32:13
 * software: GoLand
**/

type ServerConfig struct {
	Name          string      `mapstructure:"name"`
	Port          int         `mapstructure:"port"`
	MysqlInfo     MysqlConfig `mapstructure:"mysql"`
	RedisInfo     RedisConfig `mapstructure:"redis"`
	LogsAddress   string      `mapstructure:"logsAddress"`
	PasswordLevel int         `mapstructure:"passwordLevel"`
	JWTKey        JWTConfig   `mapstructure:"jwt"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
