"""Flask CLI 命令 - 爬虫相关命令"""
import click
from flask import Blueprint

from app.spider.client import CSDNSpider

crawl_bp = Blueprint('crawl', __name__)


@crawl_bp.cli.command('info')
@click.option('--user-id', default=None, help='CSDN 用户 ID')
def crawl_info(user_id):
    """爬取用户信息"""
    click.echo('开始爬取用户信息...')
    try:
        spider = CSDNSpider(user_id)
        success = spider.crawl_info()
        if success:
            click.echo('用户信息爬取成功')
        else:
            click.echo('用户信息爬取失败')
    except Exception as e:
        click.echo(f'用户信息爬取失败: {e}', err=True)


@crawl_bp.cli.command('categories')
@click.option('--user-id', default=None, help='CSDN 用户 ID')
def crawl_categories(user_id):
    """爬取分类信息"""
    click.echo('开始爬取分类信息...')
    try:
        spider = CSDNSpider(user_id)
        spider.crawl_categorize()
        click.echo('分类信息爬取成功')
    except Exception as e:
        click.echo(f'分类信息爬取失败: {e}', err=True)


@crawl_bp.cli.command('articles')
@click.option('--user-id', default=None, help='CSDN 用户 ID')
def crawl_articles(user_id):
    """爬取文章信息"""
    click.echo('开始爬取文章信息...')
    try:
        spider = CSDNSpider(user_id)
        spider.crawl_articles()
        click.echo('文章信息爬取成功')
    except Exception as e:
        click.echo(f'文章信息爬取失败: {e}', err=True)


@crawl_bp.cli.command('all')
@click.option('--user-id', default=None, help='CSDN 用户 ID')
def crawl_all(user_id):
    """爬取所有信息"""
    click.echo('开始爬取所有信息...')
    try:
        spider = CSDNSpider(user_id)
        click.echo('1. 爬取用户信息...')
        spider.crawl_info()
        click.echo('2. 爬取分类信息...')
        spider.crawl_categorize()
        click.echo('3. 爬取文章信息...')
        spider.crawl_articles()
        click.echo('所有信息爬取成功')
    except Exception as e:
        click.echo(f'爬取失败: {e}', err=True)
