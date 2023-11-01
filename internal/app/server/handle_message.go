package server

import (
	"qqbot-reconstruction/internal/pkg/util"
	"qqbot-reconstruction/internal/pkg/variable"
)

// songMessage
// @description: 查询歌曲信息
// @param info 歌名
// @return string cq码
func (send *Send)songMessage(info string) *Send {
	song := send.queryCloudSong(info).Result
	if song.SongCount != 0 {
		res := util.MusicCQ(((song.Songs)[0]).ID, "163")
		msg := (*send).Params.(*variable.SendMsg)
		msg.Message = res
	}
	return send
}
