<template>
  <div class="chart-container">
    <h3 class="chart-title">各类型文章占比情况</h3>
    <div ref="chartRef" class="chart"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import * as echarts from 'echarts'
import type { EChartsOption } from 'echarts'
import type { CategoryData } from '@/types'

interface Props {
  data: CategoryData[]
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'chart-click': [params: { name: string }]
}>()

const chartRef = ref<HTMLElement | null>(null)
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

const updateChart = () => {
  if (!chartInstance || !props.data.length) return

  const chartData = props.data.map((item) => ({
    value: item.value,
    name: item.name,
  }))

  const colorList = ['#3b82f6', '#10b981', '#8b5cf6', '#f59e0b', '#ec4899', '#06b6d4', '#f97316']

  const option: EChartsOption = {
    backgroundColor: 'transparent',
    color: colorList,
    title: {
      left: 'center',
    },
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)',
      backgroundColor: 'rgba(255, 255, 255, 0.98)',
      borderColor: '#e5e7eb',
      textStyle: {
        color: '#1a1a1a',
      },
    },
    legend: {
      show: false,
    },
    series: [
      {
        name: '分类统计',
        type: 'pie',
        radius: ['30%', '70%'],
        data: chartData,
        label: {
          show: true,
          position: 'outside',
          formatter: '{b}: {d}%',
          color: '#6b7280',
        },
        labelLine: {
          show: true,
          lineStyle: {
            color: '#e5e7eb',
          },
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
    emit('chart-click', { name: params.name })
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
</style>
