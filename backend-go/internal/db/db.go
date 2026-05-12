// 文件作用：数据库入口层：负责规范化 SQLite DSN、打开连接并执行自动建表。
package db

import (
	"errors"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// Open 根据配置打开 SQLite 数据库连接。
func Open(databaseURL string) (*gorm.DB, error) {
	dsn := NormalizeSQLiteDSN(databaseURL)
	if strings.TrimSpace(dsn) == "" {
		return nil, errors.New("database dsn cannot be empty")
	}

	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}

// NormalizeSQLiteDSN 将 Python 风格的 SQLite 地址转换为 Go 驱动可识别的 DSN。
func NormalizeSQLiteDSN(databaseURL string) string {
	const sqlitePrefix = "sqlite:///"

	if strings.HasPrefix(databaseURL, sqlitePrefix) {
		return strings.TrimPrefix(databaseURL, sqlitePrefix)
	}

	return databaseURL
}

// AutoMigrate 自动创建并同步项目需要的数据库表。
func AutoMigrate(gdb *gorm.DB) error {
	return gdb.AutoMigrate(&Info{}, &Categorize{}, &Article{})
}
