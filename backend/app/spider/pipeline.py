"""数据管道 - 处理和存储抓取的数据"""
import datetime as dt

from app import db
from app.models.info import Info
from app.models.categorize import Categorize
from app.models.article import Article


class CSDNPipeline:
    """CSDN 数据管道"""

    def __init__(self):
        pass

    def process_info(self, info_data):
        """处理用户信息"""
        existing = Info.query.filter_by(author_name=info_data['author_name']).first()
        if existing:
            existing.date = info_data['date']
            existing.head_img = info_data['head_img']
            existing.code_age = info_data['code_age']
            existing.article_num = info_data['article_num']
            existing.fans_num = info_data['fans_num']
            existing.visit_num = info_data['visit_num']
            existing.like_num = info_data['like_num']
            existing.comment_num = info_data['comment_num']
            existing.collect_num = info_data['collect_num']
            existing.share_num = info_data['share_num']
            existing.rank = info_data['rank']
            existing.level = info_data['level']
            existing.score = info_data['score']
            print("更新用户信息到数据库")
        else:
            new_info = Info(**info_data)
            db.session.add(new_info)
            print("添加新用户信息到数据库")

        db.session.commit()
        print("成功保存用户信息到数据库")

    def process_categories(self, categories):
        """处理分类信息"""
        for category in categories:
            existing_categorize = Categorize.query.filter_by(href=category['href']).first()
            if existing_categorize:
                existing_categorize.categorize = category['categorize']
                existing_categorize.categorize_id = category['categorize_id']
                existing_categorize.column_num = category['column_num']
                existing_categorize.num_span = category.get('num_span', 0)
                existing_categorize.article_num = category.get('article_num', 0)
                existing_categorize.read_num = category.get('read_num', 0)
                existing_categorize.collect_num = category.get('collect_num', 0)
                print(f"更新分类信息: {category['categorize']}")
            else:
                new_categorize = Categorize(**category)
                db.session.add(new_categorize)
                print(f"添加新的分类信息: {category['categorize']}")

            db.session.commit()

        print("所有的分类信息已经保存到数据库")

    def process_articles(self, articles):
        """处理文章信息"""
        for article in articles:
            existing_article = Article.query.filter_by(url=article['url']).first()

            if existing_article:
                existing_article.title = article['title']
                existing_article.date = article['date']
                existing_article.read_num = article['read_num']
                existing_article.comment_num = article['comment_num']
                existing_article.type = article['type']
            else:
                new_article = Article(**article)
                db.session.add(new_article)

        db.session.commit()
        print(f"成功保存 {len(articles)} 篇文章到数据库")

    def process_article_item(self, item):
        """处理单篇文章"""
        existing_article = Article.query.filter_by(url=item['url']).first()

        if existing_article:
            existing_article.title = item['title']
            existing_article.date = item['date']
            existing_article.read_num = item['read_num']
            existing_article.comment_num = item['comment_num']
            existing_article.type = item['type']
        else:
            new_article = Article(**item)
            db.session.add(new_article)

        db.session.commit()
