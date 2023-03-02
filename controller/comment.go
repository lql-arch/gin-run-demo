package controller

import (
	"douSheng/cmd/comment/kitex_gen/api"
	"douSheng/cmd/rpc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func CommentAction(c *gin.Context) {
	var resp *api.CommentResponse
	var err error
	token := c.Query("token")
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 0, 64)

	if actionType == 1 {
		text := c.Query("comment_text")

		resp, err = rpc.AddCommentAction(c, token, int32(actionType), text, videoId)
	} else if actionType == 2 {
		commentId, _ := strconv.ParseInt(c.Query("comment_id"), 0, 64)

		resp, err = rpc.DeleteCommentAction(c, token, int32(actionType), commentId, videoId)
	}

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func CommentList(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 0, 64)
	token := c.Query("token")

	resp, err := rpc.CommentList(c, videoId, token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
