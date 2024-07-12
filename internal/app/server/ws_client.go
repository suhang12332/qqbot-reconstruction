package server

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/togettoyou/wsc"
	"qqbot-reconstruction/internal/app/commons"
	"qqbot-reconstruction/internal/pkg/log"
	"qqbot-reconstruction/internal/pkg/util"
	"qqbot-reconstruction/internal/pkg/variable"
	"time"
)

var ws *wsc.Wsc

// WS
// @description: ws配置
func Start() {
	// 加载注册表
	path := variable.GetConfigWd() + "plugins.yml"
	result := variable.ReadConfigs(path, &variable.PluginsConfig{})
	// 加载插件
	pluginEngine := initPluginEngine(result)
	// 热加载
	go util.WatchFile(path, func(e fsnotify.Event) {
		cfg := variable.ReadConfigs(path, &variable.PluginsConfig{})
		pluginEngine.Init(cfg, nil)
	})

	done := make(chan bool)
	ws = wsc.New(variable.Urls.Ws)
	// 可自定义配置，不使用默认配置
	ws.SetConfig(&wsc.Config{
		// 写超时
		WriteWait: 30 * time.Second,
		// 支持接受的消息最大长度，默认512字节
		MaxMessageSize: 4096,
		// 最小重连时间间隔
		MinRecTime: 2 * time.Second,
		// 最大重连时间间隔
		MaxRecTime: 60 * time.Second,
		// 每次重连失败继续重连的时间间隔递增的乘数因子，递增到最大重连时间间隔为止
		RecFactor: 1.5,
		// 消息发送缓冲池大小，默认256
		MessageBufferSize: 1024,
	})
	// 设置回调处理
	ws.OnConnected(func() {
		log.Infof("WS链接🤝成功👌")
		// 连接成功后，测试每30秒发送消息
		go func() {
			t := time.NewTicker(30 * time.Second)
			for {
				select {
				case <-t.C:
					err := ws.SendTextMessage("hello")
					if err == wsc.CloseErr {
						return
					}
				}
			}
		}()
	})
	ws.OnConnectError(func(err error) {
		log.Error("WS链接失败: ", err.Error())
	})
	ws.OnDisconnected(func(err error) {
		log.Info("WS断开链接: ", err.Error())
	})
	ws.OnClose(func(code int, text string) {
		log.Infof(fmt.Sprintf("WS关闭: %d,%s", code, text))
		done <- true
	})
	ws.OnSentError(func(err error) {
		log.Error("回复失败: ", err.Error())
	})
	ws.OnTextMessageReceived(func(message string) {
		//receiveMessage(message)
		if rv := pluginEngine.HandleMessage(message); rv != nil {
			SendQMessage(rv)
		}
	})
	go ws.Connect()
	for {
		select {
		case <-done:
			return
		}
	}
}

// SendQMessage
// @description: 发送消息
// @param c websocket指针
// @param message 消息
func SendQMessage(send *string) {
	err := ws.SendTextMessage(*send)
	if err == wsc.CloseErr {
		return
	}
}

func initPluginRegistry(plugins []variable.PluginInfo) *commons.PluginRegistry {
	pluginRegistry := commons.NewPluginRegistry()
	// 插件自动扫描
	pluginRegistry.PluginScanner(plugins)
	return pluginRegistry
}

func initPluginEngine(plugins *variable.PluginsConfig) *commons.PluginEngine {
	pluginEngine := commons.NewPluginEngine()
	pluginEngine.Init(plugins, initPluginRegistry(plugins.Plugins))

	return pluginEngine
}
