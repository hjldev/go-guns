package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-guns/model"
	"go-guns/service/system"
	"go-guns/tools"
	"net/http"
	"strconv"
)

// @Summary 分页部门列表数据
// @Description 分页列表
// @Tags 部门
// @Param name query string false "name"
// @Param id query string false "id"
// @Param position query string false "position"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/deptList [get]
// @Security Bearer
func GetDeptList(c *gin.Context) {
	var dept model.SysDept
	dept.DeptName = c.Request.FormValue("deptName")
	dept.Status = c.Request.FormValue("status")
	dept.DeptId, _ = strconv.Atoi(c.Request.FormValue("deptId"))
	result, err := system.SetDept(dept)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	tools.R(c, result)
}

func GetDeptTree(c *gin.Context) {
	var dept model.SysDept
	dept.DeptName = c.Request.FormValue("deptName")
	dept.Status = c.Request.FormValue("status")
	dept.DeptId, _ = strconv.Atoi(c.Request.FormValue("deptId"))
	result, err := system.SetDept(dept)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	tools.R(c, result)
}

// @Summary 部门列表数据
// @Description 获取JSON
// @Tags 部门
// @Param deptId path string false "deptId"
// @Param position query string false "position"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept/{deptId} [get]
// @Security Bearer
func GetDept(c *gin.Context) {
	var dept model.SysDept
	dept.DeptId, _ = strconv.Atoi(c.Param("deptId"))
	dept.DataScope = tools.GetUserIdStr(c)
	result, err := system.GetDept(dept)
	tools.HasError(err, 404, "page not fount")
	tools.R(c, result)
}

// @Summary 添加部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param data body system.SysDept true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dept [post]
// @Security Bearer
func InsertDept(c *gin.Context) {
	var data model.SysDept
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "", 500)
	data.CreateBy = tools.GetUserIdStr(c)
	err = system.CreateDept(data)
	tools.HasError(err)
	tools.R(c, data)
}

// @Summary 修改部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body system.SysDept true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dept [put]
// @Security Bearer
func UpdateDept(c *gin.Context) {
	var data model.SysDept
	err := c.BindJSON(&data)
	tools.HasError(err, "", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	err = system.UpdateDept(data)
	tools.HasError(err)
	tools.R(c, data)
}

// @Summary 删除部门
// @Description 删除数据
// @Tags 部门
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dept/{id} [delete]
func DeleteDept(c *gin.Context) {
	var data model.SysDept
	id, err := strconv.Atoi(c.Param("id"))
	data.DeptId = id
	_, err = system.DeleteDept(data)
	tools.HasError(err, 500, "删除失败")
	tools.R(c)
}

func GetDeptTreeRoleselect(c *gin.Context) {
	var Dept model.SysDept
	var SysRole model.SysRole
	id, err := strconv.Atoi(c.Param("roleId"))
	SysRole.RoleId = id
	result, err := system.SetDeptLabel(Dept)
	tools.HasError(err)
	menuIds := make([]int, 0)
	if id != 0 {
		menuIds, err = system.GetRoleDeptId(SysRole)
		tools.HasError(err, "抱歉未找到相关信息", -1)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        200,
		"depts":       result,
		"checkedKeys": menuIds,
	})
}
