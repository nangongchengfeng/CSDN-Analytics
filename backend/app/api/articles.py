"""文章数据 API"""
from datetime import datetime
from collections import defaultdict
from flask import Blueprint, jsonify, request

from app.models.article import Article

bp = Blueprint('articles', __name__)


def get_articles():
    """获取所有文章数据并进行预处理"""
    try:
        articles = Article.query.all()

        result = []
        for article in articles:
            article_date = datetime.strptime(article.date, '%Y-%m-%d %H:%M:%S')
            article_dict = {
                'id': article.id,
                'url': article.url,
                'title': article.title,
                'date': article.date,
                'read_num': article.read_num,
                'comment_num': article.comment_num,
                'type': article.type,
                'date_day': article_date.date(),
                'date_month': article_date.strftime('%Y年%m月'),
                'weekday': article_date.weekday(),
                'year': article_date.year,
                'month': article_date.month,
                'quarter': get_quarter(article_date.month),
                'week': article_date.isocalendar()[1],
            }
            result.append(article_dict)

        return result
    except Exception as e:
        print(f"Error in get_articles: {str(e)}")
        return []


def get_quarter(month):
    """根据月份返回季度"""
    if 1 <= month <= 3:
        return "第一季度"
    elif 4 <= month <= 6:
        return "第二季度"
    elif 7 <= month <= 9:
        return "第三季度"
    elif 10 <= month <= 12:
        return "第四季度"


@bp.route('/articles', methods=['GET'])
def get_articles_list():
    """获取文章列表"""
    try:
        articles = get_articles()

        filter_type = request.args.get('type', None)
        filter_quarter = request.args.get('quarter', None)
        filter_year = request.args.get('year', None)
        filter_week = request.args.get('week', None)
        filter_day = request.args.get('day', None)

        if filter_type:
            articles = [article for article in articles if article['type'] == filter_type]

        if filter_quarter:
            articles = [article for article in articles if article['quarter'] == filter_quarter]

        if filter_year:
            articles = [article for article in articles if article['year'] == int(filter_year)]

        if filter_week:
            articles = [article for article in articles if article['week'] == int(filter_week)]

        if filter_day:
            articles = [article for article in articles if article['weekday'] == int(filter_day) - 1]

        seen_urls = set()
        unique_articles = []
        for article in articles:
            if article['url'] not in seen_urls:
                seen_urls.add(article['url'])
                unique_articles.append(article)

        unique_articles.sort(key=lambda x: x['date'], reverse=True)

        result = []
        for article in unique_articles[:100]:
            result.append({
                'title': article['title'],
                'url': article['url'],
                'date': article['date'],
                'type': article['type']
            })

        return jsonify({
            'code': 200,
            'msg': '操作成功',
            'data': result
        })
    except Exception as e:
        return jsonify({
            'code': 500,
            'msg': f'获取文章列表失败: {str(e)}'
        }), 500


@bp.route('/articles/crawl', methods=['POST'])
def crawl_articles():
    """抓取文章信息"""
    try:
        from app.spider.client import CSDNSpider

        spider = CSDNSpider()
        spider.crawl_articles()

        return jsonify({
            'code': 200,
            'msg': '抓取成功',
            'data': None
        })
    except Exception as e:
        return jsonify({
            'code': 500,
            'msg': f'抓取文章信息失败: {str(e)}'
        }), 500


@bp.route('/heatmap/<int:year>', methods=['GET'])
def get_heatmap(year):
    """
    获取指定年份的文章发布热力图数据
    :param year: 指定的年份
    :return: 热力图数据
    """
    try:
        data = get_articles()
        data_filtered = [article for article in data if article['year'] == year]

        heatmap_dict = defaultdict(int)
        for article in data_filtered:
            week = article['week']
            weekday = article['weekday']
            key = (week, weekday)
            heatmap_dict[key] += 1

        heatmap_data = []
        for (week, weekday), count in heatmap_dict.items():
            heatmap_data.append([week - 1, weekday, count])

        weeks = sorted(set(week for week, _ in heatmap_dict.keys()))
        weekdays = sorted(set(weekday for _, weekday in heatmap_dict.keys()))

        complete_heatmap_data = []
        for week in range(1, max(weeks) + 1 if weeks else 0):
            for weekday in range(7):
                count = heatmap_dict.get((week, weekday), 0)
                complete_heatmap_data.append([week - 1, weekday, count])

        if not weeks:
            max_weeks = 52
        else:
            max_weeks = max(weeks)

        result = {
            'data': complete_heatmap_data,
            'xAxis': [f'第{i}周' for i in range(1, max_weeks + 1)],
            'yAxis': ["星期{}".format(i + 1) if i != 6 else "星期日" for i in range(7)]
        }

        return jsonify({
            'code': 200,
            'msg': '操作成功',
            'data': result
        })
    except Exception as e:
        return jsonify({
            'code': 500,
            'msg': f'获取热力图数据失败: {str(e)}'
        }), 500


@bp.route('/years', methods=['GET'])
def get_years():
    """获取所有年份列表"""
    try:
        articles = get_articles()
        years = sorted({str(article["year"]) for article in articles}, reverse=True)
        return jsonify({
            'code': 200,
            'msg': '操作成功',
            'data': years
        })
    except Exception as e:
        return jsonify({
            'code': 500,
            'msg': f'获取年份列表失败: {str(e)}'
        }), 500

