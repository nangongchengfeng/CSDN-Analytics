// 文件作用：爬虫命令行入口：加载配置并组装 crawl 子命令。
package main

import (
	"log"

	"csdn-analytics/backend-go/internal/app"
	"csdn-analytics/backend-go/internal/cli"
	"csdn-analytics/backend-go/internal/config"
	"csdn-analytics/backend-go/internal/spider"
	"github.com/joho/godotenv"
)

// main 是爬虫命令行的程序入口。
func main() {
	_ = godotenv.Load()

	cfg := config.Load()
	root := cli.NewRootCommand(func(userID string) cli.SpiderRunner {
		service, err := app.NewSpiderService(cfg, userID)
		if err != nil {
			log.Fatal(err)
		}
		return service
	})

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}

var _ cli.SpiderRunner = (*spider.Service)(nil)
