package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"go-guns/app/model"
	"go-guns/service/system"
	"go-guns/tools"
	"net/http"
	"strconv"
)

// @Summary 列表用户信息数据
// @Description 获取JSON
// @Tags 用户
// @Param username query string false "username"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/sysUserList [get]
// @Security Bearer
func GetSysUserList(c *gin.Context) {
	var data model.SysUser
	var err error
	var pageSize = 10
	var pageIndex = 1

	size := c.Request.FormValue("pageSize")
	if size != "" {
		pageSize, err = strconv.Atoi(size)
	}

	index := c.Request.FormValue("pageIndex")
	if index != "" {
		pageIndex, err = strconv.Atoi(index)
	}

	c.Bind(&data)

	data.Username = c.Request.FormValue("username")
	data.Status = c.Request.FormValue("status")
	data.Phone = c.Request.FormValue("phone")

	postId := c.Request.FormValue("postId")
	data.PostId, _ = strconv.Atoi(postId)

	deptId := c.Request.FormValue("deptId")
	data.DeptId, _ = strconv.Atoi(deptId)

	userId := tools.GetUserId(c)
	data.DataScope = strconv.Itoa(userId)
	result, count, err := system.GetUserPage(data, pageSize, pageIndex)
	tools.HasError(err, "", -1)

	pageData := tools.PageData{
		List:      result,
		Count:     count,
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}

	tools.R(c, pageData)
}

// @Summary 获取用户
// @Description 获取JSON
// @Tags 用户
// @Param userId path int true "用户编码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sysUser/{userId} [get]
// @Security Bearer
func GetSysUser(c *gin.Context) {
	var sysUser model.SysUser
	sysUser.UserId, _ = strconv.Atoi(c.Param("userId"))
	result, err := system.GetUserInfo(sysUser)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	var sysRole model.SysRole
	var post model.SysPost
	roles, err := system.GetRoleList(sysRole)
	posts, err := system.GetPostList(post)

	postIds := make([]int, 0)
	postIds = append(postIds, result.PostId)

	roleIds := make([]int, 0)
	roleIds = append(roleIds, result.RoleId)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    result,
		"postIds": postIds,
		"roleIds": roleIds,
		"roles":   roles,
		"posts":   posts,
	})
}

// @Summary 获取个人中心用户
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/profile [get]
// @Security Bearer
func GetSysUserProfile(c *gin.Context) {
	var sysUser model.SysUser
	userId := tools.GetUserId(c)
	sysUser.UserId = userId
	result, err := system.GetUserInfo(sysUser)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	var sysRole model.SysRole
	var post model.SysPost
	var dept model.SysDept
	//获取角色列表
	roles, err := system.GetRoleList(sysRole)
	//获取职位列表
	posts, err := system.GetPostList(post)
	//获取部门列表
	dept.DeptId = result.DeptId
	dept, err = system.GetDept(dept)

	postIds := make([]int, 0)
	postIds = append(postIds, result.PostId)

	roleIds := make([]int, 0)
	roleIds = append(roleIds, result.RoleId)

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    result,
		"postIds": postIds,
		"roleIds": roleIds,
		"roles":   roles,
		"posts":   posts,
		"dept":    dept,
	})

}

// @Summary 获取用户角色和职位
// @Description 获取JSON
// @Tags 用户
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sysUser [get]
// @Security Bearer
func GetSysUserInit(c *gin.Context) {
	var sysRole model.SysRole
	var post model.SysPost
	roles, err := system.GetRoleList(sysRole)
	posts, err := system.GetPostList(post)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	mp := make(map[string]interface{}, 2)
	mp["roles"] = roles
	mp["posts"] = posts
	tools.R(c, mp)
}

// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body models.SysUser true "用户数据"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/sysUser [post]
func InsertSysUser(c *gin.Context) {
	var sysuser model.SysUser
	err := c.BindWith(&sysuser, binding.JSON)
	tools.HasError(err, "非法数据格式", 500)

	sysuser.CreateBy = tools.GetUserIdStr(c)
	id, err := system.InsertUser(sysuser)
	tools.HasError(err, "添加失败", 500)
	tools.R(c, id)
}

// @Summary 修改用户数据
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body models.SysUser true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "修改失败"}"
// @Router /api/v1/sysuser/{userId} [put]
func UpdateSysUser(c *gin.Context) {
	var data model.SysUser
	err := c.Bind(&data)
	tools.HasError(err, "数据解析失败", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	err = system.UpdateUser(data)
	tools.HasError(err, "修改失败", 500)
	tools.R(c, data)
}

// @Summary 删除用户数据
// @Description 删除数据
// @Tags 用户
// @Param userId path int true "userId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/sysuser/{userId} [delete]
func DeleteSysUser(c *gin.Context) {
	var data model.SysUser
	data.UpdateBy = tools.GetUserIdStr(c)
	ids := tools.IdsStrToIdsIntGroup("userId", c)
	result, err := system.BatchDeleteUser(ids)
	tools.HasError(err, "删除失败", 500)
	tools.R(c, result)
}

// @Summary 修改头像
// @Description 获取JSON
// @Tags 用户
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/user/profileAvatar [post]
func InsetSysUserAvatar(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	guid := uuid.New().String()
	filePath := "static/uploadfile/" + guid + ".jpg"
	for _, file := range files {
		// 上传文件至指定目录
		_ = c.SaveUploadedFile(file, filePath)
	}
	sysuser := model.SysUser{}
	sysuser.UserId = tools.GetUserId(c)
	sysuser.Avatar = "/" + filePath
	sysuser.UpdateBy = tools.GetUserIdStr(c)
	system.UpdateUser(sysuser)
	tools.R(c, filePath)
}

func SysUserUpdatePwd(c *gin.Context) {
	var pwd model.SysUserPwd
	err := c.Bind(&pwd)
	tools.HasError(err, "数据解析失败", 500)
	sysuser := model.SysUser{}
	sysuser.UserId = tools.GetUserId(c)
	sysuser.Password = pwd.NewPassword
	system.SetUserPwd(sysuser)
	tools.R(c)
}

func GetInfo(c *gin.Context) {

	var roles = make([]string, 1)
	jwtAuth := tools.GetJwtAuth(c)
	roles[0] = jwtAuth.Role

	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"

	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"

	u := system.GetUserById(jwtAuth.UserId)

	RoleMenu := model.SysRoleMenu{}
	RoleMenu.RoleId = u.RoleId

	var mp = make(map[string]interface{})
	mp["roles"] = roles
	if jwtAuth.Role == "admin" {
		mp["permissions"] = permissions
		mp["buttons"] = buttons
	} else {
		list, _ := system.GetPermisRoleMenu(RoleMenu)
		mp["permissions"] = list
		mp["buttons"] = list
	}

	sysuser := model.SysUser{}
	sysuser.UserId = jwtAuth.UserId

	mp["introduction"] = " am a super administrator"

	mp["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	if u.Avatar != "" {
		mp["avatar"] = u.Avatar
	}
	mp["userName"] = u.NickName
	mp["userId"] = u.UserId
	mp["deptId"] = u.DeptId
	mp["name"] = u.NickName

	tools.R(c, mp)
}
