<template>
  <div class="article-list-container">
    <div class="header-section">
      <h3 class="chart-title">文章列表</h3>
      <div v-if="hasFilter" class="filter-info">
        <span class="filter-label">{{ filterText }}</span>
        <button class="clear-btn" @click="handleClear">清除筛选</button>
      </div>
    </div>
    <div class="article-list">
      <table v-if="articles.length" class="article-table">
        <thead>
          <tr>
            <th>标题</th>
            <th>类型</th>
            <th>日期</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(article, index) in articles" :key="index">
            <td>
              <a :href="article.url" target="_blank" rel="noopener noreferrer">{{ article.title }}</a>
            </td>
            <td>{{ article.type }}</td>
            <td>{{ article.date }}</td>
          </tr>
        </tbody>
      </table>
      <div v-else class="no-data">暂无文章数据</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { Article } from '@/types'

interface Props {
  articles: Article[]
  currentFilter?: { type?: string; year?: string; quarter?: string; week?: number; weekday?: number }
}

const props = withDefaults(defineProps<Props>(), {
  currentFilter: () => ({}),
})

const emit = defineEmits<{
  'clear-filter': []
}>()

const hasFilter = computed(() => {
  const { type, year, quarter, week, weekday } = props.currentFilter
  return !!(type || year || quarter || (week !== undefined) || (weekday !== undefined))
})

const filterText = computed(() => {
  const { type, year, quarter, week, weekday } = props.currentFilter
  const parts: string[] = []

  if (type) {
    parts.push(`类型: ${type}`)
  }
  if (year && (quarter || (week === undefined && weekday === undefined))) {
    parts.push(`${year}年`)
  }
  if (quarter) {
    parts.push(quarter)
  }
  if (week !== undefined && weekday !== undefined) {
    const weekdayNames = ['星期一', '星期二', '星期三', '星期四', '星期五', '星期六', '星期日']
    parts.push(`${year}年 · 第${week}周 · ${weekdayNames[weekday]}`)
  } else if (week !== undefined) {
    parts.push(`${year}年 · 第${week}周`)
  } else if (weekday !== undefined) {
    const weekdayNames = ['星期一', '星期二', '星期三', '星期四', '星期五', '星期六', '星期日']
    parts.push(weekdayNames[weekday])
  }

  return parts.join(' · ')
})

const handleClear = () => {
  emit('clear-filter')
}
</script>

<style scoped>
.article-list-container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
}

.header-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.chart-title {
  color: var(--accent-color);
  font-weight: 600;
  font-size: 1rem;
  margin: 0;
}

.filter-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.filter-label {
  color: var(--text-secondary);
  font-size: 0.85rem;
}

.clear-btn {
  background-color: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  border: 1px solid rgba(226, 232, 240, 0.8);
  color: var(--accent-color);
  padding: 6px 16px;
  border-radius: 10px;
  font-size: 0.85rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

.clear-btn:hover {
  background-color: var(--accent-color);
  color: white;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.25);
}

.article-list {
  flex: 1;
  overflow: auto;
}

.article-table {
  width: 100%;
  border-collapse: collapse;
  text-align: left;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  font-size: 13px;
}

.article-table th {
  padding: 10px 12px;
  border-bottom: 2px solid rgba(226, 232, 240, 0.8);
  background-color: rgba(248, 250, 252, 0.6);
  color: var(--text-secondary);
  font-size: 13px;
  font-weight: 600;
}

.article-table td {
  padding: 10px 12px;
  border-bottom: 1px solid rgba(226, 232, 240, 0.6);
  color: var(--text-primary);
}

.article-table tr:hover {
  background-color: rgba(248, 250, 252, 0.8);
}

.article-table a {
  text-decoration: none;
  color: var(--accent-color);
  font-size: 14px;
  font-weight: 500;
  transition: color 0.2s;
}

.article-table a:hover {
  text-decoration: underline;
  color: #2563eb;
}

.no-data {
  text-align: center;
  padding: 40px;
  color: var(--text-secondary);
}
</style>
