package controller

import (
	"douSheng/class"
	"douSheng/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 0, 64)
	actionType, _ := strconv.Atoi(c.Query("action_type"))

	if _, exist := sql.FindUser(token); !exist {
		c.JSON(http.StatusOK, class.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	if err := sql.FavoriteAction(token, videoID, actionType); err != nil {
		c.JSON(http.StatusOK, class.Response{StatusCode: 1, StatusMsg: "关注失败"})
		return
	}

	c.JSON(http.StatusOK, class.Response{StatusCode: 0})
}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		list, nextTime := sql.ReadVideos(time.Now().Unix(), token)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: class.Response{
				StatusCode: 0,
			},
			VideoList: list,
			NextTime:  nextTime,
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: class.Response{
				StatusCode: 0,
			},
			VideoList: sql.ReadFavoriteVideos(token),
		})
	}
}
