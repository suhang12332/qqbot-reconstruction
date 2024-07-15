package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"qqbot-reconstruction/internal/pkg/log"
	"qqbot-reconstruction/internal/pkg/variable"
)

var Database *DB

type DB struct {
	db *gorm.DB
}

func NewDB() *DB {
	dbInstance, err := gorm.Open(mysql.Open(variable.Urls.Mysql), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("数据库链接失败: %s", err.Error())
	}
	sql, _ := dbInstance.DB()
	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)

	return &DB{db: dbInstance}
}

func (d *DB) InsertMessage(receive *variable.TReceive, sender *variable.TSender) {
	d.db.Create(sender)
	receive.Sender = sender.ID
	d.db.Create(receive)
}

func (d *DB) GenDialogHistory(GroupId string, period variable.Period) []variable.DialogHistory {
	var results []variable.DialogHistory

	subQuery := d.db.Table("t_senders").
		Select("id").
		Where("user_id = t_receives.user_id").
		Order("id DESC").
		Limit(1)

	d.db.Table("t_receives").
		Select("t_receives.time, t_receives.raw_message, t_senders.card").
		Joins("JOIN t_senders ON t_receives.user_id = t_senders.user_id").
		Where("t_senders.id = (?) AND t_receives.group_id = (?) AND t_receives.time BETWEEN (?) AND (?)", subQuery, GroupId, period.Begin, period.End).
		Scan(&results)

	return results
}

// GetCardById 通过qq号获取群昵称
func (d *DB) GetCardById(id string) string {
	var card string

	err := d.db.Table("t_senders").
		Select("card").
		Where("user_id = ?", id).
		Order("id DESC").
		Limit(1).
		Row().Scan(&card)
	if err != nil {
		log.Errorf(err.Error())
	}

	return card
}

func init() {
	Database = NewDB()
}
