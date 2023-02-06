// Package global 应用的全局配置允许在任意地方使用
package global

import (
	ut "github.com/go-playground/universal-translator"

	"github.com/gzwillyy/mini/internal/pkg/config"
)

var (
	// Trans 翻译器语言包
	Trans ut.Translator
	// ServerConfig 应用的配置信息 (指针类型 其他地方会（需要）改变它)
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	//全局微服务srv grpc客户端
	//UserSrvClient proto.UserClient
)
