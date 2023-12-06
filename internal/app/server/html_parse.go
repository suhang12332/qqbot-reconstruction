package server

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"qqbot-reconstruction/internal/pkg/handle"
	"qqbot-reconstruction/internal/pkg/util"
	"qqbot-reconstruction/internal/pkg/variable"
	"strings"
)

func magnet(document *goquery.Document) []byte {
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
	return []byte(util.ParseJson(build.String()))
}
