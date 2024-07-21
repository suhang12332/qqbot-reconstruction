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

type MagnetPlugin struct {
	name      string
	keyword   string
	status    bool
	whitelist []string
	args      []string
	scope     []string
}

const magnetResult = "磁力搜索结果"

func (m *MagnetPlugin) SetName(name string) {
	m.name = name
}

func (m *MagnetPlugin) GetKeyword() string {
	return m.keyword
}

func (m *MagnetPlugin) SetKeyword(keyword string) {
	m.keyword = keyword
}

func (m *MagnetPlugin) GetName() string {
	return m.name
}

func (m *MagnetPlugin) GetStatus() bool {
	return m.status
}

func (m *MagnetPlugin) SetStatus(status bool) {
	m.status = status
}
func (m *MagnetPlugin) SetArgs(args []string) {
	m.args = args
}

func (m *MagnetPlugin) GetArgs() []string {
	return m.args
}
func (m *MagnetPlugin) Execute(receive *message.Receive) *message.Send {
	args := strings.Split(receive.RawMessage, " ")
	if len(args) <= 1 {
		return receive.NoArgsTips()
	}
	send := receive.InitSend(true)
	if result, b := m.Query(args[1]); b {
		data := result.Rows
		messages := make([]variable.Messages, len(data))
		for key, value := range data {
			util.ParseCiliDanDan(&(value.Name))
			magnet := fmt.Sprintf("magnet:?xt=urn:btih:%s", value.InfoHash)
			value := fmt.Sprintf("%s %s", value.Name, magnet)
			util.ParseMessage(&value)
			messages[key] = variable.Messages{
				Type: variable.NODE,
				Data: variable.GroupFowardData{
					Name:    magnetResult,
					Uin:     receive.UserID,
					Content: value,
				},
			}
		}
		send.ForwardMsg(messages)
		return send
	}
	return receive.RequestFail()
}

func (m *MagnetPlugin) GetWhiteList() []string {
	return m.whitelist
}
func (m *MagnetPlugin) Help(receive *message.Receive) *message.Send {
	return receive.Tips(util.ParseHelpTips("查询磁力链接", `查询磁力链接,使用 "/资源 查询的名称"`, `/资源 蜘蛛侠`, util.ParseHelp(m.scope)))
}

func (m *MagnetPlugin) SetWhiteList(whiteList []string) {
	m.whitelist = whiteList
}
func (m *MagnetPlugin) SetScope(scope []string) {
	m.scope = scope
}

func (m *MagnetPlugin) GetScope() []string {
	return m.scope
}

func (m *MagnetPlugin) Query(info string) (variable.MagnetResponse, bool) {
	escape := url.QueryEscape(info)
	urls := fmt.Sprintf(variable.Urls.Magnet, escape)
	_, v, b := api.Fetch(http.MethodGet, urls, nil, &variable.MagnetResponse{}, map[string]string{
		"Host":    "www.cilidandan5.com",
		"Referer": fmt.Sprintf("https://www.cilidandan5.com/so/search/%s", escape),
	}, variable.JSON, false, nil, false, nil)
	return *v, b
}
