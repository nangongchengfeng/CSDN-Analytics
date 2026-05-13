// 文件作用：HTTP 路由层：注册 Gin 中间件并挂载全部 /api 路由。
package http

import (
	stdhttp "net/http"

	"csdn-analytics/backend-go/internal/config"
	"csdn-analytics/backend-go/internal/http/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Dependencies 收集路由层运行所依赖的仓储与服务。
type Dependencies struct {
	// InfoRepo 提供用户信息读取能力。
	InfoRepo handlers.InfoRepository
	// CategorizeRepo 提供分类信息读取能力。
	CategorizeRepo handlers.CategorizeRepository
	// ArticleRepo 提供文章信息读取能力。
	ArticleRepo handlers.ArticleRepository
	// Spider 提供抓取任务执行能力。
	Spider handlers.SpiderService
}

// NewRouter 创建 Gin 路由并注册全部 API。
func NewRouter(cfg config.Config, deps Dependencies) *gin.Engine {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	// 注册 Gin 中间件。
	router.Use(gin.Logger(), gin.Recovery())
	// 注册 CORS 中间件。
	router.Use(cors.New(cors.Config{
		AllowOrigins: cfg.CORSOrigins,
		AllowMethods: []string{stdhttp.MethodGet, stdhttp.MethodPost, stdhttp.MethodPut, stdhttp.MethodDelete, stdhttp.MethodOptions},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))
	// 注册 API 路由。
	handlerSet := handlers.New(deps.InfoRepo, deps.CategorizeRepo, deps.ArticleRepo, deps.Spider)
	api := router.Group("/api")
	{
		api.GET("/info", handlerSet.GetInfo)
		api.POST("/info/crawl", handlerSet.CrawlInfo)
		api.GET("/quarter", handlerSet.GetQuarterStats)
		api.GET("/read", handlerSet.GetReadStats)
		api.GET("/pie", handlerSet.GetPieData)
		api.GET("/categorize", handlerSet.GetPieData)
		api.POST("/categorize/crawl", handlerSet.CrawlCategorize)
		api.GET("/articles", handlerSet.GetArticlesList)
		api.POST("/articles/crawl", handlerSet.CrawlArticles)
		api.GET("/heatmap/:year", handlerSet.GetHeatmap)
		api.GET("/years", handlerSet.GetYears)
	}

	return router
}
