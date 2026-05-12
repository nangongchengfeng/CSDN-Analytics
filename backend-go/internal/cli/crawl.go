// 文件作用：命令行定义层：声明 crawl/info/categories/articles/all 子命令。
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type SpiderRunner interface {
	CrawlInfo() (bool, error)
	CrawlCategorize() error
	CrawlArticles() error
}

type SpiderFactory func(userID string) SpiderRunner

// NewRootCommand 创建 crawl 根命令并挂载全部子命令。
func NewRootCommand(factory SpiderFactory) *cobra.Command {
	root := &cobra.Command{
		Use:   "crawl",
		Short: "Crawl CSDN analytics data",
	}

	root.AddCommand(newInfoCommand(factory))
	root.AddCommand(newCategoriesCommand(factory))
	root.AddCommand(newArticlesCommand(factory))
	root.AddCommand(newAllCommand(factory))

	return root
}

// newInfoCommand 创建抓取用户信息的子命令。
func newInfoCommand(factory SpiderFactory) *cobra.Command {
	var userID string
	cmd := &cobra.Command{
		Use: "info",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.OutOrStdout(), "开始爬取用户信息...")
			ok, err := factory(userID).CrawlInfo()
			if err != nil {
				return err
			}
			if !ok {
				return fmt.Errorf("抓取失败")
			}
			fmt.Fprintln(cmd.OutOrStdout(), "用户信息爬取成功")
			return nil
		},
	}
	cmd.Flags().StringVar(&userID, "user-id", "", "CSDN 用户 ID")
	return cmd
}

// newCategoriesCommand 创建抓取分类信息的子命令。
func newCategoriesCommand(factory SpiderFactory) *cobra.Command {
	var userID string
	cmd := &cobra.Command{
		Use: "categories",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.OutOrStdout(), "开始爬取分类信息...")
			if err := factory(userID).CrawlCategorize(); err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "分类信息爬取成功")
			return nil
		},
	}
	cmd.Flags().StringVar(&userID, "user-id", "", "CSDN 用户 ID")
	return cmd
}

// newArticlesCommand 创建抓取文章信息的子命令。
func newArticlesCommand(factory SpiderFactory) *cobra.Command {
	var userID string
	cmd := &cobra.Command{
		Use: "articles",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.OutOrStdout(), "开始爬取文章信息...")
			if err := factory(userID).CrawlArticles(); err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "文章信息爬取成功")
			return nil
		},
	}
	cmd.Flags().StringVar(&userID, "user-id", "", "CSDN 用户 ID")
	return cmd
}

// newAllCommand 创建按顺序抓取全部数据的子命令。
func newAllCommand(factory SpiderFactory) *cobra.Command {
	var userID string
	cmd := &cobra.Command{
		Use: "all",
		RunE: func(cmd *cobra.Command, args []string) error {
			runner := factory(userID)
			fmt.Fprintln(cmd.OutOrStdout(), "开始爬取所有信息...")
			fmt.Fprintln(cmd.OutOrStdout(), "1. 爬取用户信息...")
			ok, err := runner.CrawlInfo()
			if err != nil {
				return err
			}
			if !ok {
				return fmt.Errorf("抓取失败")
			}
			fmt.Fprintln(cmd.OutOrStdout(), "2. 爬取分类信息...")
			if err := runner.CrawlCategorize(); err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "3. 爬取文章信息...")
			if err := runner.CrawlArticles(); err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "所有信息爬取成功")
			return nil
		},
	}
	cmd.Flags().StringVar(&userID, "user-id", "", "CSDN 用户 ID")
	return cmd
}
