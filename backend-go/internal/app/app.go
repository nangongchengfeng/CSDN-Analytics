// 文件作用：应用装配层：负责创建仓储、Spider 服务和 HTTP 路由。
package app

import (
	"csdn-analytics/backend-go/internal/config"
	dbpkg "csdn-analytics/backend-go/internal/db"
	httpapi "csdn-analytics/backend-go/internal/http"
	"csdn-analytics/backend-go/internal/repository"
	"csdn-analytics/backend-go/internal/spider"

	"github.com/gin-gonic/gin"
)

// New 组装完整的 Gin 应用，并注入数据库仓储与抓取服务。
func New(cfg config.Config) (*gin.Engine, error) {
	infoRepo, categorizeRepo, articleRepo, err := buildRepositories(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	return httpapi.NewRouter(cfg, httpapi.Dependencies{
		InfoRepo:       infoRepo,
		CategorizeRepo: categorizeRepo,
		ArticleRepo:    articleRepo,
		Spider:         spider.NewService(cfg, infoRepo, categorizeRepo, articleRepo),
	}), nil
}

// NewSpiderService 创建可供 CLI 复用的抓取服务实例。
func NewSpiderService(cfg config.Config, userID string) (*spider.Service, error) {
	infoRepo, categorizeRepo, articleRepo, err := buildRepositories(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if userID != "" {
		cfg.CSDNUserID = userID
		cfg.CSDNBlogURL = "https://blog.csdn.net/" + userID
	}

	return spider.NewService(cfg, infoRepo, categorizeRepo, articleRepo), nil
}

// buildRepositories 打开数据库并构造三类基础仓储。
func buildRepositories(databaseURL string) (*repository.InfoRepository, *repository.CategorizeRepository, *repository.ArticleRepository, error) {
	gdb, err := dbpkg.Open(databaseURL)
	if err != nil {
		return nil, nil, nil, err
	}

	if err := dbpkg.AutoMigrate(gdb); err != nil {
		return nil, nil, nil, err
	}

	return repository.NewInfoRepository(gdb), repository.NewCategorizeRepository(gdb), repository.NewArticleRepository(gdb), nil
}
