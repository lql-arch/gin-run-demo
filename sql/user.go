package sql

import (
	"douSheng/class"
	"gorm.io/gorm"
)

// FindUser 根据token取出User
func FindUser(token string) (class.User, bool) {
	var user class.User
	result := db.Where("token = ?", token).Find(&user).RowsAffected
	if result == 0 {
		return user, false
	}

	return user, true
}

// FindUserIdSequence 取数据库中id的最大值(用于添加新数据),而不是拥有的数量
func FindUserIdSequence() int64 {
	var user class.User
	db.Select("id").Where("").Limit(1).Order("id DESC").Find(&user)
	return user.Id
}

func Info() map[string]class.User {
	ans := make(map[string]class.User)

	var users []class.User
	db.Where("").Find(&users)

	for _, user := range users {
		ans[user.Token] = user
	}

	return ans
}

// InsertUser 添加数据到数据库
func InsertUser(users []class.User) ([]int64, error) {
	result := db.Create(&users)

	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	var ids []int64

	for _, user := range users {
		ids = append(ids, user.Id)
	}

	return ids, nil
}

func GetUserByToken(token string) class.User {
	var user class.User
	db.Where("token = ?", token).Find(&user)

	return user
}

func CheckUser(userId int64, token string) error {

	err := db.Where("id = ? AND token = ?", userId, token).Find(&class.User{}).Error

	return err
}

func GetUserIdByVideoId(id int64) class.User {
	var user class.User
	var userVideo class.UserVideoFavorite

	db.Select("token").Where("video_id = ?", id).Find(&userVideo)
	db.Select("id").Where("token = ?", userVideo.Token).Find(&user)

	return user
}

func InsertUserPublicCount(id int64, flag bool) {

	if flag { // 增加
		db.Model(&class.User{}).Where("id = ?", id).Update("work_count", gorm.Expr("work_count + 1"))
	} else { // 删除
		db.Model(&class.User{}).Where("id = ?", id).Update("work_count", gorm.Expr("work_count - 1"))
	}

}
