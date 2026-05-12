// 文件作用：文章领域服务：负责文章预处理、筛选、聚合和统计计算。
package service

import (
	"sort"
	"strconv"
	"time"

	dbpkg "csdn-analytics/backend-go/internal/db"
)

type ArticleRepository interface {
	ListAll() ([]dbpkg.Article, error)
}

// ArticleView 表示附带派生时间字段的文章视图对象。
type ArticleView struct {
	ID         uint
	URL        string
	Title      string
	Date       string
	ReadNum    int64
	CommentNum int64
	Type       string
	Weekday    int
	Year       int
	Month      int
	Quarter    string
	Week       int
}

// ArticleListItem 表示文章列表接口返回的精简对象。
type ArticleListItem struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Date  string `json:"date"`
	Type  string `json:"type"`
}

// HeatmapResult 表示热力图接口返回的数据结构。
type HeatmapResult struct {
	Data  [][]int  `json:"data"`
	XAxis []string `json:"xAxis"`
	YAxis []string `json:"yAxis"`
}

// ReadResult 表示阅读统计接口返回的数据结构。
type ReadResult struct {
	Labels []string `json:"labels"`
	Reads  []int64  `json:"reads"`
	Counts []int    `json:"counts"`
}

// ArticleService 负责文章相关的筛选、预处理和聚合统计。
type ArticleService struct {
	repo ArticleRepository
}

// NewArticleService 创建文章领域服务实例。
func NewArticleService(repo ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

// ProcessedArticles 将原始文章记录转换为带派生字段的视图列表。
func (s *ArticleService) ProcessedArticles() ([]ArticleView, error) {
	items, err := s.repo.ListAll()
	if err != nil {
		return nil, err
	}

	result := make([]ArticleView, 0, len(items))
	for _, item := range items {
		// Python 版本会直接按固定时间格式解析；这里保持相同约束。
		parsed, err := time.Parse("2006-01-02 15:04:05", item.Date)
		if err != nil {
			continue
		}
		_, isoWeek := parsed.ISOWeek()

		result = append(result, ArticleView{
			ID:         item.ID,
			URL:        item.URL,
			Title:      item.Title,
			Date:       item.Date,
			ReadNum:    item.ReadNum,
			CommentNum: item.CommentNum,
			Type:       item.Type,
			Weekday:    pythonWeekday(parsed),
			Year:       parsed.Year(),
			Month:      int(parsed.Month()),
			Quarter:    quarterFromMonth(int(parsed.Month())),
			Week:       isoWeek,
		})
	}

	return result, nil
}

// FilteredArticleList 按查询参数过滤文章并生成前端列表数据。
func (s *ArticleService) FilteredArticleList(filterType, filterQuarter, filterYear, filterWeek, filterDay string) ([]ArticleListItem, error) {
	articles, err := s.ProcessedArticles()
	if err != nil {
		return nil, err
	}

	filtered := make([]ArticleView, 0, len(articles))
	for _, article := range articles {
		if filterType != "" && article.Type != filterType {
			continue
		}
		if filterQuarter != "" && article.Quarter != filterQuarter {
			continue
		}
		if filterYear != "" && article.Year != atoiDefault(filterYear, -1) {
			continue
		}
		if filterWeek != "" && article.Week != atoiDefault(filterWeek, -1) {
			continue
		}
		if filterDay != "" && article.Weekday != atoiDefault(filterDay, -1)-1 {
			continue
		}
		filtered = append(filtered, article)
	}

	seen := make(map[string]struct{}, len(filtered))
	unique := make([]ArticleView, 0, len(filtered))
	for _, article := range filtered {
		// 与 Python 版本一致：按 URL 去重，保留第一次遇到的记录。
		if _, ok := seen[article.URL]; ok {
			continue
		}
		seen[article.URL] = struct{}{}
		unique = append(unique, article)
	}

	sort.Slice(unique, func(i, j int) bool {
		return unique[i].Date > unique[j].Date
	})

	// 前端列表只展示最近 100 条，保持与原后端一致。
	if len(unique) > 100 {
		unique = unique[:100]
	}

	result := make([]ArticleListItem, 0, len(unique))
	for _, article := range unique {
		result = append(result, ArticleListItem{
			Title: article.Title,
			URL:   article.URL,
			Date:  article.Date,
			Type:  article.Type,
		})
	}

	return result, nil
}

// QuarterStats 统计每年各季度的文章数量。
func (s *ArticleService) QuarterStats() ([]map[string]any, error) {
	articles, err := s.ProcessedArticles()
	if err != nil {
		return nil, err
	}

	byYear := map[int]map[string]int{}
	for _, article := range articles {
		if _, ok := byYear[article.Year]; !ok {
			byYear[article.Year] = map[string]int{}
		}
		byYear[article.Year][article.Quarter]++
	}

	years := make([]int, 0, len(byYear))
	for year := range byYear {
		years = append(years, year)
	}
	sort.Ints(years)

	result := make([]map[string]any, 0, len(years))
	for _, year := range years {
		row := map[string]any{"product": strconv.Itoa(year)}
		for _, quarter := range []string{"第一季度", "第二季度", "第三季度", "第四季度"} {
			row[quarter] = byYear[year][quarter]
		}
		result = append(result, row)
	}

	return result, nil
}

// ReadStats 统计每个分类的阅读量与文章数量。
func (s *ArticleService) ReadStats() (ReadResult, error) {
	articles, err := s.ProcessedArticles()
	if err != nil {
		return ReadResult{}, err
	}

	// stat 用于暂存单个分类的文章数量和累计阅读量。
	type stat struct {
		count int
		reads int64
	}

	byType := map[string]stat{}
	order := make([]string, 0)
	for _, article := range articles {
		current, ok := byType[article.Type]
		if !ok {
			order = append(order, article.Type)
		}
		current.count++
		current.reads += article.ReadNum
		byType[article.Type] = current
	}

	result := ReadResult{
		Labels: make([]string, 0, len(order)),
		Reads:  make([]int64, 0, len(order)),
		Counts: make([]int, 0, len(order)),
	}
	for _, label := range order {
		result.Labels = append(result.Labels, label)
		result.Reads = append(result.Reads, byType[label].reads)
		result.Counts = append(result.Counts, byType[label].count)
	}

	return result, nil
}

// Heatmap 统计指定年份的发文热力图矩阵。
func (s *ArticleService) Heatmap(year int) (HeatmapResult, error) {
	articles, err := s.ProcessedArticles()
	if err != nil {
		return HeatmapResult{}, err
	}

	heatmap := map[[2]int]int{}
	maxWeek := 0
	for _, article := range articles {
		if article.Year != year {
			continue
		}
		key := [2]int{article.Week, article.Weekday}
		heatmap[key]++
		if article.Week > maxWeek {
			maxWeek = article.Week
		}
	}

	data := make([][]int, 0)
	if maxWeek > 0 {
		// 热力图需要补齐缺失的周/星期组合，否则前端坐标轴会错位。
		for week := 1; week <= maxWeek; week++ {
			for weekday := 0; weekday < 7; weekday++ {
				data = append(data, []int{week - 1, weekday, heatmap[[2]int{week, weekday}]})
			}
		}
	}

	if maxWeek == 0 {
		maxWeek = 52
	}

	xAxis := make([]string, 0, maxWeek)
	for i := 1; i <= maxWeek; i++ {
		xAxis = append(xAxis, "第"+strconv.Itoa(i)+"周")
	}

	return HeatmapResult{
		Data:  data,
		XAxis: xAxis,
		YAxis: []string{"星期1", "星期2", "星期3", "星期4", "星期5", "星期6", "星期日"},
	}, nil
}

// Years 返回文章数据中出现过的年份列表。
func (s *ArticleService) Years() ([]string, error) {
	articles, err := s.ProcessedArticles()
	if err != nil {
		return nil, err
	}

	set := map[string]struct{}{}
	for _, article := range articles {
		set[strconv.Itoa(article.Year)] = struct{}{}
	}

	years := make([]string, 0, len(set))
	for year := range set {
		years = append(years, year)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(years)))
	return years, nil
}

// pythonWeekday 将 Go 的星期值转换为 Python weekday 语义。
func pythonWeekday(t time.Time) int {
	return (int(t.Weekday()) + 6) % 7
}

// quarterFromMonth 根据月份返回中文季度名称。
func quarterFromMonth(month int) string {
	switch {
	case month >= 1 && month <= 3:
		return "第一季度"
	case month >= 4 && month <= 6:
		return "第二季度"
	case month >= 7 && month <= 9:
		return "第三季度"
	default:
		return "第四季度"
	}
}

// atoiDefault 尝试解析整数，失败时返回给定默认值。
func atoiDefault(value string, fallback int) int {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
