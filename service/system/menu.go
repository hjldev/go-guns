package system

import (
	"errors"
	"go-guns/database"
	"go-guns/model"
	"strconv"
)

func GetMenu(id int) (Menu model.SysMenu, err error) {

	if err = database.Db.First(&Menu, id).Error; err != nil {
		return
	}
	return
}

func SetMenu(e model.SysMenu) (m []model.SysMenu, err error) {
	menulist, err := GetMenuPage(e)

	m = make([]model.SysMenu, 0)
	for i := 0; i < len(menulist); i++ {
		if menulist[i].ParentId != 0 {
			continue
		}
		menusInfo := DiguiMenu(&menulist, menulist[i])

		m = append(m, menusInfo)
	}
	return
}

func DiguiMenu(menulist *[]model.SysMenu, menu model.SysMenu) model.SysMenu {
	list := *menulist

	min := make([]model.SysMenu, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != list[j].ParentId {
			continue
		}
		mi := model.SysMenu{}
		mi.MenuId = list[j].MenuId
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.Action = list[j].Action
		mi.Permission = list[j].Permission
		mi.ParentId = list[j].ParentId
		mi.NoCache = list[j].NoCache
		mi.Breadcrumb = list[j].Breadcrumb
		mi.Component = list[j].Component
		mi.Sort = list[j].Sort
		mi.Visible = list[j].Visible
		mi.CreatedAt = list[j].CreatedAt
		mi.Children = []model.SysMenu{}

		if mi.MenuType != "F" {
			ms := DiguiMenu(menulist, mi)
			min = append(min, ms)

		} else {
			min = append(min, mi)
		}

	}
	menu.Children = min
	return menu
}

func SetMenuLabel(e model.SysMenu) (m []model.SysMenuLabel, err error) {
	menulist, err := GetMenuList(e)

	m = make([]model.SysMenuLabel, 0)
	for i := 0; i < len(menulist); i++ {
		if menulist[i].ParentId != 0 {
			continue
		}
		e := model.SysMenuLabel{}
		e.Id = menulist[i].MenuId
		e.Label = menulist[i].Title
		menusInfo := DiguiMenuLabel(&menulist, e)

		m = append(m, menusInfo)
	}
	return
}

func DiguiMenuLabel(menulist *[]model.SysMenu, menu model.SysMenuLabel) model.SysMenuLabel {
	list := *menulist

	min := make([]model.SysMenuLabel, 0)
	for j := 0; j < len(list); j++ {

		if menu.Id != list[j].ParentId {
			continue
		}
		mi := model.SysMenuLabel{}
		mi.Id = list[j].MenuId
		mi.Label = list[j].Title
		mi.Children = []model.SysMenuLabel{}
		if list[j].MenuType != "F" {
			ms := DiguiMenuLabel(menulist, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}

	}
	menu.Children = min
	return menu
}

func SetMenuRole(rolename string) (m []model.SysMenu, err error) {

	menulist, err := GetMenuByRoleName(rolename)

	m = make([]model.SysMenu, 0)
	for i := 0; i < len(menulist); i++ {
		if menulist[i].ParentId != 0 {
			continue
		}
		menusInfo := DiguiMenu(&menulist, menulist[i])

		m = append(m, menusInfo)
	}
	return
}

func GetMenuByName(name string) (Menus []model.SysMenu, err error) {
	table := database.Db.Table(model.SysMenuTableName)
	if name != "" {
		table = table.Where("menu_name = ?", name)
	}
	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}
	return
}

func GetMenuByRoleName(rolename string) (Menus []model.SysMenu, err error) {
	table := database.Db.Table(model.SysMenuTableName).Select("sys_menu.*").Joins("left join sys_role_menu on sys_role_menu.menu_id=sys_menu.menu_id")
	table = table.Where("sys_role_menu.role_name=? and menu_type in ('M','C')", rolename)
	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}
	return
}

func GetMenuList(e model.SysMenu) (Menus []model.SysMenu, err error) {
	table := database.Db.Table(model.SysMenuTableName)
	if e.MenuName != "" {
		table = table.Where("menu_name = ?", e.MenuName)
	}
	if e.Path != "" {
		table = table.Where("path = ?", e.Path)
	}
	if e.Action != "" {
		table = table.Where("action = ?", e.Action)
	}
	if e.MenuType != "" {
		table = table.Where("menu_type = ?", e.MenuType)
	}

	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}
	return
}

func GetMenuPage(e model.SysMenu) (Menus []model.SysMenu, err error) {
	table := database.Db.Table(model.SysMenuTableName)
	if e.MenuName != "" {
		table = table.Where("menu_name = ?", e.MenuName)
	}
	if e.Title != "" {
		table = table.Where("title = ?", e.Title)
	}
	if e.Visible != "" {
		table = table.Where("visible = ?", e.Visible)
	}
	if e.MenuType != "" {
		table = table.Where("menu_type = ?", e.MenuType)
	}

	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}
	return
}

func CreateMenu(e model.SysMenu) (id int, err error) {
	result := database.Db.Create(&e)
	if result.Error != nil {
		err = result.Error
		return
	}
	err = InitPaths(e)
	if err != nil {
		return
	}
	id = e.MenuId
	return
}

func InitPaths(menu model.SysMenu) (err error) {
	parentMenu := model.SysMenu{}
	if menu.ParentId != 0 {
		database.Db.Where("menu_id = ?", menu.ParentId).First(&parentMenu)
		if parentMenu.Paths == "" {
			err = errors.New("父级paths异常，请尝试对当前节点父级菜单进行更新操作！")
			return
		}
		menu.Paths = parentMenu.Paths + "/" + strconv.Itoa(menu.MenuId)
	} else {
		menu.Paths = "/0/" + strconv.Itoa(menu.MenuId)
	}
	database.Db.Model(&menu).Where("menu_id = ?", menu.MenuId).Update("paths", menu.Paths)
	return
}

func UpdateMenu(e model.SysMenu) (err error) {
	if err = database.Db.Model(&e).Updates(&e).Error; err != nil {
		return
	}
	err = InitPaths(e)
	if err != nil {
		return
	}
	return
}

func DeleteMenu(id int) (success bool, err error) {
	if err = database.Db.Delete(&model.SysMenu{}, id).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}
