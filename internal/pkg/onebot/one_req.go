package onebot

import (
    "fmt"
    "github.com/LagrangeDev/LagrangeGo/client"
    "github.com/LagrangeDev/LagrangeGo/message"
    "html"
    msg "qqbot-reconstruction/internal/app/message"
    "qqbot-reconstruction/internal/pkg/cache"
)

func ParseGroupMsg(client *client.QQClient, group *message.GroupMessage) *msg.Receive {
    cache.GroupMessageLru.Add(group.Id, group)
    return &msg.Receive{
        Time:        int(group.Time),
        SelfID:      int64(client.Uin),
        PostType:    "message",
        MessageType: "group",
        SubType:     "normal",
        MessageID:   int(group.Id),
        GroupId:     int(group.GroupCode),
        UserID:      int(group.Sender.Uin),
        RawMessage:  MsgToRawMsg(client, group.Elements),
        Sender: msg.SenderMessage{
            UserID:   int(group.Sender.Uin),
            Nickname: group.Sender.Nickname,
            Card:     group.Sender.CardName,
        },
    }
}

func ParsePrivateMsg(client *client.QQClient, private *message.PrivateMessage) *msg.Receive {
    cache.PrivateMessageLru.Add(private.Id, private)
    return &msg.Receive{
        Time:        int(private.Time),
        SelfID:      int64(client.Uin),
        PostType:    "message",
        MessageType: "group",
        SubType:     "normal",
        MessageID:   int(private.Id),
        UserID:      int(private.Sender.Uin),
        RawMessage:  MsgToRawMsg(client, private.Elements),
        Sender: msg.SenderMessage{
            UserID:   int(private.Sender.Uin),
            Nickname: private.Sender.Nickname,
            Card:     private.Sender.CardName,
        },
    }
}


func MsgToRawMsg(client *client.QQClient, messageChain []message.IMessageElement) string {
    result := ""
    for _, element := range messageChain {
        switch elem := element.(type) {
        case *message.TextElement:
            result += elem.Content
        case *message.ImageElement:
            result += fmt.Sprintf(`<image image_id="%s" url="%s"/>`, html.EscapeString(elem.ImageId), html.EscapeString(elem.Url))
        case *message.FaceElement:
            result += fmt.Sprintf(`<face id="%d"/>`, elem.FaceID)
        case *message.VoiceElement:
            result += fmt.Sprintf(`<voice url="%s"/>`, html.EscapeString(elem.Url))
        case *message.ReplyElement:
            result += fmt.Sprintf(`<reply time="%d" sender="%d" raw_message="%s" reply_seq="%d"/>`, elem.Time, elem.SenderUin, html.EscapeString(MsgToRawMsg(client, elem.Elements)), elem.ReplySeq)
        }
    }
    return result
}

