package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

func CrudAll(db *gorm.DB, info request.PageInfo, list interface{}) (int64, error) {
	var total int64
	if info.Page ==0{
		info.Page = 1
	}

	if info.PageSize == 0{
		info.PageSize = 10
	}

	offset := info.PageSize * (info.Page - 1)
	err := db.Limit(info.PageSize).Offset(offset).Find(list).Error
	if err != nil {
		return 0, err
	}
	err = db.Offset(-1).Count(&total).Error
	if err != nil {
		return 0, err
	}

	return total, err
}
