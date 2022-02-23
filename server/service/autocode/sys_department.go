package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
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

func (SysDpService *SysDepartmentService)GetDepartmentTreeMap() (err error, treeMap map[int][]response.DepartmentTreeResponse) {
	var allDepartments []autocode.SysDepartment
	treeMap = make(map[int][]response.DepartmentTreeResponse)
	err = global.GVA_DB.Find(&allDepartments).Error
	if err!=nil{
		return err, treeMap
	}
	for _, v := range allDepartments{
		treeMap[v.ParentId] = append(treeMap[v.ParentId], response.DepartmentTreeResponse{
			SysDepartment: v,
		})
	}
	return err, treeMap
}

func (SysDpService *SysDepartmentService)FillDepartmentChildren(treeMap map[int][]response.DepartmentTreeResponse, dr *response.DepartmentTreeResponse)(err error){
	dr.Children = treeMap[int(dr.ID)]
	for i, _ := range dr.Children{
		err = SysDpService.FillDepartmentChildren(treeMap, &dr.Children[i])
	}
	return err
}


func (SysDpService *SysDepartmentService)GetDepartmentTree() (departments []response.DepartmentTreeResponse,err error) {
	err, treeMap := SysDpService.GetDepartmentTreeMap()
	if err!=nil{
		return
	}
	departments = treeMap[0]
	for i, _ := range departments{
		err = SysDpService.FillDepartmentChildren(treeMap, &departments[i])
	}
	return departments, err
}

