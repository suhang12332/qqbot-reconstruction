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
func (receive *Receive) SearchSong() {
	initSend(receive).
		songMessage(receive)
}

// SearchSong
// @description: 搜素云盘(阿里云盘)
func (receive *Receive) searchAli() {
	initSend(receive).
		aliMessage(receive)
}

func (receive *Receive) searchMagnet() {
	initSend(receive).
		magnetMessage(receive)
}

// initSend
// @description: 初始化消息
func initSend(receive *Receive) *Send {
	send := Send{}
	send.assembleMessage(false, receive, false, assembleSendMsg)
	return &send
}

// assembleSendMsg
// @description: 组装发送消息
func assembleSendMsg(isSpae bool, receive *Receive, isForward bool) any {
	if isForward && receive.MessageType == variable.GROUPMESSGAE {
		return &variable.SendGroupForwardMsg{
			GroupID: (*receive).GroupId,
		}
	} else {
		return &variable.SendMsg{
			MessageType: (*receive).MessageType,
			UserId:      (*receive).UserID,
			GroupId:     (*receive).GroupId,
			AutoEscape:  isSpae,
		}
	}

}

// assembleMessage
// @description: 组装消息
func (send *Send) assembleMessage(isSpae bool, receive *Receive, isForward bool, processor func(isSpae bool, receive *Receive, isForward bool) any) *Send {
	send.Params = processor(isSpae, receive, isForward)
	if isForward && receive.MessageType == variable.GROUPMESSGAE {
		send.Action = variable.Actions.SendGroupForwardMsg
	} else {
		send.Action = variable.Actions.SendMsg
	}
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
