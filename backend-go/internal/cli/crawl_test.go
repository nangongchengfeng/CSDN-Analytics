// 文件作用：CLI 测试：验证 crawl 命令是否暴露了与 Python 对齐的子命令。
package cli

import (
	"testing"

	"github.com/spf13/cobra"
)

// TestCrawlCLISupportsPythonEquivalentSubcommands 验证 CLI 子命令与 Python 版本保持一致。
func TestCrawlCLISupportsPythonEquivalentSubcommands(t *testing.T) {
	root := NewRootCommand(fakeSpiderFactory)

	requireCommandExists(t, root, "info")
	requireCommandExists(t, root, "categories")
	requireCommandExists(t, root, "articles")
	requireCommandExists(t, root, "all")
}

// requireCommandExists 断言指定名称的子命令已经被成功注册。
func requireCommandExists(t *testing.T, root *cobra.Command, name string) {
	t.Helper()

	for _, cmd := range root.Commands() {
		if cmd.Name() == name {
			return
		}
	}

	t.Fatalf("command %s not found", name)
}

// fakeSpiderFactory 返回测试用的 SpiderRunner 实现。
func fakeSpiderFactory(string) SpiderRunner {
	return fakeSpiderRunner{}
}

// fakeSpiderRunner 是用于 CLI 测试的空实现。
type fakeSpiderRunner struct{}

// CrawlInfo 在测试中模拟用户信息抓取成功。
func (fakeSpiderRunner) CrawlInfo() (bool, error) { return true, nil }

// CrawlCategorize 在测试中模拟分类抓取成功。
func (fakeSpiderRunner) CrawlCategorize() error { return nil }

// CrawlArticles 在测试中模拟文章抓取成功。
func (fakeSpiderRunner) CrawlArticles() error { return nil }
