package util

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "io"
    "qqbot-reconstruction/internal/pkg/log"
    "qqbot-reconstruction/internal/pkg/variable"
    "strings"
)

// ParseHtml
// @description: 读取http的响应的html结果
// @return: docment结果
func ParseHtml(body io.ReadCloser) *goquery.Document {
    reader, err := goquery.NewDocumentFromReader(body)
    if err != nil {
        log.Infof(err.Error())
    }
    return reader
}

// ParseJson
// @description: 用于构成json响应
// @return: json响应
func ParseJson(data string) string {
    if strings.HasSuffix(data, ",") {
        data = strings.TrimRight(data, ",")
    }
    return fmt.Sprintf("{\"code\": \"200\", \"msg\": \"操作成功\", \"data\": [%s]}", data)
}

// ParseMessage
// @description: 用于格式化消息的转义符
// @return: 结果
func ParseMessage(mes *string) {
    if strings.Contains(*mes, "[") || strings.Contains(*mes, "]") || strings.Contains(*mes, "&") || strings.Contains(*mes, ",") {
        all := strings.ReplaceAll(*mes, "&", "&amp")
        all1 := strings.ReplaceAll(all, "[", "&#91")
        all2 := strings.ReplaceAll(all1, "]", "&#93")
        res := strings.ReplaceAll(all2, ",", "&#44")
        *mes = res
    }
}

func TextParseImg(info string) string {
    return PictureCQ(info)
}

func ParseHelpTips(fun string, desc string, example string, scope string) string {
    return "🙏说明一下用法🤭\n" + fmt.Sprintf("功能: %s\n描述: %s\n例如: %s\n范围: %s\n", fun, desc, example, scope) + "byd 你个老登儿,给我好好看🫵 "
}

func ParseMessageType(info string) string {
    switch info {
    case variable.PRIVATEMESSAGE:
        return variable.PRIVATEMESSAGEZH
    case variable.GROUPMESSGAE:
        return variable.GROUPMESSGAEZH
    default:
        return variable.UNKNOWNMESSAGEZH
    }
}