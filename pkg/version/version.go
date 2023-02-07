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

// version
package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/gosuri/uitable"
)

var (
	// GitVersion 是语义化的版本号
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate 是 ISO8601 格式的构建时间, $(date -u +'%Y-%m-%dT%H:%M:%SZ') 命令的输出.
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit 是 Git 的 SHA1 值，$(git rev-parse HEAD) 命令的输出.
	GitCommit = "$Format:%H$"
	// GitTreeState 代表构建时 Git 仓库的状态，可能的值有：clean, dirty.
	GitTreeState = ""
)

type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String 返回人性化的版本信息字符串.
func (info Info) String() string {
	if s, err := info.Text(); err == nil {
		return string(s)
	}

	return info.GitVersion
}

// ToJSON 以 JSON 格式返回版本信息.
func (info Info) ToJSON() string {
	s, _ := json.Marshal(info)

	return string(s)
}

// Text 将版本信息编码为 UTF-8 格式的文本，并返回.
func (info Info) Text() ([]byte, error) {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("gitVersion:", info.GitVersion)
	table.AddRow("gitCommit:", info.GitCommit)
	table.AddRow("gitTreeState:", info.GitTreeState)
	table.AddRow("buildDate:", info.BuildDate)
	table.AddRow("goVersion:", info.GoVersion)
	table.AddRow("compiler:", info.Compiler)
	table.AddRow("platform:", info.Platform)

	return table.Bytes(), nil
}

// Get 返回详尽的代码库版本信息，用来标明二进制文件由哪个版本的代码构建.
func Get() Info {
	// 以下变量通常由 -ldflags 进行设置
	return Info{
		GitVersion:   GitVersion,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
