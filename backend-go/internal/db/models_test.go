// 文件作用：数据库测试：验证 SQLite DSN 处理和表名兼容性。
package db

import (
	"path/filepath"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestNormalizeSQLiteDSN 验证 SQLite DSN 会被正确规范化。
func TestNormalizeSQLiteDSN(t *testing.T) {
	normalized := NormalizeSQLiteDSN("sqlite:///csdn_analytics.db")
	require.Equal(t, filepath.ToSlash("csdn_analytics.db"), filepath.ToSlash(normalized))
}

// TestAutoMigrateUsesPythonCompatibleTableNames 验证建表结果与 Python 版本表名一致。
func TestAutoMigrateUsesPythonCompatibleTableNames(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "test.db")

	gdb, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	require.NoError(t, err)

	require.NoError(t, AutoMigrate(gdb))
	require.True(t, tableExists(t, gdb, "info"))
	require.True(t, tableExists(t, gdb, "categorize"))
	require.True(t, tableExists(t, gdb, "article"))
}

// tableExists 用于判断指定表是否已经被成功创建。
func tableExists(t *testing.T, gdb *gorm.DB, tableName string) bool {
	t.Helper()

	return gdb.Migrator().HasTable(tableName)
}
