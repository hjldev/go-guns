package system

import (
	"fmt"
	"go-guns/app/model"
	"go-guns/boot"
)

func InsertRoleDept(roleId int, deptIds []int) (bool, error) {
	//ORM不支持批量插入所以需要拼接 sql 串
	sql := "INSERT INTO `sys_role_dept` (`role_id`,`dept_id`) VALUES "

	for i := 0; i < len(deptIds); i++ {
		if len(deptIds)-1 == i {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,%d);", roleId, deptIds[i])
		} else {
			sql += fmt.Sprintf("(%d,%d),", roleId, deptIds[i])
		}
	}
	boot.Db.Exec(sql)

	return true, nil
}

func DeleteRoleDept(roleId int) (bool, error) {
	if err := boot.Db.Table("sys_role_dept").Where("role_id = ?", roleId).Delete(&model.SysRoleDept{}).Error; err != nil {
		return false, err
	}
	var role model.SysRole
	if err := boot.Db.First(&role, roleId).Error; err != nil {
		return false, err
	}

	return true, nil

}
