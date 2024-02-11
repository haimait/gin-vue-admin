<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
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
      
        <el-form-item label="标题" prop="title">
         <el-input v-model="searchInfo.title" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="文章作者" prop="author">
         <el-input v-model="searchInfo.author" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="参赛单位" prop="company">
         <el-input v-model="searchInfo.company" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="手机号" prop="phone">
         <el-input v-model="searchInfo.phone" placeholder="搜索条件" />

        </el-form-item>
            <el-form-item label="isLink" prop="isLink">
            <el-select v-model="searchInfo.isLink" clearable placeholder="请选择">
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
            </el-form-item>
        <el-form-item label="类型" prop="type">
         <el-input v-model="searchInfo.type" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="状态 1:待审核 2:审核通过" prop="status">
            
             <el-input v-model.number="searchInfo.status" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="前台显示 1:是 2:否" prop="publish">
            
             <el-input v-model.number="searchInfo.publish" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="获取奖项 5:无奖项 1:一等奖 2:二等奖 3:三等奖 4:特等奖" prop="awards">
            
             <el-input v-model.number="searchInfo.awards" placeholder="搜索条件" />

        </el-form-item>
            <el-form-item label="置顶 1:是 2:否" prop="top">
            <el-select v-model="searchInfo.top" clearable placeholder="请选择">
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
            </el-form-item>
        <el-form-item label="审核人" prop="auditorId">
            
             <el-input v-model.number="searchInfo.auditorId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="所属分类" prop="categoryId">
            
             <el-input v-model.number="searchInfo.categoryId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="标题" prop="title" width="120" />
        <el-table-column align="left" label="副标题" prop="subtitle" width="120" />
        <el-table-column align="left" label="关键字" prop="keywords" width="120" />
        <el-table-column align="left" label="文章描述" prop="description" width="120" />
        <el-table-column align="left" label="文章作者" prop="author" width="120" />
        <el-table-column align="left" label="参赛单位" prop="company" width="120" />
        <el-table-column align="left" label="手机号" prop="phone" width="120" />
        <el-table-column align="left" label="文章来源" prop="source" width="120" />
                      <el-table-column label="文章内容" width="200">
                         <template #default="scope">
                            [富文本内容]
                         </template>
                      </el-table-column>
                    <el-table-column label="扩展属性" width="200">
                        <template #default="scope">
                             <div class="file-list">
                               <el-tag v-for="file in scope.row.extend" :key="file.uid">{{file.name}}</el-tag>
                             </div>
                        </template>
                    </el-table-column>
        <el-table-column align="left" label="封面" prop="thumb" width="120" />
        <el-table-column align="left" label="isLink" prop="isLink" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.isLink) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Link" prop="link" width="120" />
        <el-table-column align="left" label="类型" prop="type" width="120" />
        <el-table-column sortable align="left" label="排序" prop="sortOrder" width="120" />
        <el-table-column align="left" label="状态 1:待审核 2:审核通过" prop="status" width="120" />
        <el-table-column align="left" label="前台显示 1:是 2:否" prop="publish" width="120" />
        <el-table-column align="left" label="展示次数" prop="viewCount" width="120" />
        <el-table-column align="left" label="投票次数" prop="ticketCount" width="120" />
        <el-table-column align="left" label="获取奖项 5:无奖项 1:一等奖 2:二等奖 3:三等奖 4:特等奖" prop="awards" width="120" />
        <el-table-column align="left" label="置顶 1:是 2:否" prop="top" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.top) }}</template>
        </el-table-column>
        <el-table-column align="left" label="审核人" prop="auditorId" width="120" />
        <el-table-column align="left" label="审核结果" prop="auditorFailed" width="120" />
        <el-table-column align="left" label="所属分类" prop="categoryId" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateArticleFunc(scope.row)">变更</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="标题:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入标题" />
            </el-form-item>
            <el-form-item label="副标题:"  prop="subtitle" >
              <el-input v-model="formData.subtitle" :clearable="true"  placeholder="请输入副标题" />
            </el-form-item>
            <el-form-item label="关键字:"  prop="keywords" >
              <el-input v-model="formData.keywords" :clearable="true"  placeholder="请输入关键字" />
            </el-form-item>
            <el-form-item label="文章描述:"  prop="description" >
              <el-input v-model="formData.description" :clearable="true"  placeholder="请输入文章描述" />
            </el-form-item>
            <el-form-item label="文章作者:"  prop="author" >
              <el-input v-model="formData.author" :clearable="true"  placeholder="请输入文章作者" />
            </el-form-item>
            <el-form-item label="参赛单位:"  prop="company" >
              <el-input v-model="formData.company" :clearable="true"  placeholder="请输入参赛单位" />
            </el-form-item>
            <el-form-item label="手机号:"  prop="phone" >
              <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入手机号" />
            </el-form-item>
            <el-form-item label="文章来源:"  prop="source" >
              <el-input v-model="formData.source" :clearable="true"  placeholder="请输入文章来源" />
            </el-form-item>
            <el-form-item label="文章内容:"  prop="content" >
              <RichEdit v-model="formData.content"/>
            </el-form-item>
            <el-form-item label="扩展属性:"  prop="extend" >
                <SelectFile v-model="formData.extend" />
            </el-form-item>
            <el-form-item label="封面:"  prop="thumb" >
              <el-input v-model="formData.thumb" :clearable="true"  placeholder="请输入封面" />
            </el-form-item>
            <el-form-item label="isLink:"  prop="isLink" >
              <el-switch v-model="formData.isLink" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="Link:"  prop="link" >
              <el-input v-model="formData.link" :clearable="true"  placeholder="请输入Link" />
            </el-form-item>
            <el-form-item label="类型：article/page/video:"  prop="type" >
                <el-select v-model="formData.type" placeholder="请选择类型：article/page/video" style="width:100%" :clearable="true" >
                   <el-option v-for="item in ['article','page','video']" :key="item" :label="item" :value="item" />
                </el-select>
            </el-form-item>
            <el-form-item label="排序:"  prop="sortOrder" >
              <el-input v-model.number="formData.sortOrder" :clearable="true" placeholder="请输入排序" />
            </el-form-item>
            <el-form-item label="状态 1:待审核 2:审核通过:"  prop="status" >
              <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入状态 1:待审核 2:审核通过" />
            </el-form-item>
            <el-form-item label="前台显示 1:是 2:否:"  prop="publish" >
              <el-input v-model.number="formData.publish" :clearable="true" placeholder="请输入前台显示 1:是 2:否" />
            </el-form-item>
            <el-form-item label="展示次数:"  prop="viewCount" >
              <el-input v-model.number="formData.viewCount" :clearable="true" placeholder="请输入展示次数" />
            </el-form-item>
            <el-form-item label="投票次数:"  prop="ticketCount" >
              <el-input v-model.number="formData.ticketCount" :clearable="true" placeholder="请输入投票次数" />
            </el-form-item>
            <el-form-item label="获取奖项 5:无奖项 1:一等奖 2:二等奖 3:三等奖 4:特等奖:"  prop="awards" >
              <el-input v-model.number="formData.awards" :clearable="true" placeholder="请输入获取奖项 5:无奖项 1:一等奖 2:二等奖 3:三等奖 4:特等奖" />
            </el-form-item>
            <el-form-item label="置顶 1:是 2:否:"  prop="top" >
              <el-switch v-model="formData.top" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="审核人:"  prop="auditorId" >
              <el-input v-model.number="formData.auditorId" :clearable="true" placeholder="请输入审核人" />
            </el-form-item>
            <el-form-item label="审核结果:"  prop="auditorFailed" >
              <el-input v-model="formData.auditorFailed" :clearable="true"  placeholder="请输入审核结果" />
            </el-form-item>
            <el-form-item label="所属分类:"  prop="categoryId" >
              <el-input v-model.number="formData.categoryId" :clearable="true" placeholder="请输入所属分类" />
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

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="标题">
                        {{ formData.title }}
                </el-descriptions-item>
                <el-descriptions-item label="副标题">
                        {{ formData.subtitle }}
                </el-descriptions-item>
                <el-descriptions-item label="关键字">
                        {{ formData.keywords }}
                </el-descriptions-item>
                <el-descriptions-item label="文章描述">
                        {{ formData.description }}
                </el-descriptions-item>
                <el-descriptions-item label="文章作者">
                        {{ formData.author }}
                </el-descriptions-item>
                <el-descriptions-item label="参赛单位">
                        {{ formData.company }}
                </el-descriptions-item>
                <el-descriptions-item label="手机号">
                        {{ formData.phone }}
                </el-descriptions-item>
                <el-descriptions-item label="文章来源">
                        {{ formData.source }}
                </el-descriptions-item>
                <el-descriptions-item label="文章内容">
                        [富文本内容]
                </el-descriptions-item>
                <el-descriptions-item label="扩展属性">
                        <div class="fileBtn" v-for="(item,index) in formData.extend" :key="index">
                          <el-button type="primary" text bg @click="onDownloadFile(item.url)">
                            <el-icon style="margin-right: 5px"><Download /></el-icon>
                            {{ item.name }}
                          </el-button>
                        </div>
                </el-descriptions-item>
                <el-descriptions-item label="封面">
                        {{ formData.thumb }}
                </el-descriptions-item>
                <el-descriptions-item label="isLink">
                    {{ formatBoolean(formData.isLink) }}
                </el-descriptions-item>
                <el-descriptions-item label="Link">
                        {{ formData.link }}
                </el-descriptions-item>
                <el-descriptions-item label="类型：article/page/video">
                        {{ formData.type }}
                </el-descriptions-item>
                <el-descriptions-item label="排序">
                        {{ formData.sortOrder }}
                </el-descriptions-item>
                <el-descriptions-item label="状态 1:待审核 2:审核通过">
                        {{ formData.status }}
                </el-descriptions-item>
                <el-descriptions-item label="前台显示 1:是 2:否">
                        {{ formData.publish }}
                </el-descriptions-item>
                <el-descriptions-item label="展示次数">
                        {{ formData.viewCount }}
                </el-descriptions-item>
                <el-descriptions-item label="投票次数">
                        {{ formData.ticketCount }}
                </el-descriptions-item>
                <el-descriptions-item label="获取奖项 5:无奖项 1:一等奖 2:二等奖 3:三等奖 4:特等奖">
                        {{ formData.awards }}
                </el-descriptions-item>
                <el-descriptions-item label="置顶 1:是 2:否">
                    {{ formatBoolean(formData.top) }}
                </el-descriptions-item>
                <el-descriptions-item label="审核人">
                        {{ formData.auditorId }}
                </el-descriptions-item>
                <el-descriptions-item label="审核结果">
                        {{ formData.auditorFailed }}
                </el-descriptions-item>
                <el-descriptions-item label="所属分类">
                        {{ formData.categoryId }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createArticle,
  deleteArticle,
  deleteArticleByIds,
  updateArticle,
  findArticle,
  getArticleList
} from '@/api/article'
import { getUrl } from '@/utils/image'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'Article'
})

// 自动化生成的字典（可能为空）以及字段
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
        categoryId: 0,
        })


// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               content : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               categoryId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
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
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    if (searchInfo.value.isLink === ""){
        searchInfo.value.isLink=null
    }
    if (searchInfo.value.top === ""){
        searchInfo.value.top=null
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
  const table = await getArticleList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteArticleFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteArticleByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateArticleFunc = async(row) => {
    const res = await findArticle({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rearticle
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteArticleFunc = async (row) => {
    const res = await deleteArticle({ ID: row.ID })
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
  const res = await findArticle({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rearticle
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          title: '',
          subtitle: '',
          keywords: '',
          description: '',
          author: '',
          company: '',
          phone: '',
          source: '',
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
          categoryId: 0,
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
        title: '',
        subtitle: '',
        keywords: '',
        description: '',
        author: '',
        company: '',
        phone: '',
        source: '',
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
        categoryId: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const downloadFile = (url) => {
    window.open(getUrl(url), '_blank')
}

</script>

<style>

.file-list{
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.fileBtn{
  margin-bottom: 10px;
}

.fileBtn:last-child{
  margin-bottom: 0;
}

</style>
