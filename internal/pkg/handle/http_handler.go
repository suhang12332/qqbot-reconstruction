package handle

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"io"
	"net/http"
	"os/exec"
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
func HttpHandler[T any](method string, url string, params map[string]string, t *T, header map[string]string, returnType string, isBrowser bool, fn func(*goquery.Document) []byte) T {
	var respByte []byte
	var err error

	if !isBrowser {
		var result *resty.Response
		client := resty.New().R().SetFormData(params).SetHeaders(header)
		switch method {
		case http.MethodGet:
			result, err = client.Post(url)
			break
		case http.MethodPost:
			result, err = client.Get(url)
			break
		}
		if err != nil {
			log.Error("查找资源失败: ", err)
		} else {
			respByte = []byte(result.String())
		}

	} else {
		respByte, err = exec.Command("curl", url).CombinedOutput()
	}

	if returnType == variable.HTML {
		reader := io.NopCloser(bytes.NewReader(respByte))
		respByte = fn(util.ParseHtml(reader))
	}
	err = json.Unmarshal(respByte, t)
	if err != nil {
		log.Error("返回的信息转换struct失败", err)
	}
	return *t
}
