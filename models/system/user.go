package system

import (
	"errors"
	"go-guns/database"
	"go-guns/models"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

// User
type UserDto struct {
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
	Password string `gorm:"size:128" json:"password"`
}

type User struct {
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

	models.BaseModel
	LoginM
}

func (User) TableName() string {
	return "sys_user"
}

type SysUserPwd struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type SysUserPage struct {
	User
	DeptName string `gorm:"-" json:"deptName"`
}

type SysUserView struct {
	User
	RoleName string `gorm:"column:role_name"  json:"role_name"`
}

// 获取用户数据
func (e *User) Get() (SysUserView SysUserView, err error) {

	table := database.Db.Table(e.TableName()).Select([]string{"sys_user.*", "sys_role.role_name"})
	table = table.Joins("left join sys_role on sys_user.role_id=sys_role.role_id")
	if e.UserId != 0 {
		table = table.Where("user_id = ?", e.UserId)
	}

	if e.Username != "" {
		table = table.Where("username = ?", e.Username)
	}

	if e.Password != "" {
		table = table.Where("password = ?", e.Password)
	}

	if e.RoleId != 0 {
		table = table.Where("role_id = ?", e.RoleId)
	}

	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}

	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}

	if err = table.First(&SysUserView).Error; err != nil {
		return
	}

	SysUserView.Password = ""
	return
}

func (e *User) GetUserInfo() (SysUserView SysUserView, err error) {

	table := database.Db.Table(e.TableName()).Select([]string{"sys_user.*", "sys_role.role_name"})
	table = table.Joins("left join sys_role on sys_user.role_id=sys_role.role_id")
	if e.UserId != 0 {
		table = table.Where("user_id = ?", e.UserId)
	}

	if e.Username != "" {
		table = table.Where("username = ?", e.Username)
	}

	if e.Password != "" {
		table = table.Where("password = ?", e.Password)
	}

	if e.RoleId != 0 {
		table = table.Where("role_id = ?", e.RoleId)
	}

	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}

	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}

	if err = table.First(&SysUserView).Error; err != nil {
		return
	}
	return
}

func (e *User) GetList() (SysUserView []SysUserView, err error) {

	table := database.Db.Table(e.TableName()).Select([]string{"sys_user.*", "sys_role.role_name"})
	table = table.Joins("left join sys_role on sys_user.role_id=sys_role.role_id")
	if e.UserId != 0 {
		table = table.Where("user_id = ?", e.UserId)
	}

	if e.Username != "" {
		table = table.Where("username = ?", e.Username)
	}

	if e.Password != "" {
		table = table.Where("password = ?", e.Password)
	}

	if e.RoleId != 0 {
		table = table.Where("role_id = ?", e.RoleId)
	}

	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}

	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}

	if err = table.Find(&SysUserView).Error; err != nil {
		return
	}
	return
}

func (e *User) GetPage(pageSize int, pageIndex int) ([]SysUserPage, int, error) {
	var doc []SysUserPage
	table := database.Db.Select("sys_user.*,sys_dept.dept_name").Table(e.TableName())
	table = table.Joins("left join sys_dept on sys_dept.dept_id = sys_user.dept_id")

	if e.Username != "" {
		table = table.Where("username = ?", e.Username)
	}
	if e.Status != "" {
		table = table.Where("sys_user.status = ?", e.Status)
	}

	if e.Phone != "" {
		table = table.Where("sys_user.phone = ?", e.Phone)
	}

	if e.DeptId != 0 {
		table = table.Where("sys_user.dept_id in (select dept_id from sys_dept where dept_path like ? )", "%"+strconv.Itoa(e.DeptId)+"%")
	}

	// 数据权限控制
	dataPermission := new(DataPermission)
	dataPermission.UserId, _ = strconv.Atoi(e.DataScope)
	table, err := dataPermission.GetDataScope("sys_user", table)
	if err != nil {
		return nil, 0, err
	}
	var count int64

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("sys_user.deleted_at IS NULL").Count(&count)
	return doc, int(count), nil
}

//加密
func (e *User) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

//添加
func (e User) Insert() (id int, err error) {
	if err = e.Encrypt(); err != nil {
		return
	}

	// check 用户名
	var count int64
	database.Db.Table(e.TableName()).Where("username = ?", e.Username).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}

	//添加数据
	if err = database.Db.Table(e.TableName()).Create(&e).Error; err != nil {
		return
	}
	id = e.UserId
	return
}

//修改
func (e *User) Update(id int) (update User, err error) {
	if e.Password != "" {
		if err = e.Encrypt(); err != nil {
			return
		}
	}
	if err = database.Db.Table(e.TableName()).First(&update, id).Error; err != nil {
		return
	}
	if e.RoleId == 0 {
		e.RoleId = update.RoleId
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = database.Db.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (e *User) BatchDelete(id []int) (Result bool, err error) {
	if err = database.Db.Table(e.TableName()).Where("user_id in (?)", id).Delete(&User{}).Error; err != nil {
		return
	}
	Result = true
	return
}

// todo 修改密码
func (e *User) SetPwd(pwd SysUserPwd) (Result bool, err error) {
	return true, nil
}
