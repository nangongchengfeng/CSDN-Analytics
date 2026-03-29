# CSDN Analytics 生产部署文档

## 📋 目录
- [环境要求](#环境要求)
- [一、后端部署](#一后端部署)
- [二、前端部署](#二前端部署)
- [三、Nginx 配置](#三nginx-配置)
- [四、常见问题](#四常见问题)

---

## 环境要求

| 组件 | 版本要求 |
|------|----------|
| Python | >= 3.8 (推荐 3.9+) |
| Node.js | >= 18 |
| Nginx | >= 1.18 |
| MySQL | >= 5.7 (可选) |

---

## 一、后端部署

### 1. 安装 uv（推荐，一站式 Python 管理）

uv 可以同时管理 Python 版本和依赖，非常方便！

```bash
# 一键安装 uv
curl -LsSf https://astral.sh/uv/install.sh | sh

# 重新加载配置
source ~/.bashrc

# 验证
uv --version
```

### 2. 使用 uv 安装 Python

```bash
# 查看可用的 Python 版本
uv python list

# 安装 Python 3.10（推荐，稳定且新）
uv python install 3.10

# 或安装 Python 3.9
uv python install 3.9

# 查看已安装的 Python
uv python find

# 在当前项目固定使用特定 Python 版本
cd /var/www/CSDN-Analytics/backend
uv python pin 3.10

# 验证
uv python --version
```

---

### （备选）手动编译安装 Python

如果不想用 uv，可以手动编译：

```bash
# 检查 Python 版本
python3 --version

# 如果版本 < 3.8，需要升级（CentOS 7 示例）
# 安装编译依赖
sudo yum install -y gcc openssl-devel bzip2-devel libffi-devel zlib-devel wget make

# 下载并编译 Python 3.9
cd /tmp
wget https://www.python.org/ftp/python/3.9.18/Python-3.9.18.tgz
tar xzf Python-3.9.18.tgz
cd Python-3.9.18
./configure --enable-optimizations
sudo make altinstall

# 验证
python3.9 --version
```

### 3. 部署后端代码

```bash
# 克隆或上传项目
cd /var/www
git clone https://github.com/nangongchengfeng/CSDN-Analytics.git
cd CSDN-Analytics/backend

# （可选）固定项目使用的 Python 版本
uv python pin 3.10

# 创建虚拟环境（自动使用已 pin 的 Python 版本）
uv venv

# 激活虚拟环境
source .venv/bin/activate

# 验证 Python 版本
python --version

# 安装依赖
uv pip install -r requirements.txt
```

---

### uv Python 管理常用命令

```bash
# 查看可安装的 Python 版本
uv python list

# 安装指定版本
uv python install 3.9
uv python install 3.10
uv python install 3.11

# 查看已安装的 Python
uv python find

# 在项目中固定 Python 版本
uv python pin 3.10

# 卸载 Python 版本
uv python uninstall 3.9

# 查看当前 Python 版本
uv python --version
```

### 4. 配置环境变量

```bash
# 复制配置模板
cp .env.example .env

# 编辑配置
nano .env
```

`.env` 配置示例：

```env
# Flask 配置
FLASK_APP=run.py
FLASK_ENV=production
FLASK_DEBUG=False
FLASK_HOST=0.0.0.0
FLASK_PORT=5000

# 数据库配置（SQLite 简单部署）
DATABASE_URL=sqlite:///csdn_analytics.db

# 或使用 MySQL
# DATABASE_URL=mysql+pymysql://user:password@localhost:3306/csdn_analytics?charset=utf8mb4

# CSDN 配置
CSDN_USER_ID=heian_99

# 应用配置（生产环境一定要修改！）
SECRET_KEY=your-random-secret-key-here-change-this
SQLALCHEMY_TRACK_MODIFICATIONS=False
```

### 5. 采集数据

```bash
# 确保在虚拟环境中
source .venv/bin/activate

# 采集所有数据
uv run flask crawl all

# 或分步采集
uv run flask crawl info        # 用户信息
uv run flask crawl categories  # 分类信息
uv run flask crawl articles    # 文章列表
```

### 6. 启动后端服务

#### 方式一：直接启动（测试用）

```bash
uv run flask run --host=0.0.0.0 --port=5000
```

#### 方式二：使用 systemd（生产推荐）

创建服务文件 `/etc/systemd/system/csdn-analytics.service`：

```ini
[Unit]
Description=CSDN Analytics Backend
After=network.target

[Service]
User=www-data
WorkingDirectory=/var/www/CSDN-Analytics/backend
Environment="PATH=/var/www/CSDN-Analytics/backend/.venv/bin"
ExecStart=/var/www/CSDN-Analytics/backend/.venv/bin/python run.py
Restart=always

[Install]
WantedBy=multi-user.target
```

启动服务：

```bash
sudo systemctl daemon-reload
sudo systemctl enable csdn-analytics
sudo systemctl start csdn-analytics
sudo systemctl status csdn-analytics
```

查看日志：

```bash
sudo journalctl -u csdn-analytics -f
```

---

## 二、前端部署

### 1. 构建前端

```bash
cd /var/www/CSDN-Analytics/frontend

# 安装依赖
npm install

# 配置生产环境 API 地址
nano .env.production
```

`.env.production` 配置（关键！）：

```env
# 前后端同域名时用这个（推荐）
VITE_APP_API_BASE_URL=/api

# 前后端不同域名时用这个
# VITE_APP_API_BASE_URL=https://your-backend-domain.com/api
```

构建：

```bash
npm run build
```

构建完成后，会生成 `dist/` 目录。

---

## 三、Nginx 配置

### 1. 完整 Nginx 配置

创建配置文件 `/etc/nginx/sites-available/csdn-analytics`：

```nginx
server {
    listen 80;
    server_name dash.ownit.top;

    # 前端静态文件
    location / {
        root /var/www/CSDN-Analytics/frontend/dist;
        try_files $uri $uri/ /index.html;
        index index.html;

        # 静态资源缓存
        location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg|woff|woff2|ttf|eot)$ {
            expires 1y;
            add_header Cache-Control "public, immutable";
        }
    }

    # 后端 API 反向代理
    location /api {
        # ⚠️ 关键：去掉 /api 前缀，避免 /api/api/ 问题
        rewrite ^/api/(.*)$ /$1 break;

        proxy_pass http://127.0.0.1:5000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # 超时设置
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }
}
```

### 2. 启用站点

```bash
# 创建软链接
sudo ln -s /etc/nginx/sites-available/csdn-analytics /etc/nginx/sites-enabled/

# 测试配置
sudo nginx -t

# 重启 Nginx
sudo systemctl reload nginx
```

### 3. 配置 SSL（推荐）

```bash
# 安装 Certbot
sudo apt install -y certbot python3-certbot-nginx
# 或 CentOS
sudo yum install -y certbot python3-certbot-nginx

# 获取证书
sudo certbot --nginx -d dash.ownit.top
```

Certbot 会自动修改 Nginx 配置添加 SSL。

---

## 四、常见问题

### Q1: 前端访问 `/api/api/info` 404？

**原因**：Nginx 没有去掉 `/api` 前缀。

**解决**：在 Nginx 配置中添加：

```nginx
location /api {
    rewrite ^/api/(.*)$ /$1 break;  # 关键行
    proxy_pass http://127.0.0.1:5000;
    ...
}
```

### Q2: greenlet 安装失败？

**解决**：

```bash
# 安装编译依赖
sudo yum install -y gcc python3-devel

# 单独安装 greenlet
uv pip install greenlet==3.1.1

# 再安装其他依赖
uv pip install -r requirements.txt
```

### Q3: 后端启动失败？

**检查步骤**：

```bash
# 1. 检查端口是否被占用
netstat -tlnp | grep 5000

# 2. 查看后端日志
sudo journalctl -u csdn-analytics -n 50

# 3. 手动启动测试
cd /var/www/CSDN-Analytics/backend
source .venv/bin/activate
python run.py
```

### Q4: 前端刷新页面 404？

**原因**：Vue SPA 需要回退到 index.html。

**解决**：确保 Nginx 有这行：

```nginx
try_files $uri $uri/ /index.html;
```

### Q5: 如何更新代码？

```bash
# 拉取最新代码
cd /var/www/CSDN-Analytics
git pull

# 更新后端
cd backend
source .venv/bin/activate
uv pip install -r requirements.txt
uv run flask crawl all
sudo systemctl restart csdn-analytics

# 更新前端
cd ../frontend
npm install
npm run build
sudo systemctl reload nginx
```

---

## 📞 获取帮助

- 查看后端日志：`sudo journalctl -u csdn-analytics -f`
- 查看 Nginx 日志：`sudo tail -f /var/log/nginx/error.log`
- 项目地址：https://github.com/nangongchengfeng/CSDN-Analytics
