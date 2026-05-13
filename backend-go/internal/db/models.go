// 文件作用：数据库模型定义：声明与 Python 版一致的 info、categorize、article 表结构。
package db

// Info 对应用户信息表。
type Info struct {
	// ID 是主键。
	ID uint `gorm:"column:id;primaryKey;autoIncrement"`
	// Date 是最近一次抓取时间。
	Date string `gorm:"column:date"`
	// HeadImg 是头像地址。
	HeadImg string `gorm:"column:head_img"`
	// AuthorName 是作者名称。
	AuthorName string `gorm:"column:author_name"`
	// CodeAge 是码龄文案。
	CodeAge string `gorm:"column:code_age"`
	// ArticleNum 是原创文章数量。
	ArticleNum int `gorm:"column:article_num"`
	// FansNum 是粉丝数量。
	FansNum int `gorm:"column:fans_num"`
	// LikeNum 是点赞数量。
	LikeNum int `gorm:"column:like_num"`
	// CommentNum 是评论数量。
	CommentNum int `gorm:"column:comment_num"`
	// CollectNum 是收藏数量。
	CollectNum int `gorm:"column:collect_num"`
	// ShareNum 是分享数量。
	ShareNum int `gorm:"column:share_num"`
	// VisitNum 是总访问量。
	VisitNum int `gorm:"column:visit_num"`
	// Rank 是博主排名。
	Rank int `gorm:"column:rank"`
	// Level 是博主等级文案。
	Level string `gorm:"column:level"`
	// Score 是博主积分。
	Score int `gorm:"column:score"`
}

// TableName 返回 Info 模型绑定的数据库表名。
func (Info) TableName() string {
	return "info"
}

// Categorize 对应博客分类信息表。
type Categorize struct {
	// ID 是主键。
	ID uint `gorm:"column:id;primaryKey;autoIncrement"`
	// Href 是分类详情页链接。
	Href string `gorm:"column:href"`
	// Categorize 是分类名称。
	Categorize string `gorm:"column:categorize"`
	// CategorizeID 是分类 ID。
	CategorizeID int64 `gorm:"column:categorize_id"`
	// ColumnNum 是分类入口显示的专栏文章数。
	ColumnNum int64 `gorm:"column:column_num"`
	// NumSpan 是订阅数量。
	NumSpan int64 `gorm:"column:num_span"`
	// ArticleNum 是详情页统计的文章数量。
	ArticleNum int64 `gorm:"column:article_num"`
	// ReadNum 是详情页统计的总阅读量。
	ReadNum int64 `gorm:"column:read_num"`
	// CollectNum 是详情页统计的总收藏量。
	CollectNum int64 `gorm:"column:collect_num"`
}

// TableName 返回 Categorize 模型绑定的数据库表名。
func (Categorize) TableName() string {
	return "categorize"
}

// Article 对应文章信息表。
type Article struct {
	// ID 是主键。
	ID uint `gorm:"column:id;primaryKey;autoIncrement"`
	// URL 是文章链接。
	URL string `gorm:"column:url"`
	// Title 是文章标题。
	Title string `gorm:"column:title"`
	// Date 是发布时间字符串。
	Date string `gorm:"column:date"`
	// ReadNum 是阅读量。
	ReadNum int64 `gorm:"column:read_num"`
	// CommentNum 是评论数。
	CommentNum int64 `gorm:"column:comment_num"`
	// Type 是文章所属分类。
	Type string `gorm:"column:type"`
}

// TableName 返回 Article 模型绑定的数据库表名。
func (Article) TableName() string {
	return "article"
}
