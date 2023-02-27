package sql

import (
	"douSheng/class"
	"gorm.io/gorm"
	"log"
)

func FavoriteAction(token string, videoID int64, actionType int) (err error) {
	var change = class.UserVideoFavorite{
		FavoriteState: actionType,
	}

	if actionType == 1 {
		if !FindFavorite(token, videoID) {
			err = db.Create(&class.UserVideoFavorite{
				Token:         token,
				VideoId:       videoID,
				FavoriteState: actionType,
			}).Error
		} else {
			err = db.Model(&class.UserVideoFavorite{}).
				Where("token = ? AND video_id = ?", token, videoID).Updates(change).Error
		}

		if err != nil {
			return err
		}

		// 视频点赞
		err = db.Model(&class.Video{Id: videoID}).
			Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error

		// 个人点赞
		tmp, _ := FindUser(token)
		db.Model(&class.User{}).Where("id = ?", tmp.Id).Update("favorited_count", gorm.Expr("favorited_count + 1"))
		tmp = GetUserIdByVideoId(videoID)
		db.Model(&class.User{}).Where("id = ?", tmp.Id).Update("total_favorited", gorm.Expr("total_favorited + 1"))
	} else if actionType == 2 {
		if !FindFavorite(token, videoID) {
			log.Println("favorite_video doesn't exist")
		}
		// 视频点赞
		err = db.Model(&class.Video{Id: videoID}).
			Update("favorite_count", gorm.Expr(" `favorite_count` - ?", 1)).Error

		// 个人点赞
		tmp, _ := FindUser(token)
		db.Model(&class.User{}).Where("id = ?", tmp.Id).Update("favorited_count", gorm.Expr("favorited_count - 1"))
		tmp = GetUserIdByVideoId(videoID)
		db.Model(&class.User{}).Where("id = ?", tmp.Id).Update("total_favorited", gorm.Expr("total_favorited - 1"))

		if err != nil {
			return err
		}

		// 修改关注视频状态
		err = db.Model(&class.UserVideoFavorite{}).
			Where("token = ? AND video_id = ?", token, videoID).Updates(change).Error
	}

	return err
}

func FindFavorite(token string, videoID int64) bool {
	result := db.Where("token = ? AND video_id = ?", token, videoID).Find(&class.UserVideoFavorite{})
	if result.RowsAffected == 0 {
		return false
	}
	return true
}
