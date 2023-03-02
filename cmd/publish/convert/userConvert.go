package convert

import (
	"douSheng/cmd/class"
	"douSheng/cmd/publish/kitex_gen/api"
)

func UserConvert(user class.User) *api.User {
	return &api.User{
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
