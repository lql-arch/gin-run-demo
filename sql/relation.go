package sql

import (
	"douSheng/class"
	"gorm.io/gorm"
	"log"
)

// 登录后必须置空
var follows map[int64]struct{}

func init() {
	follows = make(map[int64]struct{})
}

// Reset 置空关注状态
func Reset() {
	follows = make(map[int64]struct{})
}

// RelationAction 根据state判断添加或者取消关注
func RelationAction(myToken string, toUserId int64, state int) error {
	var user class.User

	err := db.Where("token = ?", myToken).Find(&user).Error

	if err != nil {
		return err
	}

	var relation = class.Relation{
		MyId:        user.Id,
		OtherUserId: toUserId,
		State:       state,
	}

	var result *gorm.DB

	if IsRelationExist(relation) {
		result = db.Where(&class.Relation{
			MyId:        user.Id,
			OtherUserId: toUserId,
		}).Updates(&class.Relation{
			State: state,
		})
		err = result.Error
	} else {
		result = db.Create(&relation)
		err = result.Error
	}

	if result.RowsAffected != 0 {
		return err
	}

	if state == 1 {
		db.Model(&class.User{}).Where("id = ?", user.Id).Update("follow_count", gorm.Expr("follow_count + ?", 1))
		db.Model(&class.User{}).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1))
	} else if state == 2 {
		db.Model(&class.User{}).Where("id = ?", user.Id).Update("follow_count", gorm.Expr("follow_count - ?", 1))
		db.Model(&class.User{}).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - ?", 1))
	}

	return nil
}

func IsRelationExist(relation class.Relation) bool {

	result := db.Model(&relation).Where("my_id = ? AND other_user_id = ? ", relation.MyId, relation.OtherUserId).Find(nil)

	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func FindFollowUsers(userId int64, token string) (users []class.User) {
	if err := CheckUser(userId, token); err != nil {
		log.Println(err)
		return nil
	}

	db.Select("`user`.*").
		Joins("left join relation r on user.id = r.other_user_id").
		Where("my_id = ? and r.state = ?", userId, 1).Find(&users)

	Reset()
	for i := range users {
		follows[users[i].Id] = struct{}{}
		users[i].IsFollow = true
	}

	return users
}

func FindFollowerUsers(userId int64, token string) (users []class.User) {
	_ = FindFollowUsers(userId, token)

	if err := CheckUser(userId, token); err != nil {
		log.Println(err)
		return nil
	}

	db.Select("`user`.*").
		Joins("left join relation r on user.id = r.my_id").
		Where("other_user_id = ?", userId).Find(&users)

	for i, user := range users { // 必须先使用FindFollowUsers(),才能获取喜爱信息
		if _, ok := follows[user.Id]; ok {
			users[i].IsFollow = true
		} else {
			users[i].IsFollow = false
		}
		follows[users[i].Id] = struct{}{}
	}

	return users
}

// FindFriends 只有我关注且关注我的才能看见
func FindFriends(userId int64, token string) (userFriends []class.FriendUser) {
	var followUsers []class.User
	//var followerUsers []class.User
	// follows用于避免关注者和粉丝重复
	follows, followUsers = class.UserSetByUserSlice(userId, token, FindFollowUsers)
	followers, _ := class.UserSetByUserSlice(userId, token, FindFollowerUsers)

	// 我的关注者
	for _, user := range followUsers {
		if _, ok := followers[user.Id]; !ok { // 判断是否关注我
			continue
		}

		var message class.Message
		db.Where("(my_id = ? and to_user_id = ?) or (my_id = ? and to_user_id = ?)", userId, user.Id, user.Id, userId).Order("create_at DESC").Limit(1).Find(&message)

		var friendMessage class.FriendUser
		friendMessage.User = user
		friendMessage.Message = message.Message
		if message.MyId == userId { //1 => 当前请求用户发送的消息
			friendMessage.MsgType = 1
		} else { //0 => 当前请求用户接收的消息
			friendMessage.MsgType = 0
		}

		userFriends = append(userFriends, friendMessage)
	}

	// 我的粉丝,只有发信息才会出现(已抛弃)
	//for _, user := range followerUsers {
	//	if _, ok := follows[user.Id]; ok { // 判断是否已经出现
	//		continue
	//	}
	//
	//	var message class.Message
	//	db.Where("(my_id = ? and to_user_id = ?) or (my_id = ? and to_user_id = ?)", userId, user.Id, user.Id, userId).Order("create_at DESC").Limit(1).Find(&message)
	//
	//	if message.Message == "" { // 没有消息就跳过
	//		continue
	//	}
	//
	//	var friendMessage class.FriendUser
	//	friendMessage.Message = message.Message
	//	friendMessage.User = user
	//	if message.MyId == userId { //1 => 当前请求用户发送的消息
	//		friendMessage.MsgType = 1
	//	} else { //0 => 当前请求用户接收的消息
	//		friendMessage.MsgType = 0
	//	}
	//
	//	userFriends = append(userFriends, friendMessage)
	//}
	return
}
