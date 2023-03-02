package sql

import (
	"douSheng/cmd/class"
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

// InsertUser 添加数据到数据库
func InsertUser(users *[]class.User) ([]class.User, error) {
	result := db.Create(&users)

	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return *users, nil
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

func InsertUserPublicCount(id int64, flag bool) error {
	var err error
	if flag { // 增加
		err = db.Model(&class.User{}).Where("id = ?", id).Update("work_count", gorm.Expr("work_count + 1")).Error
	} else { // 删除
		err = db.Model(&class.User{}).Where("id = ?", id).Update("work_count", gorm.Expr("work_count - 1")).Error
	}
	return err
}
