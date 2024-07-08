package server

import (
    "fmt"
    "net/http"
    "net/url"
    "qqbot-reconstruction/internal/pkg/api"
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
    return api.Fetch(http.MethodGet, urls, nil, &variable.CloudSong{}, header, variable.JSON, false, nil, false, nil)
}

// queryCloudSong
// @description: 获取阿里云盘
// @param info 搜索信息
// @return variable.CloudSong 阿里云盘结构体
func (send *Send) queryAliDriver(info string) variable.AliResponse {
    urls := fmt.Sprintf(variable.Urls.Ali, url.QueryEscape(info))
    header := make(map[string]string)
    header["Origin"] = "https://www.upyunso.com"
    header["Referer"] = "https://www.upyunso.com/"
    header["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15"
    header["Host"] = "upapi.juapp9.com"
    return api.Fetch(http.MethodGet, urls, nil, &variable.AliResponse{}, header, variable.JSON, false, nil, true, decodeBase64)
}

// queryCloudSong
// @description: 获取磁力链接
// @param info 歌名
// @return variable.CloudSong 磁力结构体
func (send *Send) queryMagnet(info string) variable.MagnetResult {
    urls := fmt.Sprintf(variable.Urls.Magnet, url.QueryEscape(info))
    return api.Fetch(http.MethodGet, urls, nil, &variable.MagnetResult{}, nil, variable.HTML, false, magnet, false, nil)
}

