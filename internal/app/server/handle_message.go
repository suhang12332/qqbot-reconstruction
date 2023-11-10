package server

import (
	"fmt"
	"github.com/deatil/go-cryptobin/cryptobin/crypto"
	"qqbot-reconstruction/internal/pkg/util"
	"qqbot-reconstruction/internal/pkg/variable"
	"strings"
)

// songMessage
// @description: 查询歌曲信息
// @param info 歌名
// @return string cq码
func (send *Send) songMessage(info string) *Send {
	song := send.queryCloudSong(info).Result
	if song.SongCount != 0 {
		res := util.MusicCQ(((song.Songs)[0]).ID, "163")
		((*send).Params.(*variable.SendMsg)).Message = res
	}
	return send
}

// aliMessage
// @description: 查询歌曲信息
// @param info 歌名
// @return string cq码
func (send *Send) aliMessage(info string, messageType string) *Send {
	aliInfos := send.queryAliDriver(info).Data.List
	messages := make([]variable.Messages, len(aliInfos))
	for key, value := range aliInfos {
		decrypt := crypto.FromBase64String(value.URL).SetKey("4OToScUFOaeVTrHE").SetIv("9CLGao1vHKqm17Oz").Aes().CBC().PKCS7Padding().Decrypt().ToString()
		result := fmt.Sprintf("%s%s", strings.ReplaceAll(strings.ReplaceAll(value.Name, "<span style=\"color: red;\">", ""), "</span>", ""), decrypt)
		messages[key] = variable.Messages{
			Type: "node",
			Data: variable.GroupFowardData{
				Name:    "阿里云盘搜索结果",
				Uin:     variable.QQ,
				Content: result,
			},
		}
	}
	if messageType == variable.GROUPMESSGAE {
		((*send).Params.(*variable.SendGroupForwardMsg)).Messages = messages
	} else {
		((*send).Params.(*variable.SendPrivateForwardMsg)).Messages = messages
	}

	return send
}
