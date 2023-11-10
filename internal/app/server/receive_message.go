package server

import (
	"encoding/json"
	"fmt"
	"qqbot-reconstruction/internal/pkg/log"
	"strings"
)

func receiveMessage(message string) {
	switch {
	case strings.Contains(message, `post_type":"message"`):
		handleReceiveMessage(message, &Receive{}).switchFunction()
	}
}

func handleReceiveMessage[T any](message string, t *T) *T {
	bytes := []byte(message)
	json.Unmarshal(bytes, t)
	return t
}

func (receive *Receive) switchFunction() {
	receive.printfMessage()
	split := strings.Split(receive.RawMessage, " ")
	if len(split) == 2 {
		switch split[0] {
		case "点歌":
			receive.SearchSong(split[1])
			break
		case "云盘":
			receive.searchAli(split[1])
			break
		}
	}
}

func (receive *Receive) printfMessage() {
	receiveMsg := (*receive).RawMessage
	if (*receive).MessageType == "group" {
		groupId := (*receive).GroupId
		if groupId != 0 {
			card := (*receive).Sender.Card
			if card == "" {
				card = (*receive).Sender.Nickname
			}
			log.Infof(fmt.Sprintf("收到消息: %s", fmt.Sprintf("群(%d)内 '%s':  %s", groupId, card, receiveMsg)))
		}
	} else {
		log.Infof(fmt.Sprintf("收到消息: %s", fmt.Sprintf("'%s'': %s", (*receive).Sender.Nickname, receiveMsg)))
	}
	//插入消息数据
	//	repository.MessageInsert(receive)
}
