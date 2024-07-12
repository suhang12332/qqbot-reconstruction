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
    status    bool
    whitelist []string
}

func (m *MagnetPlugin) Execute(receive *message.Receive) *message.Send {
    send := receive.InitSend(true)
    data := m.query(strings.Split(receive.RawMessage, " ")[1]).Data
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
func (a *MagnetPlugin) Help(receive *message.Receive) *message.Send {
    send := receive.InitSend(false)
    ((*send).Params.(*variable.SendMsg)).Message = "给傻逼说明一下用法🤭"
    return send
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
