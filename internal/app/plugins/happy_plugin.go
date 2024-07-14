package plugins

import (
    "qqbot-reconstruction/internal/app/message"
    "qqbot-reconstruction/internal/pkg/server"
    "qqbot-reconstruction/internal/pkg/util"
    "qqbot-reconstruction/internal/pkg/variable"
    "strconv"
    "strings"
)

type HappyPlugin struct {
    name      string
    keyword   string
    status    bool
    whitelist []string
    args      []string
    scope     []string
}

const happyResult = "写真搜索结果"

func (h *HappyPlugin) SetName(name string) {
    h.name = name
}

func (h *HappyPlugin) GetKeyword() string {
    return h.keyword
}

func (h *HappyPlugin) SetKeyword(keyword string) {
    h.keyword = keyword
}

func (h *HappyPlugin) GetName() string {
    return h.name
}

func (h *HappyPlugin) GetStatus() bool {
    return h.status
}

func (h *HappyPlugin) SetStatus(status bool) {
    h.status = status
}
func (h *HappyPlugin) SetArgs(args []string) {
    h.args = args
}

func (h *HappyPlugin) GetArgs() []string {
    return h.args
}
func (h *HappyPlugin) Execute(receive *message.Receive) *message.Send {
    args := strings.Split(receive.RawMessage, " ")
    if len(args) <= 1 {
        return receive.NoArgsTips()
    }
    length, err := strconv.Atoi(args[1])
    if err != nil {
        return receive.NoArgsTips()
    }
    send := receive.InitSend(true)
    if result, b := server.Infos(length); b {
        messages := make([]variable.Messages, len(result))
        for key, value := range result {

            messages[key] = variable.Messages{
                Type: variable.NODE,
                Data: variable.GroupFowardData{
                    Name:    magnetResult,
                    Uin:     receive.UserID,
                    Content: util.PictureCQ(strings.Replace(value, `https://jmtp.mediavorous.com/storage/article`, `http:127.0.0.1:8081/happy`, 1)),
                },
            }
        }
        send.ForwardMsg(messages)
        return send
    }
    return receive.RequestFail()
}

func (h *HappyPlugin) GetWhiteList() []string {
    return h.whitelist
}
func (h *HappyPlugin) Help(receive *message.Receive) *message.Send {
    return receive.Tips(util.ParseHelpTips("查找指定数量的写真", `查询指定数量的写真,数量越大时间越久;使用 "/色图 数量"`, `/色图 10`, util.ParseHelp(h.scope)))
}
func (h *HappyPlugin) SetScope(scope []string) {
    h.scope = scope
}

func (h *HappyPlugin) GetScope() []string {
    return h.scope
}

func (h *HappyPlugin) SetWhiteList(whiteList []string) {
    h.whitelist = whiteList
}
