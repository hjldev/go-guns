package system

import (
	"go-guns/database"
	"go-guns/model"
)

func CreatePost(e model.SysPost) error {
	result := database.Db.Create(&e)
	if result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

func GetPost(e model.SysPost) (model.SysPost, error) {
	var doc model.SysPost

	table := database.Db
	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}
	if e.PostName != "" {
		table = table.Where("post_name = ?", e.PostName)
	}
	if e.PostCode != "" {
		table = table.Where("post_code = ?", e.PostCode)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func GetPostList(e model.SysPost) ([]model.SysPost, error) {
	var doc []model.SysPost

	table := database.Db
	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}
	if e.PostName != "" {
		table = table.Where("post_name = ?", e.PostName)
	}
	if e.PostCode != "" {
		table = table.Where("post_code = ?", e.PostCode)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func GetPostPage(e model.SysPost, pageSize int, pageIndex int) ([]model.SysPost, int, error) {
	var doc []model.SysPost

	table := database.Db
	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}
	if e.PostName != "" {
		table = table.Where("post_name = ?", e.PostName)
	}
	if e.PostCode != "" {
		table = table.Where("post_code = ?", e.PostCode)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	var count int64

	if err := table.Order("sort").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, int(count), nil
}

func UpdatePost(e model.SysPost) (err error) {

	if err = database.Db.Model(&e).Updates(&e).Error; err != nil {
		return
	}
	return
}

func DeletePost(id int) (success bool, err error) {
	if err = database.Db.Delete(&model.SysPost{}, id).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

func BatchDeletePost(ids []int) (Result bool, err error) {
	if err = database.Db.Delete(&model.SysPost{}, ids).Error; err != nil {
		return
	}
	Result = true
	return
}
