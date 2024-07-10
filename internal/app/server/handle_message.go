package server

import (
    "fmt"
    "qqbot-reconstruction/internal/pkg/util"
    "qqbot-reconstruction/internal/pkg/variable"
    "strings"
)

// songMessage
// @description: 查询歌曲信息
// @param info 歌名
// @return string cq码
func (send *Send) songMessage(receive *Receive) {
    song := send.queryCloudSong(strings.Split(receive.RawMessage, " ")[1]).Result
    if song.SongCount != 0 {
        // [CQ:music,type=custom,url=http://baidu.com,audio=http://baidu.com/1.mp3,title=音乐标题]
        res := util.MusicCQ(((song.Songs)[0]).ID, ((song.Songs)[0]).Name)
        ((*send).Params.(*variable.SendMsg)).Message = res
    }
    send.sendMessage()
}

// aliMessage
// @description: 查询歌曲信息
// @param info 歌名
// @return string cq码
func (send *Send) aliMessage(receive *Receive) {
    aliInfos := send.queryAliDriver(strings.Split(receive.RawMessage, " ")[1]).Result.Items

    messages := make([]variable.Messages, len(aliInfos)-3)
    for key, value := range aliInfos {
        if key <= 2 {
            continue
        }
        result := fmt.Sprintf("%s %s", value.Title, value.PageURL)
        messages[key-3] = variable.Messages{
            Type: "node",
            Data: variable.GroupFowardData{
                Name:    "阿里云盘搜索结果",
                Uin:     variable.QQ,
                Content: result,
            },
        }
    }
    ((*send).Params.(*variable.SendPrivateForwardMsg)).Messages = messages
    send.sendMessage()
}

func (send *Send) happyMessage(receive *Receive) {
    //    data := send.queryMagnet(strings.Split(receive.RawMessage, " ")[1]).Data
    //    messages := make([]variable.Messages, len(data))
    //    for key, value := range data {
    //        m := value.Magnet
    //        replace := strings.Replace(m, `//btsow.click/magnet/detail/hash/`, "", len(m))
    //        value := fmt.Sprintf("%s %s", fmt.Sprintf("%s %s", value.Title, value.Size), replace)
    //        util.ParseMessage(&value)
    //        messages[key] = variable.Messages{
    //            Type: "node",
    //            Data: variable.GroupFowardData{
    //                Name:    "磁力搜索结果",
    //                Uin:     variable.QQ,
    //                Content: value,
    //            },
    //        }
    //    }
    //    infos := Infos()
    infos := []string{
        "https://jmtp.mediavorous.com/storage/article/8339/6389e605e983a.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6063956e.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e605ca89f.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6054a3c0.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6050053e.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e604cd29a.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6056d445.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e605804ab.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6063032f.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e606db298.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e60698051.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e60659889.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e606b2354.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e606ceb13.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6071b701.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e626381aa.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6267f63a.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e625d1eb3.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e625c8919.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e62585c6e.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e625478fa.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e626a90eb.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e626bd454.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e626b16d1.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e625cbd62.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6269a8f7.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6279a572.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e626869bf.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e626af6fb.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e626a65d3.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e63366132.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e63392567.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e633a98c0.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e633e1b6b.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6343b69d.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e634aedb5.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e63497529.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e63470693.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6345e209.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6342258f.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e634808cf.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e63416a5c.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e634566a8.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e63492761.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e633dea7a.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e642740ca.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e642ab627.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e6429e28d.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e642a2f1e.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e642a309b.jpg",
        "https://jmtp.mediavorous.com/storage/article/8339/6389e64298a7d.jpg"}
    for _, value := range infos {
        result := util.PictureCQ(strings.Replace(value, "https://jmtp.mediavorous.com/storage/article", "http://127.0.0.1:8081/happy", 1))
        ((*send).Params.(*variable.SendMsg)).Message = result
        send.sendMessage()
    }
}

func (send *Send) magnetMessage(receive *Receive) {
    data := send.queryMagnet(strings.Split(receive.RawMessage, " ")[1]).Data
    messages := make([]variable.Messages, len(data))
    for key, value := range data {
        m := value.Magnet
        replace := strings.Replace(m, `//btsow.click/magnet/detail/hash/`, "", len(m))
        value := fmt.Sprintf("%s %s", fmt.Sprintf("%s %s", value.Title, value.Size), replace)
        util.ParseMessage(&value)
        messages[key] = variable.Messages{
            Type: "node",
            Data: variable.GroupFowardData{
                Name:    "磁力搜索结果",
                Uin:     variable.QQ,
                Content: value,
            },
        }
    }
    ((*send).Params.(*variable.SendPrivateForwardMsg)).Messages = messages
    send.sendMessage()
}

//func (send *Send) getAliDriverUrl(value variable.FileInfo) string {
//	decrypt := crypto.FromBase64String(value.URL).SetKey("4OToScUFOaeVTrHE").SetIv("9CLGao1vHKqm17Oz").Aes().CBC().PKCS7Padding().Decrypt().ToString()
//	result := fmt.Sprintf("%s%s", strings.ReplaceAll(strings.ReplaceAll(value.Name, "<span style=\"color: red;\">", ""), "</span>", ""), decrypt)
//	return result
//}
