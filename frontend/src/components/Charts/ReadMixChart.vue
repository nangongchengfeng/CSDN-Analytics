<template>
  <div class="chart-container">
    <h3 class="chart-title">各类型文章阅读情况</h3>
    <div ref="chartRef" class="chart"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import * as echarts from 'echarts'
import type { EChartsOption } from 'echarts'
import type { ReadData } from '@/types'

interface Props {
  data: ReadData
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
  if (!chartInstance || !props.data.labels?.length) return

  const option: EChartsOption = {
    backgroundColor: 'transparent',
    color: ['#3b82f6', '#f59e0b'],
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow',
      },
      backgroundColor: 'rgba(255, 255, 255, 0.98)',
      borderColor: '#e5e7eb',
      textStyle: {
        color: '#1a1a1a',
      },
    },
    legend: {
      data: ['文章数', '阅读量'],
      textStyle: {
        color: '#6b7280',
      },
    },
    grid: {
      top: '10%',
      bottom: '25%',
      right: '10%',
    },
    xAxis: [
      {
        type: 'category',
        axisTick: {
          alignWithLabel: true,
        },
        axisLabel: {
          rotate: 90,
          interval: 0,
          color: '#6b7280',
          formatter: (value: string) => {
            return value.length > 15 ? value.substring(0, 15) + '\n' + value.substring(10) : value
          },
          fontSize: 10,
        },
        axisLine: {
          lineStyle: {
            color: '#e5e7eb',
          },
        },
        data: props.data.labels,
      },
    ],
    yAxis: [
      {
        type: 'value',
        name: '文章数',
        position: 'left',
        axisLine: {
          show: true,
          lineStyle: {
            color: '#3b82f6',
          },
        },
        axisLabel: {
          formatter: '{value}',
          color: '#6b7280',
        },
        splitLine: {
          lineStyle: {
            type: 'dashed',
            color: '#f3f4f6',
          },
        },
      },
      {
        type: 'value',
        name: '阅读量',
        position: 'right',
        alignTicks: true,
        axisLine: {
          show: true,
          lineStyle: {
            color: '#f59e0b',
          },
        },
        axisLabel: {
          formatter: '{value}',
          color: '#6b7280',
        },
        splitLine: {
          show: false,
        },
      },
    ],
    series: [
      {
        name: '文章数',
        type: 'bar',
        data: props.data.counts,
        yAxisIndex: 0,
        barWidth: '40%',
        itemStyle: {
          color: '#3b82f6',
          barBorderRadius: [4, 4, 0, 0],
        },
      },
      {
        name: '阅读量',
        type: 'line',
        data: props.data.reads,
        yAxisIndex: 1,
        smooth: true,
        lineStyle: {
          color: '#f59e0b',
          width: 3,
        },
        itemStyle: {
          color: '#f59e0b',
        },
      },
    ],
  }

  chartInstance.setOption(option)

  // 添加点击事件
  chartInstance.off('click')
  chartInstance.on('click', (params: any) => {
    const name = props.data.labels[params.dataIndex]
    emit('chart-click', { name })
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
