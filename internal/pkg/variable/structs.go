package variable

// Action
// 发送消息的结构体
type Action struct {
	SendPrivateMsg        string
	SendGroupMsg          string
	SendMsg               string
	DeleteMsg             string
	SetGroupKick          string
	SetGroupBan           string
	SetGroupWholeBan      string
	SetGroupAdmin         string
	SetGroupCard          string
	SetGroupName          string
	SetGroupLeave         string
	SetGroupSpecialTitle  string
	SetFriendAddRequest   string
	SetGroupAddRequest    string
	GetLoginInfo          string
	GetStrangerInfo       string
	GetFriendList         string
	GetGroupInfo          string
	GetGroupList          string
	GetGroupMemberInfo    string
	GetGroupMemberList    string
	GetGroupHonorInfo     string
	CanSendImage          string
	CanSendRecord         string
	GetVersionInfo        string
	SetRestart            string
	SendGroupForwardMsg   string
	SendPrivateForwardMsg string
}

// ApiUrl
// @description: 定义Api信息结构体
type ApiUrl struct {
	Ws            string
	Magnet        string
	Ali           string
	Mysql         string
	CloudSong     string
	Reply         string
	Weather       string
	Joke          string
	Hot           string
	Love          string
	TaoShow       string
	CosShow       string
	MoYu          string
	Book          string
	QqSong        string
	OpenAiMessage string
	OPenAiPhoto   string
	Verb          string
}

// Sender
// @description: 用户信息结构体
type Sender struct {
	Age      int    `json:"age"`
	Area     string `json:"area"`
	Card     string `json:"card"`
	Level    string `json:"level"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
	Sex      string `json:"sex"`
	Title    string `json:"title"`
	UserID   int    `json:"user_id"`
}

// SendMsg
// @description: 发送消息参数
type SendMsg struct {
	MessageType string `json:"message_type"`
	UserId      int    `json:"user_id"`
	GroupId     int    `json:"group_id"`
	Message     string `json:"message"`
	AutoEscape  bool   `json:"auto_escape"`
}

// DeleteMsg
// @description: 撤回消息
type DeleteMsg struct {
	MessageID int `json:"message_id"`
}

// SetFriendAddRequest
// @description: 加好友
type SetFriendAddRequest struct {
	Flag    string `json:"flag"`
	Approve bool   `json:"approve"`
	Remark  string `json:"remark"`
}

// SetGroupAddRequest
// @description: 加群
type SetGroupAddRequest struct {
	Flag    string `json:"flag"`
	Type    string `json:"type"`
	Approve bool   `json:"approve"`
	Reason  string `json:"reason"`
}

// CloudSong
// @description: 网易云歌曲结构体
type CloudSong struct {
	Result struct {
		Songs []struct {
			Name        string        `json:"name"`
			ID          int           `json:"id"`
			Position    int           `json:"position"`
			Alias       []interface{} `json:"alias"`
			Status      int           `json:"status"`
			Fee         int           `json:"fee"`
			CopyrightID int           `json:"copyrightId"`
			Disc        string        `json:"disc"`
			No          int           `json:"no"`
			Artists     []struct {
				Name      string        `json:"name"`
				ID        int           `json:"id"`
				PicID     int           `json:"picId"`
				Img1V1ID  int           `json:"img1v1Id"`
				BriefDesc string        `json:"briefDesc"`
				PicURL    string        `json:"picUrl"`
				Img1V1URL string        `json:"img1v1Url"`
				AlbumSize int           `json:"albumSize"`
				Alias     []interface{} `json:"alias"`
				Trans     string        `json:"trans"`
				MusicSize int           `json:"musicSize"`
			} `json:"artists"`
			Album struct {
				Name        string `json:"name"`
				ID          int    `json:"id"`
				IDStr       string `json:"idStr"`
				Type        string `json:"type"`
				Size        int    `json:"size"`
				PicID       int64  `json:"picId"`
				BlurPicURL  string `json:"blurPicUrl"`
				CompanyID   int    `json:"companyId"`
				Pic         int64  `json:"pic"`
				PicURL      string `json:"picUrl"`
				PublishTime int64  `json:"publishTime"`
				Description string `json:"description"`
				Tags        string `json:"tags"`
				Company     string `json:"company"`
				BriefDesc   string `json:"briefDesc"`
				Artist      struct {
					Name      string        `json:"name"`
					ID        int           `json:"id"`
					PicID     int           `json:"picId"`
					Img1V1ID  int           `json:"img1v1Id"`
					BriefDesc string        `json:"briefDesc"`
					PicURL    string        `json:"picUrl"`
					Img1V1URL string        `json:"img1v1Url"`
					AlbumSize int           `json:"albumSize"`
					Alias     []interface{} `json:"alias"`
					Trans     string        `json:"trans"`
					MusicSize int           `json:"musicSize"`
				} `json:"artist"`
				Songs           []interface{} `json:"songs"`
				Alias           []string      `json:"alias"`
				Status          int           `json:"status"`
				CopyrightID     int           `json:"copyrightId"`
				CommentThreadID string        `json:"commentThreadId"`
				Artists         []struct {
					Name      string        `json:"name"`
					ID        int           `json:"id"`
					PicID     int           `json:"picId"`
					Img1V1ID  int           `json:"img1v1Id"`
					BriefDesc string        `json:"briefDesc"`
					PicURL    string        `json:"picUrl"`
					Img1V1URL string        `json:"img1v1Url"`
					AlbumSize int           `json:"albumSize"`
					Alias     []interface{} `json:"alias"`
					Trans     string        `json:"trans"`
					MusicSize int           `json:"musicSize"`
				} `json:"artists"`
				PicIDStr string `json:"picId_str"`
			} `json:"album"`
			Starred         bool          `json:"starred"`
			Popularity      float64       `json:"popularity"`
			Score           int           `json:"score"`
			StarredNum      int           `json:"starredNum"`
			Duration        int           `json:"duration"`
			PlayedNum       int           `json:"playedNum"`
			DayPlays        int           `json:"dayPlays"`
			HearTime        int           `json:"hearTime"`
			Ringtone        interface{}   `json:"ringtone"`
			Crbt            interface{}   `json:"crbt"`
			Audition        interface{}   `json:"audition"`
			CopyFrom        string        `json:"copyFrom"`
			CommentThreadID string        `json:"commentThreadId"`
			RtURL           interface{}   `json:"rtUrl"`
			Ftype           int           `json:"ftype"`
			RtUrls          []interface{} `json:"rtUrls"`
			Copyright       int           `json:"copyright"`
			Rtype           int           `json:"rtype"`
			Rurl            interface{}   `json:"rurl"`
			Mp3URL          string        `json:"mp3Url"`
			Mvid            int           `json:"mvid"`
			BMusic          struct {
				Name        interface{} `json:"name"`
				ID          int         `json:"id"`
				Size        int         `json:"size"`
				Extension   string      `json:"extension"`
				Sr          int         `json:"sr"`
				DfsID       int         `json:"dfsId"`
				Bitrate     int         `json:"bitrate"`
				PlayTime    int         `json:"playTime"`
				VolumeDelta float64     `json:"volumeDelta"`
			} `json:"bMusic"`
			HMusic struct {
				Name        interface{} `json:"name"`
				ID          int         `json:"id"`
				Size        int         `json:"size"`
				Extension   string      `json:"extension"`
				Sr          int         `json:"sr"`
				DfsID       int         `json:"dfsId"`
				Bitrate     int         `json:"bitrate"`
				PlayTime    int         `json:"playTime"`
				VolumeDelta float64     `json:"volumeDelta"`
			} `json:"hMusic"`
			MMusic struct {
				Name        interface{} `json:"name"`
				ID          int         `json:"id"`
				Size        int         `json:"size"`
				Extension   string      `json:"extension"`
				Sr          int         `json:"sr"`
				DfsID       int         `json:"dfsId"`
				Bitrate     int         `json:"bitrate"`
				PlayTime    int         `json:"playTime"`
				VolumeDelta float64     `json:"volumeDelta"`
			} `json:"mMusic"`
			LMusic struct {
				Name        interface{} `json:"name"`
				ID          int         `json:"id"`
				Size        int         `json:"size"`
				Extension   string      `json:"extension"`
				Sr          int         `json:"sr"`
				DfsID       int         `json:"dfsId"`
				Bitrate     int         `json:"bitrate"`
				PlayTime    int         `json:"playTime"`
				VolumeDelta float64     `json:"volumeDelta"`
			} `json:"lMusic"`
		} `json:"songs"`
		SongCount int `json:"songCount"`
	} `json:"result"`
	Code int `json:"code"`
}

// ReceiveMessage
// @description: 接收消息
type ReceiveMessage struct {
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	Time        int    `json:"time"`
	SelfID      int64  `json:"self_id"`
	SubType     string `json:"sub_type"`
	Sender      Sender `json:"sender"`
	MessageID   int    `json:"message_id"`
	UserID      int    `json:"user_id"`
	TargetID    int64  `json:"target_id"`
	Message     string `json:"business"`
	MessageSeq  int    `json:"message_seq"`
	RawMessage  string `json:"raw_message"`
	Font        int    `json:"font"`
	GroupId     int    `json:"group_id"`
}

// SendMessage
// @description: 发送消息
type SendMessage struct {
	Action string `json:"action"`
	Params any    `json:"params"`
	Echo   string `json:"echo"`
}

// AliResult
// @description: 阿里云盘搜索结果结构体
type AliResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List  []FileInfo `json:"list"`
	} `json:"data"`
}

type AliResponse struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Result struct {
		Items []struct {
			Title   string `json:"title"`
			Content []struct {
				Title string `json:"title"`
				Geshi string `json:"geshi"`
				Size  string `json:"size"`
			} `json:"content"`
			PageURL       string `json:"page_url"`
			ID            string `json:"id"`
			Path          string `json:"path"`
			AvailableTime string `json:"available_time"`
			InsertTime    string `json:"insert_time"`
		} `json:"items"`
		Count string `json:"count"`
	} `json:"result"`
}
type FileInfo struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	URL         string      `json:"url"`
	Type        interface{} `json:"type"`
	From        string      `json:"from"`
	Content     interface{} `json:"content"`
	GmtCreate   string      `json:"gmtCreate"`
	GmtShare    interface{} `json:"gmtShare"`
	FileCount   int         `json:"fileCount"`
	CreatorID   string      `json:"creatorId"`
	CreatorName string      `json:"creatorName"`
	FileInfos   []struct {
		Category      interface{} `json:"category"`
		FileExtension interface{} `json:"fileExtension"`
		FileID        string      `json:"fileId"`
		FileName      string      `json:"fileName"`
		Type          string      `json:"type"`
	} `json:"fileInfos"`
}

// Messages
// @description: 群聊消息转发
type Messages struct {
	Type string          `json:"type"`
	Data GroupFowardData `json:"data"`
}

// GroupFowardData
// @description: 群聊消息转发
type GroupFowardData struct {
	Name    string `json:"name"`
	Uin     int    `json:"uin"`
	Content string `json:"content"`
}

// SendGroupForwardMsg
// @description: 群聊消息发送
type SendGroupForwardMsg struct {
	GroupID  int        `json:"group_id"`
	Messages []Messages `json:"messages"`
}

// SendPrivateForwardMsg
// @description: 私聊消息发送
type SendPrivateForwardMsg struct {
	UserID   int        `json:"user_id"`
	Messages []Messages `json:"messages"`
}


// MagnetResult
// @description: 磁力信息结构体
type MagnetResult struct {
	Code string `json:"code"`
	Data []MagnetData `json:"data"`
	Msg string `json:"msg"`
}

type MagnetData struct {
	Title  string `json:"title"`
	Size   string `json:"size"`
	Magnet string `json:"magnet"`
}

