// 文件作用：Article 仓储：封装 article 表的列表查询与按 url 更新写入。
package repository

import (
	dbpkg "csdn-analytics/backend-go/internal/db"
	"gorm.io/gorm"
)

// ArticleRepository 封装对 article 表的数据库访问。
type ArticleRepository struct {
	db *gorm.DB
}

// NewArticleRepository 创建 article 表仓储实例。
func NewArticleRepository(gdb *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: gdb}
}

// ListAll 读取全部文章记录。
func (r *ArticleRepository) ListAll() ([]dbpkg.Article, error) {
	var items []dbpkg.Article
	err := r.db.Order("id asc").Find(&items).Error
	return items, err
}

// Upsert 按 url 更新或新增文章记录。
func (r *ArticleRepository) Upsert(item dbpkg.Article) error {
	var existing dbpkg.Article
	err := r.db.Where("url = ?", item.URL).First(&existing).Error
	if err == nil {
		item.ID = existing.ID
		return r.db.Model(&existing).Updates(item).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return r.db.Create(&item).Error
}
