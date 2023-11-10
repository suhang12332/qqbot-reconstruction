package server

import (
	"encoding/json"
	"fmt"
	"qqbot-reconstruction/internal/pkg/log"
	"qqbot-reconstruction/internal/pkg/variable"
	"strings"
)

type (
	Receive variable.ReceiveMessage
	Send    variable.SendMessage
)

// SearchSong
// @description: 搜素歌曲(网易云)
func (receive *Receive) SearchSong(info string) {
	initSend().
		assembleMessage(variable.Actions.SendMsg, false, receive, variable.NORMALESSAGE, assembleSendMsg, "").
		songMessage(info).
		sendMessage()
}

// SearchSong
// @description: 搜素云盘(阿里云盘)
func (receive *Receive) searchAli(info string) {
	send := initSend()
	if receive.MessageType == variable.PRIVATEMESSAGE {
		send = send.assembleMessage(variable.Actions.SendMsg, false, receive, variable.NORMALESSAGE, assembleSendMsg, "").aliMessage(info,variable.NORMALESSAGE)
	}else {
		send = send.assembleMessage(variable.Actions.SendGroupForwardMsg, false, receive, variable.GROUPMESSGAE, assembleSendMsg, "").aliMessage(info,variable.GROUPMESSGAE)
	}
	marshal, _ := json.Marshal(send)
	fmt.Println(string(marshal))
	send.sendMessage()
}

// initSend
// @description: 初始化消息
func initSend() *Send {
	return &Send{}
}

// assembleSendMsg
// @description: 组装发送消息
func assembleSendMsg(isSpae bool, receive *Receive, messageType string, info ...string) any {
	messages := make([]variable.Messages, 0)
	if messageType == variable.GROUPMESSGAE || messageType == variable.PRIVATEMESSAGE {
		json.Unmarshal([]byte(info[0]), &messages)
	}
	switch messageType {
	case variable.GROUPMESSGAE:
		return &variable.SendGroupForwardMsg{
			GroupID:  (*receive).GroupId,
			Messages: messages,
		}
	default:
		return &variable.SendMsg{
			MessageType: (*receive).MessageType,
			UserId:      (*receive).UserID,
			GroupId:     (*receive).GroupId,
			Message:     info[0],
			AutoEscape:  isSpae,
		}
	}

}

// assembleMessage
// @description: 组装消息
func (send *Send) assembleMessage(action string, isSpae bool, receive *Receive, messageType string,processor func(isSpae bool, receive *Receive, messageType string, info ...string) any, info ...string) *Send {
	send.Params = processor(isSpae, receive, messageType, info...)
	send.Action = action
	return send
}

func (send *Send) normalMessage() {

}

func (send *Send) groupMessage() {
	
}

// sendMessage
// @description: qq发送消息
// @param send websocket链接指针

func (send *Send) sendMessage() {
	marshal, err := json.Marshal(*send)
	switch send.Action {
	case variable.Actions.SendMsg:
		log.Info("回复消息: ", strings.ReplaceAll((*send).Params.(*variable.SendMsg).Message, "\n", "\t"))
	}
	if err != nil {
		log.Error("消息回复失败: ", err)
	}
	result := string(marshal)
	sendQMessage(&result)
}
