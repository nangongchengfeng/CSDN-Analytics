// 文件作用：配置加载层：读取环境变量并生成与 Python 兼容的运行配置。
package config

import (
	"os"
	"strconv"
	"strings"
)

// Config 保存服务运行所需的全部配置项。
type Config struct {
	// Host 是 HTTP 服务监听地址。
	Host string
	// Port 是 HTTP 服务监听端口。
	Port int
	// Debug 控制 Gin 是否以调试模式运行。
	Debug bool
	// DatabaseURL 保存数据库连接地址，默认使用 SQLite。
	DatabaseURL string
	// CSDNUserID 是默认抓取的 CSDN 用户 ID。
	CSDNUserID string
	// CSDNBlogURL 是抓取入口页地址。
	CSDNBlogURL string
	// SpiderRetryTimes 表示抓取失败后的最大重试次数。
	SpiderRetryTimes int
	// SpiderRetryDelay 表示每次重试之间的等待秒数。
	SpiderRetryDelay int
	// SpiderTimeout 表示单次 HTTP 请求超时时间，单位秒。
	SpiderTimeout int
	// CORSOrigins 定义允许访问后端的前端来源列表。
	CORSOrigins []string
}

// Load 从当前进程环境变量中加载配置。
func Load() Config {
	return LoadFromEnv(os.Getenv)
}

// LoadFromEnv 使用传入的 getenv 函数构建配置，便于测试与复用。
func LoadFromEnv(getenv func(string) string) Config {
	// CSDN 博客地址允许直接覆盖；如果未显式配置，则按用户 ID 自动拼接。
	userID := getOrDefault(getenv, "CSDN_USER_ID", "heian_99")
	blogURL := getOrDefault(getenv, "CSDN_BLOG_URL", "https://blog.csdn.net/"+userID)

	return Config{
		Host:             getOrDefault(getenv, "FLASK_HOST", "127.0.0.1"),
		Port:             getIntOrDefault(getenv, "FLASK_PORT", 5000),
		Debug:            getBoolOrDefault(getenv, "FLASK_DEBUG", true),
		DatabaseURL:      getOrDefault(getenv, "DATABASE_URL", "sqlite:///csdn_analytics.db"),
		CSDNUserID:       userID,
		CSDNBlogURL:      blogURL,
		SpiderRetryTimes: getIntOrDefault(getenv, "SPIDER_RETRY_TIMES", 3),
		SpiderRetryDelay: getIntOrDefault(getenv, "SPIDER_RETRY_DELAY", 5),
		SpiderTimeout:    getIntOrDefault(getenv, "SPIDER_TIMEOUT", 10),
		CORSOrigins:      defaultCORSOrigins(),
	}
}

// getOrDefault 读取字符串配置，缺失时返回默认值。
func getOrDefault(getenv func(string) string, key, fallback string) string {
	if value := strings.TrimSpace(getenv(key)); value != "" {
		return value
	}

	return fallback
}

// getIntOrDefault 读取整数配置，解析失败时返回默认值。
func getIntOrDefault(getenv func(string) string, key string, fallback int) int {
	value := strings.TrimSpace(getenv(key))
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}

// getBoolOrDefault 读取布尔配置，解析失败时返回默认值。
func getBoolOrDefault(getenv func(string) string, key string, fallback bool) bool {
	value := strings.TrimSpace(getenv(key))
	if value == "" {
		return fallback
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}

	return parsed
}

// defaultCORSOrigins 返回前端开发环境使用的默认跨域白名单。
func defaultCORSOrigins() []string {
	return []string{
		"http://localhost:5173",
		"http://localhost:5174",
		"http://127.0.0.1:5173",
		"http://127.0.0.1:5174",
	}
}
