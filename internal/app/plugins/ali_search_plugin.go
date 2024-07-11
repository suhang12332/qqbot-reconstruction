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
    status    bool
    whitelist []string
}

func (a *AliSearchPlugin) Execute(receive *message.Receive) *message.Send {
    send := receive.InitSend(true)
    aliInfos := a.query(strings.Split(receive.RawMessage, " ")[1]).Result.Items

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
