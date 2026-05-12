// 文件作用：配置测试：验证默认配置值是否与 Python 后端保持一致。
package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestLoadFromEnvDefaultsMatchPython 验证默认配置与 Python 后端约定保持一致。
func TestLoadFromEnvDefaultsMatchPython(t *testing.T) {
	cfg := LoadFromEnv(func(string) string { return "" })

	require.Equal(t, "127.0.0.1", cfg.Host)
	require.Equal(t, 5000, cfg.Port)
	require.Equal(t, "sqlite:///csdn_analytics.db", cfg.DatabaseURL)
	require.Equal(t, "heian_99", cfg.CSDNUserID)
	require.Equal(t, "https://blog.csdn.net/heian_99", cfg.CSDNBlogURL)
	require.Equal(t, 3, cfg.SpiderRetryTimes)
	require.Equal(t, 5, cfg.SpiderRetryDelay)
	require.Equal(t, 10, cfg.SpiderTimeout)
	require.True(t, cfg.Debug)
	require.Contains(t, cfg.CORSOrigins, "http://localhost:5173")
	require.Contains(t, cfg.CORSOrigins, "http://127.0.0.1:5174")
}
