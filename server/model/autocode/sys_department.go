// 自动生成模板SysDepartment
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysDepartment 结构体
// 如果含有time.Time 请自行import time包
type SysDepartment struct {
      global.GVA_MODEL
      Code  string `json:"code" form:"code" gorm:"column:code;comment:编号;size:192;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:部门名称;size:192;"`
      ParentId  *int `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:上级部门;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:描述;size:1024;"`
}


// TableName SysDepartment 表名
func (SysDepartment) TableName() string {
  return "SysDepartment"
}

