package plugins

import (
	"qqbot-reconstruction/internal/app/message"
	"qqbot-reconstruction/internal/pkg/variable"
)

type TestSubscribePlugin struct {
	name         string
	keyword      string
	status       bool
	subscribable bool
	whitelist    []string
	args         []string
	scope        []string
}

func (t *TestSubscribePlugin) SetSubscribable(b bool) {
	t.subscribable = b
}

func (t *TestSubscribePlugin) Subscribable() bool {
	return t.subscribable
}

func (t *TestSubscribePlugin) SetName(name string) {
	t.name = name
}

func (t *TestSubscribePlugin) GetKeyword() string {
	return t.keyword
}

func (t *TestSubscribePlugin) SetKeyword(keyword string) {
	t.keyword = keyword
}

func (t *TestSubscribePlugin) GetName() string {
	return t.name
}

func (t *TestSubscribePlugin) GetStatus() bool {
	return t.status
}

func (t *TestSubscribePlugin) SetStatus(status bool) {
	t.status = status
}

func (t *TestSubscribePlugin) SetArgs(args []string) {
	t.args = args
}

func (t *TestSubscribePlugin) GetArgs() []string {
	return t.args
}

func (t *TestSubscribePlugin) Execute(receive *message.Receive) *message.Send {
	send := receive.InitSend(false)
	send.Params.(*variable.SendMsg).Message = "消息订阅测试"
	return send
}

func (t *TestSubscribePlugin) GetWhiteList() []string {
	return t.whitelist
}

func (t *TestSubscribePlugin) SetWhiteList(whiteList []string) {
	t.whitelist = whiteList
}
func (t *TestSubscribePlugin) SetScope(scope []string) {
	t.scope = scope
}

func (t *TestSubscribePlugin) GetScope() []string {
	return t.scope
}
func (t *TestSubscribePlugin) Help(receive *message.Receive) *message.Send {
	return receive.Tips("给傻逼说明一下用法🤭")
}
