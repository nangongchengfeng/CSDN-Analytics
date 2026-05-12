// 文件作用：抓取接口测试：验证 /api/*/crawl 的成功与失败响应契约。
package handlers_test

import (
	"errors"
	"net/http"
	"testing"

	"csdn-analytics/backend-go/internal/config"
	httpapi "csdn-analytics/backend-go/internal/http"
	"csdn-analytics/backend-go/internal/repository"
	"github.com/stretchr/testify/require"
)

// TestCrawlEndpointsPreservePythonContract 验证抓取接口的成功与失败响应契约。
func TestCrawlEndpointsPreservePythonContract(t *testing.T) {
	gdb := openHTTPTestDB(t)
	seedHTTPData(t, gdb)

	t.Run("success", func(t *testing.T) {
		router := httpapi.NewRouter(config.LoadFromEnv(func(string) string { return "" }), httpapi.Dependencies{
			InfoRepo:       repository.NewInfoRepository(gdb),
			CategorizeRepo: repository.NewCategorizeRepository(gdb),
			ArticleRepo:    repository.NewArticleRepository(gdb),
			Spider:         fakeSpiderService{},
		})

		for _, path := range []string{"/api/info/crawl", "/api/categorize/crawl", "/api/articles/crawl"} {
			resp := performRequest(router, http.MethodPost, path)
			require.Equal(t, http.StatusOK, resp.Code)
			require.JSONEq(t, `{"code":200,"msg":"抓取成功","data":null}`, resp.Body.String())
		}
	})

	t.Run("failure", func(t *testing.T) {
		router := httpapi.NewRouter(config.LoadFromEnv(func(string) string { return "" }), httpapi.Dependencies{
			InfoRepo:       repository.NewInfoRepository(gdb),
			CategorizeRepo: repository.NewCategorizeRepository(gdb),
			ArticleRepo:    repository.NewArticleRepository(gdb),
			Spider:         fakeSpiderService{infoErr: errors.New("boom"), categorizeErr: errors.New("boom"), articleErr: errors.New("boom")},
		})

		resp := performRequest(router, http.MethodPost, "/api/info/crawl")
		require.Equal(t, http.StatusInternalServerError, resp.Code)
		require.Contains(t, resp.Body.String(), "抓取用户信息失败")

		resp = performRequest(router, http.MethodPost, "/api/categorize/crawl")
		require.Equal(t, http.StatusInternalServerError, resp.Code)
		require.Contains(t, resp.Body.String(), "抓取分类信息失败")

		resp = performRequest(router, http.MethodPost, "/api/articles/crawl")
		require.Equal(t, http.StatusInternalServerError, resp.Code)
		require.Contains(t, resp.Body.String(), "抓取文章信息失败")
	})
}

// fakeSpiderService 是抓取接口测试使用的假实现。
type fakeSpiderService struct {
	infoErr       error
	categorizeErr error
	articleErr    error
}

// CrawlInfo 在测试中模拟用户信息抓取结果。
func (f fakeSpiderService) CrawlInfo() (bool, error) {
	if f.infoErr != nil {
		return false, f.infoErr
	}
	return true, nil
}

// CrawlCategorize 在测试中模拟分类抓取结果。
func (f fakeSpiderService) CrawlCategorize() error {
	return f.categorizeErr
}

// CrawlArticles 在测试中模拟文章抓取结果。
func (f fakeSpiderService) CrawlArticles() error {
	return f.articleErr
}
