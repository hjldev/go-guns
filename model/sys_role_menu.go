package model

var SysRoleMenuTableName = "sys_role_menu"

type SysRoleMenu struct {
	RoleId   int    `gorm:""`
	MenuId   int    `gorm:""`
	RoleName string `gorm:"size:128)"`
	CreateBy string `gorm:"size:128)"`
	UpdateBy string `gorm:"size:128)"`
}

type MenuPath struct {
	Path string `json:"path"`
}
