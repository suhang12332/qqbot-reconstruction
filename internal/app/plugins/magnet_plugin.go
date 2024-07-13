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
	args  []string
}

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
	data := m.query(args[1]).Data
	messages := make([]variable.Messages, len(data))
	for key, value := range data {
		m := value.Magnet
		replace := strings.Replace(m, `//btsow.click/magnet/detail/hash/`, "", len(m))
		value := fmt.Sprintf("%s %s", fmt.Sprintf("%s %s", value.Title, value.Size), replace)
		util.ParseMessage(&value)
		messages[key] = variable.Messages{
			Type: "node",
			Data: variable.GroupFowardData{
				Name:    "磁力搜索结果",
				Uin:     variable.QQ,
				Content: value,
			},
		}
	}
	((*send).Params.(*variable.SendPrivateForwardMsg)).Messages = messages
	return send
}

func (m *MagnetPlugin) GetWhiteList() []string {
	return m.whitelist
}
func (m *MagnetPlugin) Help(receive *message.Receive) *message.Send {
	return receive.Tips("给傻逼说明一下用法🤭")
}

func (m *MagnetPlugin) SetWhiteList(whiteList []string) {
	m.whitelist = whiteList
}

func (m *MagnetPlugin) query(info string) variable.MagnetResult {
	urls := fmt.Sprintf(variable.Urls.Magnet, url.QueryEscape(info))
	result := &variable.MagnetResult{}
	api.Fetch(http.MethodGet, urls, nil, result, nil, variable.HTML, false, api.Magnet, false, nil)
	return *result
}
