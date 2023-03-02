package convert

import (
	"context"
	"douSheng/cmd/class"
	"douSheng/cmd/publish/kitex_gen/api"
)

// VideoListConvert class.video -> api.video
func VideoListConvert(ctx context.Context, videoList []*class.Video) ([]*api.Video, error) {
	newVideoList := make([]*api.Video, 0)
	for _, video := range videoList {
		video := VideoConvert(ctx, video)
		if video == nil {
			continue
		}
		newVideoList = append(newVideoList, video)
	}
	//fmt.Printf("type:%T\n", newVideoList)
	return newVideoList, nil
}

func VideoConvert(ctx context.Context, video *class.Video) *api.Video {
	return &api.Video{
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
