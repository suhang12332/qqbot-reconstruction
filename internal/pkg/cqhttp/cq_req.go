package cqhttp

import (
    hex2 "encoding/hex"
    "fmt"
    "github.com/LagrangeDev/LagrangeGo/client"
    "github.com/LagrangeDev/LagrangeGo/message"
    "hash/crc32"
    message2 "qqbot-reconstruction/internal/app/message"
    "strconv"

    "strings"
)

func PrivateMessageEvent(client *client.QQClient, m *message.PrivateMessage) *message2.Receive {
    source := message.Source{
        SourceType: message.SourcePrivate,
        PrimaryID:  int64(m.Sender.Uin),
    }
    cqm := toStringMessage(m.Elements, source)
    return &message2.Receive{
        Time:        int(m.Time),
        SelfID:      int64(client.Uin),
        PostType:    "message",
        MessageType: "group",
        SubType:     "normal",
        MessageID:   int(m.Id),
        UserID:      int(m.Sender.Uin),
        RawMessage:  cqm,
        Sender: message2.SenderMessage{
            UserID:   int(m.Sender.Uin),
            Nickname: m.Sender.Nickname,
            Card:     m.Sender.CardName,
        },
    }
}

func toStringMessage(m []message.IMessageElement, source message.Source) string {
    elems := toElements(m, source)
    var sb strings.Builder
    for _, elem := range elems {
        elem.WriteCQCodeTo(&sb)
    }
    return sb.String()
}
func replyID(r *message.ReplyElement, source message.Source) int32 {
    id := source.PrimaryID
    seq := r.ReplySeq
    if r.GroupUin != 0 {
        id = int64(r.GroupUin)
    }
    // 私聊时，部分（不确定）的账号会在 ReplyElement 中带有 GroupID 字段。
    // 这里需要判断是由于 “直接回复” 功能，GroupID 为触发直接回复的来源那个群。
    if source.SourceType == message.SourcePrivate && (int64(r.SenderUin) == source.PrimaryID || int64(r.GroupUin) == source.PrimaryID || r.GroupUin == 0) {
        // 私聊似乎腾讯服务器有bug?
        seq = uint32(int32(uint16(seq)))
        id = int64(r.SenderUin)
    }
    return ToGlobalID(id, seq)
}

func ToGlobalID(code int64, msgID uint32) int32 {
    return int32(crc32.ChecksumIEEE([]byte(fmt.Sprintf("%d-%d", code, msgID))))
}

func toElements(e []message.IMessageElement, source message.Source) (r []Element) {
    // TODO: support OneBot V12
    type pair = Pair // simplify code
    type pairs = []pair

    r = make([]Element, 0, len(e))
    m := &message.SendingMessage{Elements: e}
    reply := m.FirstOrNil(func(e message.IMessageElement) bool {
        _, ok := e.(*message.ReplyElement)
        return ok
    })

    if reply != nil {
        replyElem := reply.(*message.ReplyElement)
        id := replyID(replyElem, source)
        elem := Element{
            Type: "reply",
            Data: pairs{
                {K: "id", V: strconv.FormatInt(int64(id), 10)},
            },
        }
        r = append(r, elem)
    }

    for _, elem := range e {
        var m Element
        switch o := elem.(type) {

        case *message.TextElement:
            m = Element{
                Type: "text",
                Data: pairs{
                    {K: "text", V: o.Content},
                },
            }
        case *message.LightAppElement:
            m = Element{
                Type: "json",
                Data: pairs{
                    {K: "data", V: o.Content},
                },
            }
        case *message.AtElement:
            target := "all"
            if o.TargetUin != 0 {
                target = strconv.FormatUint(uint64(o.TargetUin), 10)
            }
            m = Element{
                Type: "at",
                Data: pairs{
                    {K: "qq", V: target},
                },
            }
        case *message.FaceElement:
            m = Element{
                Type: "face",
                Data: pairs{
                    {K: "id", V: strconv.FormatInt(int64(o.FaceID), 10)},
                },
            }
        case *message.ImageElement:
            data := pairs{
                {K: "file", V: hex2.EncodeToString(o.Md5) + ".image"},
                {K: "url", V: o.Url},
            }
            m = Element{
                Type: "image",
                Data: data,
            }
        case *message.VoiceElement:
            m = Element{
                Type: "record",
                Data: pairs{
                    {K: "file", V: o.Name},
                    {K: "url", V: o.Url},
                },
            }
        case *message.ShortVideoElement:
            m = Element{
                Type: "video",
                Data: pairs{
                    {K: "file", V: o.Name},
                    {K: "url", V: o.Url},
                },
            }
        default:
            continue
        }
        r = append(r, m)
    }
    return
}
