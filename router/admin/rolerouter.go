package admin

import (
	"github.com/gin-gonic/gin"
	"go-guns/app/controller/system"
)

func InitRoleRouter(g *gin.RouterGroup) {
	registerPageRouter(g)
	registerBaseRouter(g)
	registerDeptRouter(g)
	//registerDictRouter(g)
	registerSysUserRouter(g)
	registerRoleRouter(g)
	registerUserCenterRouter(g)
	registerPostRouter(g)
	registerMenuRouter(g)
}

func registerPageRouter(g *gin.RouterGroup) {
	g.GET("/deptList", system.GetDeptList)
	g.GET("/deptTree", system.GetDeptTree)
	g.GET("/sysUserList", system.GetSysUserList)
	g.GET("/rolelist", system.GetRoleList)
	g.GET("/postlist", system.GetPostList)
	g.GET("/menulist", system.GetMenuList)
}

func registerBaseRouter(g *gin.RouterGroup) {

	g.GET("/getinfo", system.GetInfo)
	g.GET("/menurole", system.GetMenuRole)
	g.PUT("/roledatascope", system.UpdateRoleDataScope)
	g.GET("/roleMenuTreeselect/:roleId", system.GetMenuTreeRoleselect)
	g.GET("/roleDeptTreeselect/:roleId", system.GetDeptTreeRoleselect)

	// todo 退出登录
	//g.POST("/logout", handler.LogOut)
	g.GET("/menuids", system.GetMenuIDS)

}

func registerDeptRouter(g *gin.RouterGroup) {
	dept := g.Group("/dept")
	{
		dept.GET("/:deptId", system.GetDept)
		dept.POST("", system.InsertDept)
		dept.PUT("", system.UpdateDept)
		dept.DELETE("/:id", system.DeleteDept)
	}
}

//func registerDictRouter(g *gin.RouterGroup) {
//	dicts := g.Group("/dict")
//	{
//		dicts.GET("/datalist", dict.GetDictDataList)
//		dicts.GET("/typelist", dict.GetDictTypeList)
//		dicts.GET("/typeoptionselect", dict.GetDictTypeOptionSelect)
//
//		dicts.GET("/data/:dictCode", dict.GetDictData)
//		dicts.POST("/data", dict.InsertDictData)
//		dicts.PUT("/data/", dict.UpdateDictData)
//		dicts.DELETE("/data/:dictCode", dict.DeleteDictData)
//
//		dicts.GET("/type/:dictId", dict.GetDictType)
//		dicts.POST("/type", dict.InsertDictType)
//		dicts.PUT("/type", dict.UpdateDictType)
//		dicts.DELETE("/type/:dictId", dict.DeleteDictType)
//	}
//}

func registerSysUserRouter(g *gin.RouterGroup) {
	sysuser := g.Group("/sysUser")
	{
		sysuser.GET("/:userId", system.GetSysUser)
		sysuser.GET("/", system.GetSysUserInit)
		sysuser.POST("", system.InsertSysUser)
		sysuser.PUT("", system.UpdateSysUser)
		sysuser.DELETE("/:userId", system.DeleteSysUser)
	}
}

func registerRoleRouter(v1 *gin.RouterGroup) {
	role := v1.Group("/role")
	{
		role.GET("/:roleId", system.GetRole)
		role.POST("", system.InsertRole)
		role.PUT("", system.UpdateRole)
		role.DELETE("/:roleId", system.DeleteRole)
	}
}

func registerUserCenterRouter(g *gin.RouterGroup) {
	user := g.Group("/user")
	{
		user.GET("/profile", system.GetSysUserProfile)
		user.POST("/avatar", system.InsetSysUserAvatar)
		user.PUT("/pwd", system.SysUserUpdatePwd)
	}
}

func registerPostRouter(g *gin.RouterGroup) {
	post := g.Group("/post")
	{
		post.GET("/:postId", system.GetPost)
		post.POST("", system.InsertPost)
		post.PUT("", system.UpdatePost)
		post.DELETE("/:postId", system.DeletePost)
	}
}

func registerMenuRouter(g *gin.RouterGroup) {
	menu := g.Group("/menu")
	{
		menu.GET("/:id", system.GetMenu)
		menu.POST("", system.InsertMenu)
		menu.PUT("", system.UpdateMenu)
		menu.DELETE("/:id", system.DeleteMenu)
	}
}

//func registerPublicRouter(v1 *gin.RouterGroup) {
//	p := v1.Group("/public")
//	{
//		p.POST("/uploadFile", public.UploadFile)
//	}
//}
