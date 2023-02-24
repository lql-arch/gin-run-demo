package controller

import (
	"douSheng/class"
	"douSheng/setting"
	"douSheng/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	class.Response
	CommentList []class.JsonComment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	class.Response
	Comment class.JsonComment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType, _ := strconv.Atoi(c.Query("action_type"))

	user, exist := sql.FindUser(token)

	if exist {
		text := c.Query("comment_text")
		videoId, _ := strconv.ParseInt(c.Query("video_id"), 0, 64)

		comment := class.Comment{
			GormComment: class.GormComment{
				Author:  user,
				UserId:  user.Id,
				Content: text,
				VideoId: videoId,
				Type:    actionType,
			},
			CreateDate: time.Now().Unix(),
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
				Comment: class.JsonComment{
					GormComment: comment.GormComment,
					CreateDate:  setting.CommentTimeString(comment.CreateDate),
				},
			})

			return
		} else if actionType == 2 { // 删除评论
			comment.Id, _ = strconv.ParseInt(c.Query("comment_id"), 0, 64)
			fmt.Println(comment.Id)
			// 删除comment到数据库
			if _, err := sql.ReviseComment(comment); err != nil {
				c.JSON(http.StatusOK, class.Response{StatusCode: 1, StatusMsg: "删除失败"})
				return
			}

			c.JSON(http.StatusOK, class.Response{StatusCode: 0})
			return
		}
	} else {
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
