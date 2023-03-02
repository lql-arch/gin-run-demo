package controller

import (
	"douSheng/cmd/class"
	"douSheng/cmd/rpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 0, 64)
	actionType, _ := strconv.Atoi(c.Query("action_type"))

	resp, err := rpc.MessageAction(c, token, content, toUserId, int32(actionType))

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}

func MessageChat(c *gin.Context) { // 只返回最新的数据,
	token := c.Query("token")
	preMsgTime, _ := strconv.ParseInt(c.Query("pre_msg_time"), 0, 64)
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 0, 64)

	resp, err := rpc.MessageChat(c, token, preMsgTime, toUserId)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}

// 构建myId与ToUserId形成timeToken,键值对存储最后一次读的时间,以读取最新消息(停用)
func timeToken(myId int64, ToUserId string) string {
	id := strconv.FormatInt(myId, 10)
	return fmt.Sprintf("%s_%s", id, ToUserId)
}
