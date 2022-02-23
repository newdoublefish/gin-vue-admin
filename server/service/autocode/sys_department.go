package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type SysDepartmentService struct {
}

// CreateSysDepartment 创建SysDepartment记录
// Author [piexlmax](https://github.com/piexlmax)
func (SysDpService *SysDepartmentService) CreateSysDepartment(SysDp autocode.SysDepartment) (err error) {
	err = global.GVA_DB.Create(&SysDp).Error
	return err
}

// DeleteSysDepartment 删除SysDepartment记录
// Author [piexlmax](https://github.com/piexlmax)
func (SysDpService *SysDepartmentService)DeleteSysDepartment(SysDp autocode.SysDepartment) (err error) {
	err = global.GVA_DB.Delete(&SysDp).Error
	return err
}

// DeleteSysDepartmentByIds 批量删除SysDepartment记录
// Author [piexlmax](https://github.com/piexlmax)
func (SysDpService *SysDepartmentService)DeleteSysDepartmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.SysDepartment{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSysDepartment 更新SysDepartment记录
// Author [piexlmax](https://github.com/piexlmax)
func (SysDpService *SysDepartmentService)UpdateSysDepartment(SysDp autocode.SysDepartment) (err error) {
	err = global.GVA_DB.Save(&SysDp).Error
	return err
}

// GetSysDepartment 根据id获取SysDepartment记录
// Author [piexlmax](https://github.com/piexlmax)
func (SysDpService *SysDepartmentService)GetSysDepartment(id uint) (err error, SysDp autocode.SysDepartment) {
	err = global.GVA_DB.Where("id = ?", id).First(&SysDp).Error
	return
}

// GetSysDepartmentInfoList 分页获取SysDepartment记录
// Author [piexlmax](https://github.com/piexlmax)
func (SysDpService *SysDepartmentService)GetSysDepartmentInfoList(info autoCodeReq.SysDepartmentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.SysDepartment{})
    var SysDps []autocode.SysDepartment
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&SysDps).Error
	return err, SysDps, total
}
