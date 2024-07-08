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

// searchSong
// @description: 搜索歌曲(网易云)
func (receive *Receive) searchSong() {
    receive.initSend(false).songMessage(receive)
}

// searchAli
// @description: 搜索云盘(阿里云盘)
func (receive *Receive) searchAli() {
    receive.initSend(true).aliMessage(receive)
}

// searchMagnet
// @description: 搜索磁力(磁力蛋蛋)
func (receive *Receive) searchMagnet() {
    receive.initSend(true).magnetMessage(receive)
}

// initSend
// @description: 初始化消息
func (receive *Receive) initSend(isForward bool) *Send {
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
    fmt.Println(result)
    sendQMessage(&result)
}
