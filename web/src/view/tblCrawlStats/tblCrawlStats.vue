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
                <el-tooltip effect="dark" content="总资源库当前在库数量" placement="top">
                  <svg-icon name="icon-tips" color="lightgray" style="width:16px;height:16px" />
                </el-tooltip>
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
          <el-card style="height: 250px; width: 420px" class="rounded-10">
            <el-row justify="space-between" align="middle">
              <el-col :span="16" class="flex items-center">
                <img src="@/assets/icons/icon-bigdeal.svg" class="h w-4 mr-2 py-2" />
                <span class="text-lg font-bold">
                  大事记
                </span>
              </el-col>
              <el-col :span="8" :push="5">
                <span @click="router.push({ name: 'e' })" style="cursor: pointer;">
                  更多...
                </span>
              </el-col>
            </el-row>
            <el-scrollbar height="160px">
              <div style="height: 180px;">
                <div v-for="log in summaryInfo.events" style="margin-bottom: 10px;">
                  <p style="font-weight: bold;margin-bottom: 5px;margin-right: 20px"><svg-icon name="icon-circle"
                      style="width:8px;height:8px;padding-right: 3px" />{{ log.title }}</p>
                  <p style="margin-right: 20px">{{ log.details }}</p>
                  <p style="text-align: right;color:gray;font-size:9pt;margin-right: 20px">
                    {{ formatTimeToStr(log.occurced_time, "yyyy-MM-dd") }}</p>
                </div>
              </div>
            </el-scrollbar>
          </el-card>
        </el-col>

      </el-space>
    </el-row>

    <el-row justify="center">
      <el-card class="radius-10 report-card" body-class="no-padding">
        <el-row justify="start" align="middle" class="padding-5 report-header" gutter="10">
          <el-col :span="4">
            <h3 style="color:#fff">
              <img src="@/assets/images/icon-menu-white.png" style="width: 11px">
              数据趋势图
            </h3>
          </el-col>

          <el-col :span="3" :push="12">
            <div class="week">
              <el-button v-auth="btnAuth.addFirstCount" @click="countForm.visable = true">添加</el-button>
              <span :class="group.g0" @click="handleGroup(0)">日</span>
              <span :class="group.g1" @click="handleGroup(1)">周</span>
              <span :class="group.g2" @click="handleGroup(2)">月</span>
            </div>
          </el-col>

          <el-col :span="5" :push="12" class="flex">
            <el-date-picker v-model="timeValue" type="daterange" :shortcuts="shortcuts" range-separator="-"
              start-placeholder="开始" end-placeholder="结束" value-format="YYYY-MM-DD" size="default" @change="dateChange"
              class="diy-time" />
          </el-col>

        </el-row>
        <div v-loading="summaryInfoLoading">
          <div class="data-tab">
            <ul class="list-inline">
              <li :class="ac0" @click="clickReportsTab(0)">
                <div class="s-img v-vague"></div>
                <div class="s-box">
                  <p>盲抓 <el-tooltip effect="dark" content="指爬虫访问社媒平台的次数, 用来抓取红人链接" placement="top">
                      <svg-icon name="icon-tips" :color="tipColors.t0" style="width:16px;height:16px" />
                    </el-tooltip></p>
                  <span>{{ formatNumber(summaryInfo.crawlCount) }}</span>
                </div>
              </li>
              <li :class="ac1" @click="clickReportsTab(1)">
                <div class="s-img v-update"></div>
                <div class="s-box">
                  <p>初步更新 <el-tooltip effect="dark" content="指经过爬虫盲抓, 红人链接去重后的数量, 该阶段需要爬虫更新红人详细数据" placement="top">
                      <svg-icon name="icon-tips" :color="tipColors.t1" style="width:16px;height:16px" />
                    </el-tooltip></p>
                  <span>{{ formatNumber(summaryInfo.validCount) }}</span>
                </div>
              </li>
              <li :class="ac2" @click="clickReportsTab(2)">
                <div class="s-img v-clean"></div>
                <div class="s-box">
                  <p>清洗 <el-tooltip effect="dark"
                      content="将初步有效资源数, 根据一定的条件(红人粉丝量，播放数，视频量，最近更新时间等)筛选得到的数据, 该阶段需要GPT识别红人类目" placement="top">
                      <svg-icon name="icon-tips" :color="tipColors.t2" style="width:16px;height:16px" />
                    </el-tooltip></p>
                  <span>{{ formatNumber(summaryInfo.cleanCount) }}</span>
                </div>
              </li>
              <li :class="ac3" @click="clickReportsTab(3)">
                <div class="s-img v-total"></div>
                <div class="s-box">
                  <p>总资源更新 <el-tooltip effect="dark" content="指总资源库红人更新的数据量，更新周期每月1次" placement="top">
                      <svg-icon name="icon-tips" :color="tipColors.t3" style="width:16px;height:16px" />
                    </el-tooltip></p>
                  <span>{{ formatNumber(summaryInfo.updateCount) }}</span>
                </div>
              </li>
            </ul>
          </div>
          <LineReports :report="report" :start-time="st_time" :end-time="ed_time" :group="group.current"
            :line-title="lineTitle" />
        </div>
      </el-card>

    </el-row>

    <el-dialog v-model="countForm.visable" title="盲抓次数" width="500" :before-close="handleClose" destroy-on-close>
      <el-form :model="countForm" label-width="auto" ref="formRef">
        <el-form-item label="日期" prop="date" :rules="{ required: true, message: '请选择日期' }">
          <el-date-picker v-model="countForm.date" type="date" value-format="YYYY-MM-DD"></el-date-picker>
        </el-form-item>
        <el-form-item label="YTB" prop="ytb">
          <el-input-number v-model="countForm.ytb" :step="2000" :min="0"></el-input-number>
        </el-form-item>
        <el-form-item label="TT" prop="tiktok">
          <el-input-number v-model="countForm.tiktok" :step="2000" :min="0"></el-input-number>
        </el-form-item>
        <el-form-item label="INS" prop="ins">
          <el-input-number v-model="countForm.ins" :step="2000" :min="0"></el-input-number>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="countForm.visable = false">取消</el-button>
          <el-button :loading="countForm.loading" type="primary" @click="submitCount(formRef)">
            确认
          </el-button>
        </div>
      </template>
    </el-dialog>

  </div>
</template>

<script setup>

import { reactive, ref, onMounted, nextTick, computed, watch } from 'vue'
import { useRouter } from "vue-router"
import { formatTimeToStr } from '@/utils/date'
import { getTotalResourceInfo, getSummaryCrawlInfo } from '@/api/tblCrawlStats'
import * as echarts from 'echarts'
import LineReports from '@/view/tblCrawlStats/LineReports/lineReports.vue'
import { formatNumber } from '@/utils/reports'
import { useBtnAuth } from '@/utils/btnAuth'
import { addFirstCount } from '@/api/tblCrawlStats'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'TblCrawlStats'
})

const btnAuth = useBtnAuth()
const router = useRouter()
const dft_end = new Date()
const dft_start = new Date()
dft_start.setTime(dft_start.getTime() - 3600 * 1000 * 24 * 30)

const ac0 = ref('wsmall active')
const ac1 = ref('wsmall')
const ac2 = ref('wsmall')
const ac3 = ref('wsmall')
const st_time = ref(formatTimeToStr(dft_start, 'yyyy-MM-dd'))
const ed_time = ref(formatTimeToStr(dft_end, 'yyyy-MM-dd'))
const timeValue = ref([formatTimeToStr(dft_start), formatTimeToStr(dft_end)])
const report = ref(0)
const tipColors = reactive({
  t0: "#8ba9c8",
  t1: "#aad6fe",
  t2: "#aad6fe",
  t3: "#aad6fe"
})

const clickReportsTab = (type) => {
  if (type === 0) {
    ac1.value = ac2.value = ac3.value = 'wsmall'
    ac0.value = 'wsmall active'
    tipColors.t0 = "#8ba9c8"
    tipColors.t1 = "#aad6fe"
    tipColors.t2 = "#aad6fe"
    tipColors.t3 = "#aad6fe"

  } else if (type === 1) {
    ac0.value = ac2.value = ac3.value = 'wsmall'
    ac1.value = 'wsmall active'
    tipColors.t0 = "#aad6fe"
    tipColors.t1 = "#8ba9c8"
    tipColors.t2 = "#aad6fe"
    tipColors.t3 = "#aad6fe"
  } else if (type === 2) {
    ac0.value = ac1.value = ac3.value = 'wsmall'
    ac2.value = 'wsmall active'
    tipColors.t0 = "#aad6fe"
    tipColors.t1 = "#aad6fe"
    tipColors.t2 = "#8ba9c8"
    tipColors.t3 = "#aad6fe"
  } else if (type === 3) {
    ac0.value = ac1.value = ac2.value = 'wsmall'
    ac3.value = 'wsmall active'
    tipColors.t0 = "#aad6fe"
    tipColors.t1 = "#aad6fe"
    tipColors.t2 = "#aad6fe"
    tipColors.t3 = "#8ba9c8"
  }
  report.value = type
}

watch([st_time, ed_time], () => {
  disPlaySummaryInfo(st_time.value, ed_time.value)
})

const lineTitle = computed(() => {
  const startDate = new Date(st_time.value);
  const endDate = new Date(ed_time.value);
  if (isNaN(startDate.getTime()) || isNaN(endDate.getTime())) {
    throw new Error('Invalid date string');
  }
  const timeDifference = Math.abs(endDate - startDate);

  const daysDifference = Math.ceil(timeDifference / (1000 * 60 * 60 * 24))
  let msg
  if (report.value === 0) {
    msg = " 盲抓次数 "
  } else if (report.value === 1) {
    msg = " 初步有效资源数 "
  } else if (report.value === 2) {
    msg = " 数据清洗 "
  } else if (report.value === 3) {
    msg = " 总资源库更新 "
  } else {
    msg = ""
  }
  return `${daysDifference}天的${msg}统计数据`
})

const formRef = ref(null)
const countForm = reactive({
  date: "",
  ins: 0,
  tiktok: 0,
  ytb: 0,
  loading: false,
  visable: false
})

const handleClose = (done) => {
  formRef.value.resetFields()
  countForm.date = ""
  done()
}

const submitCount = (form) => {
  if (!form) return
  countForm.loading = true
  form.validate(async valid => {
    if (!valid) {
      countForm.loading = false
    } else {
      const resp = await addFirstCount({ date: countForm.date, ins: countForm.ins, tiktok: countForm.tiktok, ytb: countForm.ytb })
      countForm.loading = false
      countForm.visable = false
      formRef.value.resetFields()
      ElMessage.success({ message: "已添加", plain: true })
    }
  })
}

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
  if (chat1) {
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
          x: '70%',
          y: '20%',
          itemWidth: 0,
          // padding:['0%','20%','50%','0%'], //可设定图例[距上方距离，距右方距离，距下方距离，距左方距离]
          textStyle: {
            color: '#fff',
            fontSize: 13,
          },
          selectedMode: false, //取消图例上的点击事件
          formatter: (name) => {
            return "TikTok"
          },
        },
        color: ['#FF9600', '#72B9FF'],
        series: [
          {
            alignTo: 'edge',
            type: 'pie',
            data: [
              {
                value: totalResourceInfo.Tiktok,
                name: 'Tiktok',
                label: {
                  show: true,
                  position: 'center',
                  formatter: "{d}%",
                  backgroundColor: "transparent",
                  color: "white"
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

  }

  if (chat2 === null) {
    chat2 = echarts.init(c2.value)
  }
  if (chat2) {
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
            fontSize: 13,
          },
          selectedMode: false,
          formatter: (name) => {
            // let data = ''
            // if (name === 'Youtube') {
            //   data = (Number(totalResourceInfo.Youtube / totalResourceInfo.total_count) * 100).toFixed()
            //   const text = `${name} ${data}%`
            //   return `${text}`
            // }
            return name
          },
        },

        color: ['#FF9600', '#72B9FF'],
        series: [
          {
            alignTo: 'edge',
            type: 'pie',
            data: [
              {
                value: totalResourceInfo.Youtube,
                name: 'Youtube',
                label: {
                  show: true,
                  position: 'center',
                  formatter: "{d}%",
                  backgroundColor: "transparent",
                  color: "white"
                },
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
  }

  if (chat3 === null) {
    chat3 = echarts.init(c3.value)
  }
  if (chat3) {
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
            fontSize: 13,
          },
          selectedMode: false,
          formatter: (name) => {
            return name
            // let data = ''
            // if (name === 'Instagram') {
            //   data = (Number(totalResourceInfo.Instagram / totalResourceInfo.total_count) * 100).toFixed()
            //   const text = `${name} ${data}%`
            //   return `${text}`
            // }
          },
        },
        color: ['#FF9600', '#72B9FF'],
        series: [
          {
            alignTo: 'edge',
            type: 'pie',
            data: [
              {
                value: totalResourceInfo.Instagram,
                name: 'Instagram',
                label: {
                  show: true,
                  position: 'center',
                  formatter: "{d}%",
                  backgroundColor: "transparent",
                  color: "white"
                },
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
          }
        ]
      })
  }
}

const summaryInfoLoading = ref(false)

const summaryInfo = reactive({
  crawlCount: 0,
  validCount: 0,
  cleanCount: 0,
  updateCount: 0,
  totalCount: 0,
  events: []
})

const disPlaySummaryInfo = async (st_time, ed_time) => {
  try {
    summaryInfoLoading.value = true
    const resp = await getSummaryCrawlInfo({ st_time, ed_time })
    summaryInfo.crawlCount = resp.data.crawl_count
    summaryInfo.validCount = resp.data.valid_count
    summaryInfo.cleanCount = resp.data.clean_count
    summaryInfo.updateCount = resp.data.update_count
    summaryInfo.totalCount = resp.data.total_count
    summaryInfo.events = resp.data.events
  } finally {
    summaryInfoLoading.value = false
  }
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
}



const shortcuts = [
  {
    text: '近14天',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 14)
      return [start, end]
    },
  },
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

:deep(.diy-time) {
  font-size: 14px;
  background-color: #145ba4;
  border: 0px solid red;
  color: white;
  border-radius: 10px;
  padding: 2px 8px;
  box-shadow: 0 0 0 1px transparent;
}

:deep(.diy-time .el-range-input) {
  color: white;
  border-color: transparent;
}

:deep(.diy-time .el-range-separator) {
  color: white
}

:deep(.el-range-editor:hover) {
  box-shadow: 0 0 0 1px transparent;
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
  background: linear-gradient(90deg, #38a2f4, #0c5cb6);
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

.data-tab {
  background: linear-gradient(90deg, #38a2f4, #0c5cb6);
}

.data-tab ul {
  display: flex;
  margin: 0
}

/* 后加入 */
.list-inline {
  height: 84px
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

.data-tab {
  padding-right: 20px;
  position: relative;
  z-index: 1
}

.data-tab ul {
  display: flex;
  margin: 0
}

.data-tab ul li {
  flex: 1;
  padding: 6px;
  font-size: 12px;
  cursor: pointer;
  color: #fff;
  margin-left: 20px;
  padding-left: 0px;
  border-bottom: 72px solid #0000001a;
  border-top: 0px solid transparent;
  border-left: 25px solid transparent;
  border-right: 25px solid transparent;
  color: #fff;
  display: flex;
  justify-content: center;
  transition: all 0s linear;
}

.data-tab ul li:hover {
  border-bottom: 72px solid #0000004f;
  border-left: 25px solid transparent;
  border-right: 25px solid transparent;
}

.data-tab ul .s-img {
  width: 60px;
  height: 80px;
  background-size: 47px;
  background-position: center;
  background-repeat: no-repeat
}

.data-tab ul .v-vague {
  background-image: url(@/assets/images/icon-1.png)
}

.data-tab ul .v-update {
  background-image: url(@/assets/images/icon-2.png)
}

.data-tab ul .v-clean {
  background-image: url(@/assets/images/icon-3.png)
}

.data-tab ul .v-warehouse {
  background-image: url(@/assets/images/icon-4.png)
}

.data-tab ul .v-total {
  background-image: url(@/assets/images/icon-5.png)
}

.data-tab ul .active .v-vague {
  background-image: url(@/assets/images/icon-1-v.png)
}

.data-tab ul .active .v-update {
  background-image: url(@/assets/images/icon-2-v.png)
}

.data-tab ul .active .v-clean {
  background-image: url(@/assets/images/icon-3-v.png)
}

.data-tab ul .active .v-warehouse {
  background-image: url(@/assets/images/icon-4-v.png)
}

.data-tab ul .active .v-total {
  background-image: url(@/assets/images/icon-5-v.png)
}


.data-tab ul .active,
.data-tab ul .active:hover {
  color: #1e7ad0;
  border-bottom: 72px solid #fff;
  cursor: default
}

.data-tab ul .active p {
  font-weight: bold;
}

.data-tab ul .active span {
  color: #72a3d7
}

.data-tab ul p {
  margin: 0px;
  margin-top: 18px;
  padding: 0;
  font-size: 16px;
  position: relative
}

.data-tab ul img {
  position: absolute;
  width: 48px;
  left: -53px;
}

.data-tab ul span {
  font-weight: bold;
  font-size: 14px;
  color: #a2d3ff;
}

.data-tab ul .wsmall {
  flex: none;
  width: 210px;
}

.week {
  display: flex;
  background-color: #145ba4;
  padding: 7px 12px;
  border-radius: 12px;
  font-size: 14px;
  justify-content: center;
}

.week span {
  padding: 2px 2px;
  color: #8db3db;
  border-bottom: 2px solid #145ba4;
  display: block;
  margin: 0px 8px;
  cursor: pointer;
}

.week .active {
  color: #fff;
  border-bottom: 2px solid #fff;
}
</style>
