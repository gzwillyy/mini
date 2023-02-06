// Package config 配置信息的结构体映射
package config

// JWTConfig jwt key
type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

// RedisConfig redis配置信息
type RedisConfig struct {
	Host   string `mapstructure:"host" json:"host"`
	Port   int    `mapstructure:"port" json:"port"`
	Expire int    `mapstructure:"expire" json:"expire"`
}

// ServerConfig 应用配置信息
type ServerConfig struct {
	Name      string      `mapstructure:"name" json:"name"`
	Port      int         `mapstructure:"port" json:"port"`
	Host      string      `mapstructure:"host" json:"host"`
	Debug     string      `mapstructure:"debug" json:"debug"`
	JWTInfo   JWTConfig   `mapstructure:"jwt" json:"jwt"`
	RedisInfo RedisConfig `mapstructure:"redis" json:"redis"`
}
