package model

var SysUserTableName = "sys_user"

// SysUser
type SysUserDto struct {
	// key
	IdentityKey string
	// 用户名
	UserName  string
	FirstName string
	LastName  string
	// 角色
	Role string
}

type LoginM struct {
	Username string `gorm:"size:64" json:"username"`
	Password string `gorm:"size:128" json:"-"`
}

type SysUser struct {
	UserId    int    `gorm:"primary_key;AUTO_INCREMENT"  json:"userId"` // 编码
	NickName  string `gorm:"size:128" json:"nickName"`                  // 昵称
	Phone     string `gorm:"size:11" json:"phone"`                      // 手机号
	RoleId    int    `gorm:"" json:"roleId"`                            // 角色编码
	Salt      string `gorm:"size:255" json:"salt"`                      //盐
	Avatar    string `gorm:"size:255" json:"avatar"`                    //头像
	Sex       string `gorm:"size:255" json:"sex"`                       //性别
	Email     string `gorm:"size:128" json:"email"`                     //邮箱
	DeptId    int    `gorm:"" json:"deptId"`                            //部门编码
	PostId    int    `gorm:"" json:"postId"`                            //职位编码
	CreateBy  string `gorm:"size:128" json:"createBy"`                  //
	UpdateBy  string `gorm:"size:128" json:"updateBy"`                  //
	Remark    string `gorm:"size:255" json:"remark"`                    //备注
	Status    string `gorm:"size:4;" json:"status"`
	DataScope string `gorm:"-" json:"dataScope"`
	Params    string `gorm:"-" json:"params"`

	BaseModel
	LoginM
}

func (SysUser) TableName() string {
	return "sys_user"
}

type SysUserPwd struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type SysUserPage struct {
	SysUser
	DeptName string `gorm:"-" json:"deptName"`
}

type SysUserView struct {
	SysUser
	RoleName string `gorm:"column:role_name"  json:"role_name"`
}
