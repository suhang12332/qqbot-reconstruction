package util

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "io"
    "qqbot-reconstruction/internal/pkg/log"
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

func HasKey(key string, slice []string) bool {
    found := false
    for _, value := range slice {
        if value == key {
            found = true
            break
        }
    }
    return found
}
