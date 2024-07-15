package message

import (
	"fmt"
	"qqbot-reconstruction/internal/app/db"
	"qqbot-reconstruction/internal/pkg/log"
	"qqbot-reconstruction/internal/pkg/variable"
)

// SenderMessage
// @description: 用户信息结构体
type SenderMessage struct {
	Age      int    `json:"age"`
	Area     string `json:"area"`
	Card     string `json:"card"`
	Level    string `json:"level"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
	Sex      string `json:"sex"`
	Title    string `json:"title"`
	UserID   int    `json:"user_id"`
}

// ReceiveMessage
// @description: 接收消息
type ReceiveMessage struct {
	PostType    string        `json:"post_type"`
	MessageType string        `json:"message_type"`
	Time        int           `json:"time"`
	SelfID      int64         `json:"self_id"`
	SubType     string        `json:"sub_type"`
	Sender      SenderMessage `json:"sender"`
	MessageID   int           `json:"message_id"`
	UserID      int           `json:"user_id"`
	TargetID    int64         `json:"target_id"`
	Message     string        `json:"business"`
	MessageSeq  int           `json:"message_seq"`
	RawMessage  string        `json:"raw_message"`
	Font        int           `json:"font"`
	GroupId     int           `json:"group_id"`
}

type Receive ReceiveMessage

// ParseTReceive
// @description: 用于转化客户端结构体,用于数据库入库操作
// @return: 结果
func (receive *Receive) ParseTReceive() (*variable.TReceive, *variable.TSender) {
	sender := (*receive).Sender
	tSender := variable.TSender{
		Age:      (sender).Age,
		Area:     (sender).Area,
		Card:     (sender).Card,
		Level:    (sender).Level,
		Nickname: (sender).Nickname,
		Role:     (sender).Role,
		Sex:      (sender).Sex,
		Title:    (sender).Title,
		UserID:   (sender).UserID,
	}
	tReceive := variable.TReceive{
		PostType:    (*receive).PostType,
		MessageType: (*receive).MessageType,
		Time:        (*receive).Time,
		SelfID:      (*receive).SelfID,
		SubType:     (*receive).SubType,
		MessageID:   (*receive).MessageID,
		UserID:      (*receive).UserID,
		TargetID:    (*receive).TargetID,
		Message:     (*receive).Message,
		MessageSeq:  (*receive).MessageSeq,
		RawMessage:  (*receive).RawMessage,
		ID:          tSender.ID,
		Font:        (*receive).Font,
		GroupId:     (*receive).GroupId,
	}
	return &tReceive, &tSender
}

func (receive *Receive) PrintfMessage() {
	receiveMsg := (*receive).RawMessage
	if (*receive).MessageType == variable.GROUPMESSGAE {
		groupId := (*receive).GroupId
		if groupId != 0 {
			card := (*receive).Sender.Card
			if card == "" {
				card = (*receive).Sender.Nickname
			}
			log.Infof(fmt.Sprintf("收到消息: %s", fmt.Sprintf("群(%d)内 '%s':  %s", groupId, card, receiveMsg)))
		}
	} else {
		log.Infof(fmt.Sprintf("收到消息: %s", fmt.Sprintf("'%s'': %s", (*receive).Sender.Nickname, receiveMsg)))
	}
	//插入消息数据
	tReceive, tSender := receive.ParseTReceive()
	db.Database.InsertMessage(tReceive, tSender)
}
