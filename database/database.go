package database

import (
	"go-guns/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Setup() {
	var err error
	Db, err = gorm.Open(mysql.Open(config.DbCfg.Source), &gorm.Config{})
	if err != nil {

	}
}
