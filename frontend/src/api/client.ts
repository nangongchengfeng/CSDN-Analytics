import axios from 'axios';
import type { ApiResponse, UserInfo, Article, HeatmapData, QuarterData, CategoryData, ReadData } from '../types';

// 创建 Axios 实例
const apiClient = axios.create({
  baseURL: import.meta.env.VITE_APP_API_BASE_URL || '',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// 请求拦截器
apiClient.interceptors.request.use(
  (config) => {
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
apiClient.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    console.error('API 请求失败:', error);
    return Promise.reject(error);
  }
);

// API 接口
export const api = {
  // 获取用户信息
  getUserInfo: async (): Promise<UserInfo> => {
    const response = await apiClient.get<ApiResponse<UserInfo>>('/api/info');
    return response.data.data || {} as UserInfo;
  },

  // 获取季度文章数据
  getQuarterData: async (): Promise<QuarterData[]> => {
    const response = await apiClient.get<ApiResponse<QuarterData[]>>('/api/quarter');
    return response.data.data || [];
  },

  // 获取分类统计数据
  getCategoryData: async (): Promise<CategoryData[]> => {
    const response = await apiClient.get<ApiResponse<CategoryData[]>>('/api/categorize');
    return response.data.data || [];
  },

  // 获取阅读数据
  getReadData: async (): Promise<ReadData> => {
    const response = await apiClient.get<ApiResponse<ReadData>>('/api/read');
    return response.data.data || { labels: [], counts: [], reads: [] };
  },

  // 获取热力图数据
  getHeatmapData: async (year: string): Promise<HeatmapData> => {
    const response = await apiClient.get<ApiResponse<HeatmapData>>(`/api/heatmap/${year}`);
    return response.data.data || { data: [], xAxis: [], yAxis: [] };
  },

  // 获取文章列表
  getArticles: async (params?: { type?: string; year?: string; quarter?: string; week?: string; day?: string }): Promise<Article[]> => {
    const response = await apiClient.get<ApiResponse<Article[]>>('/api/articles', { params });
    return response.data.data || [];
  },

  // 获取年份列表
  getYears: async (): Promise<string[]> => {
    const response = await apiClient.get<ApiResponse<string[]>>('/api/years');
    return response.data.data || [];
  },
};
