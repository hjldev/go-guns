package system

import (
	"errors"
	"go-guns/database"
	"go-guns/model"
	"go-guns/tools"
	"strconv"
)

func GetUserById(id int) model.SysUser {
	user := model.SysUser{}
	database.Db.First(&user, id)
	return user
}

func GetUserInfo(e model.SysUser) (SysUserView model.SysUserView, err error) {

	table := database.Db.Table(model.SysUserTableName).Select([]string{"sys_user.*", "sys_role.role_name"})
	table = table.Joins("left join sys_role on sys_user.role_id=sys_role.role_id")
	if e.UserId != 0 {
		table = table.Where("user_id = ?", e.UserId)
	}

	if e.Username != "" {
		table = table.Where("username = ?", e.Username)
	}

	if e.Password != "" {
		table = table.Where("password = ?", e.Password)
	}

	if e.RoleId != 0 {
		table = table.Where("role_id = ?", e.RoleId)
	}

	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}

	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}

	if err = table.First(&SysUserView).Error; err != nil {
		return
	}
	return
}

func GetUserList(e model.SysUser) (SysUserView []model.SysUserView, err error) {

	table := database.Db.Table(e.TableName()).Select([]string{"sys_user.*", "sys_role.role_name"})
	table = table.Joins("left join sys_role on sys_user.role_id=sys_role.role_id")
	if e.UserId != 0 {
		table = table.Where("user_id = ?", e.UserId)
	}

	if e.Username != "" {
		table = table.Where("username = ?", e.Username)
	}

	if e.Password != "" {
		table = table.Where("password = ?", e.Password)
	}

	if e.RoleId != 0 {
		table = table.Where("role_id = ?", e.RoleId)
	}

	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}

	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}

	if err = table.Find(&SysUserView).Error; err != nil {
		return
	}
	return
}

func GetUserPage(e model.SysUser, pageSize int, pageIndex int) ([]model.SysUserPage, int, error) {
	var doc []model.SysUserPage
	table := database.Db.Select("sys_user.*,sys_dept.dept_name").Table(e.TableName())
	table = table.Joins("left join sys_dept on sys_dept.dept_id = sys_user.dept_id")

	if e.Username != "" {
		table = table.Where("username = ?", e.Username)
	}
	if e.Status != "" {
		table = table.Where("sys_user.status = ?", e.Status)
	}

	if e.Phone != "" {
		table = table.Where("sys_user.phone = ?", e.Phone)
	}

	if e.DeptId != 0 {
		table = table.Where("sys_user.dept_id in (select dept_id from sys_dept where dept_path like ? )", "%"+strconv.Itoa(e.DeptId)+"%")
	}

	var count int64

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("sys_user.deleted_at IS NULL").Count(&count)
	return doc, int(count), nil
}

//添加
func InsertUser(e model.SysUser) (id int, err error) {
	pwd, err := tools.Encrypt(e.Password)
	if err != nil {
		return
	}
	e.Password = pwd
	var count int64
	database.Db.Model(&model.SysUser{}).Where("username = ?", e.Username).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}

	//添加数据
	if err = database.Db.Create(&e).Error; err != nil {
		return
	}
	id = e.UserId
	return
}

//修改
func UpdateUser(update model.SysUser) (err error) {
	if update.Password != "" {
		update.Password, err = tools.Encrypt(update.Password)
		if err != nil {
			return err
		}
	}
	if err = database.Db.First(&update, update.UserId).Error; err != nil {
		return
	}

	if err = database.Db.Updates(&update).Error; err != nil {
		return
	}
	return
}

func BatchDeleteUser(id []int) (Result bool, err error) {
	if err = database.Db.Delete(model.SysUser{}, id).Error; err != nil {
		return
	}
	Result = true
	return
}

// todo 修改密码
func SetUserPwd(e model.SysUser) (Result bool, err error) {
	return true, nil
}
