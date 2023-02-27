package controller

import (
	"douSheng/class"
	"douSheng/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//存储最新消息时间,已停用,由前端解决(取消使用)
//var tempChat map[string]int64
// 历史消息(取消使用)
//var ChatMessage map[string][]class.Message

type ChatResponse struct {
	class.Response
	MessageList []class.Message `json:"message_list"`
}

func MessageAction(c *gin.Context) {
	token := c.Query("token")
	content := c.Query("content")
	if strings.TrimSpace(content) == "" { //只有空格的或者空字符串不能发送
		c.JSON(http.StatusOK, class.Response{
			StatusCode: 1,
		})
		return
	}

	if user, exist := sql.FindUser(token); exist {
		userId := c.Query("to_user_id")
		toUserId, _ := strconv.ParseInt(userId, 0, 64)
		actionType, _ := strconv.Atoi(c.Query("action_type"))

		times := time.Now().Unix()

		if actionType == 1 { // 1-发送消息
			message := class.Message{
				MyId:     user.Id,
				Message:  content,
				ToUserId: toUserId,
				CreateAt: times,
			}

			sql.InsertMessage(message)
		}

		c.JSON(http.StatusOK, ChatResponse{
			Response: class.Response{StatusCode: 0},
		})
	} else {
		c.JSON(http.StatusOK, class.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
	}
}

func MessageChat(c *gin.Context) { // 只返回最新的数据,
	token := c.Query("token")
	preMsgTime, _ := strconv.ParseInt(c.Query("pre_msg_time"), 0, 64)

	// user 不需要使用会改变的信息
	if user, exist := FindUserToken(token); exist {
		var messages []class.Message
		userId := c.Query("to_user_id")

		toUserId, _ := strconv.ParseInt(userId, 0, 64)

		messages, preMsgTime = sql.MessageChat(user.Id, toUserId, preMsgTime)

		c.JSON(http.StatusOK, ChatResponse{
			Response:    class.Response{StatusCode: 0},
			MessageList: messages,
		})
	} else {
		c.JSON(http.StatusOK, class.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
	}
}

// 构建myId与ToUserId形成timeToken,键值对存储最后一次读的时间,以读取最新消息(停用)
func timeToken(myId int64, ToUserId string) string {
	id := strconv.FormatInt(myId, 10)
	return fmt.Sprintf("%s_%s", id, ToUserId)
}
