package system

type SysUserDepartment struct {
	SysUserId       uint `json:"sysUserId" gorm:"column:sys_user_id"`
	SysDepartmentId uint `json:"sysDepartmentId" gorm:"column:sys_department_id"`
}

func (s *SysUserDepartment) TableName() string {
	return "sys_user_department"
}
