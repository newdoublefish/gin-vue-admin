package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID             `json:"uuid" gorm:"comment:用户UUID"` // 用户UUID
	EmployeeID  string                `json:"employeeID" gorm:"comment:工号"`
	Username    string                `json:"userName" gorm:"comment:用户登录名"`                                                        // 用户登录名
	Password    string                `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName    string                `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	SideMode    string                `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	HeaderImg   string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor   string                `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor string                `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	AuthorityId string                `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	PositionId  uint                  `json:"positionId" gorm:"default:0; comment:岗位ID"`
	Authority   SysAuthority          `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority        `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Departments []SysUserDepartment   `json:"departments"`
	Position    autocode.AutoPosition `json:"position"`
	StaffType   uint                  `json:"staffType" gorm:"default:1;comment:用户类型"`   //1 regular 正式工,2 temporary 临时工,
	StaffStatus uint                  `json:"staffStatus" gorm:"default:1;comment:用户状态"` //1 employed 在职, 2 unemployed 离职
}
