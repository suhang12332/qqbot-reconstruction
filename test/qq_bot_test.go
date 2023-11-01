package test

import (
	"encoding/json"
	"qqbot-reconstruction/internal/app/server"
	"testing"
)

func LoadConfigs()  {
	
}


func TestLoadConfigs(t *testing.T) {
//	variable.Load()
//	client.WS()
	req := `{"post_type":"message","message_type":"group","time":1698773324,"self_id":3271835508,"sub_type":"normal","message_seq":376515,"raw_mes"点歌 多余的解释","sender":{"age":0,"area":"","card":"偶尔钓鱼（钓小鱼版）","level":"","nickname":"1552899301","role":"admin","sex":"unknown","title":"","user_id":1552899:2099192811,"anonymous":null,"group_id":635129639,"message":"点歌 多余的解释","font":0,"user_id":1552899301}`
	message := server.Receive{}
	json.Unmarshal([]byte(req),&message)
	message.SearchSong("多余的解释")
}