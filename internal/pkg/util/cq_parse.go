package util

import (
	"fmt"
	"strings"
)

func pictureCQ(pic string) string {
	return fmt.Sprintf("[CQ:image,file=%s,type=show,id=40000]", pic)
}

func MusicCQ(id int, types string) string {
	return fmt.Sprintf("[CQ:music,type=%s,id=%d]", types, id)
}

func FaceCQ(faceId int, msg string) string {
	return fmt.Sprintf("[CQ:face,id=%d] %s", faceId, msg)
}
func ALtCQ(userId int, msg string) string {
	all := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(msg, "&", "&amp;"), "[", "&#91;"), "]", "&#93;"), ",", "&#44;")
	return fmt.Sprintf("[CQ:at,qq=%d]%s\n", userId, all)
}
func VidoeCQ(msg string) string {
	return fmt.Sprintf("[CQ:video,file=%s]", msg)
}

func VoiceCQ(data string) string {
	return fmt.Sprintf("[CQ:record,%s;filetype=1&amp;voice_codec=1]", data)
}

func ReplyCQ(mid int, qq int, id int, msg string) string {
	all := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(msg, "&", "&amp;"), "[", "&#91;"), "]", "&#93;"), ",", "&#44;")
	return fmt.Sprintf("[CQ:reply,id=%d][CQ:at,qq=%d][CQ:at,qq=%d]%s", mid, qq, qq, all)
}
