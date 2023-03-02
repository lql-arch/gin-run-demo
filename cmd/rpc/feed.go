package rpc

import (
	"douSheng/Const"
	"douSheng/cmd/feed/kitex_gen/api"
	"douSheng/cmd/feed/kitex_gen/api/feed"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"log"
)

var feedClient feed.Client

// InitFeed RPC 客户端初始化
func InitFeed() {
	c, err := feed.NewClient("feed", client.WithHostPorts(Const.IP+":8888"))
	if err != nil {
		panic(err)
	}
	feedClient = c
}

// Feed 将 获取视频流操作传给 RPC Server 端,并接受响应响应.
func Feed(c *gin.Context, latestTime int64, token string) (*api.FeedList, error) {
	resp, err := feedClient.ReadVideos(c, latestTime, token)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}
	return resp, nil
}
