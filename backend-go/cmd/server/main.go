// 文件作用：Go HTTP 服务启动入口：加载环境变量、构建 Gin 应用并监听端口。
package main

import (
	"fmt"
	"log"

	"csdn-analytics/backend-go/internal/app"
	"csdn-analytics/backend-go/internal/config"
	"github.com/joho/godotenv"
)

// main 是 HTTP 服务的程序入口。
func main() {
	_ = godotenv.Load()

	cfg := config.Load()
	router, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Printf("starting go backend on %s", addr)

	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
