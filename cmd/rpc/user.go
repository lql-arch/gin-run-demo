package rpc

import (
	"context"
	"douSheng/Const"
	"douSheng/cmd/user/kitex_gen/api"
	"douSheng/cmd/user/kitex_gen/api/userinfo"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"log"
)

var userClient userinfo.Client

// InitUser RPC 客户端初始化
func InitUser() {
	c, err := userinfo.NewClient("user", client.WithHostPorts(Const.IP+":8889"))
	if err != nil {
		panic(err)
	}
	userClient = c
}

// Register 用户注册事件发送给RPC服务器
func Register(c *gin.Context, username string, password string) (*api.UserResponse, error) {

	resp, err := userClient.Register(c, username, password)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}

	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}

func Login(ctx context.Context, username string, password string) (resp *api.UserResponse, err error) {

	resp, err = userClient.Login(ctx, username, password)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}

	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}

func UserInfo(ctx context.Context, token string, userId int64) (resp *api.UserResponse, err error) {
	resp, err = userClient.UserInfo(ctx, token, userId)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}

	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}
