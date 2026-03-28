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

  const colorList = ['#60a5fa', '#34d399', '#a78bfa', '#fbbf24', '#f472b6']

  const option: EChartsOption = {
    backgroundColor: 'transparent',
    color: colorList,
    legend: {
      textStyle: {
        color: '#64748b',
      },
    },
    tooltip: {
      trigger: 'axis',
      backgroundColor: 'rgba(255, 255, 255, 0.95)',
      borderColor: 'rgba(226, 232, 240, 0.8)',
      borderWidth: 1,
      textStyle: {
        color: '#1e293b',
      },
      boxShadow: '0 4px 20px rgba(0, 0, 0, 0.08)',
    },
    dataset: {
      dimensions: ['product', ...dimensions],
      source,
    },
    xAxis: {
      type: 'category',
      axisLabel: {
        color: '#64748b',
      },
      axisLine: {
        lineStyle: {
          color: 'rgba(226, 232, 240, 0.6)',
        },
      },
      axisTick: {
        lineStyle: {
          color: 'rgba(226, 232, 240, 0.6)',
        },
      },
    },
    yAxis: {
      axisLabel: {
        color: '#64748b',
      },
      axisLine: {
        lineStyle: {
          color: 'rgba(226, 232, 240, 0.6)',
        },
      },
      axisTick: {
        lineStyle: {
          color: 'rgba(226, 232, 240, 0.6)',
        },
      },
      splitLine: {
        lineStyle: {
          color: 'rgba(241, 245, 249, 0.8)',
          type: 'dashed',
        },
      },
    },
    series: dimensions.map((dim, index) => ({
      type: 'bar',
      name: dim,
      itemStyle: {
        borderRadius: [6, 6, 0, 0],
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
  color: #0f172a;
  font-weight: 700;
  margin-bottom: 14px;
  font-size: 1.15rem;
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

.chart {
  flex: 1;
  min-height: 200px;
}
</style>
