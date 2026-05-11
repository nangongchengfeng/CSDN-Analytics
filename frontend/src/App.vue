<template>
  <div class="dashboard-container">
    <div class="dashboard-glow dashboard-glow-primary"></div>
    <div class="dashboard-glow dashboard-glow-accent"></div>

    <!-- 顶部：用户信息 + 统计数据（同一个卡片） -->
    <div class="top-section">
      <MacWindow class="top-card">
        <div class="top-content">
          <UserProfile :user-info="userInfo" />
          <StatsGrid :user-info="userInfo" />
        </div>
      </MacWindow>
    </div>

    <div class="charts-container">
      <div class="chart-row overview-row">
        <MacWindow class="chart-window chart-window-quarter">
          <template #default>
            <QuarterlyBarChart :data="quarterData" @chart-click="handleQuarterClick" />
          </template>
        </MacWindow>

        <MacWindow class="chart-window chart-window-category">
          <template #default>
            <CategoryPieChart :data="categoryData" @chart-click="handleCategoryClick" />
          </template>
        </MacWindow>

        <MacWindow class="chart-window chart-window-read">
          <template #default>
            <ReadMixChart :data="readData" @chart-click="handleReadClick" />
          </template>
        </MacWindow>
      </div>

      <div class="chart-row insight-row">
        <MacWindow class="chart-window chart-window-heatmap">
          <template #default>
            <HeatmapChart
              :data="heatmapData"
              :years="years"
              :selected-year="selectedYear"
              @year-change="handleYearChange"
              @chart-click="handleHeatmapClick"
            />
          </template>
        </MacWindow>

        <MacWindow class="chart-window chart-window-articles">
          <template #default>
            <ArticleList
              :articles="articles"
              :current-filter="currentFilter"
              @clear-filter="handleClearFilter"
            />
          </template>
        </MacWindow>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import MacWindow from '@/components/Layout/MacWindow.vue'
import UserProfile from '@/components/Dashboard/UserProfile.vue'
import StatsGrid from '@/components/Dashboard/StatsGrid.vue'
import QuarterlyBarChart from '@/components/Charts/QuarterlyBarChart.vue'
import CategoryPieChart from '@/components/Charts/CategoryPieChart.vue'
import ReadMixChart from '@/components/Charts/ReadMixChart.vue'
import HeatmapChart from '@/components/Charts/HeatmapChart.vue'
import ArticleList from '@/components/Charts/ArticleList.vue'
import { api } from '@/api/client'
import type { UserInfo, QuarterData, CategoryData, ReadData, HeatmapData, Article } from '@/types'

// 响应式数据
const userInfo = ref<UserInfo>({
  author_name: '',
  article_num: 0,
  fans_num: 0,
  like_num: 0,
  collect_num: 0,
  visit_num: 0,
  rank: 0,
  share_num: 0,
  code_age: 0,
})

const quarterData = ref<QuarterData[]>([])
const categoryData = ref<CategoryData[]>([])
const readData = ref<ReadData>({ labels: [], counts: [], reads: [] })
const heatmapData = ref<HeatmapData>({ data: [], xAxis: [], yAxis: [] })
const articles = ref<Article[]>([])
const years = ref<string[]>([])
const selectedYear = ref<string>('2024')
const currentFilter = ref<{ type?: string; year?: string; quarter?: string; week?: number; weekday?: number }>({})

// 初始化数据
onMounted(async () => {
  await Promise.all([
    fetchUserInfo(),
    fetchQuarterData(),
    fetchCategoryData(),
    fetchReadData(),
    fetchYears(),
  ])
  // 等待年份列表加载完成后，设置默认年份为最新的，并加载对应数据
  if (years.value.length > 0) {
    selectedYear.value = years.value[0]
    await Promise.all([
      fetchHeatmapData(selectedYear.value),
      fetchArticles(),
    ])
  }
})

// 获取用户信息
const fetchUserInfo = async () => {
  try {
    const data = await api.getUserInfo()
    userInfo.value = data
  } catch (error) {
    console.error('获取用户信息失败:', error)
  }
}

// 获取季度数据
const fetchQuarterData = async () => {
  try {
    const data = await api.getQuarterData()
    quarterData.value = data
  } catch (error) {
    console.error('获取季度数据失败:', error)
  }
}

// 获取分类数据
const fetchCategoryData = async () => {
  try {
    const data = await api.getCategoryData()
    categoryData.value = data
  } catch (error) {
    console.error('获取分类数据失败:', error)
  }
}

// 获取阅读数据
const fetchReadData = async () => {
  try {
    const data = await api.getReadData()
    readData.value = data
  } catch (error) {
    console.error('获取阅读数据失败:', error)
  }
}

// 获取年份列表
const fetchYears = async () => {
  try {
    const data = await api.getYears()
    years.value = data
  } catch (error) {
    console.error('获取年份列表失败:', error)
  }
}

// 获取热力图数据
const fetchHeatmapData = async (year: string) => {
  try {
    const data = await api.getHeatmapData(year)
    heatmapData.value = data
  } catch (error) {
    console.error('获取热力图数据失败:', error)
  }
}

// 获取文章列表（调用后端API）
const fetchArticles = async (params?: { type?: string; year?: string; quarter?: string; week?: string; day?: string }) => {
  try {
    const data = await api.getArticles(params)
    articles.value = data
  } catch (error) {
    console.error('获取文章列表失败:', error)
  }
}

// 处理年份变化
const handleYearChange = async (year: string) => {
  selectedYear.value = year
  await fetchHeatmapData(year)
}

// 处理季度图表点击
const handleQuarterClick = async (params: { year: string; quarter: string }) => {
  currentFilter.value = { year: params.year, quarter: params.quarter }
  await fetchArticles({ year: params.year, quarter: params.quarter })
}

// 处理分类图表点击
const handleCategoryClick = async (params: { name: string }) => {
  currentFilter.value = { type: params.name }
  await fetchArticles({ type: params.name })
}

// 处理阅读量图表点击
const handleReadClick = async (params: { name: string }) => {
  currentFilter.value = { type: params.name }
  await fetchArticles({ type: params.name })
}

// 处理热力图点击
const handleHeatmapClick = async (params: { year: string; week: number; weekday: number }) => {
  currentFilter.value = { year: params.year, week: params.week, weekday: params.weekday }
  // 注意：后端 API 使用 week 和 day 参数（day 是 1-7）
  await fetchArticles({
    year: params.year,
    week: params.week.toString(),
    day: (params.weekday + 1).toString(),
  })
}

// 清除筛选
const handleClearFilter = async () => {
  currentFilter.value = {}
  await fetchArticles()
}
</script>

<style scoped>
.dashboard-container {
  position: relative;
  display: flex;
  flex-direction: column;
  height: 100vh;
  padding: 24px;
  gap: 18px;
  overflow: hidden;
  isolation: isolate;
}

.dashboard-container::before {
  content: '';
  position: absolute;
  inset: 0;
  z-index: -2;
  background:
    linear-gradient(rgba(59, 130, 246, 0.06) 1px, transparent 1px),
    linear-gradient(90deg, rgba(59, 130, 246, 0.06) 1px, transparent 1px);
  background-size: 42px 42px;
  mask-image: linear-gradient(180deg, rgba(0, 0, 0, 0.7), transparent 72%);
  pointer-events: none;
}

.dashboard-glow {
  position: absolute;
  z-index: -1;
  width: 38vw;
  height: 38vw;
  border-radius: 999px;
  filter: blur(26px);
  opacity: 0.7;
  pointer-events: none;
}

.dashboard-glow-primary {
  top: -24vw;
  right: -10vw;
  background: radial-gradient(circle, rgba(59, 130, 246, 0.22), transparent 62%);
}

.dashboard-glow-accent {
  left: -20vw;
  bottom: -24vw;
  background: radial-gradient(circle, rgba(245, 158, 11, 0.16), transparent 64%);
}

.top-section {
  height: clamp(126px, 14.5vh, 156px);
  flex-shrink: 0;
  min-height: 0;
}

.top-content {
  display: flex;
  gap: 18px;
  height: 100%;
  align-items: center;
  position: relative;
}

.top-content > *:first-child {
  flex: 0 0 27%;
  min-width: 280px;
}

.top-content > *:last-child {
  flex: 1;
  min-width: 0;
}

.charts-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 18px;
  min-height: 0;
}

.chart-row {
  display: grid;
  gap: 18px;
  min-height: 0;
}

.overview-row {
  grid-template-columns: minmax(0, 1.15fr) minmax(280px, 0.85fr) minmax(0, 1.25fr);
  flex: 0 0 42%;
}

.insight-row {
  grid-template-columns: minmax(0, 1.55fr) minmax(340px, 0.95fr);
  flex: 1;
}

.chart-window {
  min-width: 0;
  min-height: 0;
}

.top-card {
  position: relative;
}

.top-card::before {
  content: '';
  position: absolute;
  inset: 0;
  background:
    radial-gradient(circle at 13% 12%, rgba(59, 130, 246, 0.18), transparent 26%),
    linear-gradient(90deg, rgba(30, 64, 175, 0.08), transparent 58%);
  pointer-events: none;
}

:deep(.window-content) {
  height: 100%;
  padding: 20px;
}

@media (max-width: 1280px) {
  .dashboard-container {
    height: auto;
    min-height: 100vh;
    overflow: auto;
  }

  .top-section {
    height: auto;
  }

  .top-content,
  .chart-row {
    grid-template-columns: 1fr;
  }

  .top-content {
    display: grid;
  }

  .top-content > *:first-child {
    min-width: 0;
  }

  .overview-row,
  .insight-row {
    flex: none;
  }

  .chart-window {
    min-height: 360px;
  }
}

@media (max-width: 768px) {
  .dashboard-container {
    padding: 14px;
  }

  :deep(.window-content) {
    padding: 16px;
  }
}
</style>
