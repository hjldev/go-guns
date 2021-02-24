package boot

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Driver string
	Source string
}

var Db *gorm.DB

func InitDatabase(cfg *viper.Viper) {
	dbCfg := Database{
		Driver: cfg.GetString("driver"),
		Source: cfg.GetString("source"),
	}

	dbCfg.Setup()
}

func (database *Database) Setup() {
	var err error
	Db, err = gorm.Open(mysql.Open(database.Source), &gorm.Config{})
	if err != nil {

	}
}
