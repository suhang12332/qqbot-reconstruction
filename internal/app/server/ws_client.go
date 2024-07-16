package client

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/togettoyou/wsc"
    "qqbot-reconstruction/internal/app/commons"
    ms "qqbot-reconstruction/internal/app/message"
    "qqbot-reconstruction/internal/pkg/log"
    "qqbot-reconstruction/internal/pkg/util"
    "qqbot-reconstruction/internal/pkg/variable"
    "strings"
    "time"
)

var ws *wsc.Wsc

// Start
// @description: ws配置
func Start(engine commons.PluginEngine) {

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
        log.Error("WS链接失败: %s", err.Error())
    })
    ws.OnDisconnected(func(err error) {
        log.Info("WS断开链接: %s", err.Error())
    })
    ws.OnClose(func(code int, text string) {
        log.Infof(fmt.Sprintf("WS关闭: %d,%s", code, text))
        done <- true
    })
    ws.OnSentError(func(err error) {
        log.Error("回复失败: %s", err.Error())
    })
    ws.OnTextMessageReceived(func(message string) {
        util.SafeGo(func() {
            if strings.Contains(message, `post_type":"message"`) {
                rcv := &ms.Receive{}
                if err := json.Unmarshal([]byte(message), rcv); err != nil {
                    log.Errorf("接收消息转换失败!")
                }
                if rv := engine.HandleMessage(rcv); rv != nil {
                    SendQMessage(rv)
                }
            }

            // 实现其他功能
        })


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
    if errors.Is(err, wsc.CloseErr) {
        return
    }
}
