package server

import (
	"encoding/json"
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
		assembleMessage(variable.Actions.SendMsg, false, receive, assembleSendMsg, "").
		songMessage(info).
		sendMessage()
}

// SearchSong
// @description: 搜素云盘(阿里云盘)
func (receive *Receive) searchAli(info string) {
	initSend().
		assembleMessage(variable.Actions.SendMsg, false, receive, assembleSendMsg, "").
		songMessage(info).
		sendMessage()
}

// initSend
// @description: 初始化消息
func initSend() *Send {
	return &Send{}
}

// assembleSendMsg
// @description: 组装发送消息
func assembleSendMsg(isSpae bool, receive *Receive, info ...string) *variable.SendMsg {
	return &variable.SendMsg{
		MessageType: (*receive).MessageType,
		UserId:      (*receive).UserID,
		GroupId:     (*receive).GroupId,
		Message:     info[0],
		AutoEscape:  isSpae,
	}
}

// assembleMessage
// @description: 组装消息
func (send *Send) assembleMessage(action string, isSpae bool, receive *Receive, processor func(isSpae bool, receive *Receive, info ...string) *variable.SendMsg, info ...string) *Send {
	(*send).Params = processor(isSpae, receive, info...)
	(*send).Action = action
	return send
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
