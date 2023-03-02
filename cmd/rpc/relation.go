package rpc

import (
	"douSheng/Const"
	"douSheng/cmd/relation/kitex_gen/api"
	"douSheng/cmd/relation/kitex_gen/api/relationfunc"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"log"
)

var relationClient relationfunc.Client

// InitRelation RPC 客户端初始化
func InitRelation() {
	c, err := relationfunc.NewClient("relation", client.WithHostPorts(Const.IP+":8893"))
	if err != nil {
		panic(err)
	}
	relationClient = c
}

func RelationAction(ctx *gin.Context, token string, toUserId int64, state int32) (resp *api.RelationResponse, err error) {
	resp, err = relationClient.RelationAction(ctx, token, toUserId, state)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}

// FollowList implements the RelationFuncImpl interface.
func FollowList(ctx *gin.Context, userId int64, token string) (resp *api.RelationResponse, err error) {
	resp, err = relationClient.FollowList(ctx, userId, token)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}

// FollowerList implements the RelationFuncImpl interface.
func FollowerList(ctx *gin.Context, userId int64, token string) (resp *api.RelationResponse, err error) {
	resp, err = relationClient.FollowerList(ctx, userId, token)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}

// FriendList implements the RelationFuncImpl interface.
func FriendList(ctx *gin.Context, userId int64, token string) (resp *api.RelationResponse, err error) {
	resp, err = relationClient.FriendList(ctx, userId, token)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}
