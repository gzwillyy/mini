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

package verflag

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/pflag"

	"github.com/gzwillyy/mini/pkg/version"
)

// 定义一些常量.
const (
	VersionFalse versionValue = 0
	VersionTrue  versionValue = 1
	VersionRaw   versionValue = 2
)

const (
	strRawVersion   = "raw"
	versionFlagName = "version"
)

type versionValue int

var versionFlag = Version(versionFlagName, VersionFalse, "Print version information and quit.")

func (v *versionValue) IsBoolFlag() bool {
	return true
}

func (v *versionValue) Get() interface{} {
	return v
}

// Set 实现了 pflag.Value 接口中的 Set 方法.
func (v *versionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw

		return nil
	}
	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}

	return err
}

// Type 实现了 pflag.Value 接口中的 Type 方法.
func (v *versionValue) Type() string {
	return "version"
}

// String 实现了 pflag.Value 接口中的 String 方法.
func (v *versionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}

	return fmt.Sprintf("%v", bool(*v == VersionTrue))
}

// VersionVar 定义了一个具有指定名称和用法的标志.
func VersionVar(p *versionValue, name string, value versionValue, usage string) {
	*p = value
	pflag.Var(p, name, usage)
	// `--version` 等价于 `--version=true`
	pflag.Lookup(name).NoOptDefVal = "true"
}

// Version 包装了 VersionVar 函数.
func Version(name string, value versionValue, usage string) *versionValue {
	p := new(versionValue)
	VersionVar(p, name, value, usage)

	return p
}

// AddFlags 在任意 FlagSet 上注册这个包的标志，这样它们指向与全局标志相同的值.
func AddFlags(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(versionFlagName))
}

// PrintAndExitIfRequested 将检查是否传递了 `--version` 标志，如果是，则打印版本并退出.
func PrintAndExitIfRequested() {
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v\n", version.Get())
		os.Exit(0)
	} else if *versionFlag == VersionTrue {
		fmt.Printf("%s\n", version.Get())
		os.Exit(0)
	}
}
