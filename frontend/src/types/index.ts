// 全局类型定义
export interface ApiResponse<T = any> {
  code: number;
  message: string;
  data: T;
}

// 用户信息类型
export interface UserInfo {
  author_name: string;
  article_num: number;
  fans_num: number;
  like_num: number;
  collect_num: number;
  visit_num: number;
  rank: number;
  share_num: number;
  code_age: number;
}

// 文章类型
export interface Article {
  title: string;
  type: string;
  date: string;
  url: string;
}

// 热力图数据类型
export interface HeatmapData {
  data: number[][];
  xAxis: string[];
  yAxis: string[];
}

// 季度文章数据类型
export interface QuarterData {
  category: string;
  [key: string]: string | number;
}

// 分类统计数据类型
export interface CategoryData {
  name: string;
  value: number;
}

// 阅读数据类型
export interface ReadData {
  labels: string[];
  counts: number[];
  reads: number[];
}
