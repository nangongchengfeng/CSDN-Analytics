// 文件作用：HTTP 处理层：实现读接口和抓取接口的控制器逻辑。
package handlers

import (
	"errors"
	"net/http"
	"strconv"

	dbpkg "csdn-analytics/backend-go/internal/db"
	"csdn-analytics/backend-go/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InfoRepository 约束 handler 读取用户信息所需的能力。
type InfoRepository interface {
	// First 返回第一条用户信息记录。
	First() (dbpkg.Info, error)
}

// CategorizeRepository 约束 handler 读取分类信息所需的能力。
type CategorizeRepository interface {
	// ListAll 返回全部分类记录。
	ListAll() ([]dbpkg.Categorize, error)
}

// ArticleRepository 约束 handler 读取文章信息所需的能力。
type ArticleRepository interface {
	// ListAll 返回全部文章记录。
	ListAll() ([]dbpkg.Article, error)
}

// SpiderService 约束 handler 触发抓取任务所需的能力。
type SpiderService interface {
	// CrawlInfo 抓取用户信息。
	CrawlInfo() (bool, error)
	// CrawlCategorize 抓取分类信息。
	CrawlCategorize() error
	// CrawlArticles 抓取文章信息。
	CrawlArticles() error
}

// Handler 封装 HTTP 控制器依赖，并提供各个接口处理函数。
type Handler struct {
	infoRepo       InfoRepository
	categorizeRepo CategorizeRepository
	articleService *service.ArticleService
	spider         SpiderService
}

// New 创建 HTTP 处理器集合。
func New(infoRepo InfoRepository, categorizeRepo CategorizeRepository, articleRepo ArticleRepository, spider SpiderService) *Handler {
	return &Handler{
		infoRepo:       infoRepo,
		categorizeRepo: categorizeRepo,
		articleService: service.NewArticleService(articleRepo),
		spider:         spider,
	}
}

// GetInfo 返回用户基础信息。
func (h *Handler) GetInfo(c *gin.Context) {
	info, err := h.infoRepo.First()
	if err != nil {
		// 这里和 Flask 版本保持一致：没有数据时返回 200 + data: null，而不是 404。
		if errors.Is(err, gorm.ErrRecordNotFound) {
			success(c, nil)
			return
		}
		failure(c, http.StatusInternalServerError, "获取用户信息失败: "+err.Error())
		return
	}

	success(c, gin.H{
		"id":          info.ID,
		"date":        info.Date,
		"head_img":    info.HeadImg,
		"author_name": info.AuthorName,
		"code_age":    info.CodeAge,
		"article_num": info.ArticleNum,
		"fans_num":    info.FansNum,
		"like_num":    info.LikeNum,
		"comment_num": info.CommentNum,
		"collect_num": info.CollectNum,
		"share_num":   info.ShareNum,
		"visit_num":   info.VisitNum,
		"rank":        info.Rank,
		"level":       info.Level,
		"score":       info.Score,
	})
}

// GetQuarterStats 返回按年份和季度统计的文章数量。
func (h *Handler) GetQuarterStats(c *gin.Context) {
	data, err := h.articleService.QuarterStats()
	if err != nil {
		failure(c, http.StatusInternalServerError, "获取季度统计失败: "+err.Error())
		return
	}
	success(c, data)
}

// GetReadStats 返回各分类的阅读量与文章数量统计。
func (h *Handler) GetReadStats(c *gin.Context) {
	data, err := h.articleService.ReadStats()
	if err != nil {
		failure(c, http.StatusInternalServerError, "获取阅读量统计失败: "+err.Error())
		return
	}
	success(c, data)
}

// GetPieData 返回分类饼图所需的数据结构。
func (h *Handler) GetPieData(c *gin.Context) {
	items, err := h.categorizeRepo.ListAll()
	if err != nil {
		failure(c, http.StatusInternalServerError, "获取饼图数据失败: "+err.Error())
		return
	}

	result := make([]gin.H, 0, len(items))
	for _, item := range items {
		result = append(result, gin.H{
			"value": item.ArticleNum,
			"name":  item.Categorize,
		})
	}

	success(c, result)
}

// GetArticlesList 返回按条件筛选后的文章列表。
func (h *Handler) GetArticlesList(c *gin.Context) {
	// 查询参数名称完全复用 Python 版本，前端可以无缝切到 Go 后端。
	data, err := h.articleService.FilteredArticleList(
		c.Query("type"),
		c.Query("quarter"),
		c.Query("year"),
		c.Query("week"),
		c.Query("day"),
	)
	if err != nil {
		failure(c, http.StatusInternalServerError, "获取文章列表失败: "+err.Error())
		return
	}
	success(c, data)
}

// GetHeatmap 返回指定年份的文章发布热力图数据。
func (h *Handler) GetHeatmap(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		failure(c, http.StatusInternalServerError, "获取热力图数据失败: invalid year")
		return
	}

	data, err := h.articleService.Heatmap(year)
	if err != nil {
		failure(c, http.StatusInternalServerError, "获取热力图数据失败: "+err.Error())
		return
	}
	success(c, data)
}

// GetYears 返回数据库中存在的年份列表。
func (h *Handler) GetYears(c *gin.Context) {
	data, err := h.articleService.Years()
	if err != nil {
		failure(c, http.StatusInternalServerError, "获取年份列表失败: "+err.Error())
		return
	}
	success(c, data)
}

// CrawlInfo 触发一次用户信息抓取。
func (h *Handler) CrawlInfo(c *gin.Context) {
	successful, err := h.spider.CrawlInfo()
	if err != nil {
		failure(c, http.StatusInternalServerError, "抓取用户信息失败: "+err.Error())
		return
	}
	if !successful {
		failure(c, http.StatusInternalServerError, "抓取失败")
		return
	}
	crawlSuccess(c)
}

// CrawlCategorize 触发一次分类信息抓取。
func (h *Handler) CrawlCategorize(c *gin.Context) {
	if err := h.spider.CrawlCategorize(); err != nil {
		failure(c, http.StatusInternalServerError, "抓取分类信息失败: "+err.Error())
		return
	}
	crawlSuccess(c)
}

// CrawlArticles 触发一次文章信息抓取。
func (h *Handler) CrawlArticles(c *gin.Context) {
	if err := h.spider.CrawlArticles(); err != nil {
		failure(c, http.StatusInternalServerError, "抓取文章信息失败: "+err.Error())
		return
	}
	crawlSuccess(c)
}

// success 输出与 Python 版一致的成功 JSON。
func success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "操作成功",
		"data": data,
	})
}

// failure 输出与 Python 版一致的失败 JSON。
func failure(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"code": status,
		"msg":  message,
	})
}

// crawlSuccess 输出抓取接口专用的成功响应。
func crawlSuccess(c *gin.Context) {
	// 抓取接口单独使用“抓取成功”文案，以对齐原有 Flask 返回值。
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "抓取成功",
		"data": nil,
	})
}
