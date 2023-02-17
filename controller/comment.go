package controller

import (
	"douSheng/class"
	"douSheng/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	class.Response
	CommentList []class.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	class.Response
	Comment class.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType, _ := strconv.Atoi(c.Query("action_type"))

	user, exist := sql.FindUser(token)

	if exist {
		text := c.Query("comment_text")
		videoId, _ := strconv.Atoi(c.Query("video_id"))

		times := time.Now().Unix()

		comment := class.Comment{
			Author:     user,
			UserToken:  token,
			Content:    text,
			CreateDate: times,
			VideoId:    videoId,
			Type:       actionType,
		}

		if actionType == 1 { // 发布评论
			// 添加comment到数据库
			id, err := sql.ReviseComment(comment)
			if err != nil {
				c.JSON(http.StatusOK, class.Response{StatusCode: 1, StatusMsg: "发布失败" + err.Error()})
				return
			}

			comment.Id = id

			c.JSON(http.StatusOK, CommentActionResponse{
				Response: class.Response{StatusCode: 0},
				Comment:  comment,
			})

			return
		} else if actionType == 2 { // 删除评论
			commentId, _ := strconv.Atoi(c.Query("comment_id"))
			comment.Id = int64(commentId)
			// 删除comment到数据库
			if _, err := sql.ReviseComment(comment); err != nil {
				c.JSON(http.StatusOK, class.Response{StatusCode: 1, StatusMsg: "删除失败"})
				return
			}

			c.JSON(http.StatusOK, class.Response{StatusCode: 0})
			return
		}
	} else {
		fmt.Println(1)
		c.JSON(http.StatusOK, class.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	c.JSON(http.StatusOK, class.Response{StatusCode: 1, StatusMsg: "未知错误"})
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	token := c.Query("token")

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    class.Response{StatusCode: 0},
		CommentList: sql.FindComments(videoId, token),
	})
}
