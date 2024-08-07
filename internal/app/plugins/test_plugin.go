package plugins

import (
    "qqbot-reconstruction/internal/app/message"
    "qqbot-reconstruction/internal/pkg/variable"
    "strings"
)

type TestPlugin struct {
    name      string
    keyword   string
    status    bool
    whitelist []string
    args      []string
    scope     []string
}

func (t *TestPlugin) SetName(name string) {
    t.name = name
}

func (t *TestPlugin) GetKeyword() string {
    return t.keyword
}

func (t *TestPlugin) SetKeyword(keyword string) {
    t.keyword = keyword
}

func (t *TestPlugin) GetName() string {
    return t.name
}

func (t *TestPlugin) GetStatus() bool {
    return t.status
}

func (t *TestPlugin) SetStatus(status bool) {
    t.status = status
}

func (t *TestPlugin) SetArgs(args []string) {
    t.args = args
}

func (t *TestPlugin) GetArgs() []string {
    return t.args
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
func (t *TestPlugin) SetScope(scope []string) {
    t.scope = scope
}

func (t *TestPlugin) GetScope() []string {
    return t.scope
}
func (t *TestPlugin) Help(receive *message.Receive) *message.Send {
    return receive.Tips("给傻逼说明一下用法🤭")
}
