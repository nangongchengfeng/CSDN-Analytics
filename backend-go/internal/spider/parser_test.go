// 文件作用：解析测试：验证 HTML 解析结果与 Python 版字段语义一致。
package spider

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParseUserInfoExtractsPythonCompatibleFields 验证用户信息解析字段与 Python 语义一致。
func TestParseUserInfoExtractsPythonCompatibleFields(t *testing.T) {
	html := `
<div class="user-profile-head-info-r-c">
  <li><div class="user-profile-statistics-name">原创</div><div class="user-profile-statistics-num">12</div></li>
  <li><div class="user-profile-statistics-name">粉丝</div><div class="user-profile-statistics-num">34</div></li>
  <li><div class="user-profile-statistics-name">总访问量</div><div class="user-profile-statistics-num">56</div></li>
  <li><div class="user-profile-statistics-name">排名</div><div class="user-profile-statistics-num">78</div></li>
  <li><div class="user-profile-statistics-name">等级</div><div class="user-profile-statistics-num">L6</div></li>
  <li><div class="user-profile-statistics-name">积分</div><div class="user-profile-statistics-num">90</div></li>
</div>
<ul class="aside-common-box-achievement">
  <li><div><span>11</span>点赞</div></li>
  <li><div><span>22</span>评论</div></li>
  <li><div><span>33</span>收藏</div></li>
  <li><div><span>44</span>分享</div></li>
</ul>
<div class="person-code-age"><span>5年</span></div>
<div class="user-profile-head-name"><div>heian_99</div></div>
<div class="user-profile-avatar"><img src="https://img.example/avatar.png"></div>`

	parsed, err := ParseUserInfo(html)
	require.NoError(t, err)
	require.Equal(t, "heian_99", parsed.AuthorName)
	require.Equal(t, "5年", parsed.CodeAge)
	require.Equal(t, "https://img.example/avatar.png", parsed.HeadImg)
	require.Equal(t, "12", parsed.Statistics["原创"])
	require.Equal(t, "11", parsed.Achievements["点赞"])
}

// TestParseCategoriesAndDetails 验证分类列表和详情页解析逻辑。
func TestParseCategoriesAndDetails(t *testing.T) {
	categoryHTML := `
<a class="special-column-name" href="https://blog.csdn.net/category_123.html">
  Go专栏
  <span class="special-column-num">40篇</span>
</a>`
	detailHTML := `
<div class="column_operating">
  <span class="column-subscribe-num">订阅 18</span>
  <span class="mumber-color">0</span>
  <span class="mumber-color">25</span>
  <span class="mumber-color">300</span>
  <span class="mumber-color">9</span>
</div>`

	categories, err := ParseCategories(categoryHTML)
	require.NoError(t, err)
	require.Len(t, categories, 1)
	require.Equal(t, "Go专栏40篇", categories[0].Categorize)
	require.Equal(t, "123", categories[0].CategorizeID)
	require.Equal(t, 40, categories[0].ColumnNum)

	details, err := ParseCategoryDetails(detailHTML)
	require.NoError(t, err)
	require.Equal(t, 18, details.SubscribeNum)
	require.Equal(t, 25, details.ArticleNum)
	require.Equal(t, 300, details.ReadNum)
	require.Equal(t, 9, details.CollectNum)
}

// TestParseArticles 验证文章列表解析逻辑。
func TestParseArticles(t *testing.T) {
	html := `
<ul class="column_article_list">
  <li>
    <a target="_blank" href="https://example.com/a1"></a>
    <h2 class="title">Go intro</h2>
    <span class="status">2024-01-01 10:00:00 · 发布</span>
    <span class="status">阅读数 10</span>
    <span class="status">评论数 2</span>
  </li>
</ul>`

	articles, err := ParseArticles(html, "Go")
	require.NoError(t, err)
	require.Len(t, articles, 1)
	require.Equal(t, "https://example.com/a1", articles[0].URL)
	require.Equal(t, "Go_intro", articles[0].Title)
	require.Equal(t, "2024-01-01 10:00:00", articles[0].Date)
	require.Equal(t, 10, articles[0].ReadNum)
	require.Equal(t, 2, articles[0].CommentNum)
	require.Equal(t, "Go", articles[0].Type)
}
