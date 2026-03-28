"""CSDN 爬虫客户端"""
import datetime as dt
import random
import re
import time

import requests
from bs4 import BeautifulSoup

from app import db
from app.models.info import Info
from app.models.categorize import Categorize
from app.models.article import Article


class CSDNSpider:
    """CSDN 爬虫类"""

    def __init__(self, user_id=None):
        from app.config import DevelopmentConfig
        self.config = DevelopmentConfig()
        self.user_id = user_id or self.config.CSDN_USER_ID
        self.blog_url = self.config.CSDN_BLOG_URL
        self.headers = {
            'User-Agent': 'Mozilla/5.0 (MSIE 10.0; Windows NT 6.1; Trident/5.0)',
            'referer': 'https://passport.csdn.net/login',
        }
        self.retry_times = self.config.SPIDER_RETRY_TIMES
        self.retry_delay = self.config.SPIDER_RETRY_DELAY
        self.timeout = self.config.SPIDER_TIMEOUT

    def to_int(self, s):
        """字符串转换为整数"""
        try:
            return int(s.replace(',', '').strip())
        except:
            return 0

    def crawl_info(self):
        """获取大屏第一列信息数据并存储到数据库"""
        retry_count = 0

        while retry_count < self.retry_times:
            try:
                print(f"尝试抓取用户信息 (第 {retry_count + 1}/{self.retry_times} 次)...")
                resp = requests.get(self.blog_url, headers=self.headers, timeout=self.timeout)
                resp.raise_for_status()
                now = dt.datetime.now().strftime("%Y-%m-%d %H:%M:%S")

                soup = BeautifulSoup(resp.text, 'lxml')
                user_info_container = soup.find('div', class_='user-profile-head-info-r-c')

                if not user_info_container:
                    raise ValueError("没有找到用户信息数据源，请检查CSDN的HTML页面结构。")

                extracted_statistics = {}
                list_items = user_info_container.find_all('li')

                for item in list_items:
                    name_div = item.find('div', class_='user-profile-statistics-name')
                    num_div = item.find('div', class_='user-profile-statistics-num')

                    if name_div and num_div:
                        label = name_div.get_text(strip=True)
                        value = num_div.get_text(strip=True)
                        extracted_statistics[label] = value

                print("用户统计数据:", extracted_statistics)

                data_info_list = soup.find_all('ul', class_='aside-common-box-achievement')
                if not data_info_list:
                    raise ValueError("没有找到data_info数据源，请检查CSDN的HTML页面结构。")

                target_keywords = {
                    "点赞": "点赞",
                    "评论": "评论",
                    "收藏": "收藏",
                    "分享": "分享"
                }
                extracted_achievements = {}

                for ul_element in data_info_list:
                    list_items = ul_element.find_all('li')
                    for li_item in list_items:
                        text_container_div = li_item.find('div')
                        if not text_container_div:
                            continue

                        num_span = text_container_div.find('span')
                        if num_span:
                            value_str = num_span.get_text(strip=True)
                            full_text_content = text_container_div.get_text(strip=True)

                            for keyword, label_name in target_keywords.items():
                                if keyword in full_text_content:
                                    extracted_achievements[label_name] = value_str
                                    break

                print("成就统计数据:", extracted_achievements)

                code_age_div = soup.find('div', class_='person-code-age')
                code_age_text = None
                if code_age_div:
                    span_tag = code_age_div.find('span')
                    if span_tag:
                        code_age_text = span_tag.get_text(strip=True)
                        print(f"抓取到的码龄：{code_age_text}")
                else:
                    print("未找到码龄所在的 div 元素。")

                author_name = soup.find('div', class_='user-profile-head-name').find('div').get_text(strip=True)
                print("作者姓名:", author_name)

                img_tag_css = soup.select_one('div.user-profile-avatar img')
                image_url_css = None
                if img_tag_css:
                    image_url_css = img_tag_css.get('src')
                    print(f"抓取到的头像链接：{image_url_css}")
                else:
                    print("未找到用户头像所在的 img 元素。")
                print("用户信息抓取成功！")

                info = {
                    'date': now,
                    'head_img': image_url_css,
                    'author_name': author_name,
                    'code_age': code_age_text or '',
                    'article_num': self.to_int(extracted_statistics.get('原创', '0')),
                    'fans_num': self.to_int(extracted_statistics.get('粉丝', '0')),
                    'visit_num': self.to_int(extracted_statistics.get('总访问量', '0')),
                    'like_num': self.to_int(extracted_achievements.get('点赞', '0')),
                    'comment_num': self.to_int(extracted_achievements.get('评论', '0')),
                    'collect_num': self.to_int(extracted_achievements.get('收藏', '0')),
                    'share_num': self.to_int(extracted_achievements.get('分享', '0')),
                    'rank': self.to_int(extracted_statistics.get('排名', '0')),
                    'level': extracted_statistics.get('等级', ''),
                    'score': self.to_int(extracted_statistics.get('积分', '0')),
                }
                print("解析后的数据:", info)

                self.save_info(info)
                break

            except Exception as e:
                retry_count += 1
                print(f"抓取用户信息过程中发生错误 (尝试 {retry_count}/{self.retry_times} 次): {str(e)}")
                if retry_count < self.retry_times:
                    print(f"等待 {self.retry_delay} 秒后重试...")
                    time.sleep(self.retry_delay)
                else:
                    print("达到最大重试次数，抓取失败。")
                    return False

        return retry_count < self.retry_times

    def save_info(self, info):
        """保存用户信息到数据库"""
        existing = Info.query.filter_by(author_name=info['author_name']).first()
        if existing:
            existing.date = info['date']
            existing.head_img = info['head_img']
            existing.code_age = info['code_age']
            existing.article_num = info['article_num']
            existing.fans_num = info['fans_num']
            existing.visit_num = info['visit_num']
            existing.like_num = info['like_num']
            existing.comment_num = info['comment_num']
            existing.collect_num = info['collect_num']
            existing.share_num = info['share_num']
            existing.rank = info['rank']
            existing.level = info['level']
            existing.score = info['score']
            print("更新用户信息到数据库")
        else:
            new_info = Info(**info)
            db.session.add(new_info)
            print("添加新用户信息到数据库")

        db.session.commit()
        print("成功保存用户信息到数据库")

    def crawl_categorize(self):
        """爬取分类信息"""
        try:
            response = requests.get(self.blog_url, headers=self.headers, timeout=self.timeout)
            response.raise_for_status()

            soup = BeautifulSoup(response.content, "lxml")
            spans = soup.find_all('a', attrs={'class': 'special-column-name'})
            if not spans:
                print("No categories found.")
                return

            for span in spans:
                try:
                    href = span.get('href')
                    if not href:
                        print("Warning: Missing href for a span.")
                        continue

                    blog_column = span.text.strip()
                    blog_id = href.split("_")[-1].split(".")[0]

                    num_span = span.find('span', class_='special-column-num')
                    if num_span:
                        blogs_column_num = int(re.findall(r'\d+', num_span.text)[0])
                    else:
                        blogs_column_num = 0

                    try:
                        detail_response = requests.get(href, headers=self.headers, timeout=self.timeout)
                        detail_response.raise_for_status()

                        detail_soup = BeautifulSoup(detail_response.text, 'html.parser')
                        column_operating_div = detail_soup.find('div', {'class': 'column_operating'})

                        if column_operating_div:
                            subscribe_num = int(re.findall(r'\d+', column_operating_div.find('span', {
                                'class': 'column-subscribe-num'}).text)[0]) if column_operating_div.find('span', {
                                'class': 'column-subscribe-num'}) else 0
                            mumber_spans = column_operating_div.find_all('span', {'class': 'mumber-color'})

                            article_num = int(mumber_spans[1].text) if len(mumber_spans) > 1 else 0
                            read_num = int(mumber_spans[2].text) if len(mumber_spans) > 2 else 0
                            collect_num = int(mumber_spans[3].text) if len(mumber_spans) > 3 else 0
                        else:
                            subscribe_num = article_num = read_num = collect_num = 0
                    except Exception as e:
                        print(f"Error retrieving details for {href}: {e}")
                        subscribe_num = article_num = read_num = collect_num = 0

                    existing_categorize = Categorize.query.filter_by(href=href).first()
                    if existing_categorize:
                        existing_categorize.categorize = blog_column
                        existing_categorize.categorize_id = blog_id
                        existing_categorize.column_num = blogs_column_num
                        existing_categorize.num_span = subscribe_num
                        existing_categorize.article_num = article_num
                        existing_categorize.read_num = read_num
                        existing_categorize.collect_num = collect_num
                        print(f"更新分类信息: {blog_column}")
                    else:
                        new_categorize = Categorize(
                            href=href,
                            categorize=blog_column,
                            categorize_id=blog_id,
                            column_num=blogs_column_num,
                            num_span=subscribe_num,
                            article_num=article_num,
                            read_num=read_num,
                            collect_num=collect_num
                        )
                        db.session.add(new_categorize)
                        print(f"添加新的分类信息: {blog_column}")

                    db.session.commit()

                    time.sleep(random.uniform(1, 2))

                except Exception as inner_e:
                    print(f"Error processing category {span}: {inner_e}")
                    continue

            print("所有的分类信息已经保存到数据库")
        except Exception as e:
            print(f"错误日志: get_categorize: {e}")

    def get_blog_columns(self):
        """从 Categorize 表获取博客专栏信息"""
        blog_columns = db.session.query(
            Categorize.href,
            Categorize.categorize,
            Categorize.categorize_id,
            Categorize.article_num
        ).all()

        return [
            [column.href, column.categorize, str(column.categorize_id), str(column.article_num)]
            for column in blog_columns
        ]

    def append_blog_info(self, blog_column_url, blog_column_name, blogs):
        reply = requests.get(url=blog_column_url, headers=self.headers, timeout=self.timeout)
        blog_span = BeautifulSoup(reply.content, "lxml")
        blogs_list = blog_span.find_all('ul', attrs={'class': 'column_article_list'})

        for arch_blog_info in blogs_list:
            blogs_list = arch_blog_info.find_all('li')
            for blog_info in blogs_list:
                blog_url = blog_info.find('a', attrs={'target': '_blank'})['href']
                blog_title = blog_info.find('h2', attrs={'class': "title"}).get_text().strip().replace(" ", "_").replace(
                    '/', '_')
                statuses = blog_info.find_all("span", class_="status")
                three_status = []
                for index, status in enumerate(statuses):
                    if index == 0:
                        time_str = status.text.split('·')[0]
                        time_str = time_str.strip()
                        three_status.append(time_str)
                    else:
                        time_str = status.text.split('·')[0]
                        num = int(re.findall(r'\d+', time_str)[0])
                        three_status.append(num)

                blog_dict = {'url': blog_url, 'title': blog_title, 'date': three_status[0], 'read_num': three_status[1],
                             'comment_num': three_status[2], 'type': blog_column_name}
                blogs.append(blog_dict)
        return blogs

    def crawl_articles(self):
        """主函数：抓取并保存博客数据到数据库"""
        blog_columns = self.get_blog_columns()

        blogs = []

        for blog_column in blog_columns:
            blog_column_url = blog_column[0]
            blog_column_name = blog_column[1]
            blog_column_id = blog_column[2]
            blog_column_num = int(blog_column[3])
            print(f"正在处理专栏: {blog_column_name} , 文章数量: {blog_column_num}", "url:", blog_column_url, "id:",
                  blog_column_id)

            if blog_column_num > 40:
                page_num = round(blog_column_num / 40)

                for i in range(page_num, 0, -1):
                    url_str = blog_column_url.split('.html')[0]
                    page_url = url_str + '_' + str(i) + '.html'
                    blogs = self.append_blog_info(page_url, blog_column_name, blogs)

                blogs = self.append_blog_info(blog_column_url, blog_column_name, blogs)
            else:
                blogs = self.append_blog_info(blog_column_url, blog_column_name, blogs)

        print("--------------------------------------------------------------------------")
        print(f"已抓取 {len(blogs)} 篇文章")
        try:
            for blog in blogs:
                existing_article = Article.query.filter_by(url=blog['url']).first()

                if existing_article:
                    existing_article.title = blog['title']
                    existing_article.date = blog['date']
                    existing_article.read_num = blog['read_num']
                    existing_article.comment_num = blog['comment_num']
                    existing_article.type = blog['type']
                else:
                    new_article = Article(
                        url=blog['url'],
                        title=blog['title'],
                        date=blog['date'],
                        read_num=blog['read_num'],
                        comment_num=blog['comment_num'],
                        type=blog['type']
                    )
                    db.session.add(new_article)

            db.session.commit()
            print("文章数据已保存到数据库")
        except Exception as e:
            print(f"错误信息: {e}")
