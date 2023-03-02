package handler

import (
	"context"
	"douSheng/cmd/favorite/convert"
	api "douSheng/cmd/favorite/kitex_gen/api"
	"douSheng/sql"
	"fmt"
	"log"
)

// FavoriteImpl implements the last service interface defined in the IDL.
type FavoriteImpl struct{}

// FavoriteAction implements the FavoriteImpl interface.
func (s *FavoriteImpl) FavoriteAction(ctx context.Context, token string, videoID int64, actionType int32) (resp *api.FavoriteResponse, err error) {
	log.Println("FavoriteAction")
	if _, exist := sql.FindUser(token); !exist {
		return &api.FavoriteResponse{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		}, fmt.Errorf("用户不存在")
	}

	if err := sql.FavoriteAction(token, videoID, int(actionType)); err != nil {
		return &api.FavoriteResponse{
			StatusCode: 1,
			StatusMsg:  "关注失败",
		}, fmt.Errorf("关注失败")
	}

	return &api.FavoriteResponse{
		StatusCode: 0,
	}, nil
}

// FavoriteList implements the FavoriteImpl interface.
func (s *FavoriteImpl) FavoriteList(ctx context.Context, token string) (resp *api.FavoriteResponse, err error) {
	log.Println("FavoriteList")
	if token == "" {
		return &api.FavoriteResponse{
			StatusCode: 1,
			StatusMsg:  "token不存在",
			Videos:     nil,
		}, fmt.Errorf("token不存在")
	}

	videos, _ := convert.VideoListConvert(ctx, sql.ReadFavoriteVideos(token))

	return &api.FavoriteResponse{
		StatusCode: 0,
		StatusMsg:  "",
		Videos:     videos,
	}, nil

}
