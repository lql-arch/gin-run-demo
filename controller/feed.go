package controller

import (
	"douSheng/class"
	"douSheng/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	class.Response
	VideoList []class.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	token := c.Query("token")
	latestTime, _ := strconv.ParseInt(c.Query("latest_time"), 0, 64)
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}

	list, nextTime := sql.ReadVideos(latestTime, token)

	if nextTime == 0 {
		nextTime = time.Now().Unix()
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  class.Response{StatusCode: 0},
		VideoList: list,
		NextTime:  nextTime,
	})
}
