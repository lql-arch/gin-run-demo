package controller

import (
	"douSheng/Const"
	"douSheng/class"
	"douSheng/service"
	"douSheng/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
	token, _ = GenerateToken(ReplaceToken(username), ReplaceToken(password))
	token = Substring(token, 100)

	if _, exist := sql.FindUser(token); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: class.Response{
				StatusCode: 1,
				StatusMsg:  "User already exist",
			},
		})
	} else {
		newUser := []class.User{
			{
				Name:            username,
				Token:           token,
				FollowCount:     0,
				FollowerCount:   0,
				BackgroundImage: Const.ServiceUrl + "/jpg/bronya.jpg",
				Signature:       "这个人啥都没有",
				Avatar:          Const.ServiceUrl + "/jpg/bronya.jpg",
				TotalFavorite:   0,
				WorkCount:       0,
				FavoriteCount:   0,
			},
		}
		ids, _ := sql.InsertUser(newUser)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: class.Response{StatusCode: 0},
			UserId:   ids[0],
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	token, _ = GenerateToken(ReplaceToken(username), ReplaceToken(password))
	token = Substring(token, 100)

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

// ReplaceToken 对token进行简单的替换
func ReplaceToken(token string) string {
	var str strings.Builder

	for i, x := range token {
		t := x - int32(i)
		if t <= 0 {
			t = x + int32(i)
		}
		str.WriteByte(byte(t))
	}

	return str.String()
}
