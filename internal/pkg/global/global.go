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
