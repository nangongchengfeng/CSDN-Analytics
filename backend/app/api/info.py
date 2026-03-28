"""用户信息 API"""
from flask import Blueprint, jsonify

from app.models.info import Info
from app.models.categorize import Categorize
from app.spider.client import CSDNSpider

bp = Blueprint('info', __name__)


@bp.route('/info', methods=['GET'])
def get_info():
    """获取用户信息"""
    try:
        info = Info.query.first()
        if info:
            return jsonify({
                'code': 200,
                'msg': '操作成功',
                'data': info.to_dict()
            })
        return jsonify({
            'code': 200,
            'msg': '操作成功',
            'data': None
        })
    except Exception as e:
        return jsonify({
            'code': 500,
            'msg': f'获取用户信息失败: {str(e)}'
        }), 500


@bp.route('/info/crawl', methods=['POST'])
def crawl_info():
    """抓取用户信息"""
    try:
        spider = CSDNSpider()
        success = spider.crawl_info()
        if success:
            return jsonify({
                'code': 200,
                'msg': '抓取成功',
                'data': None
            })
        return jsonify({
            'code': 500,
            'msg': '抓取失败'
        }), 500
    except Exception as e:
        return jsonify({
            'code': 500,
            'msg': f'抓取用户信息失败: {str(e)}'
        }), 500




@bp.route('/categorize/crawl', methods=['POST'])
def crawl_categorize():
    """抓取分类信息"""
    try:
        spider = CSDNSpider()
        spider.crawl_categorize()
        return jsonify({
            'code': 200,
            'msg': '抓取成功',
            'data': None
        })
    except Exception as e:
        return jsonify({
            'code': 500,
            'msg': f'抓取分类信息失败: {str(e)}'
        }), 500
