package server

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "os"
    "qqbot-reconstruction/internal/pkg/log"
    "qqbot-reconstruction/internal/pkg/variable"
    "strings"
)

type CountResult struct {
    Label string
    Count int
}
type UserCountResult struct {
    Label  string
    UserId string
    Count  int
}

var (
    userCr = make([]UserCountResult, 0)
    cr     = make([]CountResult, 0)
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

// MessageInsert
// @description: 插入聊天消息
// @param message 消息
func (receive *Receive) MessageInsert() {
    tReceive, tSender := receive.ParseTReceive()
    db.Create(tSender)
    tReceive.Sender = tSender.ID
    db.Create(tReceive)
}

// MessageStatistics
// @description: 消息统计
// @param message 消息
// @return string 消息条数
func (receive *Receive) MessageStatistics() string {
    var msg string
    if (*receive).MessageType == "group" {
        row, err := db.Model(&(variable.TReceive{})).Select("if(t_senders.card ='',t_senders.nickname,t_senders.card) as card,count(t_senders.card) as count").Joins("left join t_senders on t_receives.sender = t_senders.id").Where(fmt.Sprintf("t_receives.group_id = '%d'", receive.GroupId)).Group("if(card ='',nickname,card)").Order("count desc").Rows()
        if err != nil {
            log.Error("查询失败: ", err)
        }
        defer row.Close()
        for row.Next() {
            var info variable.SpeechesNumber
            err := row.Scan(&info.Card, &info.Count)
            if err != nil {
                log.Error("数据解析失败: ", err)
            }
            if info.Card != "Q群管家" {
                msg = fmt.Sprintf("%s%s  %d条\n", msg, info.Card, info.Count)
            }
        }
    } else {
        msg = "该功能仅支持群聊中使用"
    }
    return strings.TrimRight(msg, "\n")
}

