package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type AutoPositionService struct {
}

// CreateAutoPosition 创建AutoPosition记录
// Author [piexlmax](https://github.com/piexlmax)
func (PositionService *AutoPositionService) CreateAutoPosition(Position autocode.AutoPosition) (err error) {
	err = global.GVA_DB.Create(&Position).Error
	return err
}

// DeleteAutoPosition 删除AutoPosition记录
// Author [piexlmax](https://github.com/piexlmax)
func (PositionService *AutoPositionService)DeleteAutoPosition(Position autocode.AutoPosition) (err error) {
	err = global.GVA_DB.Delete(&Position).Error
	return err
}

// DeleteAutoPositionByIds 批量删除AutoPosition记录
// Author [piexlmax](https://github.com/piexlmax)
func (PositionService *AutoPositionService)DeleteAutoPositionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.AutoPosition{},"id in ?",ids.Ids).Error
	return err
}

// UpdateAutoPosition 更新AutoPosition记录
// Author [piexlmax](https://github.com/piexlmax)
func (PositionService *AutoPositionService)UpdateAutoPosition(Position autocode.AutoPosition) (err error) {
	err = global.GVA_DB.Save(&Position).Error
	return err
}

// GetAutoPosition 根据id获取AutoPosition记录
// Author [piexlmax](https://github.com/piexlmax)
func (PositionService *AutoPositionService)GetAutoPosition(id uint) (err error, Position autocode.AutoPosition) {
	err = global.GVA_DB.Where("id = ?", id).First(&Position).Error
	return
}

// GetAutoPositionInfoList 分页获取AutoPosition记录
// Author [piexlmax](https://github.com/piexlmax)
func (PositionService *AutoPositionService)GetAutoPositionInfoList(info autoCodeReq.AutoPositionSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.AutoPosition{})
    var Positions []autocode.AutoPosition
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&Positions).Error
	return err, Positions, total
}
