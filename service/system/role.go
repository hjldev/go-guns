package system

import (
	"errors"
	"go-guns/database"
	"go-guns/model"
)

func GetRolePage(role model.SysRole, pageSize int, pageIndex int) ([]model.SysRole, int, error) {
	var doc []model.SysRole

	table := database.Db
	if role.RoleId != 0 {
		table = table.Where("role_id = ?", role.RoleId)
	}
	if role.RoleName != "" {
		table = table.Where("role_name = ?", role.RoleName)
	}
	if role.Status != "" {
		table = table.Where("status = ?", role.Status)
	}
	if role.RoleKey != "" {
		table = table.Where("role_key = ?", role.RoleKey)
	}

	var count int64

	if err := table.Order("role_sort").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, int(count), nil
}

func GetRole(role model.SysRole) (SysRole model.SysRole, err error) {
	table := database.Db
	if role.RoleId != 0 {
		table = table.Where("role_id = ?", role.RoleId)
	}
	if role.RoleName != "" {
		table = table.Where("role_name = ?", role.RoleName)
	}
	if err = table.First(&SysRole).Error; err != nil {
		return
	}

	return
}

func GetRoleList(role model.SysRole) (SysRole []model.SysRole, err error) {
	table := database.Db
	if role.RoleId != 0 {
		table = table.Where("role_id = ?", role.RoleId)
	}
	if role.RoleName != "" {
		table = table.Where("role_name = ?", role.RoleName)
	}
	if err = table.Order("role_sort").Find(&SysRole).Error; err != nil {
		return
	}

	return
}

// 获取角色对应的菜单ids
func GetRoleMeunId(role model.SysRole) ([]int, error) {
	menuIds := make([]int, 0)
	menuList := make([]model.SysMenuIdList, 0)
	if err := database.Db.Table("sys_role_menu").Select("sys_role_menu.menu_id").Where("role_id = ? ", role.RoleId).Where(" sys_role_menu.menu_id not in(select sys_menu.parent_id from sys_role_menu LEFT JOIN sys_menu on sys_menu.menu_id=sys_role_menu.menu_id where role_id =?  and parent_id is not null)", role.RoleId).Find(&menuList).Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(menuList); i++ {
		menuIds = append(menuIds, menuList[i].MenuId)
	}
	return menuIds, nil
}

func InsertRole(role model.SysRole) (id int, err error) {
	var i int64 = 0
	database.Db.Model(&role).Where("role_name=? or role_key = ?", role.RoleName, role.RoleKey).Count(&i)
	if i > 0 {
		return 0, errors.New("角色名称或者角色标识已经存在！")
	}
	role.UpdateBy = ""
	result := database.Db.Create(&role)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = role.RoleId
	return
}

type DeptIdList struct {
	DeptId int `json:"DeptId"`
}

func GetRoleDeptId(role model.SysRole) ([]int, error) {
	deptIds := make([]int, 0)
	deptList := make([]DeptIdList, 0)
	if err := database.Db.Table("sys_role_dept").Select("sys_role_dept.dept_id").Joins("LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id").Where("role_id = ? ", role.RoleId).Where(" sys_role_dept.dept_id not in(select sys_dept.parent_id from sys_role_dept LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id where role_id =? )", role.RoleId).Find(&deptList).Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(deptList); i++ {
		deptIds = append(deptIds, deptList[i].DeptId)
	}

	return deptIds, nil
}

//修改
func UpdateRole(role model.SysRole) (err error) {
	var r model.SysRole
	if err = database.Db.First(&r, role.RoleId).Error; err != nil {
		return
	}

	if role.RoleName != "" && role.RoleName != r.RoleName {
		return errors.New("角色名称不允许修改！")
	}

	if role.RoleKey != "" && role.RoleKey != r.RoleKey {
		return errors.New("角色标识不允许修改！")
	}

	if err = database.Db.Model(&role).Updates(&role).Error; err != nil {
		return
	}
	return
}

//批量删除
func BatchDeleteRole(id []int) (Result bool, err error) {
	if err = database.Db.Delete(&model.SysRole{}, id).Error; err != nil {
		return
	}
	Result = true
	return
}
