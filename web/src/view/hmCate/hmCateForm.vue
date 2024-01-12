<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="分类层级:" prop="level">
          <el-input v-model.number="formData.level" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="父级id:" prop="parentId">
          <el-input v-model.number="formData.parentId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="分类名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否显示:" prop="isShow">
          <el-switch v-model="formData.isShow" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="图片地址:" prop="imgPath">
          <el-input v-model="formData.imgPath" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类型:" prop="type">
        <el-select v-model="formData.type" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['article','good']" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="排序:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="扩展字段:" prop="extend">
          <el-input v-model="formData.extend" :clearable="true" placeholder="请输入" />
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
  createHmCategory,
  updateHmCategory,
  findHmCategory
} from '@/api/hmCate'

defineOptions({
    name: 'HmCategoryForm'
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
            level: 0,
            parentId: 0,
            name: '',
            isShow: false,
            imgPath: '',
            sort: 0,
            extend: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               type : [{
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
      const res = await findHmCategory({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reCategory
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
               res = await createHmCategory(formData.value)
               break
             case 'update':
               res = await updateHmCategory(formData.value)
               break
             default:
               res = await createHmCategory(formData.value)
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
