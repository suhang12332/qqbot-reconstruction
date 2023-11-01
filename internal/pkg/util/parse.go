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