package server

import (
	"fmt"
	"net/http"
	"net/url"
	"qqbot-reconstruction/internal/pkg/handle"
	"qqbot-reconstruction/internal/pkg/variable"
)

// queryCloudSong
// @description: 获取网易云歌曲
// @param info 歌名
// @return variable.CloudSong 歌曲结构体
func (send *Send) queryCloudSong(info string) variable.CloudSong {

	urls := fmt.Sprintf(variable.Urls.CloudSong, url.QueryEscape(info))
	header := make(map[string]string)
	header["Cookie"] = "NMTID=00Oj2vUG0sL7HQJLEpZrByVHMaxRMUAAAGCytb4jw"
	return handle.HttpHandler(http.MethodGet, urls, "", &variable.CloudSong{}, header, variable.JSON, nil)
}

//func(document *goquery.Document) []byte {
//	var build strings.Builder
//	document.Find(".list .item").EachWithBreak(func(i int, selection *goquery.Selection) bool {
//		// 10条中断
//		if i == 10 {
//			return false
//		}
//		title := selection.Find(".info .result-title").Text()
//		size := selection.Find(".size").Text()
//		href, _ := selection.Find(".link").Attr("href")
//		resp, _ := http.Get("https://cilisousuo.com" + href)
//		defer resp.Body.Close()
//		html := util.ParseHtml(resp.Body)
//		magnet := html.Find("input.input-magnet").First().Nodes[0].Attr[3].Val
//		data := fmt.Sprintf("{\"title\": \"%s\", \"size\": \"%s\", \"magnet\": \"%s\"},", title, size, (strings.Split(magnet, "&"))[0])
//		build.WriteString(data)
//
//		return true
//	})
//	return []byte(util.ParseJson(build.String()))
//}