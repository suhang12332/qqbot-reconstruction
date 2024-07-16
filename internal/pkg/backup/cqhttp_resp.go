package onebot
//
//import (
//    "encoding/xml"
//    "github.com/LagrangeDev/LagrangeGo/client"
//    "qqbot-reconstruction/internal/pkg/variable"
//    "regexp"
//)
//
////func HandleSendPrivateMsg(cli *client.QQClient, req *onebot.SendPrivateMsgReq) *onebot.SendPrivateMsgResp {
////	miraiMsg := ProtoMsgToMiraiMsg(cli, req.Message, req.AutoEscape)
////	sendingMessage := &message.SendingMessage{Elements: miraiMsg}
////	log.Infof("Bot(%d) Private(%d) <- %s", cli.Uin, req.UserId, MiraiMsgToRawMsg(cli, miraiMsg))
////	ret, _ := cli.SendPrivateMessage(uint32(req.UserId), sendingMessage.Elements)
////	cache.PrivateMessageLru.Add(ret.Result, ret)
////	return &onebot.SendPrivateMsgResp{
////		MessageId: int32(ret.PrivateSequence),
////	}
////}
//
//type Message struct {
//
//	Type string            
//	Data map[string]string 
//}
//
//
//
//type OneBotSendMsg struct {
//    MessageType string     
//	UserId      int64      
//	GroupId     int64      
//	Message     []*Message 
//	AutoEscape  bool       
//}
//
//type OneBotSendGroupFawordMsg struct {
//    
//}
//
//type OneBotSendPrivateFawordMsg struct {
//    
//}
//
//type Send variable.SendMsg
//
//type SendMsg OneBotSendMsg
//
//func (s *Send)ParseMsg() *SendMsg {
//    return &SendMsg{
//		MessageType: s.MessageType,
//		GroupId:     int64(s.GroupId),
//		UserId: int64(s.UserId),
//		Message:     s.Message,
//		AutoEscape:  s.AutoEscape,
//	},
//}
//
//
//
//func (s *SendMsg)HandleSendMsg(cli *client.QQClient) {
//	miraiMsg := ProtoMsgToMiraiMsg(cli, req.Message, req.AutoEscape)
//	sendingMessage := &message.SendingMessage{Elements: miraiMsg}
//
//	/*if req.GroupId != 0 && req.UserId != 0 { // 临时
//		ret, _ := cli.SendTempMessage(uint32(req.GroupId), uint32(req.UserId), sendingMessage.Elements)
//		cache.PrivateMessageLru.Add(ret.PrivateSequence, ret)
//		return &onebot.SendMsgResp{
//			MessageId: int32(ret.PrivateSequence),
//		}
//	}*/
//
//	if req.GroupId != 0 { // 群
//		if g, err := cli.GetCachedGroupInfo(uint32(req.GroupId)); err != nil || g == nil {
//			log.Warnf("发送消息失败，群聊 %d 不存在", req.GroupId)
//			return nil
//		}
//		ret, _ := cli.SendGroupMessage(uint32(req.GroupId), sendingMessage.Elements)
//		if ret.GroupSequence.IsNone() {
//			config.Fragment = !config.Fragment
//			log.Warnf("发送群消息失败，可能被风控，下次发送将改变分片策略，Fragment: %+v", config.Fragment)
//			return nil
//		}
//		cache.GroupMessageLru.Add(int32(ret.GroupSequence.Unwrap()), ret)
//		return &onebot.SendMsgResp{
//			MessageId: int32(ret.GroupSequence.Unwrap()),
//		}
//	}
//
//	if req.UserId != 0 { // 私聊
//		ret, _ := cli.SendPrivateMessage(uint32(req.UserId), sendingMessage.Elements)
//		cache.PrivateMessageLru.Add(ret.Result, ret)
//		return &onebot.SendMsgResp{
//			MessageId: int32(ret.PrivateSequence),
//		}
//	}
//	log.Warnf("failed to send msg")
//	return nil
//}
//
//
//
//
//
//
//
//type Node struct {
//    XMLName xml.Name
//    Attr    []xml.Attr `xml:",any,attr"`
//}
//
//var re = regexp.MustCompile("<[\\s\\S]+?/>")
//
//func ProtoMsgToMiraiMsg(cli *client.QQClient, msgList []*onebot.Message, notConvertText bool) []message.IMessageElement {
//	containReply := false // 每条消息只能包含一个reply
//	messageChain := make([]message.IMessageElement, 0)
//	for _, protoMsg := range msgList {
//		switch protoMsg.Type {
//		case "text":
//			if notConvertText {
//				messageChain = append(messageChain, ProtoTextToMiraiText(protoMsg.Data))
//			} else {
//				text, ok := protoMsg.Data["text"]
//				if !ok {
//					log.Warnf("text不存在")
//					continue
//				}
//				messageChain = append(messageChain, RawMsgToMiraiMsg(cli, text)...) // 转换xml码
//			}
//		case "at":
//			messageChain = append(messageChain, ProtoAtToMiraiAt(protoMsg.Data))
//		case "image":
//			messageChain = append(messageChain, ProtoImageToMiraiImage(protoMsg.Data))
//		case "img":
//			messageChain = append(messageChain, ProtoImageToMiraiImage(protoMsg.Data))
//		case "friend_image":
//			messageChain = append(messageChain, ProtoPrivateImageToMiraiPrivateImage(protoMsg.Data))
//		case "friend_img":
//			messageChain = append(messageChain, ProtoPrivateImageToMiraiPrivateImage(protoMsg.Data))
//		case "record":
//			messageChain = append(messageChain, ProtoVoiceToMiraiVoice(protoMsg.Data))
//		case "face":
//			messageChain = append(messageChain, ProtoFaceToMiraiFace(protoMsg.Data))
//		case "reply":
//			if replyElement := ProtoReplyToMiraiReply(protoMsg.Data); replyElement != nil && !containReply {
//				containReply = true
//				messageChain = append([]message.IMessageElement{replyElement}, messageChain...)
//			}
//		case "sleep":
//			ProtoSleep(protoMsg.Data)
//		default:
//			log.Errorf("不支持的消息类型 %+v", protoMsg)
//		}
//	}
//	return messageChain
//}
//
//func RawMsgToMiraiMsg(cli *client.QQClient, str string) []message.IMessageElement {
//    containReply := false
//    var node Node
//    textList := re.Split(str, -1)
//    codeList := re.FindAllString(str, -1)
//    elemList := make([]message.IMessageElement, 0)
//    for len(textList) > 0 || len(codeList) > 0 {
//        if len(textList) > 0 && strings.HasPrefix(str, textList[0]) {
//            text := textList[0]
//            textList = textList[1:]
//            str = str[len(text):]
//            elemList = append(elemList, message.NewText(text))
//        }
//        if len(codeList) > 0 && strings.HasPrefix(str, codeList[0]) {
//            code := codeList[0]
//            codeList = codeList[1:]
//            str = str[len(code):]
//            err := xml.Unmarshal([]byte(code), &node)
//            if err != nil {
//                elemList = append(elemList, message.NewText(code))
//                continue
//            }
//            attrMap := make(map[string]string)
//            for _, attr := range node.Attr {
//                attrMap[attr.Name.Local] = html.UnescapeString(attr.Value)
//            }
//            switch node.XMLName.Local {
//            case "at":
//                elemList = append(elemList, ProtoAtToMiraiAt(attrMap))
//            case "img":
//                elemList = append(elemList, ProtoImageToMiraiImage(attrMap)) // TODO 为了兼容我的旧代码偷偷加的
//            case "image":
//                elemList = append(elemList, ProtoImageToMiraiImage(attrMap))
//            case "face":
//                elemList = append(elemList, ProtoFaceToMiraiFace(attrMap))
//            case "voice":
//                elemList = append(elemList, ProtoVoiceToMiraiVoice(attrMap))
//            case "record":
//                elemList = append(elemList, ProtoVoiceToMiraiVoice(attrMap))
//            case "text":
//                elemList = append(elemList, ProtoTextToMiraiText(attrMap))
//            case "reply":
//                if replyElement := ProtoReplyToMiraiReply(attrMap); replyElement != nil && !containReply {
//                    containReply = true
//                    elemList = append([]message.IMessageElement{replyElement}, elemList...)
//                }
//            case "sleep":
//                ProtoSleep(attrMap)
//            default:
//                log.Warning("不支持的类型 %s", code)
//                elemList = append(elemList, message.NewText(code))
//            }
//        }
//    }
//    return elemList
//}
//func EmptyText() *message.TextElement {
//    return message.NewText("")
//}
//func ProtoAtToMiraiAt(data map[string]string) message.IMessageElement {
//    qq, ok := data["qq"]
//    if !ok {
//        log.Warningf("atQQ不存在")
//        return EmptyText()
//    }
//    if qq == "all" {
//        return message.NewAt(0)
//    }
//    userId, err := strconv.ParseInt(qq, 10, 64)
//    if err != nil {
//        log.Warningf("atQQ不是数字")
//        return EmptyText()
//    }
//    return message.NewAt(uint32(userId))
//}
//
//func ProtoFaceToMiraiFace(data map[string]string) message.IMessageElement {
//    idStr, ok := data["id"]
//    if !ok {
//        log.Warningf("faceId不存在")
//        return EmptyText()
//    }
//    id, err := strconv.Atoi(idStr)
//    if err != nil {
//        log.Warningf("faceId不是数字")
//        return EmptyText()
//    }
//    return &message.FaceElement{
//        FaceID: uint16(id),
//    }
//}
//
//func ProtoTextToMiraiText(data map[string]string) message.IMessageElement {
//    text, ok := data["text"]
//    if !ok {
//        log.Warningf("text不存在")
//        return EmptyText()
//    }
//    return message.NewText(text)
//}
//func ProtoReplyToMiraiReply(data map[string]string) *message.ReplyElement {
//    rawMessage, hasRawMessage := data["raw_message"] // 如果存在 raw_message，按照raw_message显示
//
//    messageIdStr, ok := data["message_id"]
//    if !ok {
//        return nil
//    }
//    messageIdInt, err := strconv.Atoi(messageIdStr)
//    if err != nil {
//        return nil
//    }
//    messageId := int32(messageIdInt)
//    eventInterface, ok := cache.GroupMessageLru.Get(messageId)
//    if ok {
//        groupMessage, ok := eventInterface.(*message.GroupMessage)
//        if ok {
//            return &message.ReplyElement{
//                ReplySeq:  uint32(groupMessage.Id),
//                SenderUin: groupMessage.Sender.Uin,
//                Time: uint32(groupMessage.Time),
//                Elements: func() []message.IMessageElement {
//                    if hasRawMessage {
//                        return []message.IMessageElement{message.NewText(rawMessage)}
//                    } else {
//                        return groupMessage.Elements
//                    }
//                }(),
//            }
//        }
//    }
//    eventInterface, ok = cache.PrivateMessageLru.Get(messageId)
//    if ok {
//        privateMessage, ok := eventInterface.(*message.PrivateMessage)
//        if ok {
//            return &message.ReplyElement{
//                ReplySeq: uint32(privateMessage.Id),
//                SenderUin:   privateMessage.Sender.Uin,
//                Time: uint32(privateMessage.Time),
//                Elements: func() []message.IMessageElement {
//                    if hasRawMessage {
//                        return []message.IMessageElement{message.NewText(rawMessage)}
//                    } else {
//                        return privateMessage.Elements
//                    }
//                }(),
//            }
//        }
//    }
//    return nil
//}
//
//func ProtoSleep(data map[string]string) {
//    t, ok := data["time"]
//    if !ok {
//        log.Warningf("failed to get sleep time1")
//        return
//    }
//    ms, err := strconv.Atoi(t)
//    if err != nil {
//        log.Warning("failed to get sleep time2, %+v", err)
//        return
//    }
//    if ms > 24*3600*1000 {
//        log.Warningf("最多 sleep 24小时")
//        ms = 24 * 3600 * 1000
//    }
//    time.Sleep(time.Duration(ms) * time.Millisecond)
//}
//
//func ProtoImageToMiraiImage(data map[string]string) message.IMessageElement {
//	elem := &message.ImageElement{}
//	url, ok := data["url"]
//	if !ok {
//		url, ok = data["src"] // TODO 为了兼容我的旧代码偷偷加的
//		if !ok {
//			url, ok = data["file"]
//		}
//	}
//	if !ok {
//		log.Warningf("imageUrl不存在")
//		return EmptyText()
//	}
//	b, err := preprocessImageMessage(url)
//	if err == nil {
//		elem.Stream = bytes.NewReader(b)
//	}
//	return elem
//}
//
//func preprocessImageMessage(path string) ([]byte, error) {
//    if strings.Contains(path, "http") {
//        resp, err := http.Get(path)
//        defer resp.Body.Close()
//        if err != nil {
//            return nil, err
//        }
//        imo, err := io.ReadAll(resp.Body)
//        if err != nil {
//            return nil, err
//        }
//        return imo, nil
//    } else {
//        f, err := os.Open(path)
//        if err != nil {
//            return nil, err
//        }
//        reader, err := io.ReadAll(f)
//        if err != nil {
//            return nil, err
//        }
//        return reader, nil
//    }
//}
//
//func ProtoVoiceToMiraiVoice(data map[string]string) message.IMessageElement {
//    url, ok := data["url"]
//    if !ok {
//        url, ok = data["file"]
//    }
//    if !ok {
//        log.Warningf("recordUrl不存在")
//        return EmptyText()
//    }
//    b, err := GetBytes(url)
//    if err != nil {
//        log.Errorf("下载语音失败")
//        return EmptyText()
//    }
//    if !IsAMRorSILK(b) {
//        log.Errorf("不是amr或silk格式")
//        return EmptyText()
//    }
//    return &message.VoiceElement{Stream: bytes.NewReader(b)}
//}
//
//var (
//    HEADER_AMR  = []byte("#!AMR")
//    HEADER_SILK = []byte("\x02#!SILK_V3")
//)
//
//func IsAMRorSILK(b []byte) bool {
//    return bytes.HasPrefix(b, HEADER_AMR) || bytes.HasPrefix(b, HEADER_SILK)
//}
//
//var httpClient = http.Client{
//    Timeout: 15 * time.Second,
//}
//
//func GetBytes(url string) ([]byte, error) {
//    req, err := http.NewRequest("GET", url, nil)
//    if err != nil {
//        return nil, err
//    }
//    req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36 Edg/83.0.478.61"}
//    resp, err := httpClient.Do(req)
//    if err != nil {
//        return nil, err
//    }
//    defer resp.Body.Close()
//    body, err := io.ReadAll(resp.Body)
//    if err != nil {
//        return nil, err
//    }
//    if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
//        buffer := bytes.NewBuffer(body)
//        r, _ := gzip.NewReader(buffer)
//        defer r.Close()
//        unCom, err := ioutil.ReadAll(r)
//        return unCom, err
//    }
//    return body, nil
//}