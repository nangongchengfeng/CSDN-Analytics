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
  align-items: flex-start;
  margin-bottom: 14px;
  gap: 14px;
}

.chart-title {
  color: #0f172a;
  font-weight: 700;
  font-size: 1.15rem;
  margin: 0;
  letter-spacing: -0.01em;
  display: flex;
  align-items: center;
  gap: 8px;
}

.chart-title::before {
  content: '';
  display: inline-block;
  width: 4px;
  height: 20px;
  background: linear-gradient(135deg, var(--accent-color), #60a5fa);
  border-radius: 2px;
}

.filter-info {
  display: flex;
  align-items: center;
  gap: 10px;
  min-width: 0;
}

.filter-label {
  max-width: 220px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: var(--accent-color);
  font-size: 0.78rem;
  font-weight: 700;
  background: rgba(30, 64, 175, 0.08);
  border: 1px solid rgba(59, 130, 246, 0.14);
  border-radius: 999px;
  padding: 7px 10px;
}

.clear-btn {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.95), rgba(248, 250, 252, 0.9));
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border: 1px solid rgba(226, 232, 240, 0.9);
  color: var(--accent-color);
  padding: 8px 18px;
  border-radius: 12px;
  font-size: 0.85rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.22, 1, 0.36, 1);
  box-shadow: 0 4px 12px rgba(15, 23, 42, 0.04),
              0 2px 6px rgba(15, 23, 42, 0.02),
              0 1px 0 rgba(255, 255, 255, 0.9) inset;
}

.clear-btn:hover {
  background: linear-gradient(135deg, var(--accent-color), #60a5fa);
  color: white;
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(59, 130, 246, 0.25),
              0 4px 12px rgba(59, 130, 246, 0.15),
              0 1px 0 rgba(255, 255, 255, 0.2) inset;
  border-color: transparent;
}

.article-list {
  flex: 1;
  overflow: auto;
  border: 1px solid rgba(226, 232, 240, 0.72);
  border-radius: 16px;
  background: rgba(255, 255, 255, 0.48);
}

.article-table {
  width: 100%;
  border-collapse: collapse;
  text-align: left;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  font-size: 13px;
}

.article-table th {
  position: sticky;
  top: 0;
  z-index: 1;
  padding: 12px 14px;
  border-bottom: 2px solid rgba(226, 232, 240, 0.8);
  background: rgba(248, 250, 252, 0.92);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  color: var(--text-secondary);
  font-size: 12px;
  font-weight: 800;
  letter-spacing: 0.04em;
}

.article-table td {
  padding: 12px 14px;
  border-bottom: 1px solid rgba(226, 232, 240, 0.6);
  color: var(--text-primary);
}

.article-table tr:hover {
  background: linear-gradient(90deg, rgba(59, 130, 246, 0.08), rgba(255, 255, 255, 0.72));
}

.article-table a {
  text-decoration: none;
  color: var(--accent-color);
  font-size: 14px;
  font-weight: 700;
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
