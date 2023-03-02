package convert

import (
	"douSheng/cmd/class"
	"douSheng/cmd/relation/kitex_gen/api"
)

func UsersConvertFriends(users []*class.User) []*api.FriendUser {
	newFriends := make([]*api.FriendUser, 0)
	for _, user := range users {
		newUser := UserConvertFriend(user)
		if newUser == nil {
			continue
		}
		newFriends = append(newFriends, newUser)
	}

	return newFriends
}

func UserConvertFriend(user *class.User) *api.FriendUser {
	return &api.FriendUser{
		Id:              user.Id,
		Token:           user.Token,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorite,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
}

func FriendsConvert(users []*class.FriendUser) []*api.FriendUser {
	newFriends := make([]*api.FriendUser, 0)
	for _, user := range users {
		newUser := FriendConvert(user)
		if newUser == nil {
			continue
		}
		newFriends = append(newFriends, newUser)
	}

	return newFriends
}

func FriendConvert(user *class.FriendUser) *api.FriendUser {
	return &api.FriendUser{
		Id:              user.Id,
		Token:           user.Token,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorite,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
		Message:         user.Message,
		MsgType:         int32(user.MsgType),
	}
}
