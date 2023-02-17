package controller

import (
	"douSheng/class"
	"douSheng/service"
	"douSheng/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync/atomic"
)

type UserLoginResponse struct {
	class.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	class.Response
	User class.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := sql.FindUser(token); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: class.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		userIdSequence := sql.FindUserIdSequence()
		atomic.AddInt64(&userIdSequence, 1)
		newUser := []class.User{
			{
				Name:  username,
				Token: token,
			},
		}
		_ = sql.InsertUser(newUser)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: class.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := sql.FindUser(token); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: class.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: class.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
	go service.LoginReset()
	go LoginReset()
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 0, 64)

	user, exist := sql.FindUser(token)
	if !exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: class.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}

	if userId != user.Id {
		c.JSON(http.StatusOK, UserResponse{
			Response: class.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Response: class.Response{StatusCode: 0},
		User:     user,
	})

}
