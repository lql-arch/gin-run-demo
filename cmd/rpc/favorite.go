package rpc

import (
	"douSheng/Const"
	"douSheng/cmd/favorite/kitex_gen/api"
	"douSheng/cmd/favorite/kitex_gen/api/favorite"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"log"
)

var favoriteClient favorite.Client

// InitFavorite RPC 客户端初始化
func InitFavorite() {
	c, err := favorite.NewClient("favorite", client.WithHostPorts(Const.IP+":8890"))
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

func FavoriteList(c *gin.Context, token string) (*api.FavoriteResponse, error) {
	resp, err := favoriteClient.FavoriteList(c, token)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}

func FavoriteAction(c *gin.Context, token string, videoID int64, actionType int) (*api.FavoriteResponse, error) {
	resp, err := favoriteClient.FavoriteAction(c, token, videoID, int32(actionType))
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("服务器连接异常")
	}
	if resp.StatusCode != 0 {
		return nil, fmt.Errorf(resp.StatusMsg)
	}

	return resp, nil
}
