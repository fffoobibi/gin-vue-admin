<template>
  <div v-loading="loading">
    <div>
      <div ref="line" class="reports-line" />
    </div>
    <el-row :gutter="20" style="margin-left: 10px;margin-right: 10px; margin-top: 10px;">
      <el-col :span="8">
        <el-card shadow="hover" class="card-pie">
          <span style="font-weight: bold;">
            <svg-icon name="icon-menu" style="width: 20px; height:20px;" />
            地区占比
          </span>
          <div ref="c2" class="card-reports-pie"></div>
        </el-card>
      </el-col>

      <el-col :span="8">
        <el-card shadow="hover" class="card-pie">
          <span style="font-weight: bold;">
            <svg-icon name="icon-menu" style="width: 20px; height:20px;" />
            平台占比
          </span>
          <div ref="c3" class="card-reports-pie"></div>
        </el-card>
      </el-col>

      <el-col :span="8">
        <el-card shadow="hover" class="card-pie">
          <span style="font-weight: bold;">
            <svg-icon name="icon-menu" style="width: 20px; height:20px;" />
            类目排行
          </span>
          <div ref="c4" class="card-reports-pie" style="height: 200px"></div>
        </el-card>
      </el-col>
    </el-row>

  </div>

</template>
<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, watch, toRefs } from 'vue'
import { useWindowResize } from '@/hooks/use-windows-resize'
import { getFirstCrawlInfo, getCrawlStatsPieData, getCleanList, getTotalResourceReportsList } from '@/api/tblCrawlStats'
import { getKolCrawlInfo } from "@/api/tblKolResource"
import { groupData, formatNumber } from '@/utils/reports'

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

const dataZoom = () => {
  let startDate = new Date(props.startTime)
  let endDate = new Date(props.endTime)
  let timeDiff = endDate.getTime() - startDate.getTime()
  let daysDiff = Math.ceil(timeDiff / (1000 * 3600 * 24))
  let rs
  if (daysDiff >= 10) {
    rs = [
      {
        type: 'slider',
        show: true,
        height: 18,
        start: 0,
        end: 100,
      },
      // {
      //   type: 'inside'
      // },
    ]
  } else {
    rs = [
      {
        type: 'slider',
        show: false,
        height: 15,
        start: 50,
        end: 100,
      },
    ]
  }
  return rs
}

let chart = null
const line = ref(null)
const { startTime, endTime, report, group } = toRefs(props)
const loading = ref(true)

watch([startTime, endTime, group, report], async (newVal, oldVal) => {
  chart.showLoading({ text: '' })
  await displayLineReports()
  chart.hideLoading()
}, { immediate: false })

useWindowResize(() => {
  if (!chart) {
    return
  }
  chart.resize()
})

const renderLineChart = (legendData, xdata, ydata) => {
  if (!chart) {
    chart = echarts.init(line.value)
  }
  chart.setOption({
    legend: {
      textStyle: { fontWeight: 'bold' },
      data: legendData,
      y: 'top'
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
    dataZoom: dataZoom(),
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
    if (dataObject[dataArray[i].date]) {
      dataObject[dataArray[i].date].count += dataArray[i].count
    } else {
      dataObject[dataArray[i].date] = dataArray[i]
    }
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
    currentDate.setDate(currentDate.getDate() + 1)
  }

  const rs = groupData(parsed, group)

  rs.forEach(v => {
    keys.push(v.date)
    result.push(v.count)
  })

  return [result, keys]
}

async function displayLineReports() {
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
      renderLineChart(legendData, keys, data)
    } else if (props.report == 1) {
      // 有效资源访问数量
      const resp = await getKolCrawlInfo({ st_time: props.startTime, ed_time: props.endTime })
      const legendData = Object.keys(resp.data)
      const data = []
      let keys
      legendData.forEach(element => {
        const [d, k] = fillData(resp.data[element] ?? [], props.startTime, props.endTime, props.group)
        data.push({ name: element, type: 'line', smooth: true, data: d })
        keys = k
      })
      renderLineChart(legendData, keys, data)
    } else if (props.report == 2) {
      // 数据清洗报表
      const resp = await getCleanList({ st_time: props.startTime, ed_time: props.endTime })
      const legendData = Object.keys(resp.data)
      const data = []
      let keys
      legendData.forEach(element => {
        const [d, k] = fillData(resp.data[element] ?? [], props.startTime, props.endTime, props.group)
        data.push({ name: element, type: 'line', smooth: true, data: d })
        keys = k
      })
      renderLineChart(legendData, keys, data)
    } else if (props.report == 3) {
      // 更细次数报表
      const resp = await getTotalResourceReportsList({ st_time: props.startTime, ed_time: props.endTime })
      const legendData = Object.keys(resp.data)
      const data = []
      let keys
      legendData.forEach(element => {
        const [d, k] = fillData(resp.data[element] ?? [], props.startTime, props.endTime, props.group)
        data.push({ name: element, type: 'line', smooth: true, data: d })
        keys = k
      })
      renderLineChart(legendData, keys, data)
    }
  }
}

const c2 = ref(null)
const c3 = ref(null)
const c4 = ref(null)
let chart2 = null // 国家
let chart3 = null // 平台
let chart4 = null // 类目

const renderPieCharts = (countryData, platformData, categories, categoryData, reportType = 0) => {
  if (chart2 === null) {
    chart2 = echarts.init(c2.value)
  }
  if (chart3 === null) {
    chart3 = echarts.init(c3.value)
  }

  if (chart4 === null) {
    chart4 = echarts.init(c4.value)
  }
  chart2.setOption(
    {
      legend: {
        type: 'scroll',
        orient: 'vertical',
        x: 'right',
        // formatter: '{b}: {c} ({d}%)', // {b} 代表名称，{c} 代表值，{d} 代表百分比
      },
      tooltip: {
        trigger: "item",
        // formatter: '{b} {c} {d}%'
        formatter: function (params) {
          return params.name + params.value + params.percent + '%';
        }
      },
      series: [
        {
          type: 'pie',
          data: countryData,
          radius: ['40%', '70%'],
          labelLine: {
            show: false
          },
          emphasis: {
            label: {
              show: true,
              fontSize: 12,
              fontWeight: 'bold'
            }
          },
          label: {
            show: false,
            position: 'center',
            formatter: function (params) {
              return params.name + "\n" + params.value;
            }
          }
        },

      ]
    })

  chart3.setOption(
    {
      tooltip: {
        trigger: "item",
        // formatter: '{b}: {c}',
        formater: function (params) {
          return params.name + formatNumber(params.value)
        },
        textStyle: {
          fontWeight: 'bold'
        }
      },
      series: [
        {
          type: 'pie',
          data: platformData,
          label: {
            show: true,
            fontWeight: 'bold',
            fontSize: 14,
            formatter: function (params) {
              return params.name + " \n" + params.percent.toFixed(2) + '%';
            }
          },
        }
      ]
    })

  chart4.setOption(
    {
      // dataZoom: [
      //   {
      //     type: 'slider',
      //     show: true,
      //     yAxisIndex: [0],
      //     left: '50%',
      //     start: 29,
      //     end: 36
      //   },
      // ],
      tooltip: {
        trigger: "item",
        // formatter: '{b}: {c}',
        formater: function (params) {
          return params.name + formatNumber(params.value)
        },
        textStyle: {
          fontWeight: 'bold'
        }
      },
      xAxis: {
        type: 'value',
      },
      grid: {
        left: '5%', // 调整左边距
        right: '5%', // 调整右边距
        top: '5%',
        bottom: '5%', // 调整底边距
        containLabel: true // 自适应内容显示，防止内容溢出
      },
      yAxis: {
        axisTick: { show: false }, // y轴刻度线
        type: 'category',
        data: categories, // ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        inverse: true,
        minInterval: 1,
        boundaryGap: [0, 0.01],
        axisLabel: {
          interval: 0 // 设置轴标签显示间隔为0，表示每个标签都显示
        }
      },
      series: [
        {
          data: categoryData,  // [120, 200, 150, 80, 70, 110, 130],
          type: 'bar',
        }
      ]
    })
}

const distPlayPieChart = async () => {
  if (props.startTime && props.endTime) {
    if (props.report === 0) {
      // 初次访问次数
      const resp = await getCrawlStatsPieData({ st_time: props.startTime, ed_time: props.endTime })
      const sliceCountry = resp.data.country.slice(0, 9)
      const other = resp.data.country.slice(9).reduce((x, y) => x.count + y.count)
      sliceCountry.push({ country_name: '其他', count: other })

      const country = sliceCountry.map(v => ({ name: v.country_name, value: v.count }))
      const platform = resp.data.platform.map(v => {
        let name
        if (v.platform == 1) {
          name = 'TikTok'
        } else if (v.platform == 2) {
          name = 'Youtube'
        } else {
          name = 'Instagram'
        }
        return { name: name, value: v.count }
      })
      const categoris = []
      const categoryData = []
      resp.data.category.slice(0, 7).forEach(v => {
        categoris.push(v.name)
        categoryData.push(v.count)
      })
      renderPieCharts(country, platform, categoris, categoryData)
    }
  }

}

onMounted(async () => {
  await nextTick()
  chart = echarts.init(line.value)
  await displayLineReports()
  await distPlayPieChart()
  loading.value = false

})

onUnmounted(() => {
  if (chart) {
    chart.dispose()
    chart = null
  }
  if (chart2) {
    chart2.dispose()
    chart2 = null
  }
  if (chart3) {
    chart3.dispose()
    chart3 = null
  }
  if (chart4) {
    chart4.dispose()
    chart4 = null
  }
})
</script>

<style scoped>
.reports-line {
  background-color: #fff;
  height: 380px;
  width: 100%;
}

.card-reports-pie {
  background-color: transparent;
  height: 190px;
  width: 100%;
}

.card-pie {
  background-color: #F2F6F9;
  height: 250px;
}
</style>
