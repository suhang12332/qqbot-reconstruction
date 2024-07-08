package api

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/PuerkitoBio/goquery"
    "github.com/imroc/req/v3"
    "io"
    "net/http"
    "qqbot-reconstruction/internal/pkg/log"
    "qqbot-reconstruction/internal/pkg/util"
    "qqbot-reconstruction/internal/pkg/variable"
)

// Fetch
// @description: 用来解析三方接口的响应结果
// @param method 请求接口
// @param url 请求地址
// @param params 请求参数
// @param t 泛型参数
// @return T 泛型返回值
// @returnType 返回类型 html,json
// Error function for logging errors

// Fetch is a generic HTTP request handler
func Fetch[T any](method string, url string, params interface{}, t *T, header map[string]string, returnType string, isBrowser bool, fn func(*goquery.Document) []byte, isEncry bool, en func([]byte) []byte) T {
    var respByte []byte
    var err error
    var result *req.Response
    client := req.C()
    if isBrowser {
        client.ImpersonateSafari()
    }
    r := client.R().SetHeaders(header)
    if params != nil {
        switch v := params.(type) {
        case map[string]string:
            r.SetFormData(v)
        default:
            r.SetBodyJsonString(params.(string))
        }
    }

    switch method {
    case http.MethodGet:
        result, err = r.Get(url)
    case http.MethodPost:
        result, err = r.Post(url)
    default:
        log.Error("无效的HTTP方法", errors.New("invalid HTTP method"))
    }

    if err != nil {
        log.Error("查找资源失败: ", err)
    }

    respByte = result.Bytes()
    if isEncry && en != nil {
        respByte = en(respByte)
    }
    switch returnType {
    case variable.HTML:
        reader := io.NopCloser(bytes.NewReader(respByte))
        doc := util.ParseHtml(reader)
        respByte = fn(doc)
    case variable.JSON:
        err = json.Unmarshal(respByte, t)
        if err != nil {
            log.Error("返回的信息转换struct失败", err)
        }
    default:
        // If the returnType is not recognized, return the raw response bytes
    }

    return *t
}
