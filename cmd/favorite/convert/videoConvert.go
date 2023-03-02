package convert

import (
	"context"
	"douSheng/cmd/class"
	favorite "douSheng/cmd/favorite/kitex_gen/api"
)

// VideoListConvert class.video -> api.video
func VideoListConvert(ctx context.Context, videoList []*class.Video) ([]*favorite.Video, error) {
	newVideoList := make([]*favorite.Video, 0)
	for _, video := range videoList {
		video := VideoConvert(ctx, video)
		if video == nil {
			continue
		}
		newVideoList = append(newVideoList, video)
	}

	return newVideoList, nil
}

func VideoConvert(ctx context.Context, video *class.Video) *favorite.Video {
	return &favorite.Video{
		Id:            video.Id,
		Title:         video.Title,
		AuthorId:      video.AuthorId,
		Author:        UserConvert(video.Author),
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
		CreateAt:      video.CreateAt,
		UpdateAt:      video.UpdateAt,
	}
}
