"""CSDN Analytics 后端应用工厂"""
from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_cors import CORS

from app.config import config

db = SQLAlchemy()


def create_app(config_name='default'):
    """应用工厂函数"""
    app = Flask(__name__)
    app.config.from_object(config[config_name])
    config[config_name].init_app(app)

    db.init_app(app)

    # 启用 CORS
    CORS(app, resources={
        r"/api/*": {
            "origins": ["http://localhost:5173", "http://localhost:5174", "http://127.0.0.1:5173", "http://127.0.0.1:5174"],
            "methods": ["GET", "POST", "PUT", "DELETE", "OPTIONS"],
            "allow_headers": ["Content-Type", "Authorization"]
        }
    })

    # 注册蓝图
    from app.api.info import bp as info_bp
    from app.api.stats import bp as stats_bp
    from app.api.articles import bp as articles_bp

    app.register_blueprint(info_bp, url_prefix='/api')
    app.register_blueprint(stats_bp, url_prefix='/api')
    app.register_blueprint(articles_bp, url_prefix='/api')

    # 注册 CLI 命令
    from app.commands.crawl import crawl_bp
    app.register_blueprint(crawl_bp)

    # 创建数据库表
    with app.app_context():
        # 导入模型以确保它们被注册
        from app.models.info import Info
        from app.models.categorize import Categorize
        from app.models.article import Article
        db.create_all()

    return app
