package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// User register structure
type Register struct {
	Username      string   `json:"userName"`
	Password      string   `json:"passWord"`
	NickName      string   `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg     string   `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	AuthorityId   string   `json:"authorityId" gorm:"default:888"`
	AuthorityIds  []string `json:"authorityIds"`
	DepartmentIds []uint   `json:"departmentIds"`
	PositionId    uint     `json:"positionId"`
	EmployeeID    string   `json:"employeeID"`
	StaffType     uint     `json:"staffType"`
	StaffStatus   uint     `json:"staffStatus"`
	Gender        uint     `json:"gender"`
	Mobile        string   `json:"mobile"`
	CitizenNumber string   `json:"citizenNumber"`
}

func (r Register) ToSysUser() system.SysUser {
	return system.SysUser{Username: r.Username, PositionId: r.PositionId, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, EmployeeID: r.EmployeeID, StaffType: r.StaffType, StaffStatus: r.StaffStatus, Gender: r.Gender, Mobile: r.Mobile, CitizenNumber: r.CitizenNumber}
}

// User register structure
type UpdateUserBasicInfo struct {
	ID            uint
	NickName      string   `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg     string   `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	AuthorityIds  []string `json:"authorityIds"`
	DepartmentIds []uint   `json:"departmentIds"`
	PositionId    uint     `json:"positionId"`
	EmployeeID    string   `json:"employeeID"`
	StaffType     uint     `json:"staffType"`
	StaffStatus   uint     `json:"staffStatus"`
	Gender        uint     `json:"gender"`
	Mobile        string   `json:"mobile"`
	CitizenNumber string   `json:"citizenNumber"`
}

func (r UpdateUserBasicInfo) ToSysUser() system.SysUser {
	return system.SysUser{NickName: r.NickName, HeaderImg: r.HeaderImg, PositionId: r.PositionId, EmployeeID: r.EmployeeID, StaffType: r.StaffType, StaffStatus: r.StaffStatus, Gender: r.Gender, Mobile: r.Mobile, CitizenNumber: r.CitizenNumber}
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type OAuth struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId string `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []string `json:"authorityIds"` // 角色ID
}

type SetUserDepartments struct {
	ID            uint
	DepartmentIds []uint `json:"departmentIds"` // 角色ID
}

type UserSearch struct {
	request.PageInfo
	EmployeeID   string `json:"employeeID"`
	Username     string `json:"userName"`
	DepartmentId uint   `json:"departmentId"`
	PositionId   uint   `json:"positionId"`
}
