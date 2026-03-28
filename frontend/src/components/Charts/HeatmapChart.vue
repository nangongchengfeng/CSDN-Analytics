<template>
  <div class="chart-container">
    <h3 class="chart-title" ref="titleRef">
      {{ selectedYear }} 年 写作发布热力图
    </h3>
    <div ref="chartRef" class="chart"></div>
    <select class="year-select" v-model="selectedYear" @change="handleYearChange">
      <option
        v-for="year in years"
        :key="year"
        :value="year"
      >
        {{ year }}年
      </option>
    </select>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import * as echarts from 'echarts'
import type { EChartsOption } from 'echarts'
import type { HeatmapData } from '@/types'

interface Props {
  data: HeatmapData
  years: string[]
  selectedYear: string
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'year-change': [year: string]
  'chart-click': [params: { year: string; week: number; weekday: number }]
}>()

const chartRef = ref<HTMLElement | null>(null)
const titleRef = ref<HTMLElement | null>(null)
const selectedYear = ref(props.selectedYear)
let chartInstance: echarts.ECharts | null = null

onMounted(() => {
  if (chartRef.value) {
    chartInstance = echarts.init(chartRef.value)
    updateChart()
  }
})

watch(
  () => props.data,
  () => {
    updateChart()
  },
  { deep: true }
)

watch(
  () => props.selectedYear,
  (newYear) => {
    selectedYear.value = newYear
  }
)

const handleYearChange = () => {
  emit('year-change', selectedYear.value)
}

const updateChart = () => {
  if (!chartInstance || !props.data.data) return

  const colorScale = [
    [0, '#f3f4f6'],
    [0.2, '#dbeafe'],
    [0.4, '#93c5fd'],
    [0.6, '#3b82f6'],
    [1, '#1d4ed8'],
  ]

  const option: EChartsOption = {
    backgroundColor: 'transparent',
    tooltip: {
      position: 'top',
      backgroundColor: 'rgba(255, 255, 255, 0.98)',
      borderColor: '#e5e7eb',
      textStyle: {
        color: '#1a1a1a',
      },
      formatter: (params: any) => {
        const week = params.value[0] + 1
        const count = params.value[2]
        return `${props.data.xAxis[params.value[0]]}<br>星期: ${props.data.yAxis[params.value[1]]}<br>文章数: ${count}`
      },
    },
    grid: {
      height: '70%',
      width: '75%',
      top: '10%',
      left: '10%',
    },
    xAxis: {
      type: 'category',
      data: props.data.xAxis,
      splitArea: {
        show: true,
        lineStyle: {
          width: 1,
          type: 'dashed',
          color: '#e5e7eb',
        },
        areaStyle: {
          color: ['#f9fafb', '#ffffff'],
        },
      },
      axisLabel: {
        fontSize: 14,
        color: '#6b7280',
        rotate: 45,
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e5e7eb',
        },
      },
    },
    yAxis: {
      type: 'category',
      data: props.data.yAxis,
      splitArea: {
        show: true,
        lineStyle: {
          width: 1,
          type: 'dashed',
          color: '#e5e7eb',
        },
        areaStyle: {
          color: ['#f9fafb', '#ffffff'],
        },
      },
      axisLabel: {
        fontSize: 14,
        color: '#6b7280',
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: '#e5e7eb',
        },
      },
    },
    visualMap: {
      min: 0,
      max: props.data.data && props.data.data.length > 0
        ? Math.max(1, ...props.data.data.map((item) => item[2] || 0))
        : 1,
      calculable: true,
      orient: 'vertical',
      right: '5%',
      top: 'middle',
      inRange: {
        color: colorScale.map((item) => item[1]),
      },
      textStyle: {
        fontSize: 12,
        color: '#6b7280',
      },
    },
    series: [
      {
        name: '文章发布热力图',
        type: 'heatmap',
        data: props.data.data,
        label: {
          show: false,
        },
        emphasis: {
          itemStyle: {
            shadowBlur: 10,
            shadowColor: 'rgba(0, 0, 0, 0.15)',
          },
        },
      },
    ],
  }

  chartInstance.setOption(option)

  // 添加点击事件
  chartInstance.off('click')
  chartInstance.on('click', (params: any) => {
    const week = params.value[0] + 1
    const weekday = params.value[1]
    emit('chart-click', { year: selectedYear.value, week, weekday })
  })
}

window.addEventListener('resize', () => {
  chartInstance?.resize()
})
</script>

<style scoped>
.chart-container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  position: relative;
}

.chart-title {
  color: var(--accent-color);
  font-weight: 600;
  margin-bottom: 12px;
  font-size: 1rem;
}

.chart {
  flex: 1;
  min-height: 200px;
}

.year-select {
  position: absolute;
  top: 15px;
  right: 15px;
  background-color: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  z-index: 10;
  color: #1a1a1a;
  font-size: 12px;
  cursor: pointer;
}

.year-select:hover {
  border-color: #3b82f6;
}
</style>
