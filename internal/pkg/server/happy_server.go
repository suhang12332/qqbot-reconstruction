package server

import (
    "crypto/md5"
    "encoding/json"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/deatil/go-cryptobin/cryptobin/crypto"
    "github.com/go-resty/resty/v2"
    "io"
    "math/rand"
    "net/http"
    "qqbot-reconstruction/internal/pkg/log"
    "strconv"
    "strings"
    "time"
)

var (
    imageUrl string
    token    string
    userKey  string
)

func StartHappyServer() {
    http.HandleFunc("/happy/", func(writer http.ResponseWriter, request *http.Request) {
        urls := strings.Replace(request.URL.Path, "/happy", "https://jmtp.mediavorous.com/storage/article", 1)
        picture := getPicture(urls)
        writer.Write(picture)
    })
    http.ListenAndServe(":8081", nil)

}

type Lists struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    struct {
        Title    string `json:"title"`
        Author   string `json:"author"`
        Thumb    string `json:"thumb"`
        End      int    `json:"end"`
        Chapters []struct {
            ID          int    `json:"id"`
            Title       string `json:"title"`
            Chapter     int    `json:"chapter"`
            Description string `json:"description"`
            PublishedAt string `json:"published_at"`
        } `json:"chapters"`
    } `json:"data"`
}
type Info struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    struct {
        AudioFile        string   `json:"audio_file"`
        AudioHls         string   `json:"audio_hls"`
        Author           string   `json:"author"`
        Buy              int      `json:"buy"`
        Buys             int      `json:"buys"`
        Chapter          int      `json:"chapter"`
        CollectionNumber int      `json:"collection_number"`
        Content          string   `json:"content"`
        CreatedAt        string   `json:"created_at"`
        DeletedAt        string   `json:"deleted_at"`
        Description      string   `json:"description"`
        DownloadCount    int      `json:"download_count"`
        DownloadImages   int      `json:"download_images"`
        ID               int      `json:"id"`
        IsCollected      int      `json:"is_collected"`
        IsLike           int      `json:"is_like"`
        LikeNumber       int      `json:"like_number"`
        ModelID          int      `json:"model_id"`
        PublishedAt      string   `json:"published_at"`
        ReplyCounts      int      `json:"reply_counts"`
        SeriesCategory   int      `json:"series_category"`
        SeriesID         int      `json:"series_id"`
        Sort             int      `json:"sort"`
        SrtPath          string   `json:"srt_path"`
        Status           int      `json:"status"`
        Subtitle         string   `json:"subtitle"`
        Tags             []string `json:"tags"`
        Thumb            string   `json:"thumb"`
        Thumbnail        string   `json:"thumbnail"`
        Title            string   `json:"title"`
        UpdatedAt        string   `json:"updated_at"`
        Video            string   `json:"video"`
        VideoContent     string   `json:"video_content"`
        VideoCover       string   `json:"video_cover"`
        VideoFile        string   `json:"video_file"`
        VideoHls         string   `json:"video_hls"`
        VideoLength      string   `json:"video_length"`
        VideoName        string   `json:"video_name"`
        Views            int      `json:"views"`
        Vol              int      `json:"vol"`
    } `json:"data"`
}
type AlbumList struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    []struct {
        ID        int    `json:"id"`
        Title     string `json:"title"`
        Author    string `json:"author"`
        Thumb     string `json:"thumb"`
        Thumbnail string `json:"thumbnail"`
        End       int    `json:"end"`
        FirstAt   string `json:"first_at"`
        LatestAt  string `json:"latest_at"`
        LikeTotal string `json:"like_total"`
    } `json:"data"`
}
type BootStrap struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    struct {
        User struct {
            ID                int    `json:"id"`
            Username          string `json:"username"`
            Nickname          string `json:"nickname"`
            Sex               string `json:"sex"`
            Avatar            string `json:"avatar"`
            Mobile            string `json:"mobile"`
            UserKey           string `json:"user_key"`
            Email             string `json:"email"`
            Integral          int    `json:"integral"`
            Status            string `json:"status"`
            CommentStatus     string `json:"comment_status"`
            Desc              string `json:"desc"`
            DeviceID          int    `json:"device_id"`
            Code              string `json:"code"`
            LoginIP           string `json:"login_ip"`
            LoggedAt          string `json:"logged_at"`
            CreatedAt         string `json:"created_at"`
            UpdatedAt         string `json:"updated_at"`
            DeletedAt         string `json:"deleted_at"`
            IsVisitor         int    `json:"is_visitor"`
            IsPassExam        int    `json:"is_pass_exam"`
            CommentCount      int    `json:"comment_count"`
            Token             string `json:"token"`
            SexStatus         int    `json:"sex_status"`
            SexNextUpdateTime string `json:"sex_next_update_time"`
            SexLastTime       string `json:"sex_last_time"`
            SexUpdateDay      int    `json:"sex_update_day"`
            AfterDaysExam     int    `json:"after_days_exam"`
        } `json:"user"`
        ForwardLink []string `json:"forward_link"`
        ForwardText string   `json:"forward_text"`
        QrCode      string   `json:"qr_code"`
        Email       string   `json:"email"`
        FirstDay    string   `json:"first_day"`
        Share       struct {
            Title string `json:"title"`
            Desc  string `json:"desc"`
            URL   string `json:"url"`
            Icon  string `json:"icon"`
        } `json:"share"`
        ServerTime            string `json:"server_time"`
        QuizStatus            int    `json:"quiz_status"`
        BusinessContact       string `json:"business_contact"`
        ArticleAdsGapInterval string `json:"article_ads_gap_interval"`
        Domain                struct {
            API                  []string `json:"api"`
            Video                []string `json:"video"`
            VideoHls             []string `json:"video_hls"`
            Image                []string `json:"image"`
            ImageLow             []string `json:"image_low"`
            Audio                []string `json:"audio"`
            AudioDownload        []string `json:"audio_download"`
            T08                  []string `json:"t08"`
            Resource             []string `json:"resource"`
            Webp                 []string `json:"webp"`
            FeedbackDomain       []string `json:"feedback_domain"`
            LaunchPageURL        []string `json:"launch_page_url"`
            VipDetails           []string `json:"vip_details"`
            FlutterImg           []string `json:"flutter_img"`
            FlutterAudioDownload []string `json:"flutter_audio_download"`
            FlutterVideoDownload []string `json:"flutter_video_download"`
            FlutterEncImg        []string `json:"flutter_enc_img"`
            FlutterVideoPlay     []string `json:"flutter_video_play"`
            FlutterAPI           []string `json:"flutter_api"`
        } `json:"domain"`
        Switches struct {
            BootstrapAds int `json:"bootstrap_ads"`
            PopupAds     int `json:"popup_ads"`
        } `json:"switches"`
        Country      string `json:"country"`
        HotKeywords  string `json:"hot_keywords"`
        SmsContent   string `json:"sms_content"`
        ExpirePrompt struct {
            Switch40 int    `json:"switch_40"`
            Switch60 int    `json:"switch_60"`
            Tips40   string `json:"tips_40"`
            Tips60   string `json:"tips_60"`
        } `json:"expire_prompt"`
        CommentReview struct {
            Switch int      `json:"switch"`
            Rules  []string `json:"rules"`
        } `json:"comment_review"`
        UnpaidNoticeMessage string `json:"unpaid_notice_message"`
        Vip                 struct {
            SkipAd int `json:"skip_ad"`
        } `json:"vip"`
        Comment struct {
            Switch    int    `json:"switch"`
            Countdown int    `json:"countdown"`
            Title     string `json:"title"`
            Message   string `json:"message"`
        } `json:"comment"`
        FeedbackRestrictions struct {
            ImageLimit int `json:"image_limit"`
            VideoLimit int `json:"video_limit"`
        } `json:"feedback_restrictions"`
        ShowType            int `json:"show_type"`
        AdvertInterval      int `json:"advert_interval"`
        VideoAdvertInterval int `json:"video_advert_interval"`
        IsShowWebp          int `json:"is_show_webp"`
        BehaviorVerifyOpen  int `json:"behavior_verify_open"`
    } `json:"data"`
}

func init() {
    s := "kXnzSuPkvkVcX7LZyEYw7CGTnQfpmyci2tbSVkRMH//SRVnwLed6Q+fS61RJ9pUy/pTF79087CJMgBpzUdw5hSRDKvebhp1XzoLqx9FNuTnj4umklvp9uob8HlFm/6/TPYo4Xa9ssJFRIEsp+UwqjrhXkMvUUv8YzzOI8Iul+Nn/JU4/4Wc6VgrvhgdfnzhOE0/Di0wimDXJfwnDwl2qNry4JN0XPjy2qrDgdraWWUyZWxKeLv6IOFiqHGFy1RVnK5gDM7W4rLdPmx91XTy57XNsdTpg6RseySoLri69sTGSzXsDB5qVQnQvAiDOBeKnSlvLs4kQ9UIa7fhk7AcXB76uodOJDI/nnG1knu7Fpqs="
    bytes := post(s, "https://api.10cb1c.com/v2.5/bootstrap")
    strap := BootStrap{}
    json.Unmarshal(bytes, &strap)
    user := strap.Data.User
    token = user.Token
    userKey = user.UserKey
    imageUrl = (strap.Data.Domain.Image)[0]
}

func parseHtml(s string) []string {
    is := make([]string, 0)
    reader, _ := goquery.NewDocumentFromReader(strings.NewReader(s))
    reader.Find("img").Each(func(i int, selection *goquery.Selection) {
        src, _ := selection.Attr("src")
        all := strings.ReplaceAll(src, `\"`, "")
        s := fmt.Sprintf("https://jmtp.mediavorous.com%s", strings.ReplaceAll(all, `\`, ""))
        is = append(is, s)
    })
    return is
}
func post(body string, url string) []byte {
    client := resty.New()
    res, err := client.R().SetBody(body).SetHeaders(getHeader(body)).Post(url)
    if err != nil {
        log.Error("%v",err)
    }
    return decrypt(res.Body())
}
func getInfo(id int, name string) Info {
    req := fmt.Sprintf("id=%d", id)
    res := encrypt(req)
    bytes := post(res, "https://api.10cb1c.com/v2.5/article/detail")
    albumList := Info{}
    json.Unmarshal(bytes, &albumList)
    return albumList
}
func getListResult(id int, name string) Lists {
    req := fmt.Sprintf("series_id=%d", id)
    res := encrypt(req)
    bytes := post(res, "https://api.10cb1c.com/v2.5/series/chapters")
    albumList := Lists{}
    json.Unmarshal(bytes, &albumList)
    return albumList
}

func getList(i int) AlbumList {
    body := fmt.Sprintf("type=0&page=%d&size=31&sort=published_at", i)
    res := encrypt(body)
    bytes := post(res, "https://api.10cb1c.com/v2.5/series/album/list")
    albumList := AlbumList{}
    json.Unmarshal(bytes, &albumList)
    return albumList
}
func encrypt(s string) string {
    return crypto.FromString(s).SetKey("l*bv%Ziq000Biaog").SetIv("8597506002939249").Aes().CBC().PKCS7Padding().Encrypt().ToBase64String()
}
func decrypt(s []byte) []byte {
    return crypto.FromBase64String(string(s)).SetKey("l*bv%Ziq000Biaog").SetIv("8597506002939249").Aes().CBC().PKCS7Padding().Decrypt().ToBytes()
}

func getHeader(s string) map[string]string {

    formatInt := strconv.FormatInt(time.Now().Unix(), 10)
    m := map[string]string{
        "uuid":            "5e37046d-b683-3baa-8504-e9115ddc0830",
        "timestamp":       formatInt,
        "ip":              "0.0.0.0",
        "user-key":        userKey,
        "platform":        "1",
        "sign":            getSign(formatInt),
        "app-version":     "2.5.2",
        "Content-Type":    "application/x-www-form-urlencoded; charset=utf-8",
        "Host":            "api.10cb1c.com",
        "Connection":      "Keep-Alive",
        "Accept-Encoding": "gzip",
        "token":           token,
        "Content-Length":  strconv.Itoa(len(s)),
        "User-Agent":      "okhttp/4.3.1",
    }
    return m

}

func getSign(timex string) string {

    s := fmt.Sprintf("0.0.0.0.1.%s..5e37046d-b683-3baa-8504-e9115ddc0830", timex)
    str := fmt.Sprintf("%x", md5.Sum([]byte(s)))
    sprintf := fmt.Sprintf("%s%s", str, "m4n2hjPeYWkD6tFpqKF^3HO^h24P@idT")
    sign := fmt.Sprintf("%x", md5.Sum([]byte(sprintf)))
    return sign
}
func Infos(length int) ([]string,bool) {
    results := make([]string, 0)
    list := getList(rand.Intn(5)+1)
    if len(list.Data) == 0 {
        return nil,false
    }
    data := list.Data[rand.Intn(len(list.Data))]
    l := getListResult(data.ID, data.Title)
    if len(l.Data.Chapters)== 0 {
        return nil,false
    }
    v := l.Data.Chapters[rand.Intn(len(l.Data.Chapters))]
    info := getInfo(v.ID, v.Title)
    if info.Data.Content == "" {
        return nil,false
    }
    srcs := parseHtml(info.Data.Content)
    if len(srcs)== 0 {
        return nil,false
    }
    for k, l := range srcs {
        results = append(results, l)
        if k == length {
            break
        }
    }
    return results,true

}
func getPicture(src string) []byte {
    resp, _ := http.Get(src)
    defer resp.Body.Close()
    all, _ := io.ReadAll(resp.Body)
    s := crypto.FromBytes(all).SetKey("saIZXc4yMvq0Iz56").SetIv("kbJYtBJUECT0oyjo").Aes().CBC().PKCS7Padding().Decrypt().ToBytes()
    return s
}
