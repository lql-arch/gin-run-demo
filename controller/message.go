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

var tempChat map[string]int64

// 历史消息(取消使用)
//var ChatMessage map[string][]class.Message

func init() {
	tempChat = make(map[string]int64)
	//ChatMessage = make(map[string][]class.Message)
}

func ChatReset() {
	tempChat = make(map[string]int64)
	//ChatMessage = make(map[string][]class.Message)
}

func ChatEmpty() bool {
	if len(tempChat) == 0 {
		return true
	}
	return false
}

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
		id := timeToken(user.Id, userId)
		tempChat[id] = times

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

	// user 不需要使用会改变的信息
	if user, exist := FindUserToken(token); exist {
		var messages []class.Message
		userId := c.Query("to_user_id")

		toUserId, _ := strconv.ParseInt(userId, 0, 64)
		id := timeToken(user.Id, userId)
		recentTime := tempChat[id] // 防止自己发送消息多次显示

		messages, recentTime = sql.MessageChat(user.Id, toUserId, recentTime)
		// 历史信息(已取消使用)
		//ChatMessage[id] = append(ChatMessage[id], messages...)
		tempChat[id] = recentTime
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

// 构建myId与ToUserId形成timeToken,键值对存储最后一次读的时间,以读取最新消息
func timeToken(myId int64, ToUserId string) string {
	id := strconv.FormatInt(myId, 10)
	return fmt.Sprintf("%s_%s", id, ToUserId)
}
