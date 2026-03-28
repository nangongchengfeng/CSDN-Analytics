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
  background-color: #f9fafb;
  border: 1px solid var(--border-color);
  color: var(--accent-color);
  padding: 4px 12px;
  border-radius: 6px;
  font-size: 0.8rem;
  cursor: pointer;
  transition: all 0.2s;
}

.clear-btn:hover {
  background-color: var(--accent-color);
  color: white;
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
  font-size: 12px;
}

.article-table th {
  padding: 8px;
  border-bottom: 2px solid var(--border-color);
  background-color: #f9fafb;
  color: var(--text-secondary);
  font-size: 14px;
  font-weight: bold;
}

.article-table td {
  padding: 8px;
  border-bottom: 1px solid var(--border-color);
  color: var(--text-primary);
}

.article-table tr:hover {
  background-color: #f9fafb;
}

.article-table a {
  text-decoration: none;
  color: var(--accent-color);
  font-size: 14px;
  font-weight: 500;
}

.article-table a:hover {
  text-decoration: underline;
}

.no-data {
  text-align: center;
  padding: 40px;
  color: var(--text-secondary);
}
</style>
