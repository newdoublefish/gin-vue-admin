// 自动生成模板AutoPosition
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AutoPosition 结构体
// 如果含有time.Time 请自行import time包
type AutoPosition struct {
      global.GVA_MODEL
      Code  string `json:"code" form:"code" gorm:"column:code;comment:岗位编码;size:192;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:岗位名称;size:192;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:岗位描述;size:192;"`
}


// TableName AutoPosition 表名
func (AutoPosition) TableName() string {
  return "AutoPosition"
}

