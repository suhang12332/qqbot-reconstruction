package plugins

import (
	"fmt"
	"net/http"
	"qqbot-reconstruction/internal/app/db"
	"qqbot-reconstruction/internal/app/message"
	"qqbot-reconstruction/internal/pkg/api"
	"qqbot-reconstruction/internal/pkg/util"
	"qqbot-reconstruction/internal/pkg/variable"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// GptSummaryPlugin 使用GPT总结每日群聊的内容
type GptSummaryPlugin struct {
	name         string
	keyword      string
	status       bool
	subscribable bool
	whitelist    []string
	args         []string
	scope        []string
}

func (g *GptSummaryPlugin) SetSubscribable(b bool) {
	g.subscribable = b
}

func (g *GptSummaryPlugin) Subscribable() bool {
	return g.subscribable
}

func (g *GptSummaryPlugin) SetScope(args []string) {
	g.scope = args
}

func (g *GptSummaryPlugin) GetScope() []string {
	//TODO implement me
	return g.scope
}

func (g *GptSummaryPlugin) SetArgs(args []string) {
	g.args = args
}

func (g *GptSummaryPlugin) GetArgs() []string {
	return g.args
}

func (g *GptSummaryPlugin) Help(receive *message.Receive) *message.Send {
	return receive.Tips(util.ParseHelpTips("总结群聊的内容", `使用AI对过去24小时内的聊天内容进行总结;使用 "/总结"`, `/色图`, util.ParseHelp(g.scope)))
}

func (g *GptSummaryPlugin) Execute(receive *message.Receive) *message.Send {
	//util.Subscribe(24*time.Hour, "23:00", func() {
	//
	//})
	send := receive.InitSend(false)
	summary, ok := g.Query(g.GetMessageSummary(strconv.Itoa(receive.GroupId)))
	if ok {
		send.Params.(*variable.SendMsg).Message = summary
		return send
	}
	return nil
}

func (g *GptSummaryPlugin) GetMessageSummary(groupId string) string {
	dialogHistory := db.Database.GenDialogHistory(groupId, variable.Period{
		Begin: time.Now().Add(-24 * time.Hour).Unix(),
		End:   time.Now().Unix(),
	})
	var dialog string
	for _, m := range dialogHistory {
		dialog += Dialog2String(&m)
	}
	return dialog
}

func (g *GptSummaryPlugin) Query(dialog string) (string, bool) {
	header := make(map[string]string)
	header["Authorization"] = variable.Urls.GptKey
	body := map[string]interface{}{
		"model":           "kimi",
		"conversation_id": "cqa9u5g967u4ac8lmbn0",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": dialog,
			},
		},
		"use_search": false,
		"stream":     false,
	}
	//	jsonData, _ := json.Marshal(body)
	_, v, ok := api.Fetch(http.MethodPost, variable.Urls.Gpt, body, &variable.GPTResponse{}, header, variable.JSON, false, nil, false, nil)

	return v.Choices[0].Message.Content, ok
}

func (g *GptSummaryPlugin) GetKeyword() string {
	return g.keyword
}

func (g *GptSummaryPlugin) SetKeyword(keyword string) {
	g.keyword = keyword
}

func (g *GptSummaryPlugin) GetWhiteList() []string {
	return g.whitelist
}

func (g *GptSummaryPlugin) SetWhiteList(whiteList []string) {
	g.whitelist = whiteList
}

func (g *GptSummaryPlugin) GetStatus() bool {
	return g.status
}

func (g *GptSummaryPlugin) SetStatus(status bool) {
	g.status = status
}

func (g *GptSummaryPlugin) GetName() string {
	return g.name
}

func (g *GptSummaryPlugin) SetName(name string) {
	g.name = name
}

func Dialog2String(dialog *variable.DialogHistory) string {
	return fmt.Sprintf("[%s] %s: %s", util.Timestamp2String(dialog.Time), dialog.Card, parseCQ(dialog.RawMessage))
}

func parseCQ(message string) string {
	message = strings.TrimSpace(message)
	re := regexp.MustCompile(`\[CQ:(\w+)(?:,(\w+)=(.*?))*(,[^]]+)?\]`) // mode, type, id
	if !re.MatchString(message) {
		return message
	}

	matches := re.FindAllStringSubmatch(message, -1)

	for _, match := range matches {
		switch match[1] {
		case "at":
			message = strings.Replace(message, match[0], fmt.Sprintf("[提到%s]", db.Database.GetCardById(match[3])), -1)
		case "reply":
			message = strings.Replace(message, match[0], fmt.Sprintf("[回复%s]", db.Database.GetCardById(match[3])), -1)
		case "face":
			message = strings.Replace(message, match[0], "[表情]", -1)
		case "image":
			message = strings.Replace(message, match[0], "[图片]", -1)
		case "record":
			message = strings.Replace(message, match[0], "[语音]", -1)
		case "forward":
		case "node":
			message = strings.Replace(message, match[0], "[转发消息]", -1)
		case "redbag":
			message = strings.Replace(message, match[0], "[红包]", -1)
		case "music":
		case "share":
			message = strings.Replace(message, match[0], "[分享]", -1)
		default:
			message = strings.Replace(message, match[0], "[其他消息]", -1)
		}
	}
	return message
}
