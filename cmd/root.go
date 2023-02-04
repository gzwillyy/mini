/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd 表示在没有任何子命令的情况下调用时的基本命令
var rootCmd = &cobra.Command{
	Use:   "mini",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute 将所有子命令添加到根命令并适当地设置标志。
// 这由 main.main() 调用。 它只需要对 rootCmd 发生一次。
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 设置了程序在运行时执行的回调函数 initConfig ,initConfig用来初始化配置
	cobra.OnInitialize(initConfig)

	// 在这里您将定义标志和配置设置。
	// Cobra 支持持久性标志，如果在此处定义，
	// 对你的应用程序来说是全局的。
	// 通过以下代码给 cobrademo 应用设置了命令行选项 --config，该选项用来指定配置文件（默认值为""）：
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mini.yaml)")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.mini.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// 用来设置 viper 需要读取的配置文件（该配置文件通过 --config 参数指定）
		viper.SetConfigFile(cfgFile)
	} else {
		// 获取用户主目录，
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)    // 通过 viper.AddConfigPath(home) 将用户主目录加入到配置文件的搜索路径中
		viper.SetConfigType("yaml")  //设置配置文件格式
		viper.SetConfigName(".mini") //设置配置文件名
	}

	// 设置 viper 查找是否有跟配置文件中相匹配的环境变量，如果有，则将该环境变量的值设置为配置项的值
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in. 读取设置的配置文件
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
