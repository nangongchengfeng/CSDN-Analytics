"""CSDN Analytics 后端启动文件"""
import os
from dotenv import load_dotenv

from app import create_app

# 加载环境变量
load_dotenv()

# 获取配置环境
config_name = os.getenv('FLASK_ENV', 'development')
app = create_app(config_name)


def main():
    """主函数"""
    debug = os.getenv('FLASK_DEBUG', 'True').lower() == 'true'
    host = os.getenv('FLASK_HOST', '127.0.0.1')
    port = int(os.getenv('FLASK_PORT', '5000'))

    app.run(debug=debug, host=host, port=port)


if __name__ == '__main__':
    main()
