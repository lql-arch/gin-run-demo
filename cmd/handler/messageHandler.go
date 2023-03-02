package handler

import (
	"context"
	"douSheng/cmd/class"
	"douSheng/cmd/message/convert"
	api "douSheng/cmd/message/kitex_gen/api"
	"douSheng/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

// MessageFuncImpl implements the last service interface defined in the IDL.
type MessageFuncImpl struct{}

// MessageAction implements the MessageFuncImpl interface.
func (s *MessageFuncImpl) MessageAction(ctx context.Context, token string, content string, toUserId int64, actionType int32) (resp *api.MessageResponse, err error) {
	log.Println("MessageAction")

	if strings.TrimSpace(content) == "" { //只有空格的或者空字符串不能发送
		return &api.MessageResponse{
			StatusCode: 1,
		}, nil
	}

	user, exist := sql.FindUser(token)
	if !exist {
		return &api.MessageResponse{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		}, fmt.Errorf("用户不存在")
	}

	times := time.Now().Unix()
	if actionType == 1 { // 1-发送消息
		message := class.Message{
			MyId:     user.Id,
			Message:  content,
			ToUserId: toUserId,
			CreateAt: times,
		}

		err = sql.InsertMessage(message)
		if err != nil {
			return &api.MessageResponse{
				StatusCode: 0,
				StatusMsg:  err.Error(),
			}, err
		}
	}

	return &api.MessageResponse{
		StatusCode: 0,
	}, nil
}

// MessageChat implements the MessageFuncImpl interface.
func (s *MessageFuncImpl) MessageChat(ctx context.Context, token string, preMsgTime int64, toUserId int64) (resp *api.MessageResponse, err error) {
	log.Println("MessageChat")

	// user 不需要使用会改变的信息
	user, exist := FindUserToken(token)
	if !exist {
		return &api.MessageResponse{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		}, fmt.Errorf("用户不存在")
	}

	var messages []*class.Message
	messages, preMsgTime = sql.MessageChat(user.Id, toUserId, preMsgTime)

	return &api.MessageResponse{
		StatusCode: 0,
		Message:    convert.MessageListConvert(messages),
	}, nil
}
