package sql

import (
	"douSheng/class"
	"gorm.io/gorm"
	"log"
)

func FavoriteAction(token string, videoID, actionType int) (err error) {
	var change = class.UserVideoFavorite{
		FavoriteState: actionType,
	}

	if actionType == 1 {
		if !FindFavorite(token, videoID) {
			err = db.Create(&class.UserVideoFavorite{
				Token:         token,
				VideoId:       int64(videoID),
				FavoriteState: actionType,
			}).Error
		} else {
			err = db.Model(&class.UserVideoFavorite{}).
				Where("token = ? AND video_id = ?", token, videoID).Updates(change).Error
		}
		if err != nil {
			return err
		}
		err = db.Model(&class.Video{Id: int64(videoID)}).
			Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	} else if actionType == 2 {
		if !FindFavorite(token, videoID) {
			log.Println("favorite_video doesn't exist")
		}
		err = db.Model(&class.Video{Id: int64(videoID)}).
			Update("favorite_count", gorm.Expr(" `favorite_count` - ?", 1)).Error

		if err != nil {
			return err
		}

		err = db.Model(&class.UserVideoFavorite{}).
			Where("token = ? AND video_id = ?", token, videoID).Updates(change).Error
	}

	return err
}

func FindFavorite(token string, videoID int) bool {
	result := db.Where("token = ? AND video_id = ?", token, videoID).Find(&class.UserVideoFavorite{})
	if result.RowsAffected == 0 {
		return false
	}
	return true
}
