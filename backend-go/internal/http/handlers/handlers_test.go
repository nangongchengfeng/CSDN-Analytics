// 文件作用：读接口测试：验证 info、articles、quarter、read、heatmap、years 等契约。
package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"csdn-analytics/backend-go/internal/config"
	dbpkg "csdn-analytics/backend-go/internal/db"
	httpapi "csdn-analytics/backend-go/internal/http"
	"csdn-analytics/backend-go/internal/repository"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestReadEndpointsPreservePythonContract 验证读接口响应契约与 Python 版本一致。
func TestReadEndpointsPreservePythonContract(t *testing.T) {
	gdb := openHTTPTestDB(t)
	seedHTTPData(t, gdb)

	router := httpapi.NewRouter(config.LoadFromEnv(func(string) string { return "" }), httpapi.Dependencies{
		InfoRepo:       repository.NewInfoRepository(gdb),
		CategorizeRepo: repository.NewCategorizeRepository(gdb),
		ArticleRepo:    repository.NewArticleRepository(gdb),
	})

	t.Run("GET /api/info", func(t *testing.T) {
		resp := performRequest(router, http.MethodGet, "/api/info")
		require.Equal(t, http.StatusOK, resp.Code)

		var body map[string]any
		require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &body))
		require.Equal(t, float64(200), body["code"])
		require.Equal(t, "操作成功", body["msg"])
		data := body["data"].(map[string]any)
		require.Equal(t, "heian_99", data["author_name"])
	})

	t.Run("GET /api/articles filters and dedupe", func(t *testing.T) {
		resp := performRequest(router, http.MethodGet, "/api/articles?type=Go&year=2024&quarter=第一季度&week=1&day=1")
		require.Equal(t, http.StatusOK, resp.Code)

		var body map[string]any
		require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &body))
		items := body["data"].([]any)
		require.Len(t, items, 1)
		item := items[0].(map[string]any)
		require.Equal(t, "Go_intro_updated", item["title"])
		require.Equal(t, "Go", item["type"])
	})

	t.Run("GET /api/quarter", func(t *testing.T) {
		resp := performRequest(router, http.MethodGet, "/api/quarter")
		require.Equal(t, http.StatusOK, resp.Code)

		var body map[string]any
		require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &body))
		items := body["data"].([]any)
		require.NotEmpty(t, items)
		first := items[0].(map[string]any)
		require.Contains(t, first, "product")
		require.Contains(t, first, "第一季度")
	})

	t.Run("GET /api/read", func(t *testing.T) {
		resp := performRequest(router, http.MethodGet, "/api/read")
		require.Equal(t, http.StatusOK, resp.Code)

		var body map[string]any
		require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &body))
		data := body["data"].(map[string]any)
		require.Contains(t, data, "labels")
		require.Contains(t, data, "reads")
		require.Contains(t, data, "counts")
	})

	t.Run("GET /api/categorize and /api/pie", func(t *testing.T) {
		for _, path := range []string{"/api/categorize", "/api/pie"} {
			resp := performRequest(router, http.MethodGet, path)
			require.Equal(t, http.StatusOK, resp.Code)

			var body map[string]any
			require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &body))
			items := body["data"].([]any)
			require.Len(t, items, 2)
		}
	})

	t.Run("GET /api/heatmap/2024", func(t *testing.T) {
		resp := performRequest(router, http.MethodGet, "/api/heatmap/2024")
		require.Equal(t, http.StatusOK, resp.Code)

		var body map[string]any
		require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &body))
		data := body["data"].(map[string]any)
		require.Contains(t, data, "data")
		require.Contains(t, data, "xAxis")
		require.Contains(t, data, "yAxis")
		require.Equal(t, "星期日", data["yAxis"].([]any)[6])
	})

	t.Run("GET /api/years", func(t *testing.T) {
		resp := performRequest(router, http.MethodGet, "/api/years")
		require.Equal(t, http.StatusOK, resp.Code)

		var body map[string]any
		require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &body))
		years := body["data"].([]any)
		require.Equal(t, "2024", years[0])
	})
}

// performRequest 用于向测试路由发送一次 HTTP 请求。
func performRequest(handler http.Handler, method, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	return resp
}

// openHTTPTestDB 创建测试专用的临时 SQLite 数据库。
func openHTTPTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	dbPath := filepath.Join(t.TempDir(), "http.db")
	gdb, err := dbpkg.Open("sqlite:///" + filepath.ToSlash(dbPath))
	require.NoError(t, err)
	require.NoError(t, dbpkg.AutoMigrate(gdb))
	return gdb
}

// seedHTTPData 向测试数据库写入接口断言所需的样例数据。
func seedHTTPData(t *testing.T, gdb *gorm.DB) {
	t.Helper()

	require.NoError(t, repository.NewInfoRepository(gdb).Upsert(dbpkg.Info{
		AuthorName: "heian_99",
		Date:       "2026-05-12 12:00:00",
		VisitNum:   123,
	}))

	catRepo := repository.NewCategorizeRepository(gdb)
	require.NoError(t, catRepo.Upsert(dbpkg.Categorize{Href: "https://example.com/go", Categorize: "Go", ArticleNum: 2}))
	require.NoError(t, catRepo.Upsert(dbpkg.Categorize{Href: "https://example.com/python", Categorize: "Python", ArticleNum: 1}))

	articleRepo := repository.NewArticleRepository(gdb)
	require.NoError(t, articleRepo.Upsert(dbpkg.Article{
		URL:        "https://example.com/a1",
		Title:      "Go_intro",
		Date:       "2024-01-01 10:00:00",
		ReadNum:    10,
		CommentNum: 1,
		Type:       "Go",
	}))
	require.NoError(t, articleRepo.Upsert(dbpkg.Article{
		URL:        "https://example.com/a1",
		Title:      "Go_intro_updated",
		Date:       "2024-01-01 10:00:00",
		ReadNum:    11,
		CommentNum: 2,
		Type:       "Go",
	}))
	require.NoError(t, articleRepo.Upsert(dbpkg.Article{
		URL:        "https://example.com/a2",
		Title:      "Python_intro",
		Date:       "2024-04-02 10:00:00",
		ReadNum:    20,
		CommentNum: 3,
		Type:       "Python",
	}))
}
