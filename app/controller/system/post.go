package system

import (
	"github.com/gin-gonic/gin"
	"go-guns/app/model"
	"go-guns/service/system"
	"go-guns/tools"
	"strconv"
)

// @Summary 岗位列表数据
// @Description 获取JSON
// @Tags 岗位
// @Param postName query string false "postName"
// @Param postCode query string false "postCode"
// @Param postId query string false "postId"
// @Param status query string false "status"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [get]
// @Security Bearer
func GetPostList(c *gin.Context) {
	var data model.SysPost
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, err = strconv.Atoi(size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, err = strconv.Atoi(index)
	}

	id := c.Request.FormValue("postId")
	data.PostId, err = strconv.Atoi(id)

	data.PostCode = c.Request.FormValue("postCode")
	data.PostName = c.Request.FormValue("postName")
	data.Status = c.Request.FormValue("status")

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := system.GetPostPage(data, pageSize, pageIndex)
	tools.HasError(err, "", -1)
	pageData := tools.PageData{
		List:      result,
		Count:     count,
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}
	tools.PageR(c, pageData)
}

// @Summary 获取岗位信息
// @Description 获取JSON
// @Tags 岗位
// @Param postId path int true "postId"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post/{postId} [get]
// @Security Bearer
func GetPost(c *gin.Context) {
	var Post model.SysPost
	Post.PostId, _ = strconv.Atoi(c.Param("postId"))
	result, err := system.GetPost(Post)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	tools.R(c, result)
}

// @Summary 添加岗位
// @Description 获取JSON
// @Tags 岗位
// @Accept  application/json
// @Product application/json
// @Param data body models.Post true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/post [post]
// @Security Bearer
func InsertPost(c *gin.Context) {
	var data model.SysPost
	err := c.Bind(&data)
	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	err = system.CreatePost(data)
	tools.HasError(err, "", -1)
	tools.R(c, data)
}

// @Summary 修改岗位
// @Description 获取JSON
// @Tags 岗位
// @Accept  application/json
// @Product application/json
// @Param data body models.Post true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/post/ [put]
// @Security Bearer
func UpdatePost(c *gin.Context) {
	var data model.SysPost

	err := c.Bind(&data)
	data.UpdateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", -1)
	err = system.UpdatePost(data)
	tools.HasError(err, "", -1)
	tools.R(c, data)
}

// @Summary 删除岗位
// @Description 删除数据
// @Tags 岗位
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 500 {string} string	"{"code": 500, "message": "删除失败"}"
// @Router /api/v1/post/{postId} [delete]
func DeletePost(c *gin.Context) {
	var data model.SysPost
	data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("postId", c)
	result, err := system.BatchDeletePost(IDS)
	tools.HasError(err, "删除失败", 500)
	tools.R(c, result)
}
