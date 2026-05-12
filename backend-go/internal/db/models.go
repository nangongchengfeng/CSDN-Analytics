// 文件作用：数据库模型定义：声明与 Python 版一致的 info、categorize、article 表结构。
package db

// Info 对应用户信息表。
type Info struct {
	ID         uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Date       string `gorm:"column:date"`
	HeadImg    string `gorm:"column:head_img"`
	AuthorName string `gorm:"column:author_name"`
	CodeAge    string `gorm:"column:code_age"`
	ArticleNum int    `gorm:"column:article_num"`
	FansNum    int    `gorm:"column:fans_num"`
	LikeNum    int    `gorm:"column:like_num"`
	CommentNum int    `gorm:"column:comment_num"`
	CollectNum int    `gorm:"column:collect_num"`
	ShareNum   int    `gorm:"column:share_num"`
	VisitNum   int    `gorm:"column:visit_num"`
	Rank       int    `gorm:"column:rank"`
	Level      string `gorm:"column:level"`
	Score      int    `gorm:"column:score"`
}

// TableName 返回 Info 模型绑定的数据库表名。
func (Info) TableName() string {
	return "info"
}

// Categorize 对应博客分类信息表。
type Categorize struct {
	ID           uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Href         string `gorm:"column:href"`
	Categorize   string `gorm:"column:categorize"`
	CategorizeID int64  `gorm:"column:categorize_id"`
	ColumnNum    int64  `gorm:"column:column_num"`
	NumSpan      int64  `gorm:"column:num_span"`
	ArticleNum   int64  `gorm:"column:article_num"`
	ReadNum      int64  `gorm:"column:read_num"`
	CollectNum   int64  `gorm:"column:collect_num"`
}

// TableName 返回 Categorize 模型绑定的数据库表名。
func (Categorize) TableName() string {
	return "categorize"
}

// Article 对应文章信息表。
type Article struct {
	ID         uint   `gorm:"column:id;primaryKey;autoIncrement"`
	URL        string `gorm:"column:url"`
	Title      string `gorm:"column:title"`
	Date       string `gorm:"column:date"`
	ReadNum    int64  `gorm:"column:read_num"`
	CommentNum int64  `gorm:"column:comment_num"`
	Type       string `gorm:"column:type"`
}

// TableName 返回 Article 模型绑定的数据库表名。
func (Article) TableName() string {
	return "article"
}
