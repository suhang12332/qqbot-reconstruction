package commons

import (
	"qqbot-reconstruction/internal/app/message"
)

type Plugin interface {
	Execute(receive *message.Receive) *message.Send
	GetKeyword() string
	SetKeyword(keyword string)
	GetWhiteList() []string
	SetWhiteList(whiteList []string)
	GetStatus() bool
	SetStatus(status bool)
	GetName() string
	SetName(name string)
	SetArgs(args []string)
	GetArgs() []string
	SetScope(args []string)
	GetScope() []string
	SetSubscribable(bool)
	Subscribable() bool
	Help(receive *message.Receive) *message.Send
}
