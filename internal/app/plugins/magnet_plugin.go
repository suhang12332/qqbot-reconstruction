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
    if result, b := m.query(args[1]); b {
        data := result.Data
        messages := make([]variable.Messages, len(data))
        for key, value := range data {
            m := value.Magnet
            replace := strings.Replace(m, `//btsow.click/magnet/detail/hash/`, "", len(m))
            value := fmt.Sprintf("%s %s", fmt.Sprintf("%s %s", value.Title, value.Size), replace)
            util.ParseMessage(&value)
            messages[key] = variable.Messages{
                Type: variable.NODE,
                Data: variable.GroupFowardData{
                    Name:    magnetResult,
                    Uin:     variable.QQ,
                    Content: value,
                },
            }
        }
        ((*send).Params.(*variable.SendPrivateForwardMsg)).Messages = messages
        return send
    }
    return receive.RequestFail()
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
func (m *MagnetPlugin) SetScope(scope []string) {
    m.scope = scope
}

func (m *MagnetPlugin) GetScope() []string {
    return m.scope
}

func (m *MagnetPlugin) query(info string) (variable.MagnetResult, bool) {
    urls := fmt.Sprintf(variable.Urls.Magnet, url.QueryEscape(info))
    _, v, b := api.Fetch(http.MethodGet, urls, nil, &variable.MagnetResult{}, nil, variable.HTML, false, api.Magnet, false, nil)
    return *v, b
}
