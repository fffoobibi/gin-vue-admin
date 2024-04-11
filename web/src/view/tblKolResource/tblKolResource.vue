<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="头像" prop="avatar" width="120" />
        <el-table-column align="left" label="均观看量" prop="averageViews" width="120" />
        <el-table-column align="left" label="categoryId字段" prop="categoryId" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.categoryId) }}</template>
        </el-table-column>
        <el-table-column align="left" label="频道账号" prop="channelAccount" width="120" />
        <el-table-column align="left" label="频道ID" prop="channelId" width="120" />
        <el-table-column align="left" label="频道链接" prop="channelUrl" width="120" />
        <el-table-column align="left" label="创建时间" prop="createTime" width="120" />
        <el-table-column align="left" label="邮箱" prop="email" width="120" />
        <el-table-column align="left" label="粉丝数量" prop="fans" width="120" />
        <el-table-column align="left" label="id字段" prop="id" width="120" />
        <el-table-column align="left" label="互动率" prop="interactionRate" width="120" />
        <el-table-column align="left" label="语言" prop="language" width="120" />
        <el-table-column align="left" label="最近抓取时间" prop="lastCrawlTime" width="120" />
        <el-table-column align="left" label="最近发布时间" prop="lastPublishedTime" width="120" />
        <el-table-column align="left" label="昵称" prop="nickname" width="120" />
        <el-table-column align="left" label="其它联系方式" prop="otherContact" width="120" />
        <el-table-column align="left" label="平台：1 - TikTok 2 - YouTube 3 - Instagram" prop="platform" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.platform) }}</template>
        </el-table-column>
        <el-table-column align="left" label="地区编码" prop="regionCode" width="120" />
        <el-table-column align="left" label="红人简介" prop="signature" width="120" />
        <el-table-column align="left" label="source字段" prop="source" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.source) }}</template>
        </el-table-column>
        <el-table-column align="left" label="总点赞数" prop="totalLikes" width="120" />
        <el-table-column align="left" label="总视频数" prop="totalVideos" width="120" />
        <el-table-column align="left" label="总观看量" prop="totalViews" width="120" />
        <el-table-column align="left" label="更新时间" prop="updateTime" width="120" />
        <el-table-column align="left" label="关联tbl_scrapy_sites 字段" prop="whichServer" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.whichServer) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateTblKolResourceFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #title>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="头像:"  prop="avatar" >
              <el-input v-model="formData.avatar" :clearable="true"  placeholder="请输入头像" />
            </el-form-item>
            <el-form-item label="均观看量:"  prop="averageViews" >
              <el-input v-model.number="formData.averageViews" :clearable="true" placeholder="请输入均观看量" />
            </el-form-item>
            <el-form-item label="categoryId字段:"  prop="categoryId" >
              <el-switch v-model="formData.categoryId" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="频道账号:"  prop="channelAccount" >
              <el-input v-model="formData.channelAccount" :clearable="true"  placeholder="请输入频道账号" />
            </el-form-item>
            <el-form-item label="频道ID:"  prop="channelId" >
              <el-input v-model="formData.channelId" :clearable="true"  placeholder="请输入频道ID" />
            </el-form-item>
            <el-form-item label="频道链接:"  prop="channelUrl" >
              <el-input v-model="formData.channelUrl" :clearable="true"  placeholder="请输入频道链接" />
            </el-form-item>
            <el-form-item label="创建时间:"  prop="createTime" >
              <el-input v-model.number="formData.createTime" :clearable="true" placeholder="请输入创建时间" />
            </el-form-item>
            <el-form-item label="邮箱:"  prop="email" >
              <el-input v-model="formData.email" :clearable="true"  placeholder="请输入邮箱" />
            </el-form-item>
            <el-form-item label="粉丝数量:"  prop="fans" >
              <el-input v-model.number="formData.fans" :clearable="true" placeholder="请输入粉丝数量" />
            </el-form-item>
            <el-form-item label="id字段:"  prop="id" >
              <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
            </el-form-item>
            <el-form-item label="互动率:"  prop="interactionRate" >
              <el-input v-model.number="formData.interactionRate" :clearable="true" placeholder="请输入互动率" />
            </el-form-item>
            <el-form-item label="语言:"  prop="language" >
              <el-input v-model="formData.language" :clearable="true"  placeholder="请输入语言" />
            </el-form-item>
            <el-form-item label="最近抓取时间:"  prop="lastCrawlTime" >
              <el-input v-model.number="formData.lastCrawlTime" :clearable="true" placeholder="请输入最近抓取时间" />
            </el-form-item>
            <el-form-item label="最近发布时间:"  prop="lastPublishedTime" >
              <el-input v-model.number="formData.lastPublishedTime" :clearable="true" placeholder="请输入最近发布时间" />
            </el-form-item>
            <el-form-item label="昵称:"  prop="nickname" >
              <el-input v-model="formData.nickname" :clearable="true"  placeholder="请输入昵称" />
            </el-form-item>
            <el-form-item label="其它联系方式:"  prop="otherContact" >
              <el-input v-model="formData.otherContact" :clearable="true"  placeholder="请输入其它联系方式" />
            </el-form-item>
            <el-form-item label="平台：1 - TikTok 2 - YouTube 3 - Instagram:"  prop="platform" >
              <el-switch v-model="formData.platform" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="地区编码:"  prop="regionCode" >
              <el-input v-model.number="formData.regionCode" :clearable="true" placeholder="请输入地区编码" />
            </el-form-item>
            <el-form-item label="红人简介:"  prop="signature" >
              <el-input v-model="formData.signature" :clearable="true"  placeholder="请输入红人简介" />
            </el-form-item>
            <el-form-item label="source字段:"  prop="source" >
              <el-switch v-model="formData.source" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="总点赞数:"  prop="totalLikes" >
              <el-input v-model.number="formData.totalLikes" :clearable="true" placeholder="请输入总点赞数" />
            </el-form-item>
            <el-form-item label="总视频数:"  prop="totalVideos" >
              <el-input v-model.number="formData.totalVideos" :clearable="true" placeholder="请输入总视频数" />
            </el-form-item>
            <el-form-item label="总观看量:"  prop="totalViews" >
              <el-input v-model.number="formData.totalViews" :clearable="true" placeholder="请输入总观看量" />
            </el-form-item>
            <el-form-item label="更新时间:"  prop="updateTime" >
              <el-input v-model.number="formData.updateTime" :clearable="true" placeholder="请输入更新时间" />
            </el-form-item>
            <el-form-item label="关联tbl_scrapy_sites 字段:"  prop="whichServer" >
              <el-switch v-model="formData.whichServer" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer size="800" v-model="detailShow" :before-close="closeDetailShow" title="查看详情" destroy-on-close>
          <template #title>
             <div class="flex justify-between items-center">
               <span class="text-lg">查看详情</span>
             </div>
         </template>
        <el-descriptions :column="1" border>
                <el-descriptions-item label="头像">
                        {{ formData.avatar }}
                </el-descriptions-item>
                <el-descriptions-item label="均观看量">
                        {{ formData.averageViews }}
                </el-descriptions-item>
                <el-descriptions-item label="categoryId字段">
                    {{ formatBoolean(formData.categoryId) }}
                </el-descriptions-item>
                <el-descriptions-item label="频道账号">
                        {{ formData.channelAccount }}
                </el-descriptions-item>
                <el-descriptions-item label="频道ID">
                        {{ formData.channelId }}
                </el-descriptions-item>
                <el-descriptions-item label="频道链接">
                        {{ formData.channelUrl }}
                </el-descriptions-item>
                <el-descriptions-item label="创建时间">
                        {{ formData.createTime }}
                </el-descriptions-item>
                <el-descriptions-item label="邮箱">
                        {{ formData.email }}
                </el-descriptions-item>
                <el-descriptions-item label="粉丝数量">
                        {{ formData.fans }}
                </el-descriptions-item>
                <el-descriptions-item label="id字段">
                        {{ formData.id }}
                </el-descriptions-item>
                <el-descriptions-item label="互动率">
                        {{ formData.interactionRate }}
                </el-descriptions-item>
                <el-descriptions-item label="语言">
                        {{ formData.language }}
                </el-descriptions-item>
                <el-descriptions-item label="最近抓取时间">
                        {{ formData.lastCrawlTime }}
                </el-descriptions-item>
                <el-descriptions-item label="最近发布时间">
                        {{ formData.lastPublishedTime }}
                </el-descriptions-item>
                <el-descriptions-item label="昵称">
                        {{ formData.nickname }}
                </el-descriptions-item>
                <el-descriptions-item label="其它联系方式">
                        {{ formData.otherContact }}
                </el-descriptions-item>
                <el-descriptions-item label="平台：1 - TikTok 2 - YouTube 3 - Instagram">
                    {{ formatBoolean(formData.platform) }}
                </el-descriptions-item>
                <el-descriptions-item label="地区编码">
                        {{ formData.regionCode }}
                </el-descriptions-item>
                <el-descriptions-item label="红人简介">
                        {{ formData.signature }}
                </el-descriptions-item>
                <el-descriptions-item label="source字段">
                    {{ formatBoolean(formData.source) }}
                </el-descriptions-item>
                <el-descriptions-item label="总点赞数">
                        {{ formData.totalLikes }}
                </el-descriptions-item>
                <el-descriptions-item label="总视频数">
                        {{ formData.totalVideos }}
                </el-descriptions-item>
                <el-descriptions-item label="总观看量">
                        {{ formData.totalViews }}
                </el-descriptions-item>
                <el-descriptions-item label="更新时间">
                        {{ formData.updateTime }}
                </el-descriptions-item>
                <el-descriptions-item label="关联tbl_scrapy_sites 字段">
                    {{ formatBoolean(formData.whichServer) }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createTblKolResource,
  deleteTblKolResource,
  deleteTblKolResourceByIds,
  updateTblKolResource,
  findTblKolResource,
  getTblKolResourceList
} from '@/api/tblKolResource'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'TblKolResource'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        avatar: '',
        averageViews: 0,
        categoryId: false,
        channelAccount: '',
        channelId: '',
        channelUrl: '',
        createTime: 0,
        email: '',
        fans: 0,
        id: 0,
        interactionRate: 0,
        language: '',
        lastCrawlTime: 0,
        lastPublishedTime: 0,
        nickname: '',
        otherContact: '',
        platform: false,
        regionCode: 0,
        signature: '',
        source: false,
        totalLikes: 0,
        totalVideos: 0,
        totalViews: 0,
        updateTime: 0,
        whichServer: false,
        })


// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    if (searchInfo.value.categoryId === ""){
        searchInfo.value.categoryId=null
    }
    if (searchInfo.value.platform === ""){
        searchInfo.value.platform=null
    }
    if (searchInfo.value.source === ""){
        searchInfo.value.source=null
    }
    if (searchInfo.value.whichServer === ""){
        searchInfo.value.whichServer=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getTblKolResourceList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteTblKolResourceFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteTblKolResourceByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateTblKolResourceFunc = async(row) => {
    const res = await findTblKolResource({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.retblKolResource
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTblKolResourceFunc = async (row) => {
    const res = await deleteTblKolResource({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findTblKolResource({ id: row.id })
  if (res.code === 0) {
    formData.value = res.data.retblKolResource
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          avatar: '',
          averageViews: 0,
          categoryId: false,
          channelAccount: '',
          channelId: '',
          channelUrl: '',
          createTime: 0,
          email: '',
          fans: 0,
          id: 0,
          interactionRate: 0,
          language: '',
          lastCrawlTime: 0,
          lastPublishedTime: 0,
          nickname: '',
          otherContact: '',
          platform: false,
          regionCode: 0,
          signature: '',
          source: false,
          totalLikes: 0,
          totalVideos: 0,
          totalViews: 0,
          updateTime: 0,
          whichServer: false,
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        avatar: '',
        averageViews: 0,
        categoryId: false,
        channelAccount: '',
        channelId: '',
        channelUrl: '',
        createTime: 0,
        email: '',
        fans: 0,
        id: 0,
        interactionRate: 0,
        language: '',
        lastCrawlTime: 0,
        lastPublishedTime: 0,
        nickname: '',
        otherContact: '',
        platform: false,
        regionCode: 0,
        signature: '',
        source: false,
        totalLikes: 0,
        totalVideos: 0,
        totalViews: 0,
        updateTime: 0,
        whichServer: false,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createTblKolResource(formData.value)
                  break
                case 'update':
                  res = await updateTblKolResource(formData.value)
                  break
                default:
                  res = await createTblKolResource(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
