package client

import (
    "github.com/LagrangeDev/LagrangeGo/client"
    "github.com/LagrangeDev/LagrangeGo/client/auth"
    "github.com/LagrangeDev/LagrangeGo/message"
    "os"
    "os/signal"
    "qqbot-reconstruction/internal/app/commons"
    "qqbot-reconstruction/internal/pkg/cqhttp"
    "qqbot-reconstruction/internal/pkg/log"
    "qqbot-reconstruction/internal/pkg/onebot"
    "qqbot-reconstruction/internal/pkg/util"
    "qqbot-reconstruction/internal/pkg/variable"
    "strconv"
    "syscall"
)

func Login(engine commons.PluginEngine) {
    device := variable.Devices
    path := variable.GetConfigWd() + "../cmd/device/"
    appInfo := auth.AppList[device.Os][device.Version]
    qqclient := client.NewClient(0, device.Sign, appInfo)
    qq, err := strconv.Atoi(device.Qq)
    if err != nil {
        log.Fatal("qq格式错误: %v", err)
    }
    qqclient.UseDevice(auth.NewDeviceInfo(qq))
    data, err := os.ReadFile(path + "sig.bin")
    if err != nil {
        log.Warning("读取sig.bin文件失败: %v", err)
        log.Infof("没有登陆,请扫码....")
        qrcode, url, err := qqclient.FetchQRCodeDefault()
        if err != nil {
            log.Fatal("获取登陆二维码错误: %v", err)
        }
        util.StdOutQrCode(qrcode, url)
    } else {
        sig, err := auth.UnmarshalSigInfo(data, true)
        if err != nil {
            log.Warning("加载sig.bin文件失败: %v", err)
        } else {
            qqclient.UseSig(sig)
        }
    }
    
    err = qqclient.Login("", variable.GetConfigWd()+"../cmd/device/qrcode.png")
    if err != nil {
        log.Fatal("登陆错误: %v", err)
    }
    log.Infof("登陆成功..🎉")
    
    Serve(qqclient,engine)
    
    defer qqclient.Release()

    defer func() {
        data, err = qqclient.Sig().Marshal()
        if err != nil {
            log.Fatal("读取sig.bin文件失败: %v", err)
            return
        }
        err = os.WriteFile(path+"sig.bin", data, 0644)
        if err != nil {
            log.Fatal("更新sig.bin文件失败: %v", err)
            return
        }
        log.Infof("更新sig.bin成功!")
    }()

    mc := make(chan os.Signal, 2)
    signal.Notify(mc, os.Interrupt, syscall.SIGTERM)
    for {
        switch <-mc {
        case os.Interrupt, syscall.SIGTERM:
            return
        }
    }
    
}
func Serve(cli *client.QQClient,engine commons.PluginEngine) {
	cli.PrivateMessageEvent.Subscribe(func(client *client.QQClient, event *message.PrivateMessage) {
        util.SafeGo(func() {
            if rv := engine.HandleMessage(cqhttp.PrivateMessageEvent(client, event)); rv != nil {
                _, err := client.SendPrivateMessage(event.Sender.Uin, []message.IMessageElement{
                    message.NewText("老登儿!"),
                })
                if err != nil {
                    return
                }
            }
        })
    })
    
	cli.GroupMessageEvent.Subscribe(func(client *client.QQClient, event *message.GroupMessage) {
        util.SafeGo(func() {
            if rv := engine.HandleMessage(onebot.ParseGroupMsg(client, event)); rv != nil {
                _, err := client.SendGroupMessage(event.Sender.Uin, []message.IMessageElement{
                   message.NewText("老登儿!"),
                })
                if err != nil {
                    return
                }
            }
        })
    })
//	cli.GroupMemberJoinEvent.Subscribe(handleMemberJoinGroup)
//	cli.GroupMemberLeaveEvent.Subscribe(handleMemberLeaveGroup)
//	cli.GroupMemberLeaveEvent.Subscribe(handleLeaveGroup)
//	cli.GroupInvitedEvent.Subscribe(handleGroupInvitedRequest)
//	cli.GroupRecallEvent.Subscribe(handleGroupMessageRecalled)
//	cli.FriendRecallEvent.Subscribe(handleFriendMessageRecalled)
//	cli.GroupMuteEvent.Subscribe(handleGroupMute)
}

