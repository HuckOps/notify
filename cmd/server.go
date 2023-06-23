/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/HuckOps/notify/pkg/rbac"
	"github.com/HuckOps/notify/src/config"
	"github.com/HuckOps/notify/src/db/mongo"
	"github.com/HuckOps/notify/src/db/mysql"
	"github.com/HuckOps/notify/src/db/redis"
	"github.com/HuckOps/notify/src/server"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var fp string
var privateKeyPath string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer(privateKeyPath)
		config.InitConfig(fp, s.Restart, mongo.Mongo.Load, redis.Redis.Load, mysql.MySQL.Load)
		config.ConfigWatchDog()
		//server.Server()
		rbac.Init()

		go s.Listen()
		signChan := make(chan os.Signal, 1)
		signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)
		// 在另一个 goroutine 中监听信号通道

		go func() {
			sig := <-signChan

			fmt.Println("接收到信号:", sig)
			s.Kill()
			os.Exit(0) // 优雅地终止程序
		}()

		select {}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&fp, "config", "c", "", "config file path")
	serverCmd.Flags().StringVarP(&privateKeyPath, "private", "k", "", "Password decode key path")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
