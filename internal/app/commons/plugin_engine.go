package commons

import (
    "encoding/json"
    "fmt"
    "qqbot-reconstruction/internal/app/message"
    "qqbot-reconstruction/internal/pkg/log"
    "qqbot-reconstruction/internal/pkg/util"
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
    if registry != nil {
        e.pluginRegistry = registry
    }

    if e.count > 0 {
        loadCount := 0
        reloadCount := 0
        kwMappings := make(map[string]Plugin)
        for _, info := range plugins.Plugins {
            plugin, loaded := e.pluginRepository[info.Keyword]
            if !loaded {
                // 检查是否修改了keyword
                for k, v := range e.pluginRepository {
                    if v.GetName() == info.Name { // keyword修改
                        plugin = v
                        kwMappings[info.Keyword] = v
                        delete(e.pluginRepository, k)
                        e.count--
                        break
                    }
                } // 插件不存在
            }

            if plugin == nil { // 插件未加载 直接创建实例
                if err := e.loadPlugin(&info); err != nil {
                    log.Error(fmt.Sprintf("加载%s插件失败", info.Name), err)
                    continue
                }
                loadCount++
            } else {
                // 检查插件是否更新
                if isUpdated := updatePlugin(&plugin, &info); isUpdated {
                    reloadCount++
                }
            }
        }

        // 更新keyword
        for k, v := range kwMappings {
            e.pluginRepository[k] = v
        }
        log.Infof("已重载" + strconv.Itoa(reloadCount) + "个插件")
        log.Infof("新加载" + strconv.Itoa(loadCount) + "个插件")

        return
    }

    // 初始化插件引擎：加载插件
    for _, v := range plugins.Plugins {
        v.Status = true
        if err := e.loadPlugin(&v); err != nil {
            log.Error(fmt.Sprintf("加载%s插件失败", v.Name), err)
            continue
        }
    }

    log.Info("已加载的插件个数", strconv.Itoa(int(e.count)))
}

func (e *PluginEngine) loadPlugin(info *variable.PluginInfo) error {
    plugin, err := e.pluginRegistry.CreatePlugin(info.Name, info)
    if err != nil {
        return err
    }
    log.Infof(info.Name + "," + info.Keyword)
    e.pluginRepository[info.Keyword] = plugin
    e.count++
    return nil
}

func (e *PluginEngine) HandleMessage(msg string) *string {
    if strings.Contains(msg, `post_type":"message"`) {
        rcv := &message.Receive{}
        if err :=json.Unmarshal([]byte(msg), rcv); err!= nil {
            log.Errorf("接收消息转换失败!")
        }
        rcv.PrintfMessage()
        split := strings.Split(rcv.RawMessage, " ")
        if plugin, loaded := e.pluginRepository[split[0]]; loaded {
            wl := plugin.GetWhiteList()
			// 校验白名单
            if len(wl) != 0 && !util.In(strconv.Itoa(rcv.UserID), wl) {
                return message.Send2res(rcv.NoPermissionsTips())
            }
            if len(split) > 1 {
				// 校验帮助
                if util.HasKey(split[1], variable.Help) {
                    if rv := plugin.Help(rcv); rv != nil {
                        return message.Send2res(rv)
                    }
                }
				// 校验参数
                if len(plugin.GetArgs())!= 0 && !util.HasKey(split[1],plugin.GetArgs()) {
					return message.Send2res(rcv.NoArgsTips())
                }
            }
            if rv := plugin.Execute(rcv); rv != nil {
                return message.Send2res(rv)
            }
        }
    }
    return nil
}

func (e *PluginEngine) SetStatus(name string, status bool) {
    plugin, loaded := e.pluginRepository[name]
    if !loaded {
        log.Errorf(fmt.Sprintf("插件%s不存在", name))
    }

    if plugin.GetStatus() != status {
        plugin.SetStatus(status)
    }

}

func (e *PluginEngine) SetArgs(name string, args []string, mode string) {
    plugin, loaded := e.pluginRepository[name]
    if !loaded {
        log.Errorf(fmt.Sprintf("插件%s不存在", name))
    }

    wl := plugin.GetArgs()

    switch mode {
    case variable.ADD:
        plugin.SetArgs(util.RemoveRepeatedElement(append(wl, args...)))
    case variable.REMOVE:
        if len(wl) != 0 {
            newWl := util.RemoveElement(args, wl)
            if len(newWl) != len(wl) {
                plugin.SetArgs(newWl)
            }
        }
    }
}

func (e *PluginEngine) SetWhiteList(name string, whiteList []string, mode string) {
    plugin, loaded := e.pluginRepository[name]
    if !loaded {
        log.Errorf(fmt.Sprintf("插件%s不存在", name))
    }

    wl := plugin.GetWhiteList()

    switch mode {
    case variable.ADD:
        plugin.SetWhiteList(util.RemoveRepeatedElement(append(wl, whiteList...)))
    case variable.REMOVE:
        if len(wl) != 0 {
            newWl := util.RemoveElement(whiteList, wl)
            if len(newWl) != len(wl) {
                plugin.SetWhiteList(newWl)
            }
        }
    }
}



func updatePlugin(plugin *Plugin, info *variable.PluginInfo) bool {
    // 更新插件信息
    updated := false
    if (*plugin).GetKeyword() != info.Keyword {
        (*plugin).SetKeyword(info.Keyword)
        updated = true
    }
    if (*plugin).GetStatus() != info.Status {
        (*plugin).SetStatus(info.Status)
        updated = true
    }
    if !util.IsStringArraysEqual((*plugin).GetWhiteList(), info.Whitelist) {
        (*plugin).SetWhiteList(info.Whitelist)
        updated = true
    }
    if !util.IsStringArraysEqual((*plugin).GetArgs(), info.Args) {
        (*plugin).SetArgs(info.Args)
        updated = true
    }
    return updated
}


