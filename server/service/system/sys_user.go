package system

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/casdoor/casdoor-go-sdk/auth"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/attendant"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	utils2 "github.com/flipped-aurora/gin-vue-admin/server/utils/enums"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

type UserService struct{}

func (userService *UserService) Register(u system.SysUser) (err error, userInter system.SysUser) {

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var user system.SysUser
		if !errors.Is(tx.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
			return errors.New("用户名已注册")
		}

		if !errors.Is(tx.Where("employee_id = ?", u.EmployeeID).First(&user).Error, gorm.ErrRecordNotFound) { // 判断员工号是否注册
			return errors.New("员工号已注册")
		}
		// 否则 附加uuid 密码md5简单加密 注册
		u.Password = utils.MD5V([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err := tx.Create(&u).Error
		//if len(u.Departments) > 0{
		//	for i,_ := range u.Departments{
		//		u.Departments[i].SysUserId = u.ID
		//	}
		//	err = tx.Create(&u.Departments).Error
		//}
		return err
	})

	if err != nil {
		return err, userInter
	}

	userService.CacheSingleUserToRedis(u)

	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *UserService) Login(u *system.SysUser) (err error, userInter *system.SysUser) {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init"), nil
	}

	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authorities").Preload("Authority").First(&user).Error
	return err, &user
}

func (userService *UserService) Oauth(code string, state string) (err error, token string, userInter system.SysUser) {
	t, err := auth.GetOAuthToken(code, state)
	if err != nil {
		return
	} else {
		fmt.Printf("AccessToken:%s\n", t.AccessToken)
	}
	token = t.AccessToken

	claims, err := auth.ParseJwtToken(token)
	if err != nil {
		return
	} else {
		fmt.Printf("claims:%v\n", claims)
	}

	err = global.GVA_DB.Where("username = ?", claims.Name).Preload("Authorities").Preload("Authority").First(&userInter).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser

func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (err error, userInter *system.SysUser) {
	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info systemReq.UserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	var responseList []response.SysUserListResponse

	//TODO: 搜索逻辑

	if info.Username != "" {
		db = db.Where("sys_users.username = ?", info.Username)
	}

	if info.OriginType != 0 {
		db = db.Where("sys_users.origin_type = ?", info.OriginType)
	}

	if info.NickName != "" {
		db = db.Where("sys_users.nick_name like ?", "%"+info.NickName+"%")

	}

	if info.EmployeeID != "" {
		db = db.Where("sys_users.employee_id = ?", info.EmployeeID)
	}

	if info.PositionId != 0 {
		db = db.Where("sys_users.position_id", info.PositionId)
	}

	if info.PositionId != 0 {
		db = db.Where("sys_users.position_id", info.PositionId)
	}

	if info.AuthorityId != "" {
		db = db.Where("sys_users.authority_id", info.AuthorityId)
	}

	if info.StaffStatus != 0 {
		db = db.Where("sys_users.staff_status", info.StaffStatus)
	}

	if info.StaffType != 0 {
		db = db.Where("sys_users.staff_type", info.StaffType)
	}

	if info.DepartmentId != 0 {
		db = db.Select("sys_users.*, sys_user_department.sys_department_id as sys_department_id").Joins("left join sys_user_department on sys_user_department.sys_user_id = sys_users.id ").Where("sys_department_id = ?", info.DepartmentId)
	}

	err = db.Offset(-1).Count(&total).Error
	if err != nil {
		return
	}

	db.Select("sys_users.*")

	db = db.Limit(limit).Offset(offset)
	db = db.Preload("Departments")
	db = db.Preload("Position")
	db = db.Preload("Authorities")
	db = db.Preload("Authority")
	err = db.Find(&userList).Error

	for index, asb := range userList {
		var re response.SysUserListResponse
		re.SysUser = userList[index]
		re.OriginTypeStr = asb.OriginType.String()
		responseList = append(responseList, re)
	}

	return err, responseList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (userService *UserService) SetUserAuthority(id uint, uuid uuid.UUID, authorityId string) (err error) {
	assignErr := global.GVA_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUseAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&system.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

func (userService *UserService) UpAuthorities(tx *gorm.DB, id uint, authorityIds []string) (err error) {
	TxErr := tx.Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
	if TxErr != nil {
		return TxErr
	}
	useAuthority := []system.SysUseAuthority{}
	for _, v := range authorityIds {
		useAuthority = append(useAuthority, system.SysUseAuthority{
			id, v,
		})
	}
	TxErr = tx.Create(&useAuthority).Error
	if TxErr != nil {
		return TxErr
	}
	TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
	if TxErr != nil {
		return TxErr
	}
	// 返回 nil 提交事务
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

func (userService *UserService) SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		return userService.UpAuthorities(tx, id, authorityIds)
	})
}

func (userService *UserService) UpdateDepartments(tx *gorm.DB, id uint, departmentIds []uint) (err error) {
	TxErr := tx.Delete(&[]system.SysUserDepartment{}, "sys_user_id = ?", id).Error
	if TxErr != nil {
		return TxErr
	}
	userDepartments := []system.SysUserDepartment{}
	if len(departmentIds) > 0 {
		for _, v := range departmentIds {
			userDepartments = append(userDepartments, system.SysUserDepartment{
				id, v,
			})
		}
		TxErr = tx.Create(&userDepartments).Error
	}
	if TxErr != nil {
		return TxErr
	}
	return nil
}

func (userService *UserService) SetUserDepartments(id uint, departmentIds []uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		return userService.UpdateDepartments(tx, id, departmentIds)
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (userService *UserService) DeleteUser(id float64) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var user system.SysUser
		err := tx.Unscoped().Where("id = ?", id).First(&user).Error
		if err != nil {
			return err
		}
		err = tx.Unscoped().Delete(&user).Error
		//err = tx.Where("id = ?", id).Unscoped().Delete(&system.SysUser{}).Error
		if err != nil {
			return err
		}
		err = tx.Unscoped().Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
		if err != nil {
			return err
		}

		err = tx.Unscoped().Delete(&[]system.SysUserDepartment{}, "sys_user_id = ?", id).Error
		userService.DeleteCacheUser(user)
		return err
	})

	//var user system.SysUser
	//tx := global.GVA_DB.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	//err = tx.Where("id = ?", id).First(&user).Error
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//err = tx.Unscoped().Delete(&user).Error
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//err = tx.Unscoped().Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//err = tx.Unscoped().Delete(&[]system.SysUserDepartment{}, "sys_user_id = ?", id).Error
	//userService.DeleteCacheUser(user)
	//return tx.Commit().Error

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetUserInfo(reqUser system.SysUser) (err error, user system.SysUser) {
	err = global.GVA_DB.Updates(&reqUser).Error
	return err, reqUser
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (err error, user system.SysUser) {
	var reqUser system.SysUser
	err = global.GVA_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	return err, reqUser
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserById(id int) (err error, user *system.SysUser) {
	var u system.SysUser
	err = global.GVA_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserByUuid(uuid string) (err error, user *system.SysUser) {
	var u system.SysUser
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: resetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.MD5V([]byte("123456"))).Error
	return err
}

func (userService *UserService) UpdateBasicInfo(r systemReq.UpdateUserBasicInfo) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 加锁
		var user system.SysUser

		//err :=tx.Where("employee_id = ? and id != ? ", r.EmployeeID, r.ID).First(&user).Error
		if !errors.Is(tx.Where("employee_id = ? and id != ? ", r.EmployeeID, r.ID).First(&user).Error, gorm.ErrRecordNotFound) { // 判断员工号是否注册
			return errors.New("员工号已被占用")
		}

		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", r.ID).First(&user).Error
		if err != nil {
			return err
		}

		err = tx.Model(&user).Updates(r.ToSysUser()).Error
		if err != nil {
			return err
		}

		err = userService.UpAuthorities(tx, r.ID, r.AuthorityIds)
		if err != nil {
			return err
		}

		err = userService.UpdateDepartments(tx, r.ID, r.DepartmentIds)
		if err != nil {
			return err
		}

		userService.CacheSingleUserToRedis(user)

		return nil
	})
}

func (userService *UserService) DeleteCacheUser(reqUser system.SysUser) {
	global.GVA_REDIS.HDel(context.Background(), "users", reqUser.EmployeeID)
}

func (userService *UserService) CacheSingleUserToRedis(reqUser system.SysUser) {
	var u system.SysUser
	userMap := make(map[string]interface{})
	if err := global.GVA_DB.Where("`uuid` = ?", reqUser.UUID).Preload("Authority").Preload("Departments").Preload("Position").Preload("Authorities").First(&u).Error; err == nil {
		if data, err := json.Marshal(u); err == nil {
			userMap[u.EmployeeID] = data
		}
	}
	if len(userMap) > 0 {
		global.GVA_REDIS.HMSet(context.Background(), "users", userMap)
	}
}

func (userService *UserService) CacheUsersToRedis() {
	global.GVA_LOG.Info("start user sync task")
	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	db = db.Preload("Departments")
	db = db.Preload("Position")
	db = db.Preload("Authorities")
	db = db.Preload("Authority")
	err := db.Find(&userList).Error

	userMap := make(map[string]interface{})
	if err == nil {
		for _, user := range userList {
			if data, err := json.Marshal(user); err == nil {
				userMap[user.EmployeeID] = data
			}
		}
	}
	if len(userMap) > 0 {
		global.GVA_REDIS.HMSet(context.Background(), "users", userMap)
	}
}

//var tempAttendantSyncRestrictMap map[string]string
//
//func init()  {
//	tempAttendantSyncRestrictMap = map[string]string{
//		"00000085":"00000085", // 肖红梅
//		"00000054":"00000054", // 易香林
//		"":""
//
//	}
//}

// SyncUsersFromAttendantSystem 从考勤系统中同步用户
func (userService *UserService) SyncUsersFromAttendantSystem() {
	var Hes []attendant.HrEmployee
	err := global.GVA_ATTENDANT.Find(&Hes).Error
	if err != nil {

	}

	employeeMap := make(map[string]attendant.HrEmployee)

	for _, emp := range Hes {
		var user system.SysUser
		employeeMap[emp.EmplID] = emp
		err = global.GVA_DB.Model(&system.SysUser{}).Where("origin_type = ? and origin_code = ?", utils2.OriginTypeAttendance, emp.EmplID).First(&user).Error
		if err != nil {
			// 未找到,需要新建
			user = system.SysUser{}
			user.Password = utils.MD5V([]byte("123456"))
			user.UUID = uuid.NewV4()
			t := utils2.OriginTypeAttendance
			user.OriginType = &t
			user.OriginCode = emp.EmplID
			user.EmployeeID = fmt.Sprintf("10000%s", strings.TrimSpace(emp.CardID))
			if len(emp.EntryDate) > 0{
				t, err := time.Parse("2006-01-02", emp.EntryDate)
				if err == nil{
					user.CreatedAt = t
				}
			}
			user.NickName = emp.EmplName
			user.Username = user.EmployeeID
			// 默认正式工
			status := utils2.StaffTypeRegular
			user.StaffType = &status
			sType := utils2.StaffStatusEmployed
			user.StaffStatus = &sType
			if emp.Sex == "false" {
				user.Gender = 1
				user.HeaderImg = "http://10.0.0.33:9000/users/man.png"
			} else {
				user.Gender = 2
				user.HeaderImg = "http://10.0.0.33:9000/users/girl.png"
			}
			user.AuthorityId = "888"
			var authorities []system.SysAuthority
			authorities = append(authorities, system.SysAuthority{
				AuthorityId: user.AuthorityId,
			})
			user.Authorities = authorities
			err = global.GVA_DB.Create(&user).Error
			global.GVA_LOG.Info("sync", zap.String("用户", user.NickName))
		}
	}

	// 查找用户管理系统中的考勤系统中用户,
	var userList []system.SysUser

	err = global.GVA_DB.Model(&system.SysUser{}).Where("origin_type = 2").Find(&userList).Error
	if err == nil {
		for _, user := range userList {
			// 如果不在考勤系统中，则更新为离岗
			if emp, ok := employeeMap[user.OriginCode]; !ok {
				status := utils2.StaffStatusUnEmployed
				user.StaffStatus = &status
				global.GVA_DB.Updates(user)
			}else{
				// 如果在的话看是否有离职日期
				if len(emp.LeaveDate) > 0 {
					status := utils2.StaffStatusUnEmployed
					user.StaffStatus = &status
					global.GVA_DB.Updates(user)
				}
			}
		}
	}
}

func (userService *UserService) GetUserAttendant(query systemReq.UserAttendantQuery) (*response.UserAttendantResponse, error) {
	if query.Date == "" {
		return nil, errors.New("日期不能为空")
	}

	if query.Code == "" {
		return nil, errors.New("用户卡号不能为空")
	}

	var sysUser system.SysUser
	if err := global.GVA_DB.Model(&system.SysUser{}).Where("employee_id = ?", query.Code).First(&sysUser).Error; err != nil {
		return nil, err
	}

	if sysUser.OriginType == nil || *sysUser.OriginType != utils2.OriginTypeAttendance {
		return nil, errors.New("用户来源非erp系统")
	}

	// 判断考勤系统中是否有考勤记录
	var record attendant.HrAttendantRecord
	if err := global.GVA_ATTENDANT.Model(&attendant.HrAttendantRecord{}).Where("EmplID = ? and RecDate <= ?", sysUser.OriginCode, query.Date).First(&record).Error; err != nil {
		return nil, err
	}

	var records []attendant.HrAttendantRecord
	db := global.GVA_ATTENDANT.Model(&attendant.HrAttendantRecord{})
	err := db.Where("EmplID = ? and RecDate = ?", sysUser.OriginCode, query.Date).Find(&records).Error
	if err != nil {
		return nil, err
	}
	reLen := len(records)
	userAttendant := response.UserAttendantResponse{
		Code: query.Code,
		Date: query.Date,
	}
	if reLen > 0 {
		userAttendant.ClockIn = fmt.Sprintf("%s %s", records[0].RecDate, records[0].RecTime)
		if len(records) > 1 {

			userAttendant.ClockOut = fmt.Sprintf("%s %s", records[reLen-1].RecDate, records[reLen-1].RecTime)
		}
	} else {
		return nil, errors.New("无记录")
	}
	return &userAttendant, nil
}

func (userService *UserService) GetUserAttendantLastSyncTime() (string, error) {
	var record attendant.HrAttendantRecord
	err := global.GVA_ATTENDANT.Raw("select top 1 * FROM AtdRecord order by OperDate desc").Scan(&record).Error
	if err != nil {
		global.GVA_LOG.Error("GetUserAttendantLastSyncTime", zap.String("err", err.Error()))
		return "", err
	}
	return record.OperDate, nil
}

//func (userService *UserService) GetUserAttendantLastSyncTime() (string, error){
//	var record attendant.HrAttendantRecord
//	err := global.GVA_ATTENDANT.Raw("select * FROM AtdRecord order by OperDate desc").Scan(&record).Error
//	if err!=nil{
//		global.GVA_LOG.Error("GetUserAttendantLastSyncTime", zap.String("err", err.Error()))
//		return "", err
//	}
//	return record.OperDate, nil
//}
