package plugins

import (
	"fmt"
	"net/http"
	"net/url"
	"qqbot-reconstruction/internal/app/message"
	"qqbot-reconstruction/internal/pkg/api"
	"qqbot-reconstruction/internal/pkg/util"
	"qqbot-reconstruction/internal/pkg/variable"
	"strings"
)

type MusicPlugin struct {
	name      string
	keyword   string
	status    bool
	whitelist []string
}

func (m *MusicPlugin) Execute(receive *message.Receive) *message.Send {
	send := message.NewSendMessage(false, receive)

	song := query(strings.Split(receive.RawMessage, " ")[1]).Result
	if song.SongCount != 0 {
		// [CQ:music,type=custom,url=http://baidu.com,audio=http://baidu.com/1.mp3,title=音乐标题]
		res := util.MusicCQ(((song.Songs)[0]).ID, ((song.Songs)[0]).Name)
		((*send).Params.(*variable.SendMsg)).Message = res
		return send
	}

	return nil
}

func (m *MusicPlugin) GetWhiteList() []string {
	return m.whitelist
}

func (m *MusicPlugin) SetWhiteList(whiteList []string) {
	m.whitelist = whiteList
}

func query(info string) variable.CloudSong {
	urls := fmt.Sprintf(variable.Urls.CloudSong, url.QueryEscape(info))
	header := make(map[string]string)
	header["Cookie"] = "NMTID=00Oj2vUG0sL7HQJLEpZrByVHMaxRMUAAAGCytb4jw"
	result := &variable.CloudSong{}
	api.Fetch(http.MethodGet, urls, nil, result, header, variable.JSON, false, nil, false, nil)
	return *result
}
