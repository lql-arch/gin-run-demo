package sql

import (
	"douSheng/class"
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
