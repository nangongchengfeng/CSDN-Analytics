"""统计数据 API"""
from datetime import datetime
from collections import defaultdict
from flask import Blueprint, jsonify

from app.api.articles import get_articles

bp = Blueprint('stats', __name__)


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


@bp.route('/quarter', methods=['GET'])
def get_quarter_stats():
    """获取每年每季度的博客数量"""
    try:
        year_quarter_count = defaultdict(lambda: defaultdict(int))
        data = get_articles()

        for article in data:
            year = article["year"]
            quarter = article["quarter"]
            year_quarter_count[year][quarter] += 1

        result = []
        for year, quarters in year_quarter_count.items():
            year_data = {"product": str(year)}
            for quarter in ["第一季度", "第二季度", "第三季度", "第四季度"]:
                year_data[quarter] = quarters.get(quarter, 0)
            result.append(year_data)

        return jsonify({
            'code': 200,
            'msg': '操作成功',
            'data': result
        })
    except Exception as e:
        return jsonify({
            'code': 500,
            'msg': f'获取季度统计失败: {str(e)}'
        }), 500


@bp.route('/read', methods=['GET'])
def get_read_stats():
    """获取各类文章的阅读量和文章数量统计"""
    try:
        data = get_articles()
        type_stats = defaultdict(lambda: {'count': 0, 'reads': 0})

        for article in data:
            article_type = article["type"]
            type_stats[article_type]['count'] += 1
            type_stats[article_type]['reads'] += article["read_num"]

        result = {
            'labels': [],
            'reads': [],
            'counts': []
        }

        for type_name, stats in type_stats.items():
            result['labels'].append(type_name)
            result['reads'].append(stats['reads'])
            result['counts'].append(stats['count'])

        return jsonify({
            'code': 200,
            'msg': '操作成功',
            'data': result
        })
    except Exception as e:
        return jsonify({
            'code': 500,
            'msg': f'获取阅读量统计失败: {str(e)}'
        }), 500


@bp.route('/pie', methods=['GET'])
@bp.route('/categorize', methods=['GET'])
def get_pie_data():
    """获取饼图数据（分类统计）"""
    from app.models.categorize import Categorize

    try:
        categorize_data = Categorize.query.all()

        if categorize_data:
            pie_data = [
                {"value": item.article_num, "name": item.categorize}
                for item in categorize_data
            ]
        else:
            pie_data = []

        return jsonify({
            'code': 200,
            'msg': '操作成功',
            'data': pie_data
        })
    except Exception as e:
        return jsonify({
            'code': 500,
            'msg': f'获取饼图数据失败: {str(e)}'
        }), 500
