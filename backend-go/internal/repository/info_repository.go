// 文件作用：Info 仓储：封装 info 表的读取与按 author_name 更新写入。
package repository

import (
	dbpkg "csdn-analytics/backend-go/internal/db"
	"gorm.io/gorm"
)

// InfoRepository 封装对 info 表的数据库访问。
type InfoRepository struct {
	db *gorm.DB
}

// NewInfoRepository 创建 info 表仓储实例。
func NewInfoRepository(gdb *gorm.DB) *InfoRepository {
	return &InfoRepository{db: gdb}
}

// First 按主键顺序读取第一条用户信息记录。
func (r *InfoRepository) First() (dbpkg.Info, error) {
	var item dbpkg.Info
	err := r.db.Order("id asc").First(&item).Error
	return item, err
}

// Upsert 按 author_name 更新或新增用户信息。
func (r *InfoRepository) Upsert(info dbpkg.Info) error {
	var existing dbpkg.Info
	err := r.db.Where("author_name = ?", info.AuthorName).First(&existing).Error
	if err == nil {
		info.ID = existing.ID
		return r.db.Model(&existing).Updates(info).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return r.db.Create(&info).Error
}
