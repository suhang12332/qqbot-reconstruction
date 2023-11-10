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
// queryCloudSong
// @description: 获取网易云歌曲
// @param info 歌名
// @return variable.CloudSong 歌曲结构体
func (send *Send) queryAliDriver(info string) variable.AliResult {
	urls := fmt.Sprintf(variable.Urls.Ali, url.QueryEscape(info), 1)
	header := make(map[string]string)
	header["Cookie"] = "__51vcke__JkIGvjjs25ETn0wz=848f658e-f6b7-51ce-97be-dcb01d68cdce; __51vuft__JkIGvjjs25ETn0wz=1699359098434; __vtins__JkIGvjjs25ETn0wz=%7B%22sid%22%3A%20%22a15e7874-b35d-5893-a7d7-c11f9c9006f8%22%2C%20%22vd%22%3A%205%2C%20%22stt%22%3A%2085318%2C%20%22dr%22%3A%2034750%2C%20%22expires%22%3A%201699360983750%2C%20%22ct%22%3A%201699359183750%7D; satoken=a84dff77-1d6a-4488-8462-4cae3921a4f6; JSESSIONID=522F2CFC7F65C8DFC183032D89D9E89B; cf_clearance=8FBxddj53GJkqVd2wG.5tKZeiFFkPwnC0Xx9apFaOqE-1699359100-0-1-6a125ba.d841e4d2.51a07a5a-150.2.1699359100; __51uvsct__JkIGvjjs25ETn0wz=1"
	return handle.HttpHandler(http.MethodGet, urls, "", &variable.AliResult{}, header, variable.JSON, nil)
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