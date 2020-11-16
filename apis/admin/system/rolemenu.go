package system

import (
	"github.com/gin-gonic/gin"
	"go-guns/model"
	"go-guns/service/system"
	"go-guns/tools"
	"net/http"
)

// @Summary RoleMenu列表数据
// @Description 获取JSON
// @Tags 角色菜单
// @Param RoleId query string false "RoleId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/rolemenu [get]
// @Security Bearer
func GetRoleMenu(c *gin.Context) {
	var Rm model.SysRoleMenu
	err := c.ShouldBind(&Rm)
	result, err := system.GetRoleMenu(Rm)
	if err != nil {
		msg := "抱歉未找到相关信息"
		tools.R(c, nil, http.StatusOK, msg)
		return
	}
	tools.R(c, result)
}

type RoleMenuPost struct {
	RoleId   string
	RoleMenu []model.SysRoleMenu
}

// @Summary 删除用户菜单数据
// @Description 删除数据
// @Tags 角色菜单
// @Param id path string true "id"
// @Param menu_id query string false "menu_id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/rolemenu/{id} [delete]
func DeleteRoleMenu(c *gin.Context) {
	roleId := c.Param("id")
	menuId := c.Request.FormValue("menu_id")
	_, err := system.DeleteRoleMenuById(roleId, menuId)
	if err != nil {
		tools.R(c, nil, http.StatusOK, "删除失败")
	}
	tools.R(c)
}
