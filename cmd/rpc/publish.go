package rpc

import (
	"douSheng/Const"
	"douSheng/cmd/publish/kitex_gen/api"
	"douSheng/cmd/publish/kitex_gen/api/publish"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"log"
)

var publishClient publish.Client

// InitPublish RPC 客户端初始化
func InitPublish() {
	var err error
	publishClient, err = publish.NewClient("publish", client.WithHostPorts(Const.IP+":8891"))
	if err != nil {
		panic(err)
	}
}

func PublishList(c *gin.Context, token string) (*api.PublishResponse, error) {
	resp, err := publishClient.PublishList(c, token)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}

func PublishAction(c *gin.Context, token string, data *api.VideoData) (resp *api.PublishResponse, err error) {
	resp, err = publishClient.Publish(c, token, data)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}
