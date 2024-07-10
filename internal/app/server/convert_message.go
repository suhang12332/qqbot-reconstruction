package server

import "qqbot-reconstruction/internal/pkg/variable"

// ParseTReceive
// @description: 用于转化客户端结构体,用于数据库入库操作
// @return: 结果
func (receive *Receive) ParseTReceive() (*variable.TReceive, *variable.TSender) {
    sender := (*receive).Sender
    tSender := variable.TSender{
        Age:      (sender).Age,
        Area:     (sender).Area,
        Card:     (sender).Card,
        Level:    (sender).Level,
        Nickname: (sender).Nickname,
        Role:     (sender).Role,
        Sex:      (sender).Sex,
        Title:    (sender).Title,
        UserID:   (sender).UserID,
    }
    tReceive := variable.TReceive{
        PostType:    (*receive).PostType,
        MessageType: (*receive).MessageType,
        Time:        (*receive).Time,
        SelfID:      (*receive).SelfID,
        SubType:     (*receive).SubType,
        MessageID:   (*receive).MessageID,
        UserID:      (*receive).UserID,
        TargetID:    (*receive).TargetID,
        Message:     (*receive).Message,
        MessageSeq:  (*receive).MessageSeq,
        RawMessage:  (*receive).RawMessage,
        ID:          tSender.ID,
        Font:        (*receive).Font,
        GroupId:     (*receive).GroupId,
    }
    return &tReceive, &tSender
}
