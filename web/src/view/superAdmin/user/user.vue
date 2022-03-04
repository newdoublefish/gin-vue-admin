<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="用户组织">
          <el-cascader
            v-model="searchInfo.department"
            :options="departmentOptions"
            :props="{checkStrictly: true, label:'name',value:'id'}"
          />
        </el-form-item>
        <el-form-item label="员工卡号">
          <el-input v-model="searchInfo.employeeID" placeholder="员工卡号" />
        </el-form-item>
        <el-form-item label="用户名">
          <el-input v-model="searchInfo.userName" placeholder="用户名" />
        </el-form-item>
        <el-form-item label="岗位">
          <el-select v-model="searchInfo.positionId" clearable placeholder="请选择">
            <el-option
              v-for="item in positions"
              :key="item.ID"
              :label="`${item.code}${item.name}`"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <warning-bar title="注：右上角头像下拉可切换角色" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="addUser()">新增用户</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column align="left" label="头像" min-width="50">
          <template #default="scope">
            <CustomPic style="margin-top:8px" :pic-src="scope.row.headerImg" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="员工号" min-width="150" prop="employeeID" />
        <el-table-column align="left" label="用户名" min-width="150" prop="userName" />
        <el-table-column align="left" label="昵称" min-width="100" prop="nickName">
          <template #default="scope">
            <p v-if="!scope.row.editFlag" class="nickName">{{ scope.row.nickName }}
              <el-icon class="pointer" color="#66b1ff" @click="openEidt(scope.row)">
                <edit />
              </el-icon>
            </p>
            <p v-if="scope.row.editFlag" class="nickName">
              <el-input v-model="scope.row.nickName" />
              <el-icon class="pointer" color="#67c23a" @click="enterEdit(scope.row)">
                <check />
              </el-icon>
              <el-icon class="pointer" color="#f23c3c" @click="closeEdit(scope.row)">
                <close />
              </el-icon>
            </p>
          </template>
        </el-table-column>
        <el-table-column align="left" label="用户角色" min-width="150">
          <template #default="scope">
            <el-cascader
              v-model="scope.row.authorityIds"
              :options="authOptions"
              :show-all-levels="false"
              collapse-tags
              :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              :clearable="false"
              @visible-change="(flag)=>{changeAuthority(scope.row,flag)}"
              @remove-tag="()=>{changeAuthority(scope.row,false)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="用户组织" min-width="150">
          <template #default="scope">
            <el-cascader
              v-model="scope.row.departmentIds"
              :options="departmentOptions"
              :show-all-levels="false"
              collapse-tags
              :props="{ multiple:true,checkStrictly: true,label:'name',value:'id',disabled:'disabled',emitPath:false}"
              :clearable="false"
              @visible-change="(flag)=>{changeDepartments(scope.row,flag)}"
              @remove-tag="()=>{changeDepartments(scope.row,false)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="岗位" min-width="150">
          <template #default="scope">
            {{ scope.row.position.name }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="200">
          <template #default="scope">
            <el-button type="text" icon="magic-stick" size="mini" @click="editUserBasicInfo(scope.row)">编辑用户</el-button>
            <el-popover :visible="scope.row.visible" placement="top" width="160">
              <p>确定要删除此用户吗</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="mini" @click="deleteUser(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button type="text" icon="delete" size="mini">删除</el-button>
              </template>
            </el-popover>
            <el-button type="text" icon="magic-stick" size="mini" @click="resetPassword(scope.row)">重置密码</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="addUserDialog" custom-class="user-dialog" :title="dialogTitle">
      <el-form ref="userForm" :rules="rules" :model="userInfo" label-width="80px">
        <el-form-item label="用户名" prop="userName">
          <el-input v-model="userInfo.userName" :disabled="dialogType === 'edit'" />
        </el-form-item>
        <el-form-item label="员工号" prop="employeeID">
          <el-input v-model="userInfo.employeeID" />
        </el-form-item>
        <el-form-item v-if="dialogType === 'addUser'" label="密码" prop="password">
          <el-input v-model="userInfo.password" />
        </el-form-item>
        <el-form-item label="别名" prop="nickName">
          <el-input v-model="userInfo.nickName" />
        </el-form-item>
        <el-form-item label="性别">
          <el-select
            v-model="userInfo.gender"
            style="width:100%"
            placeholder="请选择用户性别"
            clearable
          >
            <el-option
              v-for="item in genderOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="电话" prop="mobile">
          <el-input v-model="userInfo.mobile" />
        </el-form-item>
        <el-form-item label="身份证号码" prop="citizenNumber">
          <el-input v-model="userInfo.citizenNumber" />
        </el-form-item>
        <el-form-item label="用户角色" prop="authorityId">
          <el-cascader
            v-model="userInfo.authorityIds"
            style="width:100%"
            :options="authOptions"
            :show-all-levels="true"
            :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
            :clearable="false"
          />
        </el-form-item>
        <el-form-item label="部门">
          <el-cascader
            v-model="userInfo.departmentIds"
            style="width:100%"
            :options="departmentOptions"
            :show-all-levels="true"
            :props="{ multiple:true,checkStrictly: true,label:'name',value:'id',disabled:'disabled',emitPath:false}"
            :clearable="false"
          />
        </el-form-item>
        <el-form-item label="用户岗位">
          <el-select v-model="userInfo.positionId" clearable placeholder="Select">
            <el-option
              v-for="p in positions"
              :key="p.ID"
              :label="p.name"
              :value="p.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="用户类型">
          <el-select
            v-model="userInfo.staffType"
            style="width:100%"
            placeholder="请选择用户类型"
            clearable
          >
            <el-option
              v-for="item in staffTypeOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="用户状态">
          <el-select
            v-model="userInfo.staffStatus"
            style="width:100%"
            placeholder="请选择用户状态"
            clearable
          >
            <el-option
              v-for="item in staffStatusOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="头像" label-width="80px">
          <div style="display:inline-block" @click="openHeaderChange">
            <img v-if="userInfo.headerImg" class="header-img-box" :src="(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http')?path+userInfo.headerImg:userInfo.headerImg">
            <div v-else class="header-img-box">从媒体库选择</div>
          </div>
        </el-form-item>

      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeAddUserDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterAddUserDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="userInfo" :target-key="`headerImg`" />
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = import.meta.env.VITE_BASE_API
import {
  getAutoPositionList
} from '@/api/autoPosition' //  此处请自行替换地址
import {
  getUserList,
  setUserAuthorities,
  setUserDepartments,
  updateBasicInfo,
  register,
  deleteUser
} from '@/api/user'
import {
  getSysDepartmentTree
} from '@/api/sysDepartment' //  此处请自行替换地址
import { getAuthorityList } from '@/api/authority'
import { getDict } from '@/utils/dictionary'
import infoList from '@/mixins/infoList'
import { mapGetters } from 'vuex'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import warningBar from '@/components/warningBar/warningBar.vue'
import { setUserInfo, resetPassword } from '@/api/user.js'
export default {
  name: 'Api',
  components: { CustomPic, ChooseImg, warningBar },
  mixins: [infoList],
  data() {
    return {
      dialogTitle: '新增Api',
      dialogType: 'addUser',
      listApi: getUserList,
      staffTypeOptions: [],
      staffStatusOptions: [],
      genderOptions: [],
      path: path,
      authOptions: [],
      departmentOptions: [],
      positions: [],
      addUserDialog: false,
      backNickName: '',
      userInfo: {
        username: '',
        password: '',
        nickName: '',
        headerImg: '',
        authorityId: '',
        positionId: undefined,
        authorityIds: [],
        departmentIds: [],
        staffType: undefined,
        staffStatus: undefined,
        employeeID: '',
        gender: undefined,
        mobile: '',
        citizenNumber: ''
      },
      rules: {
        employeeID: [
          { required: true, message: '请输入员工号', trigger: 'blur' },
          { min: 10, max: 10, message: '长度为10位字符', trigger: 'blur' },
        ],
        userName: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 5, message: '最低5位字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入用户密码', trigger: 'blur' },
          { min: 6, message: '最低6位字符', trigger: 'blur' }
        ],
        nickName: [
          { required: true, message: '请输入用户昵称', trigger: 'blur' }
        ],
        authorityId: [
          { required: true, message: '请选择用户角色', trigger: 'blur' }
        ],
        departmentId: [
          { required: true, message: '请选择用户组织', trigger: 'blur' }
        ],
        mobile: [
          { required: false, message: '请输入手机号', trigger: 'blur' },
          {
            pattern: /^1[3-9]\d{9}$/,
            message: '手机号格式错误'
          }
        ],
        citizenNumber: [
          { required: false, message: '身份证号不能为空', trigger: 'blur' },
          { validator: this.validID, trigger: 'blur', message: '身份证号格式错误' }
        ]
      }
    }
  },
  computed: {
    ...mapGetters('user', ['token'])
  },
  watch: {
    tableData() {
      this.setAuthorityIds()
      this.setDepartmentIds()
    }
  },
  async created() {
    await this.loadStaffOptions()
    await this.getTableData()
    console.log(this.tableData)
    const res = await getAuthorityList({ page: 1, pageSize: 999 })
    const dpRes = await getSysDepartmentTree({ page: 1, pageSize: 999 })
    this.setOptions(res.data.list, dpRes.data.list)
    const positionRes = await getAutoPositionList({ page: 1, pageSize: 999 })
    if (positionRes.code === 0) {
      this.positions = positionRes.data.list
      console.log(this.positions)
    }
  },
  methods: {
    // 身份证验证
    async validID(rule, value, callback) {
      if (!value || value.length === 0) {
        return
      }
      // 身份证号码为15位或者18位，15位时全为数字，18位前17位为数字，最后一位是校验位，可能为数字或字符X
      const reg = /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/
      if (reg.test(value)) {
        // await this.go(value.length)
        callback()
      } else {
        callback(new Error('身份证号码不正确'))
      }
    },
    // 实现自动生成生日，性别，年龄
    go(val) {
      const iden = this.baseInfo.idCardNo
      let sex = null
      let birth = null
      const myDate = new Date()
      const month = myDate.getMonth() + 1
      const day = myDate.getDate()
      let age = 0

      if (val === 18) {
        age = myDate.getFullYear() - iden.substring(6, 10) - 1
        sex = iden.substring(16, 17)
        birth = iden.substring(6, 10) + '-' + iden.substring(10, 12) + '-' + iden.substring(12, 14)
        if (iden.substring(10, 12) < month || iden.substring(10, 12) === month && iden.substring(12, 14) <= day) age++
      }
      if (val === 15) {
        age = myDate.getFullYear() - iden.substring(6, 8) - 1901
        sex = iden.substring(13, 14)
        birth = '19' + iden.substring(6, 8) + '-' + iden.substring(8, 10) + '-' + iden.substring(10, 12)
        if (iden.substring(8, 10) < month || iden.substring(8, 10) === month && iden.substring(10, 12) <= day) age++
      }

      if (sex % 2 === 0) { sex = '0' } else { sex = '1' }
      this.baseInfo.sex = sex
      this.baseInfo.age = age
      this.baseInfo.birthday = birth
      this.baseInfo.birthplace = this.area[iden.substring(0, 2)]
    },
    async loadStaffOptions() {
      this.staffTypeOptions = await getDict('staffType')
      this.staffStatusOptions = await getDict('staffStatus')
      this.genderOptions = await getDict('gender')
    },
    onSubmit() {
      if (this.searchInfo !== undefined && this.searchInfo.department !== undefined) {
        // console.log('length', this.searchInfo.department.length)
        if (this.searchInfo.department.length > 0) {
          var last = this.searchInfo.department[this.searchInfo.department.length - 1]
          this.searchInfo.departmentId = last
        }
      }

      console.log('search info:', this.searchInfo)

      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    onReset() {
      this.searchInfo = {}
    },
    resetPassword(row) {
      this.$confirm(
        '是否将此用户密码重置为123456?',
        '警告',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      ).then(async() => {
        const res = await resetPassword({
          ID: row.ID,
        })
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: res.msg,
          })
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
          })
        }
      })
    },
    setAuthorityIds() {
      this.tableData && this.tableData.forEach((user) => {
        const authorityIds = user.authorities && user.authorities.map(i => {
          return i.authorityId
        })
        user.authorityIds = authorityIds
      })
    },
    setDepartmentIds() {
      this.tableData && this.tableData.forEach((user) => {
        const departmentIds = user.departments && user.departments.map(i => {
          return i.sysDepartmentId
        })
        user.departmentIds = departmentIds
      })
    },
    openHeaderChange() {
      this.$refs.chooseImg.open()
    },
    setOptions(authData, dpData) {
      this.authOptions = []
      this.departmentOptions = []
      this.setAuthorityOptions(authData, this.authOptions)
      this.setDepartmentOptions(dpData, this.departmentOptions)
    },
    openEidt(row) {
      if (this.tableData.some(item => item.editFlag)) {
        this.$message('当前存在正在编辑的用户')
        return
      }
      this.backNickName = row.nickName
      row.editFlag = true
    },
    async enterEdit(row) {
      const res = await setUserInfo({ nickName: row.nickName, ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '设置成功'
        })
      }
      this.backNickName = ''
      row.editFlag = false
    },
    closeEdit(row) {
      row.nickName = this.backNickName
      this.backNickName = ''
      row.editFlag = false
    },
    setDepartmentOptions(dpData, optionsData) {
      dpData &&
        dpData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              id: item.ID,
              name: item.name,
              children: []
            }
            this.setDepartmentOptions(item.children, option.children)
            optionsData.push(option)
          } else {
            const option = {
              id: item.ID,
              name: item.name,
            }
            optionsData.push(option)
          }
        })
    },
    setAuthorityOptions(AuthorityData, optionsData) {
      AuthorityData &&
        AuthorityData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              children: []
            }
            this.setAuthorityOptions(item.children, option.children)
            optionsData.push(option)
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName
            }
            optionsData.push(option)
          }
        })
    },
    async deleteUser(row) {
      const res = await deleteUser({ id: row.ID })
      if (res.code === 0) {
        this.$message.success('删除成功')
        await this.getTableData()
        row.visible = false
      }
    },
    async enterAddUserDialog() {
      if (this.dialogType === 'addUser') {
        this.userInfo.authorityId = this.userInfo.authorityIds[0]
        this.$refs.userForm.validate(async valid => {
          if (valid) {
            const res = await register(this.userInfo)
            if (res.code === 0) {
              this.$message({ type: 'success', message: '创建成功' })
            }
            await this.getTableData()
            this.closeAddUserDialog()
          }
        })
      } else if (this.dialogType === 'edit') {
        this.$refs.userForm.validate(async valid => {
          if (valid) {
            const res = await updateBasicInfo(this.userInfo)
            if (res.code === 0) {
              this.$message({ type: 'success', message: '更新成功' })
            }
            await this.getTableData()
            this.closeAddUserDialog()
          }
        })
        this.closeAddUserDialog()
      }
    },
    resetForm() {
      this.userInfo = {}
      this.userInfo.userName = ''
      this.userInfo.employeeID = ''
      this.userInfo.password = ''
      this.userInfo.headerImg = ''
      this.userInfo.nickName = ''
      this.userInfo.positionId = undefined
      this.userInfo.authorityIds = []
      this.userInfo.departmentIds = []
      this.userInfo.staffType = undefined
      this.userInfo.staffStatus = undefined
      this.userInfo.gender = undefined
      this.userInfo.mobile = ''
      this.userInfo.citizenNumber = ''
    },
    closeAddUserDialog() {
      if (this.dialogType === 'addUser') {
        this.$refs.userForm.resetFields()
        this.resetForm()
      }
      this.addUserDialog = false
    },
    openDialog(type) {
      switch (type) {
        case 'addUser':
          this.dialogTitle = '新增用户'
          break
        case 'edit':
          this.dialogTitle = '编辑用户'
          break
        default:
          break
      }
      this.dialogType = type
      this.addUserDialog = true
    },
    addUser() {
      this.resetForm()
      this.openDialog('addUser')
    },
    editUserBasicInfo(row) {
      console.log(row)
      this.userInfo = row
      this.openDialog('edit')
    },
    async changeAuthority(row, flag) {
      if (flag) {
        return
      }
      this.$nextTick(async() => {
        const res = await setUserAuthorities({
          ID: row.ID,
          authorityIds: row.authorityIds
        })
        if (res.code === 0) {
          this.$message({ type: 'success', message: '角色设置成功' })
        }
      })
    },
    async changeDepartments(row, flag) {
      if (flag) {
        return
      }
      this.$nextTick(async() => {
        const res = await setUserDepartments({
          ID: row.ID,
          departmentIds: row.departmentIds
        })
        if (res.code === 0) {
          this.$message({ type: 'success', message: '组织设置成功' })
        }
      })
    },
  }
}
</script>

<style lang="scss">
.user-dialog {
  .header-img-box {
  width: 200px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
.nickName{
  display: flex;
  justify-content: flex-start;
  align-items: center;
}
.pointer{
  cursor: pointer;
  font-size: 16px;
  margin-left: 2px;
}
</style>
