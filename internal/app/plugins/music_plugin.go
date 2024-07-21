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
	args      []string
	scope     []string
}

func (m *MusicPlugin) SetName(name string) {
	m.name = name
}

func (m *MusicPlugin) GetKeyword() string {
	return m.keyword
}

func (m *MusicPlugin) SetKeyword(keyword string) {
	m.keyword = keyword
}

func (m *MusicPlugin) GetName() string {
	return m.name
}

func (m *MusicPlugin) GetStatus() bool {
	return m.status
}

func (m *MusicPlugin) SetStatus(status bool) {
	m.status = status
}

func (m *MusicPlugin) SetArgs(args []string) {
	m.args = args
}

func (m *MusicPlugin) GetArgs() []string {
	return m.args
}

func (m *MusicPlugin) Execute(receive *message.Receive) *message.Send {
	args := strings.Split(receive.RawMessage, " ")
	if len(args) <= 1 {
		return receive.NoArgsTips()
	}
	send := receive.InitSend(false)
	if result, b := m.query(args[1]); b {
		song := result.Result
		if song.SongCount != 0 {
			// [CQ:music,type=custom,url=http://baidu.com,audio=http://baidu.com/1.mp3,title=音乐标题]
			res := util.MusicCQ("163", ((song.Songs)[0]).ID)
			((*send).Params.(*variable.SendMsg)).Message = res
			return send
		}
		return receive.NoResults()
	}
	return receive.RequestFail()
}

func (m *MusicPlugin) GetWhiteList() []string {
	return m.whitelist
}

func (m *MusicPlugin) SetWhiteList(whiteList []string) {
	m.whitelist = whiteList
}
func (m *MusicPlugin) SetScope(scope []string) {
	m.scope = scope
}

func (m *MusicPlugin) GetScope() []string {
	return m.scope
}

func (m *MusicPlugin) query(info string) (variable.CloudSong, bool) {
	urls := fmt.Sprintf(variable.Urls.CloudSong, url.QueryEscape(info))
	header := make(map[string]string)
	header["Cookie"] = "NMTID=00Oj2vUG0sL7HQJLEpZrByVHMaxRMUAAAGCytb4jw"
	_, v, b := api.Fetch(http.MethodGet, urls, nil, &variable.CloudSong{}, header, variable.JSON, false, nil, false, nil)
	return *v, b
}
func (m *MusicPlugin) Help(receive *message.Receive) *message.Send {
	return receive.Tips("给傻逼说明一下用法🤭")
}
