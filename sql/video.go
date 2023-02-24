package sql

import (
	"douSheng/class"
	"douSheng/setting"
	"fmt"
	"log"
	"time"
)

func init() {
	UpdateVideoData()
}

func UpdateVideoData() {
	var err error
	err, setting.VideoIds = FindVideoId()
	if err != nil {
		log.Println(err)
	}
}

func ReadVideos(latestTime int64, token string) ([]class.Video, int64) {
	//db = getDB().Begin()
	var videos []class.Video
	nextTime := int64(0)

	// 获取自身信息
	var myUser class.User
	db.Where("token = ?", token).Find(&myUser)

	//更新videoIds
	UpdateVideoData()

	// 随机取30个视频,不足则乱序取全部
	if setting.VideoIds < 30 {
		latestTime = time.Now().Unix()
		db.Table("videos v").Preload("Author").
			Select("v.id,v.`author_id`,v.`play_url`,v.`cover_url`,v.`favorite_count`,v.`comment_count`,v.title,v.create_at,v.update_at,u.id as uid,u.name,u.follow_count,u.follower_count,u.token,u.background_image,u.avatar,u.signature,u.total_favorited,u.work_count,u.favorited_count").
			Joins("left join user u on v.author_id = u.id").
			Where("update_at <= ?", latestTime).Order("rand()").Find(&videos)
	} else {
		db.Table("videos v").Preload("Author").
			Select("v.id,v.`author_id`,v.`play_url`,v.`cover_url`,v.`favorite_count`,v.`comment_count`,v.title,v.create_at,v.update_at,u.id as uid,u.name,u.follow_count,u.follower_count,u.token,u.background_image,u.avatar,u.signature,u.total_favorited,u.work_count,u.favorited_count").
			Joins("left join user u on v.author_id = u.id").
			Where("update_at <= ?", latestTime).Order("rand()").Limit(30).Find(&videos)
	}

	fmt.Println(videos)

	for i := range videos {
		if token != "" { // 如果有token
			var userVideo class.UserVideoFavorite
			result := db.Where("token = ? and video_id = ? and favorite_state = 1", token, videos[i].Id).Find(&userVideo)
			if result.RowsAffected != 0 {
				videos[i].IsFavorite = true
			} else {
				videos[i].IsFavorite = false
			}
		}

		if videos[i].Author.Id == 0 {
			videos[i].Author.Id = videos[i].Author.Uid
		}

		if myUser.Id == videos[i].Author.Id { // 自己与目标用户关系
			videos[i].Author.IsFollow = false
		} else {
			result := db.Where("my_id = ? and other_user_id = ? and state = 1", myUser.Id, videos[i].Author.Id).Find(&class.Relation{}).RowsAffected
			if result == 0 {
				videos[i].Author.IsFollow = false
			} else {
				videos[i].Author.IsFollow = true
			}
		}

		if nextTime == 0 {
			nextTime = videos[i].UpdateAt
		} else {
			tempTime := videos[i].UpdateAt
			if tempTime < nextTime {
				nextTime = tempTime
			}
		}
	}

	return videos, nextTime
}

func ReadFavoriteVideos(token string) []class.Video {
	videoIDs := GetFavoriteVideoIDByToken(token)

	var videos []class.Video

	// 获取自身信息
	var myUser class.User
	db.Where("token = ?", token).Find(&myUser)

	for _, v := range videoIDs {
		var video class.Video
		var user class.User
		db.Where("id = ?", v.VideoId).Find(&video)
		db.Where("id = ?", video.AuthorId).Find(&user)

		var userVideo class.UserVideoFavorite
		db.Where("token = ? and video_id = ?", token, v.VideoId).Find(&userVideo)
		if userVideo.FavoriteState == 1 {
			video.IsFavorite = true
		} else {
			video.IsFavorite = false
		}

		if myUser.Id == user.Id {
			user.IsFollow = false
		} else {
			result := db.Where("my_id = ? and other_user_id = ? and state = 1", myUser.Id, user.Id).Find(&class.Relation{}).RowsAffected
			if result == 0 {
				user.IsFollow = false
			} else {
				user.IsFollow = true
			}
		}

		video.Author = user

		videos = append(videos, video)
	}

	return videos
}

func ReadPublishVideos(token string) []class.Video {
	videoIDs := GetPublicVideoIDByToken(token)

	var videos []class.Video
	var user class.User
	for _, v := range videoIDs {
		var video class.Video
		db.Where("id = ?", v.VideoId).Find(&video)
		db.Where("id = ?", video.AuthorId).Find(&user)

		var userVideo class.UserVideoFavorite
		db.Where("token = ? and video_id = ?", token, v.VideoId).Find(&userVideo)
		if userVideo.FavoriteState == 1 {
			video.IsFavorite = true
		} else {
			video.IsFavorite = false
		}
		user.IsFollow = false
		video.Author = user

		videos = append(videos, video)
	}

	return videos
}

func GetFavoriteVideoIDByToken(token string) (result []class.UserVideoFavorite) {
	db.Where("favorite_state = 1 AND token = ?", token).Find(&result)
	return
}

func FindVideoId() (error, int64) {
	var id int64
	result := db.Model(&class.Video{}).Where("").Count(&id)
	if result.Error != nil {
		db.Rollback()
		return fmt.Errorf("查询出错,%v", result.Error), 0
	}

	return nil, id
}

func GetPublicVideoIDByToken(token string) (result []class.UserVideoFavorite) {
	db.Where("public_state = 1 AND token = ?", token).Find(&result)
	return
}

func FindVideoByFile(saveFile string, user class.User) bool {
	result := db.Model(&class.Video{}).Where("play_url = ? AND author_id = ?", saveFile, user.Id).First(nil)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
