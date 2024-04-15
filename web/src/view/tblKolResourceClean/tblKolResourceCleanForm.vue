<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="头像:" prop="avatar">
          <el-input v-model="formData.avatar" :clearable="true"  placeholder="请输入头像" />
       </el-form-item>
        <el-form-item label="均观看量:" prop="averageViews">
          <el-input v-model.number="formData.averageViews" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="categoryId字段:" prop="categoryId">
          <el-switch v-model="formData.categoryId" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="频道账号:" prop="channelAccount">
          <el-input v-model="formData.channelAccount" :clearable="true"  placeholder="请输入频道账号" />
       </el-form-item>
        <el-form-item label="频道ID:" prop="channelId">
          <el-input v-model="formData.channelId" :clearable="true"  placeholder="请输入频道ID" />
       </el-form-item>
        <el-form-item label="频道链接:" prop="channelUrl">
          <el-input v-model="formData.channelUrl" :clearable="true"  placeholder="请输入频道链接" />
       </el-form-item>
        <el-form-item label="创建时间:" prop="createTime">
          <el-input v-model.number="formData.createTime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true"  placeholder="请输入邮箱" />
       </el-form-item>
        <el-form-item label="粉丝数量:" prop="fans">
          <el-input v-model.number="formData.fans" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="互动率:" prop="interactionRate">
          <el-input v-model.number="formData.interactionRate" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="语言:" prop="language">
          <el-input v-model="formData.language" :clearable="true"  placeholder="请输入语言" />
       </el-form-item>
        <el-form-item label="最近抓取时间:" prop="lastCrawlTime">
          <el-input v-model.number="formData.lastCrawlTime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="最近发布时间:" prop="lastPublishedTime">
          <el-input v-model.number="formData.lastPublishedTime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="昵称:" prop="nickname">
          <el-input v-model="formData.nickname" :clearable="true"  placeholder="请输入昵称" />
       </el-form-item>
        <el-form-item label="其它联系方式:" prop="otherContact">
          <el-input v-model="formData.otherContact" :clearable="true"  placeholder="请输入其它联系方式" />
       </el-form-item>
        <el-form-item label="平台：1 - TikTok 2 - YouTube 3 - Instagram:" prop="platform">
          <el-switch v-model="formData.platform" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="关联tbl_kol_resource_polices:" prop="policyId">
          <el-switch v-model="formData.policyId" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="地区编码:" prop="regionCode">
          <el-input v-model.number="formData.regionCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="红人简介:" prop="signature">
          <el-input v-model="formData.signature" :clearable="true"  placeholder="请输入红人简介" />
       </el-form-item>
        <el-form-item label="source字段:" prop="source">
          <el-switch v-model="formData.source" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="总点赞数:" prop="totalLikes">
          <el-input v-model.number="formData.totalLikes" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总视频数:" prop="totalVideos">
          <el-input v-model.number="formData.totalVideos" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总观看量:" prop="totalViews">
          <el-input v-model.number="formData.totalViews" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新时间:" prop="updateTime">
          <el-input v-model.number="formData.updateTime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="关联tbl_scrapy_sites 字段:" prop="whichServer">
          <el-switch v-model="formData.whichServer" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createTblKolResourceClean,
  updateTblKolResourceClean,
  findTblKolResourceClean
} from '@/api/tblKolResourceClean'

defineOptions({
    name: 'TblKolResourceCleanForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            id: 0,
            avatar: '',
            averageViews: 0,
            categoryId: false,
            channelAccount: '',
            channelId: '',
            channelUrl: '',
            createTime: 0,
            email: '',
            fans: 0,
            interactionRate: 0,
            language: '',
            lastCrawlTime: 0,
            lastPublishedTime: 0,
            nickname: '',
            otherContact: '',
            platform: false,
            policyId: false,
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

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTblKolResourceClean({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retblKolResourceClean
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createTblKolResourceClean(formData.value)
               break
             case 'update':
               res = await updateTblKolResourceClean(formData.value)
               break
             default:
               res = await createTblKolResourceClean(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
