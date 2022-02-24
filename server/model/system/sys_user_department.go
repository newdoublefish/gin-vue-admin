package system

type SysUserDepartment struct {
	SysUserId       uint   `gorm:"column:sys_user_id"`
	SysDepartmentId string `gorm:"column:sys_department_id"`
}

func (s *SysUserDepartment) TableName() string {
	return "sys_user_department"
}
