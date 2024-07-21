package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"qqbot-reconstruction/internal/pkg/log"
	"qqbot-reconstruction/internal/pkg/util"
	"qqbot-reconstruction/internal/pkg/variable"
	"strings"
)

func Magnet(document *goquery.Document) []byte {
	var build strings.Builder
	document.Find(".list .item").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		// 10条中断
		if i == 10 {
			return false
		}
		title := selection.Find(".info .result-title").Text()
		size := selection.Find(".size").Text()
		href, _ := selection.Find(".link").Attr("href")
		r, _, b := Fetch(http.MethodGet, "https://cilisousuo.com"+href, nil, &variable.MagnetData{}, nil, variable.HTML, false, func(document *goquery.Document) []byte {
			magnet := document.Find("input.input-magnet").First().Nodes[0].Attr[3].Val
			data := fmt.Sprintf("{\"title\": \"%s\", \"size\": \"%s\", \"magnet\": \"%s\"},", title, size, (strings.Split(magnet, "&"))[0])
			return []byte(data)
		}, false, nil)
		if !b {
			return false
		}
		marshal, _ := json.Marshal(r)
		build.WriteString(string(marshal))
		return true
	})
	return []byte(util.ParseJson(build.String()))
}

func CiLiLemon(document *goquery.Document) []byte {
	var build strings.Builder
	document.Find(".border-radius").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		// 10条中断
		if i == 10 {
			return false
		}
		title := selection.Find(".panel-heading .highlight").Text()
		size := selection.Find(".panel-footer .small").Text()
		href, _ := selection.Find(".link").Attr("href")
		r, _, b := Fetch(http.MethodGet, "https://cilisousuo.com"+href, nil, &variable.MagnetData{}, nil, variable.HTML, false, func(document *goquery.Document) []byte {
			magnet := document.Find("input.input-magnet").First().Nodes[0].Attr[3].Val
			data := fmt.Sprintf("{\"title\": \"%s\", \"size\": \"%s\", \"magnet\": \"%s\"},", title, size, (strings.Split(magnet, "&"))[0])
			return []byte(data)
		}, false, nil)
		if !b {
			return false
		}
		marshal, _ := json.Marshal(r)
		build.WriteString(string(marshal))
		return true
	})
	return []byte(util.ParseJson(build.String()))
}

func DecodeBase64(bytes []byte) []byte {
	decodeString, err := base64.StdEncoding.DecodeString(string(bytes))
	if err != nil {
		log.Errorf("base64解码失败")
	}
	return decodeString
}
