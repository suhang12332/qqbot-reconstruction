package plugins

import (
	"fmt"
	"net/http"
	"net/url"
	"qqbot-reconstruction/internal/app/message"
	"qqbot-reconstruction/internal/pkg/api"
	"qqbot-reconstruction/internal/pkg/util"
	"qqbot-reconstruction/internal/pkg/variable"
	"strings"
)

type AliSearchPlugin struct {
	name         string
	keyword      string
	status       bool
	subscribable bool
	whitelist    []string
	args         []string
	scope        []string
}

func (a *AliSearchPlugin) SetSubscribable(b bool) {
	a.subscribable = b
}

func (a *AliSearchPlugin) Subscribable() bool {
	return a.subscribable
}

const aliDriver = "阿里云盘搜索结果"

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
	if result, b := a.query(args[1]); b {
		aliInfos := result.Result.Items
		if len(aliInfos) <= 4 {
			return receive.NoResults()
		}
		messages := make([]variable.Messages, len(aliInfos)-4)

		for key, value := range aliInfos {
			if key <= 3 {
				continue
			}
			result := fmt.Sprintf("%s %s", value.Title, value.PageURL)
			messages[key-4] = variable.Messages{
				Type: variable.NODE,
				Data: variable.GroupFowardData{
					Name:    aliDriver,
					Uin:     receive.UserID,
					Content: result,
				},
			}
		}
		send.ForwardMsg(messages)
		return send
	}
	return receive.RequestFail()
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

func (a *AliSearchPlugin) SetScope(scope []string) {
	a.scope = scope
}

func (a *AliSearchPlugin) GetScope() []string {
	return a.scope
}
func (a *AliSearchPlugin) query(info string) (variable.AliResponse, bool) {
	urls := fmt.Sprintf(variable.Urls.Ali, url.QueryEscape(info))
	header := make(map[string]string)
	header["Origin"] = "https://www.upyunso.com"
	header["Referer"] = "https://www.upyunso.com/"
	header["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15"
	header["Host"] = "upapi.juapp9.com"
	_, v, b := api.Fetch(http.MethodGet, urls, nil, &variable.AliResponse{}, header, variable.JSON, false, nil, true, api.DecodeBase64)
	return *v, b
}

func (a *AliSearchPlugin) Help(receive *message.Receive) *message.Send {

	return receive.Tips(util.ParseHelpTips("查询阿里云盘链接", `查询阿里云盘链接,使用 "/云盘 查询的名称"`, `/云盘 蜘蛛侠`, util.ParseHelp(a.scope)))
}
