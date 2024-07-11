package message

import (
	"qqbot-reconstruction/internal/pkg/variable"
)

// SendMessage
// @description: 发送消息
type SendMessage struct {
	Action string `json:"action"`
	Params any    `json:"params"`
	Echo   string `json:"echo"`
}

type Send SendMessage

// initSend
// @description: 初始化消息
func (receive *Receive) initSend(isForward bool) *Send {
	send := Send{}
	send.assembleMessage(isForward, receive)
	return &send
}

func NewSendMessage(isForward bool, receive *Receive) *Send {
	send := Send{}
	send.assembleMessage(isForward, receive)
	return &send
}

// assembleMessage
// @description: 组装消息
func (send *Send) assembleMessage(isForward bool, receive *Receive) *Send {
	if isForward {
		send.Params = &variable.SendPrivateForwardMsg{
			UserID: variable.QQ,
		}
		send.Action = variable.Actions.SendGroupForwardMsg
	} else {
		send.Params = &variable.SendMsg{
			MessageType: receive.MessageType,
			UserId:      receive.Sender.UserID,
			GroupId:     receive.GroupId,
		}
		send.Action = variable.Actions.SendMsg
	}

	return send
}
