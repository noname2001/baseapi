package models

import (
	"baseapi/global"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	admin := global.BA_CONFIG.Mysql
	mysqlConfig := mysql.Config{
		DSN:                       admin.Username + ":" + admin.Password + "@(" + admin.Path + ")/" + admin.Dbname + "?" + admin.Config, // DSN data source name
		DefaultStringSize:         191,                                                                                                  // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                                 // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                                 // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                                 // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                                // 根据版本自动配置
	}
	var gormConfig *gorm.Config
	if admin.LogMode { //根据配置决定是否开启日志
		gormConfig = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		gormConfig = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}

	var err error
	if db, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		global.BA_LOG.Error("MySQL启动异常", err)
		os.Exit(0)
	} else {
		global.BA_DB = db
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(admin.MaxIdleConns)
		sqlDB.SetMaxOpenConns(admin.MaxOpenConns)
	}
	dbTables()
	insertData()
}

func dbTables() {
	err := db.AutoMigrate(
		&Tag{},
		&User{},
	)
	if err != nil {
		global.BA_LOG.Error("register table failed", err)
		os.Exit(0)
	}
}

func insertData() {
	var ptr = new(int)
	*ptr = 0
	for i := 0; i < 3; i++ {
		// db.Model(&user).Association("Tags").Append(&Tag{Name: randomdata.SillyName(), State: ptr})
		// db.Model(&user).Association("Followers").Append(&User{Name: randomdata.SillyName(), Password: randomdata.Alphanumeric(8)})
		// db.Create(&Tag{Name: randomdata.SillyName(), State: randomdata.Number(-1, 10)})
		// db.Create(&User{Name: randomdata.SillyName(), Password: randomdata.Alphanumeric(8)})
	}
	var user, userb User
	db.Find(&user, 238)
	db.Find(&userb, 240)
	db.Model(&user).Association("Followers").Append(&userb)
}
