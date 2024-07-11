package plugins

import (
    "qqbot-reconstruction/internal/app/message"
    "qqbot-reconstruction/internal/pkg/variable"
    "strings"
)

type TestPlugin struct {
    name      string
    status    bool
    whitelist []string
}

func (t *TestPlugin) Execute(receive *message.Receive) *message.Send {
    if strings.Split(receive.RawMessage, " ")[1] == "订阅" {
        send := receive.InitSend(false)
        send.Params.(*variable.SendMsg).Message = "/早报 退订"
        return send
    }
    return nil
}

func (t *TestPlugin) GetWhiteList() []string {
    return t.whitelist
}

func (t *TestPlugin) SetWhiteList(whiteList []string) {
    t.whitelist = whiteList
}
