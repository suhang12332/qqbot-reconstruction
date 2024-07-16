package util

import (
    "bytes"
    "github.com/mdp/qrterminal/v3"
    "image"
    "image/png"
    "os"
    "qqbot-reconstruction/internal/pkg/log"
    "qqbot-reconstruction/internal/pkg/variable"
)

func StdOutQrCode(code []byte, url string) {
    // 创建一个配置对象
    cfg := qrterminal.Config{
        Level:     qrterminal.L, // 二维码的编码级别
        Writer:    os.Stdout,    // 输出到标准输出
        BlackChar: qrterminal.BLACK,
        WhiteChar: qrterminal.WHITE,
    }

    // 将 QR 码图像保存为 PNG 格式
    path := variable.GetConfigWd() + "../cmd/device/"
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        if err :=os.Mkdir(path, os.ModePerm);err!= nil {
            log.Fatal("创建qrcode登陆文件错误: %v", err)
        }
    }
    file, err := os.OpenFile(path+"qrcode.png", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
    if err != nil {
        log.Fatal("创建qrcode登陆文件错误: %v", err)
    }
    defer file.Close()
    buffer := bytes.NewBuffer(code)
    decode, _, err := image.Decode(buffer)
    if err != nil {
        log.Fatal("创建qrcode登陆文件错误: %v", err)
    }
    if err := png.Encode(file, decode); err != nil {
        log.Fatal("创建qrcode登陆文件错误: %v", err)
    }
    qrterminal.GenerateWithConfig(url, cfg)

}
