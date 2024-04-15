<template>
  <div>
    <el-row justify="center">
      <el-space :size="30">
        <el-col>
          <el-card style="height: 250px; width: 840px; position: relative" class="total-data radius-10">
            <div class="inv-bg"></div>
            <el-row justify="center">
              <p style="color: white; font-weight: bold; margin-top: 20px;margin-bottom:20px;font-size: 12pt;">
                <svg-icon name="icon-data" color="red" style="width:16px;height:16px" />
                当前在库数量
                <svg-icon name="icon-tips" color="red" style="width:16px;height:16px" />
              </p>
            </el-row>

            <el-row justify="center">
              <el-statistic title="" :value="totalResourceInfo.total_count"
                value-style="color:white;font-weight:bold;font-size:24pt" />
            </el-row>

            <el-row justify="space-between">
              <el-col :span="8">
                <div ref="c1" class="pie" />
              </el-col>
              <el-col :span="8">
                <div ref="c2" class="pie" />
              </el-col>
              <el-col :span="8">
                <div ref="c3" class="pie" />
              </el-col>
            </el-row>

          </el-card>
        </el-col>

        <el-col>
          <el-card style="height: 250px; width: 420px" class="radius-10">
            <el-row justify="space-between" align="middle">
              <el-col :span="16">
                <h4>
                  <svg-icon name="icon-bigdeal" style="width:16px;height:16px" />大事记
                </h4>
              </el-col>
              <el-col :span="8" :push="5">
                <span @click="console.log('ffff')" style="cursor: pointer;">
                  更多...
                </span>
              </el-col>
            </el-row>
            <el-scrollbar height="150px">
              <div style="height: 180px;">
                <div v-for="log in logs" style="margin-bottom: 10px;">
                  <p style="font-weight: bold;margin-bottom: 5px;">{{ log.title }}</p>
                  <span>{{ log.detail }}</span>
                  <p style="text-align: right;color:gray;font-size:9pt">2024.4.12</p>
                </div>
              </div>
            </el-scrollbar>
          </el-card>
        </el-col>

      </el-space>
    </el-row>

    <el-row justify="center">
      <el-card class="radius-10 report-card" body-class="no-padding">
        <!-- <el-affix :offset="98"> -->
        <el-row justify="space-between" align="middle" class="padding-5 report-header">
          <el-col :span="4">
            <h3>
              <svg-icon name="icon-menu" style="width: 20px; height:20px" />
              数据趋势图
            </h3>
          </el-col>

          <el-col :offset="10" :span="2.9">
            <div class="week">
              <span :class="group.g0" @click="handleGroup(0)">日</span>
              <span :class="group.g1" @click="handleGroup(1)">周</span>
              <span :class="group.g2" @click="handleGroup(2)">月</span>
            </div>

          </el-col>

          <el-col :span="5">
            <el-date-picker v-model="timeValue" style="width: 240px" type="daterange" :shortcuts="shortcuts"
              range-separator="-" start-placeholder="开始" end-placeholder="结束" value-format="YYYY-MM-DD" size="default"
              @change="dateChange" />
          </el-col>

        </el-row>
        <div class="data-tab">
          <ul class="list-inline">
            <li :class="ac0" @click="clickReportsTab(0)">
              <p>首次抓取次数</p>
              <span>{{ formatNumber(summaryInfo.crawlCount) }}</span>
            </li>
            <li :class="ac1" @click="clickReportsTab(1)">
              <p>有效资源数</p>
              <span>{{ formatNumber(summaryInfo.validCount) }}</span>
            </li>
            <li :class="ac2" @click="clickReportsTab(2)">
              <p>数据清洗</p>
              <span>{{ formatNumber(summaryInfo.cleanCount) }}</span>
            </li>
            <li :class="ac3" @click="clickReportsTab(3)">
              <p>总资源库更新</p>
              <span>{{ formatNumber(summaryInfo.updateCount) }}</span>
            </li>
          </ul>
        </div>
        <!-- </el-affix> -->
        <p style="color:red;margin-top: 10px; padding-left: 10px;font-size: 9pt;">{{ reportTip }}</p>
        <LineReports :report="report" :start-time="st_time" :end-time="ed_time" :group="group.current" />
      </el-card>

    </el-row>

  </div>
</template>

<script setup>

// 全量引入格式化工具 请按需保留
import { reactive, ref, onMounted, nextTick, computed } from 'vue'
import { formatTimeToStr } from '@/utils/date'
import { getTotalResourceInfo, getSummaryCrawlInfo } from '@/api/tblCrawlStats'
import * as echarts from 'echarts'
import LineReports from '@/view/tblCrawlStats/LineReports/lineReports.vue'
import { formatNumber } from '@/utils/reports'

defineOptions({
  name: 'TblCrawlStats'
})

const dft_end = new Date()
const dft_start = new Date()
dft_start.setTime(dft_start.getTime() - 3600 * 1000 * 24 * 30)

const ac0 = ref('active')
const ac1 = ref('')
const ac2 = ref('')
const ac3 = ref('')
const st_time = ref(formatTimeToStr(dft_start, 'yyyy-MM-dd'))
const ed_time = ref(formatTimeToStr(dft_end, 'yyyy-MM-dd'))
const timeValue = ref([formatTimeToStr(dft_start), formatTimeToStr(dft_end)])
const report = ref(0)

const reportTip = computed(() => {
  if (report.value == 0) {
    return "*注:抓取次数指访问社媒平台的次数"
  } else if (report.value == 1) {
    return "*注:指红人链接去重后的数量"
  } else if (report.value == 2) {
    return "*注:指采集到的红人根据一定条件筛选后的数据量"
  } else if (report.value == 3) {
    return "*注:指总资源库红人更新的数据量"
  }
})

const group = reactive({
  current: 'week',
  g0: '',
  g1: 'active',
  g2: ''
})

const totalResourceInfo = reactive({
  total_count: null,
  Tiktok: null,
  Youtube: null,
  Instagram: null
})

const c1 = ref(null)
const c2 = ref(null)
const c3 = ref(null)
let chat1 = null
let chat2 = null
let chat3 = null

const setTotalPies = () => {
  if (chat1 === null) {
    chat1 = echarts.init(c1.value)
  }
  chat1.setOption(
    {
      // tooltip: {
      //   trigger: 'item',
      //   formatter: function (params) {
      //     return `${params.value} 占比%`
      //   }
      // },
      legend: {
        orient: 'vertical',
        x: 'right',
        y: '20%',
        itemWidth: 0,
        // padding:['0%','20%','50%','0%'], //可设定图例[距上方距离，距右方距离，距下方距离，距左方距离]
        textStyle: {
          color: '#fff',
          fontSize: 12,
        },
        selectedMode: false, //取消图例上的点击事件

      },
      formatter: (name) => {
        let data = ''
        if (name === 'Tiktok') {
          data = (Number(totalResourceInfo.Tiktok / totalResourceInfo.total_count) * 100).toFixed()
          const text = `${name} ${data}%`
          return `${text}`
        }
      },
      color: ['#FF9600', '#72B9FF'],
      series: [
        {
          type: 'pie',
          data: [
            {
              value: totalResourceInfo.Tiktok,
              name: 'Tiktok',
              label: {
                show: false,
              },
            },
            {
              value: totalResourceInfo.Youtube + totalResourceInfo.Instagram,
              name: ''
            }
          ],
          radius: ['40%', '70%'],
          labelLine: {
            show: false
          }
        },

      ]
    })

  if (chat2 === null) {
    chat2 = echarts.init(c2.value)
  }
  chat2.setOption(
    {
      // tooltip: {
      //   trigger: 'item',
      //   formatter: function (params) {
      //     return `${params.value} 占比%`
      //   }
      // },
      legend: {
        orient: 'vertical',
        x: 'right',
        y: '20%',
        itemWidth: 0,
        // padding:['0%','20%','50%','0%'], //可设定图例[距上方距离，距右方距离，距下方距离，距左方距离]
        textStyle: {
          color: '#fff',
          fontSize: 12,
        },
        selectedMode: false
      },
      formatter: (name) => {
        let data = ''
        if (name === 'Youtube') {
          data = (Number(totalResourceInfo.Youtube / totalResourceInfo.total_count) * 100).toFixed()
          const text = `${name} ${data}%`
          return `${text}`
        }
      },
      color: ['#FF9600', '#72B9FF'],
      series: [
        {
          type: 'pie',
          data: [
            {
              value: totalResourceInfo.Youtube,
              name: 'Youtube'
            },
            {
              value: totalResourceInfo.Tiktok + totalResourceInfo.Instagram,
              name: ''
            }
          ],
          radius: ['40%', '70%'],
          labelLine: {
            show: false
          },
          label: {
            show: false,
          },
        }
      ]
    })

  if (chat3 === null) {
    chat3 = echarts.init(c3.value)
  }
  chat3.setOption(
    {
      // tooltip: {
      //   trigger: 'item',
      //   formatter: function(params) {
      //     return `${params.value} 占比%`
      //   }
      // },
      legend: {
        orient: 'vertical',
        x: 'right',
        y: '20%',
        itemWidth: 0,
        // padding:['0%','20%','50%','0%'], //可设定图例[距上方距离，距右方距离，距下方距离，距左方距离]
        textStyle: {
          color: '#fff',
          fontSize: 12,
        },
        selectedMode: false
      },
      formatter: (name) => {
        let data = ''
        if (name === 'Instagram') {
          data = (Number(totalResourceInfo.Instagram / totalResourceInfo.total_count) * 100).toFixed()
          const text = `${name} ${data}%`
          return `${text}`
        }
      },
      color: ['#FF9600', '#72B9FF'],
      series: [
        {
          type: 'pie',
          data: [
            {
              value: totalResourceInfo.Instagram,
              name: 'Instagram'
            },
            {
              value: totalResourceInfo.Youtube + totalResourceInfo.Tiktok,
              name: ''
            }
          ],
          radius: ['40%', '70%'],
          labelLine: {
            show: false
          },
          label: {
            show: false,
          },
        }
      ]
    })
}

const summaryInfo = reactive({
  crawlCount: 0,
  validCount: 0,
  cleanCount: 0,
  updateCount: 0,
  totalCount: 0
})

const disPlaySummaryInfo = async (st_time, ed_time) => {
  const resp = await getSummaryCrawlInfo({ st_time, ed_time })
  summaryInfo.crawlCount = resp.data.crawl_count
  summaryInfo.validCount = resp.data.valid_count
  summaryInfo.cleanCount = resp.data.clean_count
  summaryInfo.updateCount = resp.data.update_count
  summaryInfo.totalCount = resp.data.total_count
}

onMounted(async () => {
  await nextTick()
  try {
    const resp = await getTotalResourceInfo()
    totalResourceInfo.total_count = resp.data.total
    totalResourceInfo.Youtube = resp.data.Youtube
    totalResourceInfo.Instagram = resp.data.Instagram
    totalResourceInfo.Tiktok = resp.data.Tiktok
    setTotalPies(resp.data)
    await disPlaySummaryInfo(st_time.value, ed_time.value)
  } catch (error) {

  }
})

const handleGroup = (val) => {
  if (val === 0) {
    group.g1 = group.g2 = ''
    group.g0 = 'active'
    group.current = 'day'
  } else if (val === 1) {
    group.g0 = group.g2 = ''
    group.g1 = 'active'
    group.current = 'week'
  } else {
    group.g0 = group.g1 = ''
    group.g2 = 'active'
    group.current = 'month'
  }
}

const dateChange = (val) => {
  st_time.value = val[0]
  ed_time.value = val[1]
  console.log(st_time.value, ed_time.value)
}

const clickReportsTab = (type) => {
  if (type === 0) {
    ac1.value = ac2.value = ac3.value = ''
    ac0.value = 'active'
  } else if (type === 1) {
    ac0.value = ac2.value = ac3.value = ''
    ac1.value = 'active'
  } else if (type === 2) {
    ac0.value = ac1.value = ac3.value = ''
    ac2.value = 'active'
  } else if (type === 3) {
    ac0.value = ac1.value = ac2.value = ''
    ac3.value = 'active'
  }
  report.value = type
}

const logs = ref([
  { title: '印度服务器添加', detail: '3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%3月17日印度新服务器新增资源数突破' },
  { title: '印度服务器添加', detail: '3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%3月17日印度新服务器新增资源数突破' },
  { title: '印度服务器添加', detail: '3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%3月17日印度新服务器新增资源数突破' },
  { title: '印度服务器添加', detail: '3月17日印度新服务器新增资源数突破1器新增资源数突破100W同比上升20%3月17日印度新服务器新增资源数突破' },

])

const shortcuts = [
  {
    text: '近30天',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
      return [start, end]
    },
  },
  {
    text: '近60天',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 60)
      return [start, end]
    },
  },
  {
    text: '近180天',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 180)
      return [start, end]
    },
  },
]

</script>

<style scoped>
:deep(.no-padding) {
  padding: 0px;
}

.report-card {
  width: 1290px;
  margin-top: 20px;
  min-height: 840px;
}

.inv-bg {
  width: 140px;
  height: 140px;
  background: url(@/assets/in-stock-bg.png) no-repeat 0 0;
  background-size: cover;
  position: absolute;
  left: 700px;
  top: 0px;
  right: 0px;
  opacity: 0.3;
}

.pie {
  background-color: transparent;
  height: 130px;
  width: 100%;
  margin-top: 5px;
}



.report-header {
  background: #F1F5F9;
}

.padding-5 {
  padding: 5px
}

.radius-10 {
  border-radius: 10px;
}

.total-data {
  background-image: linear-gradient(90deg, #38a2f4, #0c5cb6);
}

.statics-space {
  margin-left: 30px;
  margin-right: 30px;
}

.week {
  display: flex;
  background-color: #fff;
  padding: 7px 12px;
  border-radius: 12px;
}

.week span {
  padding: 2px 2px;
  color: #666;
  border-bottom: 2px solid #fff;
  display: block;
  margin: 0px 8px;
  cursor: pointer;
}

.week .active {
  color: #000;
  border-bottom: 2px solid #000;
}

.data-tab {
  background-color: #dff1ff;
}

.data-tab ul {
  display: flex;
  margin: 0
}

/* .data-tab ul li {
  background: url(@/assets/tab-bg.png) right 0px no-repeat;
  padding: 10px;
  height: 30px;
  font-size: 16px;
  min-width: 180px;
  cursor: pointer;
  padding-left: 30px;
  font-weight: bold;
  color: #51728c;
}

.data-tab ul .active {
  background-position: right -183px;
  color: #215883;
} */

.data-tab ul li {
  background: url(@/assets/tab-bg1.png) right 0px no-repeat;
  padding: 10px;
  height: 30px;
  font-size: 16px;
  min-width: 180px;
  cursor: pointer;
  padding-left: 30px;
  font-weight: bold;
  color: #51728c
}

.data-tab ul .active {
  background: url(@/assets/tab-bg2.png) right 0px no-repeat;
  color: #215883;
}

.data-tab ul p {
  margin: 0px;
  padding: 0;
}

.data-tab ul span {
  font-weight: bold;
  font-size: 14px;
  opacity: 0.9;
}

.data-cont .data-report {
  padding: 12px;
}

.list-inline {
  padding-left: 0;
  list-style: none
}

.list-inline-item {
  display: inline-block
}

.list-inline-item:not(:last-child) {
  margin-right: .5rem
}
</style>
