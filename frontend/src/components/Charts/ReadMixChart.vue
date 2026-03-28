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
    color: ['#60a5fa', '#fbbf24'],
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow',
      },
      backgroundColor: 'rgba(255, 255, 255, 0.95)',
      borderColor: 'rgba(226, 232, 240, 0.8)',
      borderWidth: 1,
      textStyle: {
        color: '#1e293b',
      },
      boxShadow: '0 4px 20px rgba(0, 0, 0, 0.08)',
    },
    legend: {
      data: ['文章数', '阅读量'],
      textStyle: {
        color: '#64748b',
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
          color: '#64748b',
          formatter: (value: string) => {
            return value.length > 15 ? value.substring(0, 15) + '\n' + value.substring(10) : value
          },
          fontSize: 10,
        },
        axisLine: {
          lineStyle: {
            color: 'rgba(226, 232, 240, 0.6)',
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
            color: '#60a5fa',
          },
        },
        axisLabel: {
          formatter: '{value}',
          color: '#64748b',
        },
        splitLine: {
          lineStyle: {
            type: 'dashed',
            color: 'rgba(241, 245, 249, 0.8)',
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
            color: '#fbbf24',
          },
        },
        axisLabel: {
          formatter: '{value}',
          color: '#64748b',
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
          color: '#60a5fa',
          borderRadius: [6, 6, 0, 0],
        },
      },
      {
        name: '阅读量',
        type: 'line',
        data: props.data.reads,
        yAxisIndex: 1,
        smooth: true,
        lineStyle: {
          color: '#fbbf24',
          width: 3,
        },
        itemStyle: {
          color: '#fbbf24',
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
