package system

import (
	"fmt"
	"go-guns/app/model"
	"go-guns/boot"
)

func GetRoleMenu(rm model.SysRoleMenu) ([]model.SysRoleMenu, error) {
	var r []model.SysRoleMenu
	table := boot.Db.Table("sys_role_menu")
	if rm.RoleId != 0 {
		table = table.Where("role_id = ?", rm.RoleId)

	}
	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func GetPermisRoleMenu(rm model.SysRoleMenu) ([]string, error) {
	var r []model.SysMenu
	table := boot.Db.Select("sys_menu.permission").Table("sys_menu").Joins("left join sys_role_menu on sys_menu.menu_id = sys_role_menu.menu_id")

	table = table.Where("role_id = ?", rm.RoleId)

	table = table.Where("sys_menu.menu_type in('F','C')")
	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	var list []string
	for i := 0; i < len(r); i++ {
		list = append(list, r[i].Permission)
	}
	return list, nil
}

func GetRoleMenuIds(rm model.SysRoleMenu) ([]model.MenuPath, error) {
	var r []model.MenuPath
	table := boot.Db.Select("sys_menu.path").Table("sys_role_menu")
	table = table.Joins("left join sys_role on sys_role.role_id=sys_role_menu.role_id")
	table = table.Joins("left join sys_menu on sys_menu.id=sys_role_menu.menu_id")
	table = table.Where("sys_role.role_name = ? and sys_menu.type=1", rm.RoleName)
	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func DeleteRoleMenu(roleId int) (bool, error) {
	tx := boot.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}

	if err := tx.Table("sys_role_dept").Where("role_id = ?", roleId).Delete(&model.SysRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	if err := tx.Table("sys_role_menu").Where("role_id = ?", roleId).Delete(&model.SysRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	var role model.SysRole
	if err := tx.Table("sys_role").Where("role_id = ?", roleId).First(&role).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	sql3 := "delete from casbin_rule where v0= '" + role.RoleKey + "';"
	if err := tx.Exec(sql3).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	if err := tx.Commit().Error; err != nil {
		return false, err
	}

	return true, nil

}

func BatchDeleteRoleMenu(roleIds []int) (bool, error) {
	tx := boot.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}

	if err := tx.Table("sys_role_menu").Where("role_id in (?)", roleIds).Delete(&model.SysRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	var role []model.SysRole
	if err := tx.Table("sys_role").Where("role_id in (?)", roleIds).Find(&role).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	sql := ""
	for i := 0; i < len(role); i++ {
		sql += "delete from casbin_rule where v0= '" + role[i].RoleName + "';"
	}
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	if err := tx.Commit().Error; err != nil {
		return false, err
	}
	return true, nil

}

func InsertRoleMenu(rm model.SysRoleMenu, roleId int, menuId []int) (bool, error) {
	var role model.SysRole
	tx := boot.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}
	if err := tx.Table("sys_role").Where("role_id = ?", roleId).First(&role).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	var menu []model.SysMenu
	if err := tx.Table("sys_menu").Where("menu_id in (?)", menuId).Find(&menu).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	//ORM不支持批量插入所以需要拼接 sql 串
	sql := "INSERT INTO `sys_role_menu` (`role_id`,`menu_id`,`role_name`) VALUES "

	sql2 := "INSERT INTO casbin_rule  (`p_type`,`v0`,`v1`,`v2`) VALUES "
	for i := 0; i < len(menu); i++ {
		if len(menu)-1 == i {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,%d,'%s');", role.RoleId, menu[i].MenuId, role.RoleKey)
			if menu[i].MenuType == "A" {
				sql2 += fmt.Sprintf("('p','%s','%s','%s');", role.RoleKey, menu[i].Path, menu[i].Action)
			}
		} else {
			sql += fmt.Sprintf("(%d,%d,'%s'),", role.RoleId, menu[i].MenuId, role.RoleKey)
			if menu[i].MenuType == "A" {
				sql2 += fmt.Sprintf("('p','%s','%s','%s'),", role.RoleKey, menu[i].Path, menu[i].Action)
			}
		}
	}
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	sql2 = sql2[0:len(sql2)-1] + ";"
	if err := tx.Exec(sql2).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	if err := tx.Commit().Error; err != nil {
		return false, err
	}
	return true, nil
}

func DeleteRoleMenuById(RoleId string, MenuID string) (bool, error) {
	table := boot.Db.Model(model.SysRoleMenu{}).Where("role_id = ?", RoleId)
	if MenuID != "" {
		table = table.Where("menu_id = ?", MenuID)
	}
	if err := table.Delete(&model.SysRoleMenu{}).Error; err != nil {
		return false, err
	}
	return true, nil

}
