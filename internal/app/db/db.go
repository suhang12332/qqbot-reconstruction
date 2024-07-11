package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"qqbot-reconstruction/internal/pkg/log"
	"qqbot-reconstruction/internal/pkg/variable"
)

var (
	userCr = make([]variable.UserCountResult, 0)
	cr     = make([]variable.CountResult, 0)
	db     *gorm.DB
)

// DB
// @description: 获取db
// @return *gorm.DB
func DB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(variable.Urls.Mysql), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Error("数据库链接失败: %s", err)
		os.Exit(1)
	}
	sql, _ := db.DB()
	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	return db
}

// init
// @description: 初始化db
func init() {
	db = DB()
}

func InsertMessage(tReceive *variable.TReceive, tSender *variable.TSender) {
	db.Create(tSender)
	tReceive.Sender = tSender.ID
	db.Create(tReceive)
}
