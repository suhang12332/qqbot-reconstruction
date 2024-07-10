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
    infos := Infos()
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
