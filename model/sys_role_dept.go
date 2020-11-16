package model

var SysRoleDeptTableName = "sys_role_dept"

type SysRoleDept struct {
	RoleId int `gorm:""`
	DeptId int `gorm:""`
}
