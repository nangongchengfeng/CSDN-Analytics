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
    [0, '#f8fafc'],
    [0.2, '#e0f2fe'],
    [0.4, '#7dd3fc'],
    [0.6, '#38bdf8'],
    [1, '#0ea5e9'],
  ]

  const option: EChartsOption = {
    backgroundColor: 'transparent',
    tooltip: {
      position: 'top',
      backgroundColor: 'rgba(255, 255, 255, 0.95)',
      borderColor: 'rgba(226, 232, 240, 0.8)',
      borderWidth: 1,
      textStyle: {
        color: '#1e293b',
      },
      boxShadow: '0 4px 20px rgba(0, 0, 0, 0.08)',
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
          color: 'rgba(226, 232, 240, 0.6)',
        },
        areaStyle: {
          color: ['rgba(248, 250, 252, 0.5)', 'rgba(255, 255, 255, 0.3)'],
        },
      },
      axisLabel: {
        fontSize: 14,
        color: '#64748b',
        rotate: 45,
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: 'rgba(226, 232, 240, 0.6)',
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
          color: 'rgba(226, 232, 240, 0.6)',
        },
        areaStyle: {
          color: ['rgba(248, 250, 252, 0.5)', 'rgba(255, 255, 255, 0.3)'],
        },
      },
      axisLabel: {
        fontSize: 14,
        color: '#64748b',
      },
      axisLine: {
        show: true,
        lineStyle: {
          color: 'rgba(226, 232, 240, 0.6)',
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
        color: '#64748b',
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
        itemStyle: {
          borderRadius: 2,
        },
        emphasis: {
          itemStyle: {
            shadowBlur: 15,
            shadowColor: 'rgba(0, 0, 0, 0.12)',
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
  background-color: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  border: 1px solid rgba(226, 232, 240, 0.8);
  border-radius: 10px;
  padding: 8px 14px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.06);
  z-index: 10;
  color: #1e293b;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.year-select:hover {
  border-color: var(--accent-color);
  box-shadow: 0 6px 20px rgba(59, 130, 246, 0.15);
}
</style>
