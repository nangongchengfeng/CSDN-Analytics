# CSDN Analytics Go Backend

并行 Go 后端，使用 Gin，目标是与现有 Python backend 保持接口和配置兼容。

## 已实现

- Gin 路由与 Python `/api/*` 接口保持一致
- 默认使用 SQLite，兼容 `sqlite:///csdn_analytics.db`
- 保持 `info`、`categorize`、`article` 三张表名和字段名
- 已迁移读接口、抓取接口和 `crawl` CLI
- 保留 Flask 兼容环境变量命名，方便平滑替换

## 环境变量

```bash
cp .env.example .env
```

默认配置：

- `FLASK_HOST=127.0.0.1`
- `FLASK_PORT=5000`
- `FLASK_DEBUG=True`
- `DATABASE_URL=sqlite:///csdn_analytics.db`
- `CSDN_USER_ID=heian_99`
- `CSDN_BLOG_URL=https://blog.csdn.net/heian_99`

## 启动服务

```bash
go mod tidy
go run ./cmd/server
```

## 运行爬虫命令

```bash
go run ./cmd/crawl info
go run ./cmd/crawl categories
go run ./cmd/crawl articles
go run ./cmd/crawl all

# 指定用户
go run ./cmd/crawl all --user-id heian_99
```

## API 列表

- `GET /api/info`
- `POST /api/info/crawl`
- `POST /api/categorize/crawl`
- `GET /api/quarter`
- `GET /api/read`
- `GET /api/pie`
- `GET /api/categorize`
- `GET /api/articles`
- `POST /api/articles/crawl`
- `GET /api/heatmap/:year`
- `GET /api/years`

## 文件说明

### 启动入口

- `cmd/server/main.go`：Go HTTP 服务启动入口，负责加载配置并启动 Gin。
- `cmd/crawl/main.go`：爬虫 CLI 启动入口，负责组装 `crawl` 命令。

### 应用装配

- `internal/app/app.go`：应用装配层，负责创建数据库仓储、Spider 服务和 HTTP 路由。
- `internal/config/config.go`：配置加载层，负责读取环境变量并生成运行配置。
- `internal/config/config_test.go`：配置测试，验证默认值与 Python 后端兼容。

### 数据库

- `internal/db/db.go`：数据库入口层，负责 SQLite DSN 规范化、打开连接和自动建表。
- `internal/db/models.go`：数据库模型定义，声明 `info`、`categorize`、`article` 三张表。
- `internal/db/models_test.go`：数据库测试，验证建表和 SQLite 路径兼容性。

### 仓储层

- `internal/repository/info_repository.go`：`info` 表仓储，负责读取和按 `author_name` upsert。
- `internal/repository/categorize_repository.go`：`categorize` 表仓储，负责列表查询和按 `href` upsert。
- `internal/repository/article_repository.go`：`article` 表仓储，负责列表查询和按 `url` upsert。
- `internal/repository/repository_test.go`：仓储测试，覆盖三类仓储的增改查行为。

### 业务服务

- `internal/service/article_service.go`：文章领域服务，负责时间预处理、筛选、去重和统计聚合。
- `internal/spider/parser.go`：HTML 解析层，从 CSDN 页面中提取结构化字段。
- `internal/spider/service.go`：抓取服务层，负责请求页面、解析结果并写回数据库。
- `internal/spider/parser_test.go`：解析测试，验证用户、分类、文章解析逻辑。

### HTTP 层

- `internal/http/router.go`：Gin 路由注册层，挂载中间件和全部 `/api/*` 接口。
- `internal/http/response.go`：统一 JSON 响应工具函数。
- `internal/http/handlers/handlers.go`：HTTP 控制器，实现读接口和抓取接口。
- `internal/http/handlers/handlers_test.go`：读接口契约测试。
- `internal/http/handlers/crawl_test.go`：抓取接口契约测试。

### CLI 层

- `internal/cli/crawl.go`：定义 `crawl info/categories/articles/all` 子命令。
- `internal/cli/crawl_test.go`：CLI 测试，验证子命令是否齐全。
