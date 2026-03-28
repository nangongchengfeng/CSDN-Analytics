<template>
  <div class="chart-container">
    <div class="chart-header">
      <h3 class="chart-title">
        {{ selectedYear }} 年 写作发布热力图
      </h3>
      <div class="year-select-wrapper">
        <span class="year-label">选择年份:</span>
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
    </div>
    <div ref="chartRef" class="chart"></div>
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
      extraCssText: 'box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);',
      formatter: (params: any) => {
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
      splitLine: {
        show: true,
        lineStyle: {
          width: 1,
          type: 'dashed',
          color: 'rgba(226, 232, 240, 0.6)',
        },
      },
      splitArea: {
        show: true,
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
      splitLine: {
        show: true,
        lineStyle: {
          width: 1,
          type: 'dashed',
          color: 'rgba(226, 232, 240, 0.6)',
        },
      },
      splitArea: {
        show: true,
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

.chart-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 14px;
  gap: 16px;
}

.chart-title {
  color: #0f172a;
  font-weight: 700;
  margin: 0;
  font-size: 1.15rem;
  letter-spacing: -0.01em;
  display: flex;
  align-items: center;
  gap: 8px;
  white-space: nowrap;
}

.chart-title::before {
  content: '';
  display: inline-block;
  width: 4px;
  height: 20px;
  background: linear-gradient(135deg, var(--accent-color), #60a5fa);
  border-radius: 2px;
}

.year-select-wrapper {
  display: flex;
  align-items: center;
  gap: 10px;
  flex-shrink: 0;
}

.year-label {
  font-size: 0.85rem;
  color: #64748b;
  font-weight: 500;
  white-space: nowrap;
}

.year-select {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.95), rgba(248, 250, 252, 0.9));
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border: 1px solid rgba(226, 232, 240, 0.9);
  border-radius: 12px;
  padding: 10px 16px;
  box-shadow: 0 6px 20px rgba(15, 23, 42, 0.06),
              0 2px 8px rgba(15, 23, 42, 0.04),
              0 1px 0 rgba(255, 255, 255, 0.9) inset;
  z-index: 10;
  color: #1e293b;
  font-size: 0.9rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.22, 1, 0.36, 1);
  outline: none;
  min-width: 100px;
}

.year-select:hover {
  border-color: var(--accent-color);
  box-shadow: 0 10px 30px rgba(59, 130, 246, 0.15),
              0 4px 12px rgba(15, 23, 42, 0.08),
              0 1px 0 rgba(255, 255, 255, 0.9) inset;
  transform: translateY(-1px);
}

.year-select:focus {
  border-color: var(--accent-color);
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1),
              0 10px 30px rgba(59, 130, 246, 0.15);
}

.year-select option {
  background: #ffffff;
  color: #1e293b;
  padding: 8px 16px;
  font-weight: 500;
}

.chart {
  flex: 1;
  min-height: 200px;
}
</style>
