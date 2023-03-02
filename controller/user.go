package controller

import (
	"douSheng/cmd/class"
	"douSheng/cmd/rpc"
	"douSheng/cmd/user/kitex_gen/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserLoginResponse struct {
	class.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	class.Response
	User api.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	register, err := rpc.Register(c, username, password)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: class.Response{
			StatusCode: 0,
		},
		UserId: register.User.Id,
		Token:  register.User.Token,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	resp, err := rpc.Login(c, username, password)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: class.Response{StatusCode: 0},
		UserId:   resp.User.Id,
		Token:    resp.User.Token,
	})

	//go sql.LoginReset()
	//go LoginReset()
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 0, 64)

	resp, err := rpc.UserInfo(c, token, userId)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Response: class.Response{StatusCode: 0},
		User:     *resp.User,
	})

}

func Errorf(err error) *UserLoginResponse { // *Response
	return &UserLoginResponse{
		Response: class.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		},
	}
}
