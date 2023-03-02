package controller

import (
	"douSheng/cmd/rpc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 0, 64)
	actionType, _ := strconv.Atoi(c.Query("action_type"))

	resp, err := rpc.FavoriteAction(c, token, videoID, actionType)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")

	resp, err := rpc.FavoriteList(c, token)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
