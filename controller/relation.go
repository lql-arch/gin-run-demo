package controller

import (
	"douSheng/cmd/class"
	"douSheng/cmd/rpc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserFriendListResponse struct {
	class.Response
	UserList []class.FriendUser `json:"user_list"`
}

type UserListResponse struct {
	class.Response
	UserList []class.User `json:"user_list"`
}

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 0, 64)
	state, _ := strconv.Atoi(c.Query("action_type"))

	resp, err := rpc.RelationAction(c, token, toUserId, int32(state))

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func FollowList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 0, 64)
	token := c.Query("token")

	resp, err := rpc.FollowList(c, userId, token)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func FollowerList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 0, 64)
	token := c.Query("token")

	resp, err := rpc.FollowerList(c, userId, token)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// FriendList 好友 : 我关注且关注我的
func FriendList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 0, 64)
	token := c.Query("token")

	resp, err := rpc.FriendList(c, userId, token)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
