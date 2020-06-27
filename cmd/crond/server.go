package crond

import (
	"fmt"
	"go-admin-demo/cache"
	"go-admin-demo/database"
	"go-admin-demo/tools"
	config2 "go-admin-demo/tools/config"
	"go-admin-demo/workflow-engine/service"
	"log"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var (
	config   string
	port     string
	mode     string
	StartCmd = &cobra.Command{
		Use:     "crond",
		Short:   "Start crond server",
		Example: "go-admin crond start config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8000", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}

func usage() {
	usageStr := `starting crond server`
	fmt.Printf("%s\n", usageStr)
}

func setup() {

	//1. 读取配置
	config2.ConfigSetup(config)
	//2. 设置日志
	tools.InitLogger()
	//3. 初始化数据库链接
	database.Setup()
	//4. 初始化redis连接
	err := cache.SetRedis(config2.RedisConfig)
	if err != nil {
		log.Panic("set redis error", err)
	}

}

func run() error {

	// 启动定时任务
	service.CronJobs()

	defer database.Eloquent.Close()

	log.Println("Enter Control + C Shutdown Server")
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	log.Println("Server exiting")
	return nil
}
