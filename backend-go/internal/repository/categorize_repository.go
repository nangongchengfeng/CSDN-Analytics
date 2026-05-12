// 文件作用：Categorize 仓储：封装 categorize 表的列表查询与按 href 更新写入。
package repository

import (
	dbpkg "csdn-analytics/backend-go/internal/db"
	"gorm.io/gorm"
)

// CategorizeRepository 封装对 categorize 表的数据库访问。
type CategorizeRepository struct {
	db *gorm.DB
}

// NewCategorizeRepository 创建 categorize 表仓储实例。
func NewCategorizeRepository(gdb *gorm.DB) *CategorizeRepository {
	return &CategorizeRepository{db: gdb}
}

// ListAll 读取全部分类记录。
func (r *CategorizeRepository) ListAll() ([]dbpkg.Categorize, error) {
	var items []dbpkg.Categorize
	err := r.db.Order("id asc").Find(&items).Error
	return items, err
}

// Upsert 按 href 更新或新增分类记录。
func (r *CategorizeRepository) Upsert(item dbpkg.Categorize) error {
	var existing dbpkg.Categorize
	err := r.db.Where("href = ?", item.Href).First(&existing).Error
	if err == nil {
		item.ID = existing.ID
		return r.db.Model(&existing).Updates(item).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return r.db.Create(&item).Error
}
