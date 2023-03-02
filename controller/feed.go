package controller

import (
	"douSheng/cmd/class"
	"douSheng/cmd/feed/kitex_gen/api"
	"douSheng/cmd/rpc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	class.Response
	VideoList []*api.Video `json:"video_list,omitempty"`
	NextTime  int64        `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	token := c.Query("token")
	latestTime, _ := strconv.ParseInt(c.Query("latest_time"), 0, 64)
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}

	//list, nextTime := sql.ReadVideos(latestTime, token)
	list, err := rpc.Feed(c, latestTime, token)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	if list.NextTime == 0 {
		list.NextTime = time.Now().Unix()
	}
	c.JSON(http.StatusOK, list)
}
