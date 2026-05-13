// 文件作用：HTML 解析层：从 CSDN 页面中提取用户、分类和文章结构化数据。
package spider

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParsedUserInfo 表示从用户主页中提取出的用户信息结构。
type ParsedUserInfo struct {
	// Statistics 保存原创、粉丝、总访问量等统计文案。
	Statistics map[string]string
	// Achievements 保存点赞、评论、收藏、分享等成就数据。
	Achievements map[string]string
	// CodeAge 是码龄文本。
	CodeAge string
	// AuthorName 是作者名称。
	AuthorName string
	// HeadImg 是头像链接。
	HeadImg string
}

// ParsedCategory 表示从主页中提取出的专栏分类结构。
type ParsedCategory struct {
	// Href 是分类详情页链接。
	Href string
	// Categorize 是分类名称。
	Categorize string
	// CategorizeID 是从链接中提取出的分类 ID。
	CategorizeID string
	// ColumnNum 是入口处展示的文章数。
	ColumnNum int
}

// ParsedCategoryDetails 表示分类详情页中的统计结构。
type ParsedCategoryDetails struct {
	// SubscribeNum 是订阅人数。
	SubscribeNum int
	// ArticleNum 是文章数量。
	ArticleNum int
	// ReadNum 是总阅读量。
	ReadNum int
	// CollectNum 是总收藏量。
	CollectNum int
}

// ParsedArticle 表示从专栏页面提取出的文章结构。
type ParsedArticle struct {
	// URL 是文章链接。
	URL string
	// Title 是文章标题。
	Title string
	// Date 是发布时间。
	Date string
	// ReadNum 是阅读量。
	ReadNum int
	// CommentNum 是评论数。
	CommentNum int
	// Type 是所属分类名称。
	Type string
}

// ParseUserInfo 从用户主页 HTML 中提取用户统计信息。
func ParseUserInfo(html string) (ParsedUserInfo, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ParsedUserInfo{}, err
	}

	container := doc.Find("div.user-profile-head-info-r-c")
	if container.Length() == 0 {
		return ParsedUserInfo{}, errors.New("没有找到用户信息数据源，请检查CSDN的HTML页面结构。")
	}

	statistics := map[string]string{}
	container.Find("li").Each(func(i int, s *goquery.Selection) {
		label := strings.TrimSpace(s.Find("div.user-profile-statistics-name").Text())
		value := strings.TrimSpace(s.Find("div.user-profile-statistics-num").Text())
		if label != "" && value != "" {
			statistics[label] = value
		}
	})

	achievements := map[string]string{}
	doc.Find("ul.aside-common-box-achievement li div").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		value := strings.TrimSpace(s.Find("span").First().Text())
		for _, keyword := range []string{"点赞", "评论", "收藏", "分享"} {
			if strings.Contains(text, keyword) {
				achievements[keyword] = value
			}
		}
	})

	return ParsedUserInfo{
		Statistics:   statistics,
		Achievements: achievements,
		CodeAge:      strings.TrimSpace(doc.Find("div.person-code-age span").First().Text()),
		AuthorName:   strings.TrimSpace(doc.Find("div.user-profile-head-name div").First().Text()),
		HeadImg:      strings.TrimSpace(attrOrEmpty(doc.Find("div.user-profile-avatar img").First(), "src")),
	}, nil
}

// ParseCategories 从主页 HTML 中提取分类列表。
func ParseCategories(html string) ([]ParsedCategory, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	result := make([]ParsedCategory, 0)
	doc.Find("a.special-column-name").Each(func(i int, s *goquery.Selection) {
		href := strings.TrimSpace(attrOrEmpty(s, "href"))
		if href == "" {
			return
		}

		text := compactText(s.Text())
		columnNum := firstInt(strings.TrimSpace(s.Find("span.special-column-num").Text()))
		categorizeID := ""
		parts := strings.Split(href, "_")
		if len(parts) > 0 {
			last := parts[len(parts)-1]
			categorizeID = strings.TrimSuffix(last, ".html")
		}

		result = append(result, ParsedCategory{
			Href:         href,
			Categorize:   text,
			CategorizeID: categorizeID,
			ColumnNum:    columnNum,
		})
	})

	return result, nil
}

// ParseCategoryDetails 从分类详情页 HTML 中提取统计信息。
func ParseCategoryDetails(html string) (ParsedCategoryDetails, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ParsedCategoryDetails{}, err
	}

	container := doc.Find("div.column_operating").First()
	if container.Length() == 0 {
		return ParsedCategoryDetails{}, nil
	}

	nums := make([]int, 0)
	container.Find("span.mumber-color").Each(func(i int, s *goquery.Selection) {
		nums = append(nums, firstInt(s.Text()))
	})

	details := ParsedCategoryDetails{
		SubscribeNum: firstInt(container.Find("span.column-subscribe-num").First().Text()),
	}
	if len(nums) > 1 {
		details.ArticleNum = nums[1]
	}
	if len(nums) > 2 {
		details.ReadNum = nums[2]
	}
	if len(nums) > 3 {
		details.CollectNum = nums[3]
	}

	return details, nil
}

// ParseArticles 从文章列表页 HTML 中提取文章数据。
func ParseArticles(html string, categoryName string) ([]ParsedArticle, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	result := make([]ParsedArticle, 0)
	doc.Find("ul.column_article_list li").Each(func(i int, s *goquery.Selection) {
		url := strings.TrimSpace(attrOrEmpty(s.Find(`a[target="_blank"]`).First(), "href"))
		if url == "" {
			return
		}

		statuses := s.Find("span.status")
		if statuses.Length() < 3 {
			return
		}

		title := strings.ReplaceAll(strings.TrimSpace(s.Find("h2.title").First().Text()), " ", "_")
		title = strings.ReplaceAll(title, "/", "_")

		result = append(result, ParsedArticle{
			URL:        url,
			Title:      title,
			Date:       strings.TrimSpace(strings.Split(statuses.Eq(0).Text(), "·")[0]),
			ReadNum:    firstInt(statuses.Eq(1).Text()),
			CommentNum: firstInt(statuses.Eq(2).Text()),
			Type:       categoryName,
		})
	})

	return result, nil
}

// attrOrEmpty 安全读取节点属性，不存在时返回空字符串。
func attrOrEmpty(s *goquery.Selection, key string) string {
	value, _ := s.Attr(key)
	return value
}

// firstInt 提取字符串中的第一个整数值。
func firstInt(input string) int {
	re := regexp.MustCompile(`\d+`)
	matched := re.FindString(input)
	if matched == "" {
		return 0
	}
	value, err := strconv.Atoi(matched)
	if err != nil {
		return 0
	}
	return value
}

// compactText 压缩多余空白并拼接文本内容。
func compactText(input string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(input)), "")
}
