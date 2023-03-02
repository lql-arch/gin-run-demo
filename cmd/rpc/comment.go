package rpc

import (
	"douSheng/Const"
	"douSheng/cmd/comment/kitex_gen/api"
	"douSheng/cmd/comment/kitex_gen/api/commentfunc"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"log"
)

var commentClient commentfunc.Client

// InitComment RPC 客户端初始化
func InitComment() {
	c, err := commentfunc.NewClient("comment", client.WithHostPorts(Const.IP+":8892"))
	if err != nil {
		panic(err)
	}
	commentClient = c
}

func CommentList(c *gin.Context, videoId int64, token string) (*api.CommentResponse, error) {
	resp, err := commentClient.CommentList(c, videoId, token)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}

func AddCommentAction(c *gin.Context, token string, actionType int32, text string, videoId int64) (resp *api.CommentResponse, err error) {
	resp, err = commentClient.AddCommentAction(c, token, actionType, text, videoId)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}

func DeleteCommentAction(c *gin.Context, token string, actionType int32, commentId int64, videoId int64) (resp *api.CommentResponse, err error) {
	resp, err = commentClient.DeleteCommentAction(c, token, actionType, commentId, videoId)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}
