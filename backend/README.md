# CSDN Analytics Backend

CSDN Analytics 后端服务，基于 Flask 3.x 构建。

## 技术栈

- Flask 3.x
- Flask-SQLAlchemy 3.1.x
- SQLite (支持 MySQL)
- BeautifulSoup4 (爬虫)
- python-dotenv (环境变量管理)

## 目录结构

```
backend/
├── .env                    # 环境变量配置 (需要自己创建)
├── .env.example            # 环境变量示例
├── .gitignore              # Git 忽略文件
├── requirements.txt        # Python 依赖
├── pyproject.toml          # 项目配置 (uv)
├── run.py                  # 应用启动文件
└── app/
    ├── __init__.py         # 应用工厂
    ├── config.py           # 配置文件
    ├── models/             # 数据模型
    │   ├── __init__.py
    │   ├── info.py         # 用户信息模型
    │   ├── categorize.py   # 分类模型
    │   └── article.py      # 文章模型
    ├── api/                # API 路由
    │   ├── __init__.py
    │   ├── info.py         # 用户信息 API
    │   ├── stats.py        # 统计数据 API
    │   └── articles.py     # 文章数据 API
    ├── commands/           # CLI 命令
    │   ├── __init__.py
    │   └── crawl.py        # 爬虫命令
    └── spider/             # 爬虫模块
        ├── __init__.py
        ├── client.py       # 爬虫客户端
        ├── parser.py       # 页面解析器
        └── pipeline.py     # 数据管道
```

## 快速开始

### 1. 安装依赖

```bash
cd backend
pip install -r requirements.txt
```

或者使用 uv：

```bash
uv pip install
```

### 2. 配置环境变量

复制 `.env.example` 为 `.env`，并根据需要修改配置：

```bash
cp .env.example .env
```

### 3. 运行应用

```bash
# 使用 flask run
flask run

# 或直接运行
python run.py
```

### 4. 运行爬虫命令

```bash
# 爬取所有数据
flask crawl all

# 单独爬取用户信息
flask crawl info

# 单独爬取分类信息
flask crawl categories

# 单独爬取文章信息
flask crawl articles

# 指定用户 ID
flask crawl all --user-id heian_99
```

## API 接口

### 用户信息

- `GET /api/info` - 获取用户信息
- `POST /api/info/crawl` - 抓取用户信息

### 分类信息

- `GET /api/categorize` - 获取分类信息
- `POST /api/categorize/crawl` - 抓取分类信息

### 统计数据

- `GET /api/quarter` - 获取季度统计数据
- `GET /api/read` - 获取阅读量统计
- `GET /api/pie` - 获取饼图数据

### 文章数据

- `GET /api/articles` - 获取文章列表
- `POST /api/articles/crawl` - 抓取文章信息
- `GET /api/heatmap/<year>` - 获取热力图数据

## 数据库配置

默认使用 SQLite，可以通过 `DATABASE_URL` 环境变量配置：

```env
# SQLite
DATABASE_URL=sqlite:///csdn_analytics.db

# MySQL
DATABASE_URL=mysql+pymysql://username:password@localhost:3306/csdn_analytics?charset=utf8mb4
```
