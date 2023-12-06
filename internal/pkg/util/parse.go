package util

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"qqbot-reconstruction/internal/pkg/handle"
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



func ParseMagnet(document *goquery.Document) []byte {
	var build strings.Builder
	document.Find(".list .item").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		// 10条中断
		if i == 10 {
			return false
		}
		title := selection.Find(".info .result-title").Text()
		size := selection.Find(".size").Text()
		href, _ := selection.Find(".link").Attr("href")
		resp := handle.HttpHandler(http.MethodGet, "https://cilisousuo.com" + href, nil, &variable.MagnetData{}, nil, variable.HTML, false, func(document *goquery.Document) []byte {
			magnet := document.Find("input.input-magnet").First().Nodes[0].Attr[3].Val
			data := fmt.Sprintf("{\"title\": \"%s\", \"size\": \"%s\", \"magnet\": \"%s\"},", title, size, (strings.Split(magnet, "&"))[0])
			return []byte(data)
		})
		marshal, _ := json.Marshal(resp)
		build.WriteString(string(marshal))
		return true
	})
	return []byte(ParseJson(build.String()))
}