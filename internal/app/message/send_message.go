package message

import (
    "encoding/json"
    "fmt"
    "qqbot-reconstruction/internal/pkg/log"
    "qqbot-reconstruction/internal/pkg/util"
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
    if receive.MessageType == variable.GROUPMESSGAE {
        info = util.ALtCQ(receive.UserID,info)
    }
    ((*send).Params.(*variable.SendMsg)).Message = info
    return send
}

func (receive *Receive) ScopeTips(name string,scope string) *Send {

    return receive.Tips(fmt.Sprintf(`"%s"功能仅在中使用%s中使用`, name, ParseMessageType(scope)))
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
        log.Info("回复消息: %s", strings.ReplaceAll((*send).Params.(*variable.SendMsg).Message, "\n", "\t"))
    }
    if err != nil {
        log.Error("消息回复失败: ", err)
    }
    result := string(marshal)

    return &result
}

func (send *Send) ForwardMsg(data []variable.Messages) {
    switch (*send).Action {
    case variable.Actions.SendGroupForwardMsg:
        ((*send).Params.(*variable.SendGroupForwardMsg)).Messages = data
        break
    case variable.Actions.SendPrivateForwardMsg:
        ((*send).Params.(*variable.SendPrivateForwardMsg)).Messages = data
        break
    }
}

func ParseMessageType(info string) string {
    switch info {
    case variable.PRIVATEMESSAGE:
        return variable.PRIVATEMESSAGEZH
    case variable.GROUPMESSGAE:
        return variable.GROUPMESSGAEZH
    default:
        return variable.UNKNOWNMESSAGEZH
    }
}