package handle

import (
	"crypto/tls"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"qqbot-reconstruction/internal/pkg/log"
	"qqbot-reconstruction/internal/pkg/util"
	"qqbot-reconstruction/internal/pkg/variable"
	"strings"
)

// HttpHandler
// @description: 用来解析三方接口的响应结果
// @param method 请求接口
// @param url 请求地址
// @param params 请求参数
// @param t 泛型参数
// @return T 泛型返回值
// @returnType 返回类型 html,json
func HttpHandler[T any](method string, url string, params string, t *T, header map[string]string, returnType string, fn func(*goquery.Document) []byte) T {
	var all []byte
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest(method, url, strings.NewReader(params))
	for v, k := range header {
		req.Header.Set(v, k)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("查找资源失败: ", err)
	}
	if resp.StatusCode != 200 || resp == nil {
		log.Errorf("资源请求失败")
	}
	defer resp.Body.Close()
	if returnType == variable.HTML {
		all = fn(util.ParseHtml(resp.Body))
	} else {
		all, _ = io.ReadAll(resp.Body)
	}
	err = json.Unmarshal(all, t)
	if err != nil {
		log.Error("返回的信息转换struct失败", err)
	}
	return *t
}
