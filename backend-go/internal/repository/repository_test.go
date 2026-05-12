// 文件作用：仓储测试：验证三类仓储的 upsert 和查询行为。
package repository

import (
	"path/filepath"
	"testing"

	dbpkg "csdn-analytics/backend-go/internal/db"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestInfoRepositoryUpsertByAuthorName 验证 info 仓储会按 author_name 执行 upsert。
func TestInfoRepositoryUpsertByAuthorName(t *testing.T) {
	gdb := openTestDB(t)
	repo := NewInfoRepository(gdb)

	require.NoError(t, repo.Upsert(dbpkg.Info{
		AuthorName: "heian_99",
		Date:       "2026-05-12 12:00:00",
		VisitNum:   10,
	}))
	require.NoError(t, repo.Upsert(dbpkg.Info{
		AuthorName: "heian_99",
		Date:       "2026-05-12 13:00:00",
		VisitNum:   99,
	}))

	info, err := repo.First()
	require.NoError(t, err)
	require.Equal(t, "2026-05-12 13:00:00", info.Date)
	require.Equal(t, 99, info.VisitNum)
}

// TestCategorizeRepositoryUpsertByHref 验证 categorize 仓储会按 href 执行 upsert。
func TestCategorizeRepositoryUpsertByHref(t *testing.T) {
	gdb := openTestDB(t)
	repo := NewCategorizeRepository(gdb)

	require.NoError(t, repo.Upsert(dbpkg.Categorize{
		Href:       "https://example.com/c1",
		Categorize: "Go",
		ArticleNum: 3,
	}))
	require.NoError(t, repo.Upsert(dbpkg.Categorize{
		Href:       "https://example.com/c1",
		Categorize: "Go",
		ArticleNum: 5,
	}))

	items, err := repo.ListAll()
	require.NoError(t, err)
	require.Len(t, items, 1)
	require.Equal(t, int64(5), items[0].ArticleNum)
}

// TestArticleRepositoryUpsertByURLAndListAll 验证 article 仓储会按 url 执行 upsert。
func TestArticleRepositoryUpsertByURLAndListAll(t *testing.T) {
	gdb := openTestDB(t)
	repo := NewArticleRepository(gdb)

	require.NoError(t, repo.Upsert(dbpkg.Article{
		URL:        "https://example.com/a1",
		Title:      "title_1",
		Date:       "2024-01-02 03:04:05",
		ReadNum:    1,
		CommentNum: 2,
		Type:       "Go",
	}))
	require.NoError(t, repo.Upsert(dbpkg.Article{
		URL:        "https://example.com/a1",
		Title:      "title_2",
		Date:       "2024-01-02 03:04:05",
		ReadNum:    9,
		CommentNum: 8,
		Type:       "Go",
	}))

	items, err := repo.ListAll()
	require.NoError(t, err)
	require.Len(t, items, 1)
	require.Equal(t, "title_2", items[0].Title)
	require.Equal(t, int64(9), items[0].ReadNum)
}

// openTestDB 创建仓储测试使用的临时数据库。
func openTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	dbPath := filepath.Join(t.TempDir(), "repo.db")
	gdb, err := dbpkg.Open("sqlite:///" + filepath.ToSlash(dbPath))
	require.NoError(t, err)
	require.NoError(t, dbpkg.AutoMigrate(gdb))

	return gdb
}
