package commons

import (
	"qqbot-reconstruction/internal/app/message"
)

type Plugin interface {
	Execute(receive *message.Receive) *message.Send
	GetWhiteList() []string
	SetWhiteList(whiteList []string)
}
