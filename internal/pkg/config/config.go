// Copyright 2023 The Mini Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
