"""CSDN 页面解析器"""
from bs4 import BeautifulSoup


class CSDNParser:
    """CSDN 页面解析器"""

    @staticmethod
    def parse_user_info(html):
        """解析用户信息"""
        soup = BeautifulSoup(html, 'lxml')
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

        code_age_div = soup.find('div', class_='person-code-age')
        code_age_text = None
        if code_age_div:
            span_tag = code_age_div.find('span')
            if span_tag:
                code_age_text = span_tag.get_text(strip=True)

        author_name = soup.find('div', class_='user-profile-head-name').find('div').get_text(strip=True)

        img_tag_css = soup.select_one('div.user-profile-avatar img')
        image_url_css = img_tag_css.get('src') if img_tag_css else None

        return {
            'statistics': extracted_statistics,
            'achievements': extracted_achievements,
            'code_age': code_age_text,
            'author_name': author_name,
            'head_img': image_url_css
        }

    @staticmethod
    def parse_categories(html):
        """解析分类信息"""
        soup = BeautifulSoup(html, 'lxml')
        spans = soup.find_all('a', attrs={'class': 'special-column-name'})
        categories = []

        for span in spans:
            href = span.get('href')
            if not href:
                continue

            blog_column = span.text.strip()
            blog_id = href.split("_")[-1].split(".")[0]

            num_span = span.find('span', class_='special-column-num')
            if num_span:
                blogs_column_num = int(num_span.text)
            else:
                blogs_column_num = 0

            categories.append({
                'href': href,
                'categorize': blog_column,
                'categorize_id': blog_id,
                'column_num': blogs_column_num
            })

        return categories

    @staticmethod
    def parse_category_details(html):
        """解析分类详情页"""
        soup = BeautifulSoup(html, 'html.parser')
        column_operating_div = soup.find('div', {'class': 'column_operating'})

        if column_operating_div:
            subscribe_num = int(column_operating_div.find('span', {
                'class': 'column-subscribe-num'}).text) if column_operating_div.find('span', {
                'class': 'column-subscribe-num'}) else 0
            mumber_spans = column_operating_div.find_all('span', {'class': 'mumber-color'})

            article_num = int(mumber_spans[1].text) if len(mumber_spans) > 1 else 0
            read_num = int(mumber_spans[2].text) if len(mumber_spans) > 2 else 0
            collect_num = int(mumber_spans[3].text) if len(mumber_spans) > 3 else 0
        else:
            subscribe_num = article_num = read_num = collect_num = 0

        return {
            'subscribe_num': subscribe_num,
            'article_num': article_num,
            'read_num': read_num,
            'collect_num': collect_num
        }

    @staticmethod
    def parse_articles(html, category_name):
        """解析文章列表"""
        soup = BeautifulSoup(html, 'lxml')
        blogs = []
        blogs_list = soup.find_all('ul', attrs={'class': 'column_article_list'})

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
                        import re
                        num = int(re.findall(r'\d+', time_str)[0])
                        three_status.append(num)

                blogs.append({
                    'url': blog_url,
                    'title': blog_title,
                    'date': three_status[0],
                    'read_num': three_status[1],
                    'comment_num': three_status[2],
                    'type': category_name
                })
        return blogs
