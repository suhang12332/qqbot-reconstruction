package message

import (
    "encoding/json"
    "qqbot-reconstruction/internal/pkg/log"
    "qqbot-reconstruction/internal/pkg/variable"
    "strings"
)

// SendMessage
// @description: 发送消息
type SendMessage struct {
    Action string `json:"action"`
    Params any    `json:"params"`
    Echo   string `json:"echo"`
}

type Send SendMessage

func (receive *Receive) InitSend(isForward bool) *Send {
    send := Send{}
    send.assembleMessage(isForward, receive)
    return &send
}

// assembleMessage
// @description: 组装消息
func (send *Send) assembleMessage(isForward bool, receive *Receive) *Send {
    if isForward {
        switch receive.MessageType {
        case variable.PRIVATEMESSAGE:
            send.Params = &variable.SendPrivateForwardMsg{
                UserID: receive.UserID,
            }
            send.Action = variable.Actions.SendPrivateForwardMsg
            break
        case variable.GROUPMESSGAE:
            send.Params = &variable.SendGroupForwardMsg{
                GroupID: receive.GroupId,
            }
            send.Action = variable.Actions.SendGroupForwardMsg
            break
        }
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
func (receive *Receive) Tips(info string) *Send {
    send := receive.InitSend(false)
    ((*send).Params.(*variable.SendMsg)).Message = info
    return send
}
func (receive *Receive) NoPermissionsTips() *Send {
    return receive.Tips(variable.Tips.Info.NoPermissions)
}

func (receive *Receive) NoArgsTips() *Send {
    return receive.Tips(variable.Tips.Info.NoArgs)
}

func (receive *Receive) NoResults() *Send {
    return receive.Tips(variable.Tips.Info.NoResults)
}

func (receive *Receive) RequestFail() *Send {
    return receive.Tips(variable.Tips.Info.RequestFail)
}

func Send2res(send *Send) *string {
    marshal, err := json.Marshal(send)
    switch send.Action {
    case variable.Actions.SendMsg:
        log.Info("回复消息: ", strings.ReplaceAll((*send).Params.(*variable.SendMsg).Message, "\n", "\t"))
    }
    if err != nil {
        log.Error("消息回复失败: ", err)
    }
    result := string(marshal)
    log.Infof(result)

    return &result
}
