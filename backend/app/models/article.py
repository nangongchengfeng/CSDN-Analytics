"""文章信息模型"""
from app import db


class Article(db.Model):
    """
    模型表示 'article' 表 - 文章表
    """
    __tablename__ = 'article'

    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    url = db.Column(db.Text)  # 文章链接
    title = db.Column(db.Text)  # 文章标题
    date = db.Column(db.Text)  # 发布日期
    read_num = db.Column(db.BigInteger)  # 阅读量
    comment_num = db.Column(db.BigInteger)  # 评论数
    type = db.Column(db.Text)  # 文章类型

    def __repr__(self):
        return f'<Article {self.title}>'

    def to_dict(self):
        """转换为字典"""
        return {
            'id': self.id,
            'url': self.url,
            'title': self.title,
            'date': self.date,
            'read_num': self.read_num,
            'comment_num': self.comment_num,
            'type': self.type
        }
