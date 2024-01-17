<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">
        <!-- <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
        <el-form-item label="分类层级" prop="level">
            
             <el-input v-model.number="searchInfo.level" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="父级id" prop="parentId">
            
             <el-input v-model.number="searchInfo.parentId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="分类名称" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
            <el-form-item label="是否显示" prop="isShow">
              <el-select v-model="searchInfo.isShow" clearable placeholder="请选择">
                  <el-option
                      key="true"
                      label="是"
                      value="true">
                  </el-option>
                  <el-option
                      key="false"
                      label="否"
                      value="false">
                  </el-option>
              </el-select>
        </el-form-item> -->
        <el-form-item label="类型:" prop="type">
          <el-select v-model="searchInfo.type" placeholder="请选择类型">
            <el-option v-for="item in typeOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <!-- <el-button icon="refresh" @click="onReset">重置</el-button> -->
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <!-- <el-button type="primary" icon="plus" @click="openDialog">新增</el-button> -->
        <el-button type="primary" icon="plus" @click="addCate(0)">增加根分类</el-button>
        <el-icon class="cursor-pointer"
          @click="toDoc('https://www.bilibili.com/video/BV1kv4y1g7nT/?p=4&vd_source=f2640257c21e3b547a790461ed94875e')">
          <VideoCameraFilled />
        </el-icon>
        <!-- <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover> -->
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange" @sort-change="sortChange">
        <!-- <el-table-column type="selection" width="55" /> -->
        <el-table-column align="left" label="ID" min-width="30" prop="ID" />

        <el-table-column align="left" label="分类名称" prop="name" width="120" />
        <el-table-column align="left" label="是否显示" prop="isShow" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isShow) == "是" ? "显示" : "隐藏" }}</template>
        </el-table-column>
        <el-table-column align="left" label="预览" width="100">
          <template #default="scope">
            <CustomPic pic-type="file" :pic-src="scope.row.imgPath" preview />
          </template>
        </el-table-column>
        <el-table-column align="left" label="类型" prop="type" width="120" />
        <el-table-column align="left" label="排序" prop="sort" width="120" />
        <!-- <el-table-column align="left" label="父级ID" prop="parentId" width="120" /> -->
        <el-table-column align="left" label="分类层级" prop="level" width="120" />
        <el-table-column align="left" label="扩展字段11" prop="extend" show-overflow-tooltip='true' width="120" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">
                {{ formatDate(scope.row.CreatedAt) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="120">
          <template #default="scope">
            <!-- <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button> -->
            <el-button type="primary" link icon="plus" @click="addCate(scope.row)">添加子菜单</el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateHmCategoryFunc(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div> -->
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type === 'create' ? '添加分类' : '修改分类'"
      destroy-on-close>
      <el-scrollbar height="500px">
        <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="85px">
          <!-- <el-form-item label="父级id:"  prop="parentId" >
              <el-input v-model.number="formData.parentId" :clearable="true" placeholder="请输入父级id" />
            </el-form-item> -->
          <el-form-item label="父节点ID" prop="parentId" style="width:98%">
            <el-cascader v-model="formData.parentId" style="width:100%" :disabled="!isEdit" :options="cateOption"
              :props="{ checkStrictly: true, label: 'name', value: 'ID', disabled: 'disabled', emitPath: false }"
              :show-all-levels="false" filterable />
          </el-form-item>
          <el-form-item label="分类名称:" prop="name" style="width:98%">
            <el-input v-model="formData.name" :clearable="true" placeholder="请输入分类名称" />
          </el-form-item>
          <el-form-item label="是否显示:" prop="isShow">
            <el-select v-model="formData.isShow" style="width:98%" placeholder="请选择类型" clearable>
              <el-option v-for="item in isShowOptions" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
          </el-form-item>


          <el-form-item label="排序:" prop="sort" style="width:98%">
            <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入排序" />
          </el-form-item>
          <el-form-item label="扩展字段:" prop="extend" style="width:98%">
            <el-input v-model="formData.extend" :clearable="true" placeholder="请输入扩展字段" />
          </el-form-item>
          <el-form-item label="类型:" prop="type">
            <el-select v-model="formData.type" placeholder="请选择类型" style="width:98%" :clearable="true" :disabled="true">
              <!-- <el-option v-for="item in ['article','good']" :key="item" :label="item" :value="item" /> -->
              <el-option v-for="item in typeOptions" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
          </el-form-item>
          <el-form-item label="图片" prop="imgPath" label-width="80px">
            <div class="flex justify-center">
              <SelectImage v-model="formData.imgPath" file-type="image" />
            </div>
          </el-form-item>

        </el-form>

      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情"
      destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
          <el-descriptions-item label="分类层级">
            {{ formData.level }}
          </el-descriptions-item>
          <el-descriptions-item label="父级id">
            {{ formData.parentId }}
          </el-descriptions-item>
          <el-descriptions-item label="分类名称">
            {{ formData.name }}
          </el-descriptions-item>
          <el-descriptions-item label="是否显示">
            {{ formatBoolean(formData.isShow) }}
          </el-descriptions-item>
          <el-descriptions-item label="图片地址">
            {{ formData.imgPath }}
          </el-descriptions-item>
          <el-descriptions-item label="类型">
            {{ formData.type }}
          </el-descriptions-item>
          <el-descriptions-item label="排序">
            {{ formData.sort }}
          </el-descriptions-item>
          <el-descriptions-item label="扩展字段">
            {{ formData.extend }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="formData" :target-key="`imgPath`" />
  </div>
</template>

<script setup>
import {
  createHmCategory,
  deleteHmCategory,
  deleteHmCategoryByIds,
  updateHmCategory,
  findHmCategory,
  getHmCategoryList,
  getHmCategoryTree
} from "@/plugin/hmcate/api/hmcate.js"

// 全量引入格式化工具 请按需保留
import { formatDate, formatBoolean} from '@/utils/format'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import SelectImage from '@/components/selectImage/selectImage.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { toDoc } from '@/utils/doc'

defineOptions({
  name: 'HmCategory'
})

const isShowOptions = ref([
  {
    label: '是',
    value: true
  },
  {
    label: '否',
    value: false
  }
])

const typeOptions = ref([
  {
    label: '文章',
    value: 'article'
  },
  {
    label: '商品',
    value: 'goods'
  }
])



// 验证规则
const rule = reactive({
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  type: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({
  type: 'article',
})
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  level: 0,
  parentId: 0,
  name: '',
  isShow: true,
  type: searchInfo.value.type,
  imgPath: '',
  sort: 0,
  extend: '',
})
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    if (searchInfo.value.isShow === "") {
      searchInfo.value.isShow = null
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

// 获取分类分页树列表查询查询
const getTableData = async () => {
  const table = await getHmCategoryTree({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()

// 获取分类分页列表查询  弃用
const getTableDatabak = async () => {
  const table = await getHmCategoryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}


// ============== 表格控制部分结束 ===============
const cateOption = ref([
  {
    ID: '0',
    title: '根菜单'
  }
])
// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
  cateOption.value = [
    {
      ID: 0,
      name: '根分类'
    }
  ]
  setCateOptions(tableData.value, cateOption.value, false)
}
const setCateOptions = (menuData, optionsData, disabled) => {
  menuData &&
    menuData.forEach(item => {
      if (item.children && item.children.length) {
        const option = {
          name: item.name,
          ID: item.ID,
          disabled: disabled || item.ID === formData.value.ID,
          children: []
        }
        setCateOptions(
          item.children,
          option.children,
          disabled || item.ID === formData.value.ID
        )
        optionsData.push(option)
      } else {
        const option = {
          name: item.name,
          ID: item.ID,
          disabled: disabled || item.ID === formData.value.ID
        }
        optionsData.push(option)
      }
    })
}


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
    deleteHmCategoryFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
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
      ids.push(item.ID)
    })
  const res = await deleteHmCategoryByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}



// 删除行
const deleteHmCategoryFunc = async (row) => {
  const res = await deleteHmCategory({ ID: row.ID })
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
const dialogTitle = ref('新增菜单')

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findHmCategory({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reCategory
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    level: 0,
    parentId: 0,
    name: '',
    isShow: true,
    type: searchInfo.value.type,
    imgPath: '',
    sort: 0,
    extend: '',
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')
// 打开弹窗
// const openDialog = () => {
//     type.value = 'create'
//     dialogFormVisible.value = true
// }
// 添加分类方法，rowData为 0则为添加根分类
const isEdit = ref(false) // 是否可以编辑
const addCate = (rowData) => {
  dialogTitle.value = '新增分类'
  type.value = 'create'
  isEdit.value = false
  formData.value.type = searchInfo.value.type
  if (rowData != "0") {
    formData.value.parentId = rowData.ID
  }
  setOptions()
  dialogFormVisible.value = true
}
// 更新行
const updateHmCategoryFunc = async (row) => {
  const res = await findHmCategory({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    isEdit.value = true
    setOptions()
    formData.value = res.data.reCategory
    dialogFormVisible.value = true
  }
}
// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    level: 0,
    parentId: 0,
    name: '',
    isShow: true,
    type: searchInfo.value.type,
    imgPath: '',
    sort: 0,
    extend: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
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
      closeDialog()
      getTableData()
    }
  })
}

</script>

<style lang="scss">

</style>
