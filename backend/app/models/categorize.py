"""分类信息模型"""
from app import db


class Categorize(db.Model):
    """
    模型表示 'categorize' 表 - 博客分类表
    """
    __tablename__ = 'categorize'

    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    href = db.Column(db.Text)  # 分类链接
    categorize = db.Column(db.Text)  # 分类名称
    categorize_id = db.Column(db.BigInteger)  # 分类ID
    column_num = db.Column(db.BigInteger)  # 专栏文章数量
    num_span = db.Column(db.BigInteger)  # 订阅数量
    article_num = db.Column(db.BigInteger)  # 文章数量
    read_num = db.Column(db.BigInteger)  # 阅读量
    collect_num = db.Column(db.BigInteger)  # 收藏数

    def __repr__(self):
        return f'<Categorize {self.categorize}>'

    def to_dict(self):
        """转换为字典"""
        return {
            'id': self.id,
            'href': self.href,
            'categorize': self.categorize,
            'categorize_id': self.categorize_id,
            'column_num': self.column_num,
            'num_span': self.num_span,
            'article_num': self.article_num,
            'read_num': self.read_num,
            'collect_num': self.collect_num
        }
