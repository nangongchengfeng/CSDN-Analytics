"""用户信息模型"""
from app import db


class Info(db.Model):
    """
    模型表示 'info' 表 - 用户信息表
    """
    __tablename__ = 'info'

    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    date = db.Column(db.Text)  # 抓取时间

    # 基础用户信息
    head_img = db.Column(db.Text)  # 头像链接
    author_name = db.Column(db.Text)  # 作者姓名
    code_age = db.Column(db.Text)  # 码龄

    # 统计数据
    article_num = db.Column(db.Integer)  # 原创文章数
    fans_num = db.Column(db.Integer)  # 粉丝数
    like_num = db.Column(db.Integer)  # 点赞数
    comment_num = db.Column(db.Integer)  # 评论数
    collect_num = db.Column(db.Integer)  # 收藏数
    share_num = db.Column(db.Integer)  # 分享数
    visit_num = db.Column(db.Integer)  # 总访问量
    rank = db.Column(db.Integer)  # 排名

    level = db.Column(db.Text)  # 等级
    score = db.Column(db.Integer)  # 积分

    def __repr__(self):
        return f'<Info {self.author_name}>'

    def to_dict(self):
        """转换为字典"""
        return {
            'id': self.id,
            'date': self.date,
            'head_img': self.head_img,
            'author_name': self.author_name,
            'code_age': self.code_age,
            'article_num': self.article_num,
            'fans_num': self.fans_num,
            'like_num': self.like_num,
            'comment_num': self.comment_num,
            'collect_num': self.collect_num,
            'share_num': self.share_num,
            'visit_num': self.visit_num,
            'rank': self.rank,
            'level': self.level,
            'score': self.score
        }
