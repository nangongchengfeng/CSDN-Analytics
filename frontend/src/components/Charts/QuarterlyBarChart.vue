<template>
  <div class="chart-container">
    <h3 class="chart-title">每季度文章写作情况</h3>
    <div ref="chartRef" class="chart"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import * as echarts from 'echarts'
import type { EChartsOption } from 'echarts'
import type { QuarterData } from '@/types'

interface Props {
  data: QuarterData[]
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'chart-click': [params: { year: string; quarter: string }]
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

  const dimensions = Object.keys(props.data[0]).filter((key) => key !== 'category' && key !== 'product')
  const source = props.data.map((item) => ({
    product: item.category || item.product,
    ...item,
  }))

  const colorList = ['#3b82f6', '#10b981', '#8b5cf6', '#f59e0b', '#ec4899']

  const option: EChartsOption = {
    backgroundColor: 'transparent',
    color: colorList,
    legend: {
      textStyle: {
        color: '#6b7280',
      },
    },
    tooltip: {
      trigger: 'axis',
      backgroundColor: 'rgba(255, 255, 255, 0.98)',
      borderColor: '#e5e7eb',
      textStyle: {
        color: '#1a1a1a',
      },
    },
    dataset: {
      dimensions: ['product', ...dimensions],
      source,
    },
    xAxis: {
      type: 'category',
      axisLabel: {
        color: '#6b7280',
      },
      axisLine: {
        lineStyle: {
          color: '#e5e7eb',
        },
      },
      axisTick: {
        lineStyle: {
          color: '#e5e7eb',
        },
      },
    },
    yAxis: {
      axisLabel: {
        color: '#6b7280',
      },
      axisLine: {
        lineStyle: {
          color: '#e5e7eb',
        },
      },
      axisTick: {
        lineStyle: {
          color: '#e5e7eb',
        },
      },
      splitLine: {
        lineStyle: {
          color: '#f3f4f6',
          type: 'dashed',
        },
      },
    },
    series: dimensions.map((dim, index) => ({
      type: 'bar',
      name: dim,
      itemStyle: {
        borderRadius: [4, 4, 0, 0],
        color: colorList[index % colorList.length],
      },
    })),
  }

  chartInstance.setOption(option)

  // 添加点击事件
  chartInstance.off('click')
  chartInstance.on('click', (params: any) => {
    // 在 dataset 模式下，从 params.value 中获取数据
    const year = Array.isArray(params.value) ? params.value[0] : params.name
    const quarter = params.seriesName
    emit('chart-click', { year, quarter })
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
