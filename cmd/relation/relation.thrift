namespace go api

include "../user/user.thrift"

struct relationResponse {
    1: i32    StatusCode    (go.tag = 'json:"status_code"')
	2: string StatusMsg     (go.tag = 'json:"status_msg,omitempty"')
    3: list<FriendUser> UserList      (go.tag = 'json:"user_list"')
}

struct FriendUser {
    1: i64 Id                 (go.tag = 'json:"id,omitempty"' )   // 用户id
    2: string Token           (go.tag = 'json:"token"')           //
    3: string Name            (go.tag = 'json:"name,omitempty"' )// 用户名称
      4: i64 FollowCount        (go.tag = 'json:"follow_count"' )// 关注总数
      5: i64 FollowerCount      (go.tag = 'json:"follower_count"' ) // 粉丝总数
      6: bool IsFollow          (go.tag = 'json:"is_follow" gorm:"-:all"' )// true-已关注，false-未关注
      7: string Avatar          (go.tag = 'json:"avatar"' ) //用户头像
      8: string BackgroundImage (go.tag = 'json:"background_image"' ) //用户个人页顶部大图
      9: string Signature       (go.tag = 'json:"signature"' )//个人简介
      10: i64 TotalFavorited    (go.tag = 'json:"total_favorited" gorm:"column:total_favorited"' )//获赞数量
      11: i64 WorkCount         (go.tag = 'json:"work_count"' )//作品数量
      12: i64 FavoriteCount     (go.tag = 'json:"favorite_count" gorm:"column:favorited_count"' )//点赞数量
      13: i64 Uid               (go.tag = 'gorm:"<-:false" json:"uid"' )
    14: string   Message (go.tag = 'json:"message" gorm:"column:message"' )
    15: i32      MsgType (go.tag = 'json:"msgType"' )
}

struct Relation {
    1: i64 MyId         (go.tag = 'json:"my_id"' )
    2: i64 OtherUserId  (go.tag = 'json:"other_user_id"' )
    3: i32 State        (go.tag = 'json:"state"' )
}

service relationFunc {
    relationResponse RelationAction(1:string token, 2: i64 toUserId ,3:i32 state)
    relationResponse FollowList(1:i64 userId , 2: string token)
    relationResponse FollowerList(1:i64 userId , 2: string token)
    relationResponse FriendList(1:i64 userId , 2: string token)
}