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

package global

import (
	ut "github.com/go-playground/universal-translator"

	"github.com/gzwillyy/mini/internal/pkg/config"
)

var (
	// Trans 翻译器语言包.
	Trans ut.Translator
	// ServerConfig 应用的配置信息 (指针类型 其他地方会（需要）改变它).
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	// 全局微服务srv grpc客户端
	// UserSrvClient proto.UserClient.
)
