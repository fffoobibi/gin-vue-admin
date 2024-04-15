<template>
  <div>
    <el-row justify="center">
      <el-card class="report-card">
        <div class="data-cont">
          <el-scrollbar>
            <ul class="data-news">

              <li v-for="item in crawlData.list" :key="item.ID" class="li-circle">
                <div class="news-title"> <svg-icon name="icon-circle" style="width: 20px; height:20px;" />{{ item.title
                  }}
                  <el-icon v-auth="btnAuth.add">
                    <Delete @click="deleteEvent(item.ID)" />
                  </el-icon>
                </div>
                <p>
                  {{ item.details }}
                </p>
                <time>{{ formatTimeToStr(item.occured_time, "yyyy-MM-dd") }}</time>
                <el-divider />
              </li>
            </ul>
          </el-scrollbar>

        </div>

        <el-row align="center" justify="space-between">
          <div style="display: flex;justify-items: center;align-items: center;">
            <p style="margin-right: 30px;color:gray">共{{ crawlData.totalCount }}条</p>
            <el-button type="primary" v-auth="btnAuth.add" @click="addCrawlEvent">添加事件</el-button>
          </div>
          <el-pagination background layout="prev, pager, next" :total="crawlData.totalCount" small
            class="center-pagination" @change="getCrawlEventsList" :page-size="crawlData.pageSize" />
        </el-row>
      </el-card>

    </el-row>
    <el-drawer title="添加大事记" v-model="drawer" direction="rtl" :before-close="handleClose" destroy-on-close>

      <el-form :model="form" label-width="80px" ref="formRef">
        <el-form-item label="标题" prop="title" :rules="{ required: true, message: '请填写标题' }">
          <el-input v-model="form.title"></el-input>
        </el-form-item>
        <el-form-item label="详情" prop="details" :rules="{ required: true, message: '请填写详情' }">
          <el-input v-model="form.details" type="textarea" :rows="4" maxlength="1024" show-word-limit></el-input>
        </el-form-item>
        <el-form-item label="日期" prop="occurced_time" :rules="{ required: true, message: '请选择日期' }">
          <el-date-picker v-model="form.occurced_time" type="date" placeholder="选择日期" format="YYYY-MM-DD" />
        </el-form-item>

      </el-form>
      <template #footer>
        <div style="flex: auto">
          <el-button @click="cancelClick">取消</el-button>
          <el-button type="primary" @click="confirmClick(formRef)" :loading="form.loading">确认</el-button>
        </div>
      </template>
    </el-drawer>

  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from "vue"
import { useBtnAuth } from '@/utils/btnAuth'
import { getTblCrawlEventsList, createTblCrawlEvents, deleteTblCrawlEvents } from "@/api/tbl_crawl_events"
import { formatTimeToStr } from "@/utils/date"
import { ElMessage, ElMessageBox } from "element-plus"

defineOptions({
  name: 'TblCrawlEvents'
})


const btnAuth = useBtnAuth()
const addCrawlEvent = () => {
  drawer.value = true
}

const drawer = ref(false)

const formRef = ref(null)
const form = reactive({
  title: '',
  details: '',
  occurced_time: '',
  loading: false
})

const handleClose = (done) => {
  formRef.value.resetFields()
  done()
}

function cancelClick() {
  drawer.value = false
}

function confirmClick(ref) {
  if (!ref) return
  form.loading = true
  ref.validate(async valid => {
    if (!valid) {
      form.loading = false
      console.log('fail')
    } else {
      const resp = await createTblCrawlEvents({ title: form.title, details: form.details, occured_time: form.occurced_time })
      form.loading = false
      if (resp.code == 0) {
        drawer.value = false
        getCrawlEventsList(1, 10)
        ElMessage.success({ message: "事件已添加", plain: true })
      }
    }
  })
}

const crawlData = reactive({
  page: 1,
  pageSize: 10,
  totalCount: null,
  list: []

})
// 获取事件列表
const getCrawlEventsList = async (currPage, pageSize) => {
  const resp = await getTblCrawlEventsList({ page: currPage, pageSize: pageSize })
  crawlData.totalCount = resp.data.total
  crawlData.list = resp.data.list
}

const deleteEvent = async (id) => {
  ElMessageBox.confirm("确认删除么?", "删除", {
    beforeClose: (action, instance, done) => {
      if (action == 'confirm') {
        instance.confirmButtonLoading = true
        deleteTblCrawlEvents({ ID: id }).then(resp => {
          done()
          if (resp.code === 0) {
            getCrawlEventsList(1, 10)
            ElMessage.success({ message: "已删除", plain: false })
          } else {
            ElMessage.error({ message: resp.msg, plain: false })
          }
        })
      } else {
        done()
      }
    },
    confirmButtonText: "确认"
  }).catch(action => {
  })
}

onMounted(async () => {
  await getCrawlEventsList(1, 10)
})

</script>

<style scoped>
:deep(.center-pagination) {
  margin: 0px;
}

.report-card {
  width: 1290px;
  max-height: calc(100vh - 124px);
}

.data-cont {
  text-align: center;
  font-weight: bold;
  font-size: 18px;
  padding-bottom: 31px;
  padding-top: 12px;
  height: calc(100vh - 264px);
}

.data-cont {
  color: #666;
  text-align: center;
  padding-top: 16px;
  font-size: 14px;
  font-weight: normal
}

.data-news {
  padding-left: 20px;
}

.data-news p {
  color: #666;
  padding-top: 6px;
  font-size: 14px;
  margin: 0px;
  text-align: left
}

.data-news li {
  margin-bottom: 28px
}

.li-circle::marker {
  color: #9f9f9f
}

.data-news li::marker {
  color: #9f9f9f
}

.data-news time {
  display: flex;
  justify-content: left;
  font-size: 12px;
  color: #666;
  position: relative;
  top: 4px;
}

/* .data-news  */
.news-title {
  font-weight: bold;
  text-align: left;
}
</style>
