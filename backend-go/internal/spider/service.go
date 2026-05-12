// 文件作用：抓取服务层：负责请求 CSDN 页面、解析数据并写入 SQLite。
package spider

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"regexp"
	"strings"
	"time"

	"csdn-analytics/backend-go/internal/config"
	dbpkg "csdn-analytics/backend-go/internal/db"
)

type InfoRepository interface {
	Upsert(dbpkg.Info) error
}

type CategorizeRepository interface {
	Upsert(dbpkg.Categorize) error
	ListAll() ([]dbpkg.Categorize, error)
}

type ArticleRepository interface {
	Upsert(dbpkg.Article) error
}

// Service 封装 CSDN 抓取流程及其依赖。
type Service struct {
	cfg            config.Config
	client         *http.Client
	infoRepo       InfoRepository
	categorizeRepo CategorizeRepository
	articleRepo    ArticleRepository
}

// NewService 创建抓取服务实例。
func NewService(cfg config.Config, infoRepo InfoRepository, categorizeRepo CategorizeRepository, articleRepo ArticleRepository) *Service {
	return &Service{
		cfg: cfg,
		client: &http.Client{
			Timeout: time.Duration(cfg.SpiderTimeout) * time.Second,
		},
		infoRepo:       infoRepo,
		categorizeRepo: categorizeRepo,
		articleRepo:    articleRepo,
	}
}

// CrawlInfo 抓取并写入用户信息。
func (s *Service) CrawlInfo() (bool, error) {
	// 用户主页就是信息抓取的主入口，先抓主页再解析用户统计信息。
	body, err := s.fetch(s.cfg.CSDNBlogURL)
	if err != nil {
		return false, err
	}

	parsed, err := ParseUserInfo(body)
	if err != nil {
		return false, err
	}

	info := dbpkg.Info{
		Date:       time.Now().Format("2006-01-02 15:04:05"),
		HeadImg:    parsed.HeadImg,
		AuthorName: parsed.AuthorName,
		CodeAge:    parsed.CodeAge,
		ArticleNum: parseInt(parsed.Statistics["原创"]),
		FansNum:    parseInt(parsed.Statistics["粉丝"]),
		VisitNum:   parseInt(parsed.Statistics["总访问量"]),
		LikeNum:    parseInt(parsed.Achievements["点赞"]),
		CommentNum: parseInt(parsed.Achievements["评论"]),
		CollectNum: parseInt(parsed.Achievements["收藏"]),
		ShareNum:   parseInt(parsed.Achievements["分享"]),
		Rank:       parseInt(parsed.Statistics["排名"]),
		Level:      parsed.Statistics["等级"],
		Score:      parseInt(parsed.Statistics["积分"]),
	}

	return true, s.infoRepo.Upsert(info)
}

// CrawlCategorize 抓取并写入分类信息。
func (s *Service) CrawlCategorize() error {
	body, err := s.fetch(s.cfg.CSDNBlogURL)
	if err != nil {
		return err
	}

	categories, err := ParseCategories(body)
	if err != nil {
		return err
	}

	for _, category := range categories {
		// 分类详情页抓取失败时退化为空字符串，保持与 Python 版“尽量继续”的策略一致。
		detailBody, err := s.fetch(category.Href)
		if err != nil {
			detailBody = ""
		}
		details, _ := ParseCategoryDetails(detailBody)

		if err := s.categorizeRepo.Upsert(dbpkg.Categorize{
			Href:         category.Href,
			Categorize:   category.Categorize,
			CategorizeID: int64(parseInt(category.CategorizeID)),
			ColumnNum:    int64(category.ColumnNum),
			NumSpan:      int64(details.SubscribeNum),
			ArticleNum:   int64(details.ArticleNum),
			ReadNum:      int64(details.ReadNum),
			CollectNum:   int64(details.CollectNum),
		}); err != nil {
			return err
		}
	}

	return nil
}

// CrawlArticles 抓取并写入文章信息。
func (s *Service) CrawlArticles() error {
	categories, err := s.categorizeRepo.ListAll()
	if err != nil {
		return err
	}

	for _, category := range categories {
		// 大于 40 篇文章的专栏会分页，分页 URL 规则与 Python 版本保持一致。
		pageURLs := buildArticlePageURLs(category.Href, int(category.ArticleNum))
		if len(pageURLs) == 0 {
			pageURLs = []string{category.Href}
		}

		for _, pageURL := range pageURLs {
			body, err := s.fetch(pageURL)
			if err != nil {
				continue
			}

			articles, err := ParseArticles(body, category.Categorize)
			if err != nil {
				continue
			}

			for _, article := range articles {
				if err := s.articleRepo.Upsert(dbpkg.Article{
					URL:        article.URL,
					Title:      article.Title,
					Date:       article.Date,
					ReadNum:    int64(article.ReadNum),
					CommentNum: int64(article.CommentNum),
					Type:       article.Type,
				}); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// fetch 带重试地请求目标页面并返回响应内容。
func (s *Service) fetch(url string) (string, error) {
	var lastErr error
	for attempt := 0; attempt < s.cfg.SpiderRetryTimes; attempt++ {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (MSIE 10.0; Windows NT 6.1; Trident/5.0)")
		req.Header.Set("referer", "https://passport.csdn.net/login")

		resp, err := s.client.Do(req)
		if err == nil && resp != nil {
			defer resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				body, readErr := io.ReadAll(resp.Body)
				if readErr != nil {
					return "", readErr
				}
				return string(body), nil
			}
			lastErr = fmt.Errorf("unexpected status: %d", resp.StatusCode)
		} else {
			lastErr = err
		}

		// 失败后按配置等待再重试，避免瞬时网络抖动直接导致整次抓取失败。
		if attempt < s.cfg.SpiderRetryTimes-1 {
			time.Sleep(time.Duration(s.cfg.SpiderRetryDelay) * time.Second)
		}
	}

	return "", lastErr
}

// buildArticlePageURLs 根据文章数量推导专栏分页地址列表。
func buildArticlePageURLs(baseURL string, articleNum int) []string {
	if articleNum <= 40 {
		return []string{baseURL}
	}

	pageNum := int(math.Round(float64(articleNum) / 40.0))
	base := strings.TrimSuffix(baseURL, ".html")
	urls := make([]string, 0, pageNum+1)
	for i := pageNum; i > 0; i-- {
		urls = append(urls, fmt.Sprintf("%s_%d.html", base, i))
	}
	urls = append(urls, baseURL)
	return urls
}

// parseInt 提取字符串中的数字并转换为整数。
func parseInt(input string) int {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(strings.ReplaceAll(input, ",", ""))
	if match == "" {
		return 0
	}
	var result int
	fmt.Sscanf(match, "%d", &result)
	return result
}
