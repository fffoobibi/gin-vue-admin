<template>
  <div>
    <div ref="echart" class="dashboard-line" />
  </div>
</template>
<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, watch, toRefs } from 'vue'
import { useWindowResize } from '@/hooks/use-windows-resize'
import { getFirstCrawlInfo } from '@/api/tblCrawlStats'
import { groupData } from '@/utils/reports'

defineOptions({
  name: 'LineReports'
})

const props = defineProps({
  report: {
    type: Number,
    required: true,
  },
  startTime: {
    type: String,
    required: true,
  },
  endTime: {
    type: String,
    required: true,
  },
  group: {
    type: String,
    required: true
  }

})

let chart = null
const echart = ref(null)
const { startTime, endTime, report, group } = toRefs(props)

watch([startTime, endTime, group, report], async(newVal, oldVal) => {
  chart.showLoading({ text: '' })
  await displayReports()
  chart.hideLoading()
}, { immediate: false })

useWindowResize(() => {
  if (!chart) {
    return
  }
  chart.resize()
})

const displayChart = (legendData, xdata, ydata) => {
  chart.setOption({
    legend: {
      textStyle: { fontWeight: 'bold' },
      data: legendData,
      y: 'bottom'
    },
    xAxis: {
      type: 'category',
      data: xdata,
    },
    yAxis: {
      type: 'value'
    },
    tooltip: {
      show: true
    },
    series: ydata,
  })
}

function fillData(dataArray, startDateStr, endDateStr, group) {
  const startDate = new Date(startDateStr)
  const endDate = new Date(endDateStr)
  const result = []
  const keys = []
  const parsed = []

  // 将数据数组转换为以日期为键的对象
  const dataObject = {}
  for (let i = 0; i < dataArray.length; i++) {
    dataObject[dataArray[i].date] = dataArray[i]
  }

  const currentDate = new Date(startDate)
  // eslint-disable-next-line no-unmodified-loop-condition
  while (currentDate <= endDate) {
    const currentDateStr = currentDate.toISOString().split('T')[0]
    if (dataObject[currentDateStr]) {
      parsed.push({ date: currentDateStr, count: dataObject[currentDateStr].count })
    } else {
      parsed.push({ date: currentDateStr, count: 0 })
    }
    parsed.push()
    currentDate.setDate(currentDate.getDate() + 1)
  }

  const rs = groupData(parsed, group)

  rs.forEach(v => {
    keys.push(v.date)
    result.push(v.count)
  })

  return [result, keys]
}

async function displayReports() {
  if (props.startTime && props.endTime) {
    if (props.report === 0) {
      const resp = await getFirstCrawlInfo({ st_time: props.startTime, ed_time: props.endTime })
      const legendData = Object.keys(resp.data)
      const data = []
      let keys
      legendData.forEach(element => {
        const [d, k] = fillData(resp.data[element] ?? [], props.startTime, props.endTime, props.group)
        data.push({ name: element, type: 'line', smooth: true, data: d })
        keys = k
      })
      displayChart(legendData, keys, data)
    }
  }
}

onMounted(async() => {
  await nextTick()
  chart = echarts.init(echart.value)
  chart.showLoading({ text: '' })
  await displayReports()
  chart.hideLoading()
})

onUnmounted(() => {
  if (!chart) {
    return
  }
  chart.dispose()
  chart = null
})
</script>

<style scoped>
.dashboard-line {
  background-color: #fff;
  height: 360px;
  width: 100%;
}
</style>
