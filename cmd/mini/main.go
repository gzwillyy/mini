package main

import (
	"os"

	// _ "go.uber.org/automaxprocs"

	"github.com/gzwillyy/mini/internal/mini"
)

// Go 程序的默认入口函数(主函数).
func main() {

	//使用 cobra 框架创建应用的cli交互 aa
	command := mini.NewMiniCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
