"""应用配置模块"""
import os
from dotenv import load_dotenv

# 加载环境变量
load_dotenv()


class Config:
    """基础配置"""
    SECRET_KEY = os.getenv('SECRET_KEY', 'dev-secret-key-change-in-production')
    SQLALCHEMY_TRACK_MODIFICATIONS = os.getenv(
        'SQLALCHEMY_TRACK_MODIFICATIONS', 'False').lower() == 'true'
    SQLALCHEMY_ENGINE_OPTIONS = {}

    # CSDN 配置
    CSDN_USER_ID = os.getenv('CSDN_USER_ID', 'heian_99')
    CSDN_BLOG_URL = os.getenv('CSDN_BLOG_URL', f'https://blog.csdn.net/{CSDN_USER_ID}')

    # 爬虫配置
    SPIDER_RETRY_TIMES = int(os.getenv('SPIDER_RETRY_TIMES', '3'))
    SPIDER_RETRY_DELAY = int(os.getenv('SPIDER_RETRY_DELAY', '5'))
    SPIDER_TIMEOUT = int(os.getenv('SPIDER_TIMEOUT', '10'))

    @staticmethod
    def init_app(app):
        """初始化应用"""
        pass


class DevelopmentConfig(Config):
    """开发环境配置"""
    DEBUG = True
    SQLALCHEMY_DATABASE_URI = os.getenv('DATABASE_URL', 'sqlite:///csdn_analytics.db')


class ProductionConfig(Config):
    """生产环境配置"""
    DEBUG = False
    SQLALCHEMY_DATABASE_URI = os.getenv('DATABASE_URL', 'sqlite:///csdn_analytics.db')


class TestingConfig(Config):
    """测试环境配置"""
    TESTING = True
    SQLALCHEMY_DATABASE_URI = os.getenv('DATABASE_URL', 'sqlite:///test_csdn_analytics.db')


config = {
    'development': DevelopmentConfig,
    'production': ProductionConfig,
    'testing': TestingConfig,
    'default': DevelopmentConfig
}
