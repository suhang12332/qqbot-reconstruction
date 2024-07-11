package commons

import (
    "encoding/json"
    "fmt"
    "qqbot-reconstruction/internal/app/message"
    "qqbot-reconstruction/internal/pkg/log"
    "qqbot-reconstruction/internal/pkg/variable"
    "strconv"
    "strings"
)

type PluginEngine struct {
    // 插件注册表
    pluginRegistry *PluginRegistry
    // 插件实例	keyword-object
    pluginRepository map[string]Plugin
    // 加载的插件计数
    count int32
}

func NewPluginEngine() *PluginEngine {
    return &PluginEngine{
        pluginRegistry:   nil,
        pluginRepository: make(map[string]Plugin),
        count:            0,
    }
}

func (e *PluginEngine) Init(plugins *variable.PluginsConfig, registry *PluginRegistry) {
    e.pluginRegistry = registry

    // 加载插件
    for _, v := range plugins.Plugins {
        plugin, err := e.pluginRegistry.CreatePlugin(v.Name, v.Whitelist)
        if err != nil {
            log.Error(fmt.Sprintf("加载%s插件失败", v.Name), err)
            continue
        }
        log.Infof(v.Name + "," + v.Keyword)
        e.pluginRepository[v.Keyword] = plugin
        e.count++
    }

    log.Info("已加载的插件个数", strconv.Itoa(int(e.count)))
}

func (e *PluginEngine) HandleMessage(msg string) *string {
    if strings.Contains(msg, `post_type":"message"`) {
        rcv := parseMessage(msg, &message.Receive{})
        rcv.PrintfMessage()

        split := strings.Split(rcv.RawMessage, " ")
        if plugin, loaded := e.pluginRepository[split[0]]; loaded {
            wl := plugin.GetWhiteList()
            if len(wl) != 0 && !in(strconv.Itoa(rcv.UserID), wl) {
                return nil
            }
            if rv := plugin.Execute(rcv); rv != nil {
                return send2res(rv)
            }
        }
    }
    return nil
}

func parseMessage[T any](message string, t *T) *T {
    bytes := []byte(message)
    json.Unmarshal(bytes, t)
    return t
}

func send2res(send *message.Send) *string {
    marshal, err := json.Marshal(send)
    switch send.Action {
    case variable.Actions.SendMsg:
        log.Info("回复消息: ", strings.ReplaceAll((*send).Params.(*variable.SendMsg).Message, "\n", "\t"))
    }
    if err != nil {
        log.Error("消息回复失败: ", err)
    }
    result := string(marshal)
    log.Infof(result)

    return &result
}

func in(target string, arr []string) bool {
    for _, element := range arr {
        if target == element {
            return true
        }
    }
    return false
}
