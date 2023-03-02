package handler

import (
	"context"
	"douSheng/cmd/class"
	"douSheng/cmd/comment/convert"
	"douSheng/cmd/comment/kitex_gen/api"
	"douSheng/setting"
	"douSheng/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

// CommentFuncImpl implements the last service interface defined in the IDL.
type CommentFuncImpl struct{}

// CommentList implements the CommentImpl interface.
func (s *CommentFuncImpl) CommentList(ctx context.Context, videoId int64, token string) (resp *api.CommentResponse, err error) {
	log.Println("CommentList")
	list, err := GetCommentList(videoId, token)

	if err != nil {
		return &api.CommentResponse{
			StatusCode: 1,
			StatusMsg:  err.Error() + ",token不存在",
		}, fmt.Errorf(err.Error() + ",token不存在")
	}

	return &api.CommentResponse{
		StatusCode:  0,
		CommentList: convert.CommentListConvert(list),
	}, nil
}

// AddCommentAction implements the CommentImpl interface.
func (s *CommentFuncImpl) AddCommentAction(ctx context.Context, token string, actionType int32, text string, videoId int64) (resp *api.CommentResponse, err error) {
	log.Println("AddCommentAction")
	user, exist := sql.FindUser(token)

	if !exist {
		return &api.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		}, fmt.Errorf("用户不存在")
	}

	comment := class.Comment{
		Author:     user,
		UserId:     user.Id,
		Content:    text,
		VideoId:    videoId,
		Type:       int(actionType),
		CreateDate: time.Now().Unix(),
	}

	if strings.TrimSpace(text) == "" { //只有空格的或者空字符串不能发送
		return &api.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "这是一条空信息",
		}, fmt.Errorf("这是一条空信息")
	}

	// 添加comment到数据库
	id, err := sql.ReviseComment(comment)

	if err != nil {
		return &api.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "发布失败:" + err.Error(),
		}, fmt.Errorf("发布失败:" + err.Error())
	}

	comment.Id = id
	comment.JSONCreateDate = setting.CommentTimeString(comment.CreateDate)

	return &api.CommentResponse{
		StatusCode: 0,
		Comment:    convert.CommentConvert(&comment),
	}, nil
}

// DeleteCommentAction implements the CommentImpl interface.
func (s *CommentFuncImpl) DeleteCommentAction(ctx context.Context, token string, actionType int32, commentId int64, videoId int64) (resp *api.CommentResponse, err error) {
	log.Println("DeleteCommentAction")
	_, exist := sql.FindUser(token)

	if !exist {
		return &api.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		}, fmt.Errorf("用户不存在")
	}

	comment := class.Comment{
		Id:      commentId,
		Type:    int(actionType),
		VideoId: videoId,
	}

	// 删除comment到数据库
	if _, err := sql.ReviseComment(comment); err != nil {
		return &api.CommentResponse{
			StatusCode: 1,
			StatusMsg:  "删除失败",
		}, fmt.Errorf("删除失败")
	}

	return &api.CommentResponse{
		StatusCode: 0,
		StatusMsg:  "删除成功",
	}, nil
}

func GetCommentList(videoId int64, token string) (list []*class.Comment, err error) {
	if _, ok := FindUserToken(token); !ok {
		return nil, fmt.Errorf("查询出错")
	}

	return sql.FindComments(videoId, token), nil
}
