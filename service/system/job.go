package system

import (
	"go-guns/database"
	"go-guns/model"
)

// 创建SysJob
func CreateJob(e model.SysJob) error {
	result := database.Db.Create(&e)
	if result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

// 获取SysJob
func GetJob(e model.SysJob) (model.SysJob, error) {
	table := database.Db.Table(model.SysJobTableName)

	if e.JobId != 0 {
		table = table.Where("job_id = ?", e.JobId)
	}

	if e.JobName != "" {
		table = table.Where("job_name like ?", "%"+e.JobName+"%")
	}

	if e.JobGroup != "" {
		table = table.Where("job_group = ?", e.JobGroup)
	}

	if e.CronExpression != "" {
		table = table.Where("cron_expression = ?", e.CronExpression)
	}

	if e.InvokeTarget != "" {
		table = table.Where("invoke_target = ?", e.InvokeTarget)
	}

	if e.Status != 0 {
		table = table.Where("status = ?", e.Status)
	}

	if err := table.First(&e).Error; err != nil {
		return e, err
	}
	return e, nil
}

// 获取SysJob带分页
func GetJobPage(e model.SysJob, pageSize int, pageIndex int) ([]model.SysJob, int, error) {
	var doc []model.SysJob

	table := database.Db.Select("*").Table(model.SysJobTableName)

	if e.JobId != 0 {
		table = table.Where("job_id = ?", e.JobId)
	}

	if e.JobName != "" {
		table = table.Where("job_name like ?", "%"+e.JobName+"%")
	}

	if e.JobGroup != "" {
		table = table.Where("job_group = ?", e.JobGroup)
	}

	if e.CronExpression != "" {
		table = table.Where("cron_expression = ?", e.CronExpression)
	}

	if e.InvokeTarget != "" {
		table = table.Where("invoke_target = ?", e.InvokeTarget)
	}

	if e.Status != 0 {
		table = table.Where("status = ?", e.Status)
	}

	var count int64

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, int(count), nil
}

func GetJobList() ([]model.SysJob, error) {
	var doc []model.SysJob

	if err := database.Db.Where("status = ?", 2).Find(&doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}

// 更新SysJob
func UpdateJob(e model.SysJob, id int) (update model.SysJob, err error) {

	if err = database.Db.Model(&e).Updates(&e).Error; err != nil {
		return
	}
	return
}

func RemoveJobAllEntryID() (err error) {

	if err = database.Db.Model(&model.SysJob{}).Updates(map[string]interface{}{"entry_id": 0}).Error; err != nil {
		return
	}
	return
}

func RemoveJobEntryID(entryID int) (err error) {

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = database.Db.Model(&model.SysJob{}).Where("entry_id = ?", entryID).Updates(map[string]interface{}{"entry_id": 0}).Error; err != nil {
		return
	}
	return
}

// 删除SysJob
func DeleteJob(id int) (success bool, err error) {
	if err = database.Db.Delete(&model.SysJob{}, id).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func BatchDeleteJob(id []int) (Result bool, err error) {
	if err = database.Db.Delete(&model.SysJob{}, id).Error; err != nil {
		return
	}
	Result = true
	return
}
