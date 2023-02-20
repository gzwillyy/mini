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

package main

import (
	"os"

	// _ "go.uber.org/automaxprocs".

	"github.com/gzwillyy/mini/internal/mini"
)

// Go 程序的默认入口函数(主函数).
func main() {
	// 使用 cobra 框架创建应用的cli交互 aa
	command := mini.NewMiniCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
