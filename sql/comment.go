package sql

import (
	"douSheng/class"
	"douSheng/setting"
	"gorm.io/gorm"
	"log"
)

func FindComments(videoId int64, token string) (comments []class.Comment) {
	users := make(map[int64]class.User)
	myUser, _ := FindUser(token)

	//gorm , id冲突
	result := db.Table("comment c").Preload("Author").
		Select("c.user_id, c.id as c_id , c.content, c.create_date, c.video_id,u.*").
		Joins("left join user u on c.user_id = u.id").
		Where("c.video_id = ?", videoId).Order("create_date").Find(&comments)

	for i := range comments {
		comments[i].JSONCreateDate = setting.CommentTimeString(comments[i].CreateDate)

		comments[i].Id = comments[i].CId

		user, ok := users[comments[i].Id] // 自己与目标用户关系
		if !ok {
			if myUser.Id == comments[i].Author.Id {
				comments[i].Author.IsFollow = false
			} else {
				result := db.Where("my_id = ? and other_user_id = ? and state = 1", myUser.Id, comments[i].Author.Id).Find(&class.Relation{}).RowsAffected
				if result == 0 {
					comments[i].Author.IsFollow = false
				} else {
					comments[i].Author.IsFollow = true
				}
			}
			users[comments[i].Author.Id] = comments[i].Author
		} else {
			comments[i].Author.IsFollow = user.IsFollow
		}
	}

	if result.Error != nil {
		log.Println(result.Error)
	}
	return comments
}

// ReviseComment 根据actionType状态添加或者删除comment到数据库
func ReviseComment(comment class.Comment) (int64, error) {
	var result *gorm.DB

	if comment.Type == 1 { // 添加
		result = db.Create(&comment)

		if result.RowsAffected == 0 {
			return comment.Id, result.Error
		}

		result = db.Model(&class.Video{}).Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count + 1"))
	} else { // 删除
		result = db.Where("id = ?", comment.Id).Delete(&class.Comment{})

		if result.RowsAffected == 0 {
			return comment.Id, result.Error
		}

		result = db.Model(&class.Video{}).Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count - 1"))
	}

	return comment.Id, result.Error
}
