package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-guns/app/model"
	"go-guns/service/system"
	"go-guns/tools"
	"net/http"
	"strconv"
)

// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/menulist [get]
// @Security Bearer
func GetMenuList(c *gin.Context) {
	var Menu model.SysMenu
	Menu.MenuName = c.Request.FormValue("menuName")
	Menu.Visible = c.Request.FormValue("visible")
	Menu.Title = c.Request.FormValue("title")
	Menu.DataScope = tools.GetUserIdStr(c)
	result, err := system.SetMenu(Menu)
	tools.HasError(err, "抱歉未找到相关信息", -1)

	tools.R(c, result)
}

// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/menu [get]
// @Security Bearer
func GetMenu(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	result, err := system.GetMenu(id)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	tools.R(c, result)
}

func GetMenuTreeRoleselect(c *gin.Context) {
	var Menu model.SysMenu
	var SysRole model.SysRole
	id, err := strconv.Atoi(c.Param("roleId"))
	SysRole.RoleId = id
	result, err := system.SetMenuLabel(Menu)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	menuIds := make([]int, 0)
	if id != 0 {
		menuIds, err = system.GetRoleMeunId(SysRole)
		tools.HasError(err, "抱歉未找到相关信息", -1)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      200,
		"menus":       result,
		"checkedKeys": menuIds,
	})
}

// @Summary 获取菜单树
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/menuTreeselect [get]
// @Security Bearer
func GetMenuTreeelect(c *gin.Context) {
	var data model.SysMenu
	result, err := system.SetMenuLabel(data)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	tools.R(c, result)
}

// @Summary 创建菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Param menuName formData string true "menuName"
// @Param Path formData string false "Path"
// @Param Action formData string true "Action"
// @Param Permission formData string true "Permission"
// @Param ParentId formData string true "ParentId"
// @Param IsDel formData string true "IsDel"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/menu [post]
// @Security Bearer
func InsertMenu(c *gin.Context) {
	var data model.SysMenu
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	data.CreateBy = tools.GetUserIdStr(c)
	result, err := system.CreateMenu(data)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	tools.R(c, result)
}

// @Summary 修改菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Param id path int true "id"
// @Param data body models.Menu true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "修改失败"}"
// @Router /api/v1/menu/{id} [put]
// @Security Bearer
func UpdateMenu(c *gin.Context) {
	var data model.SysMenu
	err2 := c.BindWith(&data, binding.JSON)
	data.UpdateBy = tools.GetUserIdStr(c)
	tools.HasError(err2, "修改失败", -1)
	err := system.UpdateMenu(data)
	tools.HasError(err, "", 501)
	tools.R(c)

}

// @Summary 删除菜单
// @Description 删除数据
// @Tags 菜单
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/menu/{id} [delete]
func DeleteMenu(c *gin.Context) {
	var data model.SysMenu
	id, err := strconv.Atoi(c.Param("id"))
	data.UpdateBy = tools.GetUserIdStr(c)
	_, err = system.DeleteMenu(id)
	tools.HasError(err, "删除失败", 500)
	tools.R(c)
}

// @Summary 根据角色名称获取菜单列表数据（左菜单使用）
// @Description 获取JSON
// @Tags 菜单
// @Param id path int true "id"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/menurole [get]
// @Security Bearer
func GetMenuRole(c *gin.Context) {
	user := tools.GetUserFromClaims(c)
	result, err := system.SetMenuRole(user.Role)
	tools.HasError(err, "获取失败", 500)
	tools.R(c, result)
}

// @Summary 获取角色对应的菜单id数组
// @Description 获取JSON
// @Tags 菜单
// @Param id path int true "id"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/menuids/{id} [get]
// @Security Bearer
func GetMenuIDS(c *gin.Context) {
	var data model.SysRoleMenu
	data.RoleName = c.GetString("role")
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := system.GetRoleMenuIds(data)
	tools.HasError(err, "获取失败", 500)
	tools.R(c, result)
}
