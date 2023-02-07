// Copyright (c) 2023 gzwillyy
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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

type LogConfig struct {
	DisableCaller     bool     `mapstructure:"disable-caller" json:"disable-caller"`
	DisableStacktrace bool     `mapstructure:"disable-stacktrace" json:"disable-stacktrace"`
	Level             string   `mapstructure:"level" json:"level"`
	Format            string   `mapstructure:"format" json:"format"`
	OutputPaths       []string `mapstructure:"output-paths" json:"output-paths"`
}

// ServerConfig 应用配置信息
type ServerConfig struct {
	Name      string      `mapstructure:"name" json:"name"`
	Host      string      `mapstructure:"host" json:"host"`
	Port      int         `mapstructure:"port" json:"port"`
	Ginmode   string      `mapstructure:"ginmode" json:"ginmode"`
	JWTInfo   JWTConfig   `mapstructure:"jwt" json:"jwt"`
	RedisInfo RedisConfig `mapstructure:"redis" json:"redis"`
	LogConfig LogConfig   `mapstructure:"log" json:"log"`
}
