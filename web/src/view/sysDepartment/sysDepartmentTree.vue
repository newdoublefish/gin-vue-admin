<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="addMenu('0')">新增根组织</el-button>
      </div>

      <!-- 由于此处菜单跟左侧列表一一对应所以不需要分页 pageSize默认999 -->
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" min-width="100" prop="ID" />
        <el-table-column align="left" label="名称" show-overflow-tooltip min-width="160" prop="name" />
        <el-table-column align="left" label="编码" show-overflow-tooltip min-width="160" prop="code" />
        <el-table-column align="left" label="描述" show-overflow-tooltip min-width="160" prop="description" />
        <el-table-column align="left" fixed="right" label="操作" width="300">
          <template #default="scope">
            <el-button
              size="small"
              type="text"
              icon="plus"
              @click="addMenu(scope.row.ID)"
            >添加子组织</el-button>
            <el-button
              size="small"
              type="text"
              icon="edit"
              @click="editMenu(scope.row.ID)"
            >编辑</el-button>
            <el-button
              size="small"
              type="text"
              icon="delete"
              @click="deleteMenu(scope.row.ID)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form ref="menuForm" :model="form" label-position="right" label-width="80px">
        <el-form-item label="组织编号:">
          <el-input v-model="form.code" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="组织名称:">
          <el-input v-model="form.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="父组织ID" style="width:30%">
          <el-cascader
            v-model="form.parentId"
            style="width:100%"
            :disabled="!isEdit"
            :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="描述:">
          <el-input
            v-model="form.description"
            :rows="4"
            type="textarea"
            placeholder="请输入"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createSysDepartment,
  deleteSysDepartment,
  // deleteSysDepartmentByIds,
  updateSysDepartment,
  findSysDepartment,
  // getSysDepartmentList,
  getSysDepartmentTree
} from '@/api/sysDepartment'
import icon from '@/view/superAdmin/menu/icon.vue'
import warningBar from '@/components/warningBar/warningBar.vue'
import { reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
const rules = reactive({
  path: [{ required: true, message: '请输入菜单name', trigger: 'blur' }],
  component: [
    { required: true, message: '请输入文件路径', trigger: 'blur' }
  ],
  'meta.title': [
    { required: true, message: '请输入菜单展示名称', trigger: 'blur' }
  ]
})
const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])
const searchInfo = ref({})
// 查询
const getTableData = async() => {
  const table = await getSysDepartmentTree({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()
// 新增参数

const form = ref({
  ID: 0,
  code: '',
  name: '',
  description: '',
  parentId: '',
})

// 删除菜单
const deleteMenu = (ID) => {
  ElMessageBox.confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteSysDepartment({ ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '已取消删除'
      })
    })
}
// 初始化弹窗内表格方法
const menuForm = ref(null)
const checkFlag = ref(false)
const initForm = () => {
  checkFlag.value = false
  menuForm.value.resetFields()
  form.value = {
    ID: 0,
    code: '',
    name: '',
    description: '',
    parentId: '',
  }
  console.log('init form', form.value)
}
// 关闭弹窗
const dialogFormVisible = ref(false)
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}
// 添加menu
const enterDialog = async() => {
  menuForm.value.validate(async valid => {
    if (valid) {
      let res
      if (isEdit.value) {
        res = await updateSysDepartment(form.value)
      } else {
        res = await createSysDepartment(form.value)
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: isEdit.value ? '编辑成功' : '添加成功!'
        })
        getTableData()
      }
      initForm()
      dialogFormVisible.value = false
    }
  })
}
const menuOption = ref([
  {
    ID: 0,
    title: '根菜单'
  }
])
const setOptions = () => {
  menuOption.value = [
    {
      ID: 0,
      title: '根目录'
    }
  ]

  setMenuOptions(tableData.value, menuOption.value, false)
  console.log(menuOption.value)
}

const setMenuOptions = (menuData, optionsData, disabled) => {
  menuData &&
        menuData.forEach(item => {
          console.log(form.value)
          if (item.children && item.children.length) {
            const option = {
              title: item.name,
              ID: item.ID,
              disabled: disabled || item.ID === form.value.ID,
              children: []
            }
            setMenuOptions(
              item.children,
              option.children,
              disabled || item.ID === form.value.ID
            )
            optionsData.push(option)
          } else {
            const option = {
              title: item.name,
              ID: item.ID,
              disabled: disabled || item.ID === form.value.ID
            }
            optionsData.push(option)
          }
        })
}
// 添加菜单方法，id为 0则为添加根菜单
const isEdit = ref(false)
const dialogTitle = ref('新增菜单')
const addMenu = (id) => {
  dialogTitle.value = '新增菜单'
  form.value.parentId = String(id)
  isEdit.value = false
  setOptions()
  dialogFormVisible.value = true
}
// 修改菜单方法
const editMenu = async(id) => {
  dialogTitle.value = '编辑菜单'
  const res = await findSysDepartment({ ID: id })
  form.value = res.data.reSysDp
  isEdit.value = true
  setOptions()
  dialogFormVisible.value = true
}
</script>

<script>
export default {
  name: 'Menus',
}
</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
.icon-column{
  display: flex;
  align-items: center;
  .el-icon{
    margin-right: 8px;
  }
}
</style>
