package handler

import (
	"context"
	"douSheng/cmd/relation/convert"
	api "douSheng/cmd/relation/kitex_gen/api"
	"douSheng/sql"
	"fmt"
	"log"
)

// RelationFuncImpl implements the last service interface defined in the IDL.
type RelationFuncImpl struct{}

// RelationAction implements the RelationFuncImpl interface.
func (s *RelationFuncImpl) RelationAction(ctx context.Context, token string, toUserId int64, state int32) (resp *api.RelationResponse, err error) {
	log.Println("RelationAction")
	user, ok := FindUserToken(token)
	// 用户是否存在
	if !ok { // 用户不存在
		return &api.RelationResponse{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		}, fmt.Errorf("用户不存在")
	}

	if user.Id == toUserId && state == 1 {
		return &api.RelationResponse{
			StatusCode: 1,
			StatusMsg:  "不能关注自己",
		}, fmt.Errorf("不能关注自己")
	}

	if err := sql.RelationAction(token, toUserId, int(state)); err != nil {
		log.Println(err)
		return &api.RelationResponse{
			StatusCode: 1,
			StatusMsg:  "关注错误",
		}, fmt.Errorf("关注错误")
	}

	return &api.RelationResponse{
		StatusCode: 0,
		StatusMsg:  "关注成功",
	}, nil
}

// FollowList implements the RelationFuncImpl interface.
func (s *RelationFuncImpl) FollowList(ctx context.Context, userId int64, token string) (resp *api.RelationResponse, err error) {
	log.Println("FollowList")
	if _, ok := FindUserToken(token); !ok {
		return &api.RelationResponse{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		}, fmt.Errorf("用户不存在")
	}

	return &api.RelationResponse{
		StatusCode: 0,
		UserList:   convert.UsersConvertFriends(sql.FindFollowUsers(userId, token)),
	}, nil
}

// FollowerList implements the RelationFuncImpl interface.
func (s *RelationFuncImpl) FollowerList(ctx context.Context, userId int64, token string) (resp *api.RelationResponse, err error) {
	log.Println("FollowerList")
	if _, ok := FindUserToken(token); !ok {
		return &api.RelationResponse{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		}, fmt.Errorf("用户不存在")
	}

	users := convert.UsersConvertFriends(sql.FindFollowerUsers(userId, token))

	return &api.RelationResponse{
		StatusCode: 0,
		UserList:   users,
	}, nil
}

// FriendList implements the RelationFuncImpl interface.
func (s *RelationFuncImpl) FriendList(ctx context.Context, userId int64, token string) (resp *api.RelationResponse, err error) {
	log.Println("FriendList")
	if _, ok := FindUserToken(token); !ok {
		return &api.RelationResponse{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		}, fmt.Errorf("用户不存在")
	}

	users := convert.FriendsConvert(sql.FindFriends(userId, token))

	return &api.RelationResponse{
		StatusCode: 0,
		UserList:   users,
	}, nil
}
