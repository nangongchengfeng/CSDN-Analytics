"""爬虫模块"""
from app.spider.client import CSDNSpider
from app.spider.parser import CSDNParser
from app.spider.pipeline import CSDNPipeline

__all__ = ['CSDNSpider', 'CSDNParser', 'CSDNPipeline']
