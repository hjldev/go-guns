package system

import (
	"errors"
	"go-guns/database"
	"go-guns/model"
	"go-guns/tools"
	"strconv"
)

func CreateDept(e model.SysDept) (err error) {
	result := database.Db.Create(&e)
	if result.Error != nil {
		err := result.Error
		return err
	}
	deptPath := "/" + strconv.Itoa(e.DeptId)
	if e.ParentId != 0 {
		var deptP model.SysDept
		database.Db.Where("dept_id = ?", e.ParentId).First(&deptP)
		deptPath = deptP.DeptPath + deptPath
	} else {
		deptPath = "/0" + deptPath
	}
	if err := database.Db.Model(&e).Where("dept_id = ?", e.DeptId).Update("deptPath", deptPath).Error; err != nil {
		err := result.Error
		return err
	}
	e.DeptPath = deptPath
	return nil
}

func GetDept(e model.SysDept) (model.SysDept, error) {

	table := database.Db.Table(model.SysDeptTableName)
	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}
	if e.DeptName != "" {
		table = table.Where("dept_name = ?", e.DeptName)
	}

	if err := table.First(&e).Error; err != nil {
		return e, err
	}
	return e, nil
}

func GetDeptList(e model.SysDept) ([]model.SysDept, error) {
	var doc []model.SysDept

	table := database.Db.Table(model.SysDeptTableName)
	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}
	if e.DeptName != "" {
		table = table.Where("dept_name = ?", e.DeptName)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	if err := table.Order("sort").Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func GetDeptPage(e model.SysDept) ([]model.SysDept, error) {
	var doc []model.SysDept

	table := database.Db.Select("*").Table(model.SysDeptTableName)
	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}
	if e.DeptName != "" {
		table = table.Where("dept_name = ?", e.DeptName)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}
	if e.DeptPath != "" {
		table = table.Where("deptPath like %?%", e.DeptPath)
	}

	if err := table.Order("sort").Find(&doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}

func SetDept(e model.SysDept) ([]model.SysDept, error) {
	list, err := GetDeptPage(e)

	m := make([]model.SysDept, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		info := DiguiSysDept(&list, list[i])

		m = append(m, info)
	}
	return m, err
}

func DiguiSysDept(deptlist *[]model.SysDept, menu model.SysDept) model.SysDept {
	list := *deptlist

	min := make([]model.SysDept, 0)
	for j := 0; j < len(list); j++ {

		if menu.DeptId != list[j].ParentId {
			continue
		}
		mi := model.SysDept{}
		mi.DeptId = list[j].DeptId
		mi.ParentId = list[j].ParentId
		mi.DeptPath = list[j].DeptPath
		mi.DeptName = list[j].DeptName
		mi.Sort = list[j].Sort
		mi.Leader = list[j].Leader
		mi.Phone = list[j].Phone
		mi.Email = list[j].Email
		mi.Status = list[j].Status
		mi.Children = []model.SysDept{}
		ms := DiguiSysDept(deptlist, mi)
		min = append(min, ms)

	}
	menu.Children = min
	return menu
}

func UpdateDept(e model.SysDept) (err error) {
	var update model.SysDept
	if err = database.Db.Where("dept_id = ?", e.DeptId).First(&update).Error; err != nil {
		return
	}

	deptPath := "/" + strconv.Itoa(e.DeptId)
	if int(e.ParentId) != 0 {
		var deptP model.SysDept
		database.Db.Where("dept_id = ?", e.ParentId).First(&deptP)
		deptPath = deptP.DeptPath + deptPath
	} else {
		deptPath = "/0" + deptPath
	}
	e.DeptPath = deptPath

	if e.DeptPath != "" && e.DeptPath != update.DeptPath {
		return errors.New("上级部门不允许修改！")
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = database.Db.Model(&update).Updates(&e).Error; err != nil {
		return
	}

	return
}

func DeleteDept(e model.SysDept) (success bool, err error) {

	user := model.SysUser{}
	user.DeptId = e.DeptId
	userlist, err := GetUserList(user)
	tools.HasError(err, 500, "")
	tools.Assert(len(userlist) <= 0, "当前部门存在用户，不能删除！", 500)

	tx := database.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err != nil {
		success = false
		return
	}

	if err = tx.Delete(&e).Error; err != nil {
		success = false
		tx.Rollback()
		return
	}
	if err = tx.Commit().Error; err != nil {
		success = false
		return
	}
	success = true

	return
}

func SetDeptLabel(dept model.SysDept) (m []model.SysDeptLabel, err error) {
	deptlist, err := GetDeptList(dept)

	m = make([]model.SysDeptLabel, 0)
	for i := 0; i < len(deptlist); i++ {
		if deptlist[i].ParentId != 0 {
			continue
		}
		e := model.SysDeptLabel{}
		e.Id = deptlist[i].DeptId
		e.Label = deptlist[i].DeptName
		deptsInfo := DiguiDeptLable(&deptlist, e)

		m = append(m, deptsInfo)
	}
	return
}

func DiguiDeptLable(deptlist *[]model.SysDept, dept model.SysDeptLabel) model.SysDeptLabel {
	list := *deptlist

	min := make([]model.SysDeptLabel, 0)
	for j := 0; j < len(list); j++ {

		if dept.Id != list[j].ParentId {
			continue
		}
		mi := model.SysDeptLabel{Id: list[j].DeptId, Label: list[j].DeptName, Children: []model.SysDeptLabel{}}
		ms := DiguiDeptLable(deptlist, mi)
		min = append(min, ms)

	}
	dept.Children = min
	return dept
}
