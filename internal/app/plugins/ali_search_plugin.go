package plugins

import (
	"fmt"
	"net/http"
	"net/url"
	"qqbot-reconstruction/internal/app/message"
	"qqbot-reconstruction/internal/pkg/api"
	"qqbot-reconstruction/internal/pkg/variable"
	"strings"
)

type AliSearchPlugin struct {
	name      string
	keyword   string
	status    bool
	whitelist []string
	args  []string
}

func (a *AliSearchPlugin) SetName(name string) {
	a.name = name
}

func (a *AliSearchPlugin) GetKeyword() string {
	return a.keyword
}

func (a *AliSearchPlugin) SetKeyword(keyword string) {
	a.keyword = keyword
}

func (a *AliSearchPlugin) GetName() string {
	return a.name
}

func (a *AliSearchPlugin) GetStatus() bool {
	return a.status
}

func (a *AliSearchPlugin) SetStatus(status bool) {
	a.status = status
}

func (a *AliSearchPlugin) Execute(receive *message.Receive) *message.Send {
    args := strings.Split(receive.RawMessage, " ")
    if len(args) <= 1 {
        return receive.NoArgsTips()
    }
    send := receive.InitSend(true)
	aliInfos := a.query(args[1]).Result.Items
	messages := make([]variable.Messages, len(aliInfos)-3)
	for key, value := range aliInfos {
		if key <= 2 {
			continue
		}
		result := fmt.Sprintf("%s %s", value.Title, value.PageURL)
		messages[key-3] = variable.Messages{
			Type: "node",
			Data: variable.GroupFowardData{
				Name:    "阿里云盘搜索结果",
				Uin:     variable.QQ,
				Content: result,
			},
		}
	}
	((*send).Params.(*variable.SendPrivateForwardMsg)).Messages = messages
	return send
}

func (a *AliSearchPlugin) GetWhiteList() []string {
	return a.whitelist
}

func (a *AliSearchPlugin) SetWhiteList(whiteList []string) {
	a.whitelist = whiteList
}

func (a *AliSearchPlugin) SetArgs(args []string) {
	a.args = args
}

func (a *AliSearchPlugin) GetArgs() []string {
	return a.args
}
func (a *AliSearchPlugin) query(info string) variable.AliResponse {
	urls := fmt.Sprintf(variable.Urls.Ali, url.QueryEscape(info))
	header := make(map[string]string)
	header["Origin"] = "https://www.upyunso.com"
	header["Referer"] = "https://www.upyunso.com/"
	header["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15"
	header["Host"] = "upapi.juapp9.com"
	result := &variable.AliResponse{}
	api.Fetch(http.MethodGet, urls, nil, result, header, variable.JSON, false, nil, true, api.DecodeBase64)
	return *result
}

func (a *AliSearchPlugin) Help(receive *message.Receive) *message.Send {
	return receive.Tips("给傻逼说明一下用法🤭")
}
