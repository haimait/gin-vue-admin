<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="分类名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="排序:" prop="order">
          <el-input v-model.number="formData.order" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否启用:" prop="enable">
          <el-select v-model="formData.enable" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in boolOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  createCategory2,
  updateCategory2,
  findCategory2
} from '@/api/category2'

defineOptions({
    name: 'Category2Form'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const boolOptions = ref([])
const formData = ref({
            name: '',
            remark: '',
            order: 0,
            enable: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCategory2({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recategory2
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    boolOptions.value = await getDictFunc('bool')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCategory2(formData.value)
               break
             case 'update':
               res = await updateCategory2(formData.value)
               break
             default:
               res = await createCategory2(formData.value)
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
