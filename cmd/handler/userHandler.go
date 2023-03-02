package handler

import (
	"context"
	"douSheng/Const"
	"douSheng/cmd/class"
	"douSheng/cmd/user/convert"
	api "douSheng/cmd/user/kitex_gen/api"
	"douSheng/sql"
	"fmt"
	"log"
)

// UserInfoImpl implements the last service interface defined in the IDL.
type UserInfoImpl struct{}

// Register implements the UserInfoImpl interface.
func (s *UserInfoImpl) Register(ctx context.Context, username string, password string) (resp *api.UserResponse, err error) {
	log.Println("Register")
	var user *api.User
	token := username + password
	token, _ = GenerateToken(ReplaceToken(username), ReplaceToken(password))
	token = Substring(token, 100)

	if _, exist := sql.FindUser(token); exist {
		return &api.UserResponse{
			StatusCode: 1,
			StatusMsg:  "用户已存在",
			Exist:      false,
			User:       nil,
		}, fmt.Errorf("用户已存在")
	}

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
	users, _ := sql.InsertUser(&newUser)
	user.Id = users[0].Id
	user.Token = users[0].Token

	return &api.UserResponse{
		StatusCode: 0,
		StatusMsg:  "",
		Exist:      true,
		User:       user,
	}, nil

}

// Login implements the UserInfoImpl interface.
func (s *UserInfoImpl) Login(ctx context.Context, username string, password string) (resp *api.UserResponse, err error) {
	log.Println("Login")
	token := username + password
	token, _ = GenerateToken(ReplaceToken(username), ReplaceToken(password))
	token = Substring(token, 100)

	user, exist := sql.FindUser(token)
	if !exist {
		return &api.UserResponse{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
			Exist:      false,
			User:       nil,
		}, fmt.Errorf("用户不存在")
	}

	return &api.UserResponse{
		StatusCode: 0,
		StatusMsg:  "",
		Exist:      true,
		User:       convert.UserConvert(user),
	}, nil
}

// UserInfo implements the UserInfoImpl interface.
func (s *UserInfoImpl) UserInfo(ctx context.Context, token string, userId int64) (resp *api.UserResponse, err error) {
	log.Println("UserInfo")
	user, exist := sql.FindUser(token)
	if !exist {
		return &api.UserResponse{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
			Exist:      false,
			User:       nil,
		}, fmt.Errorf("用户不存在")
	}

	if userId != user.Id {
		return &api.UserResponse{
			StatusCode: 1,
			StatusMsg:  "用户存在异常",
			Exist:      false,
			User:       nil,
		}, fmt.Errorf("用户存在异常")
	}

	return &api.UserResponse{
		StatusCode: 0,
		StatusMsg:  "",
		Exist:      true,
		User:       convert.UserConvert(user),
	}, nil
}
