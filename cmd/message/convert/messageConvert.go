package convert

import (
	"douSheng/cmd/class"
	"douSheng/cmd/message/kitex_gen/api"
)

func MessageListConvert(messageList []*class.Message) []*api.Message {
	newMessageList := make([]*api.Message, 0)
	for _, message := range messageList {
		newMessage := MessageConvert(message)
		if newMessage == nil {
			continue
		}
		newMessageList = append(newMessageList, newMessage)
	}

	return newMessageList
}

func MessageConvert(message *class.Message) *api.Message {
	return &api.Message{
		Id:       message.Id,
		ToUserId: message.ToUserId,
		MyId:     message.MyId,
		Message:  message.Message,
		CreateAt: message.CreateAt,
	}
}
