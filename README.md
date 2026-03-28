# CSDN博客洞察看板 🚀

## 项目简介

这是一个基于 Python Flask + Vue 3 + TypeScript + ECharts 的 CSDN 博客数据可视化项目，采用现代化玻璃态 UI 设计，提供对个人博客数据的深入洞察和交互式展示。

![QQ20260328-140524](images/QQ20260328-140524.png)



## 🌟 功能特点

1. **博客统计仪表盘**
   - 展示博主基本信息（文章数、关注数、访问量、排名、分享数、码龄等）
   - 玻璃态卡片设计，悬浮交互效果
   - 多维度数据可视化

2. **交互式图表**
   - 季度文章数量柱状图（点击筛选）
   - 博客分类饼图（点击筛选）
   - 文章阅读量混合图（点击筛选）
   - 文章发布热力图（年份切换 + 点击筛选）

3. **数据筛选与交互**
   - 支持按年份、季度、类型、周、日筛选文章
   - 图表交互式点击事件
   - 年份默认显示最新年份
   - 文章列表实时更新

## 🎨 UI 设计亮点

- **玻璃态现代风格** - backdrop-filter 毛玻璃效果
- **浅灰渐变背景** - 柔和的科技感配色
- **蓝色科技主题** - #3b82f6 主色调
- **卡片式布局** - 圆角阴影，悬浮交互
- **精致细节** - 图表标题装饰条、年份选择器、筛选按钮

## 🛠 技术栈

### 后端
- **Python 3.x**
- **Flask 3.x** - Web 框架
- **Flask-SQLAlchemy 3.1.1** - ORM
- **Flask-CORS** - 跨域支持
- **MySQL** - 数据库

### 前端
- **Vue 3** - Composition API + `<script setup>`
- **TypeScript 5.x** - 类型安全
- **Vite 6.x** - 构建工具
- **ECharts 5.x** - 图表库
- **PrimeVue** - UI 组件库

## 📦 项目结构

```bash
CSDN-Analytics/
├── backend/                 # 后端目录
│   ├── app/
│   │   ├── __init__.py     # Flask 应用工厂
│   │   ├── models.py       # 数据库模型
│   │   ├── config.py       # 配置文件
│   │   └── api/            # API 路由
│   │       ├── info.py     # 用户信息接口
│   │       ├── stats.py    # 统计数据接口
│   │       └── articles.py # 文章列表接口
│   ├── spider.py           # CSDN 数据爬虫
│   ├── extensions.py       # Flask 扩展
│   ├── requirements.txt    # Python 依赖
│   └── pyproject.toml      # Rye 项目配置
├── frontend/               # 前端目录
│   ├── src/
│   │   ├── App.vue        # 主应用组件
│   │   ├── main.ts        # 入口文件
│   │   ├── api/
│   │   │   └── client.ts  # API 客户端
│   │   ├── types/
│   │   │   └── index.ts   # TypeScript 类型定义
│   │   ├── components/
│   │   │   ├── Layout/    # 布局组件
│   │   │   ├── Dashboard/ # 仪表盘组件
│   │   │   └── Charts/    # 图表组件
│   │   └── styles/
│   │       └── global.css # 全局样式
│   ├── package.json
│   └── vite.config.ts
├── docs/                   # 文档目录
│   └── superpowers/
│       ├── specs/         # 设计规格
│       └── plans/         # 实施计划
└── README.md              # 项目说明文档
```

## 🚀 快速开始

### 环境要求

- Python 3.9+
- Node.js 18+
- MySQL 5.7+

### 后端启动

1. 克隆项目
```bash
git clone https://github.com/nangongchengfeng/CSDN-Analytics.git
cd CSDN-Analytics
```

2. 进入后端目录
```bash
cd backend
```

3. 安装依赖

**方式一：使用 pip（推荐，简单快捷）**
```bash
pip install -r requirements.txt
```

**方式二：使用 uv（超快 Python 包管理器）**
```bash
uv pip install -r requirements.txt
```

> uv 是一个用 Rust 编写的超快 Python 包管理器，兼容 pip 接口，速度快 10-100 倍。

4. 配置环境变量
复制 `.env.example` 为 `.env` 并修改配置：
```env
FLASK_ENV=development
DATABASE_URL=mysql+pymysql://user:password@localhost/csdn_analytics
```

5. 配置 CSDN 用户 ID（可选）
```bash
# 方式一：设置环境变量
export CSDN_USER_ID=heian_99

# 方式二：在爬取时通过 --user-id 参数指定
```

6. 运行爬虫采集数据
```bash
# 爬取所有数据（推荐）
flask crawl all

# 或分别爬取
flask crawl info        # 仅用户信息
flask crawl categories  # 仅分类信息
flask crawl articles    # 仅文章信息

# 指定用户 ID
flask crawl all --user-id=heian_99

# 如果使用 uv 管理虚拟环境：
# uv run flask crawl all
```

7. 启动后端服务
```bash
flask run --port=5000

# 如果使用 uv 管理虚拟环境：
# uv run flask run --port=5000
```

> 注意：数据库表会在首次启动时自动创建

后端服务运行在 http://localhost:5000

### 前端启动

1. 进入前端目录
```bash
cd frontend
```

2. 安装依赖
```bash
npm install
```

3. 启动开发服务器
```bash
npm run dev
```

前端服务运行在 http://localhost:5173

4. 构建生产版本
```bash
npm run build
```

## 🔧 配置说明

### 后端配置

- `FLASK_ENV`: 运行环境（development/production）
- `DATABASE_URL`: 数据库连接字符串
- `CORS_ORIGINS`: 允许的跨域来源

### 前端配置

- `VITE_API_URL`: 后端 API 地址（默认：http://localhost:5000）

## 📸 界面预览

![QQ20260328-140524](images/QQ20260328-140524.png)

### 玻璃态现代风格
- 浅灰渐变背景
- 蓝色科技主题
- 玻璃悬浮效果
- 精致圆角阴影

**图表交互**
- 季度柱状图点击筛选
- 热力图年份切换
- 文章列表实时更新

## 🤝 贡献指南

1. Fork 本仓库
2. 创建你的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的修改 (`git commit -m '添加了某某特性'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

## 📄 许可证

本项目采用 MIT 许可证。详见 `LICENSE` 文件。

## 🌈 作者

- **南宫乘风**
- 联系邮箱: 1794748404@qq.com
- CSDN 博客: https://blog.csdn.net/heian_99
