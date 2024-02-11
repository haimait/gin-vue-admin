<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="副标题:" prop="subtitle">
          <el-input v-model="formData.subtitle" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="关键字:" prop="keywords">
          <el-input v-model="formData.keywords" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章描述:" prop="description">
          <el-input v-model="formData.description" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章作者:" prop="author">
          <el-input v-model="formData.author" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="参赛单位:" prop="company">
          <el-input v-model="formData.company" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机号:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章来源:" prop="source">
          <el-input v-model="formData.source" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章内容:" prop="content">
          <RichEdit v-model="formData.content"/>
       </el-form-item>
        <el-form-item label="扩展属性:" prop="extend">
          <SelectFile v-model="formData.extend" />
       </el-form-item>
        <el-form-item label="封面:" prop="thumb">
          <el-input v-model="formData.thumb" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="isLink:" prop="isLink">
          <el-switch v-model="formData.isLink" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="Link:" prop="link">
          <el-input v-model="formData.link" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类型：article/page/video:" prop="type">
        <el-select v-model="formData.type" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="排序:" prop="sortOrder">
          <el-input v-model.number="formData.sortOrder" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态 1:待审核 2:审核通过:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="前台显示 1:是 2:否:" prop="publish">
          <el-input v-model.number="formData.publish" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="展示次数:" prop="viewCount">
          <el-input v-model.number="formData.viewCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="投票次数:" prop="ticketCount">
          <el-input v-model.number="formData.ticketCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="获取奖项 5:无奖项 1:一等奖 2:二等奖 3:三等奖 4:特等奖:" prop="awards">
          <el-input v-model.number="formData.awards" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="置顶 1:是 2:否:" prop="top">
          <el-switch v-model="formData.top" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="审核人:" prop="auditorId">
          <el-input v-model.number="formData.auditorId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="审核结果:" prop="auditorFailed">
          <el-input v-model="formData.auditorFailed" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="所属分类:" prop="category_id">
          <el-input v-model.number="formData.category_id" :clearable="true" placeholder="请输入" />
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
  createArticle,
  updateArticle,
  findArticle
} from '@/api/article'

defineOptions({
    name: 'ArticleForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectFile from '@/components/selectFile/selectFile.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            title: '',
            subtitle: '',
            keywords: '',
            description: '',
            author: '',
            company: '',
            phone: '',
            source: '',
            content: '',
            extend: [],
            thumb: '',
            isLink: false,
            link: '',
            sortOrder: 0,
            status: 0,
            publish: 0,
            viewCount: 0,
            ticketCount: 0,
            awards: 0,
            top: false,
            auditorId: 0,
            auditorFailed: '',
            category_id: 0,
        })
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               category_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findArticle({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rearticle
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
               res = await createArticle(formData.value)
               break
             case 'update':
               res = await updateArticle(formData.value)
               break
             default:
               res = await createArticle(formData.value)
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
