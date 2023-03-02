package sql

import (
	"douSheng/cmd/class"
	"fmt"
)

// InsertVideo 存储文件到数据库
func InsertVideo(saveFile string, user class.User, title string, timeNow int64, saveCover string) (id int64, err error) {
	var v = class.Video{
		Title:         title,
		AuthorId:      user.Id,
		PlayUrl:       saveFile,
		CoverUrl:      saveCover,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		CreateAt:      timeNow,
		UpdateAt:      timeNow,
	}

	result := db.Create(&v)

	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("err : 存储地址失败,请重试")
	}

	return v.Id, nil
}

func InsertPublic(token string, videoId int64) error {
	var userVideo = class.UserVideoFavorite{
		Token:         token,
		VideoId:       videoId,
		FavoriteState: 0,
		PublicState:   1,
	}

	result := db.Create(&userVideo)

	if result.RowsAffected == 0 {
		return fmt.Errorf("err : 存储视频关系失败")
	}

	return nil
}
