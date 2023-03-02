package rpc

import (
	"douSheng/Const"
	"douSheng/cmd/message/kitex_gen/api"
	"douSheng/cmd/message/kitex_gen/api/messagefunc"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"log"
)

var messageClient messagefunc.Client

// InitMessage RPC 客户端初始化
func InitMessage() {
	c, err := messagefunc.NewClient("message", client.WithHostPorts(Const.IP+":8894"))
	if err != nil {
		panic(err)
	}
	messageClient = c
}

func MessageAction(ctx *gin.Context, token string, content string, toUserId int64, actionType int32) (resp *api.MessageResponse, err error) {
	resp, err = messageClient.MessageAction(ctx, token, content, toUserId, actionType)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}

func MessageChat(ctx *gin.Context, token string, preMsgTime int64, toUserId int64) (resp *api.MessageResponse, err error) {
	resp, err = messageClient.MessageChat(ctx, token, preMsgTime, toUserId)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}
