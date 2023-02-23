namespace go api


struct Response {
    1: i32    status
	2: string message
}


struct Video {
	1:  i64             Id
    2:  string          Title
    3:  i64             AuthorId
    4:  User   Author
    5:  string 	        PlayUrl
    6:  string	        CoverUrl
    7:  i64	            FavoriteCount
    8:  i64	            CommentCount
    9:  bool	        IsFavorite
    10: i64	            CreateAt
    11: i64	            UpdateAt
}

//struct User {
//    1: i64  Id
//    2: string Name
//    3: i64 FollowCount
//    4: i64 FollowerCount
//    5: bool IsFollow
//    6: string token
//}

struct FriendUser {
    1: User     user
    2: string   Message
    3: i32      MsgType
}

struct Comment {
    1:i64       Id
    2:string    UserToken
    3:User      Author
    4:string    Content
    5:i64       CreateDate
    6:i64       VideoId
    7:i64       Type
    8:i64       CId
}

struct User {
  1: i64 Id // 用户id
  2: string Name  // 用户名称
  3: i64 FollowCount  // 关注总数
  4: i64 FollowerCount  // 粉丝总数
  5: bool IsFollow  // true-已关注，false-未关注
  6: string Avatar  //用户头像
  7: string BackgroundImage  //用户个人页顶部大图
  8: string Signature  //个人简介
  9: i64 TotalFavorited  //获赞数量
  10: i64 WorkCount  //作品数量
  11: i64 FavoriteCount  //点赞数量
}

struct UserVideoFavorite {
    1 : string Token
    2 : i64 VideoId
    3 : i32 FavoriteState
    4 : i32 PublicState
}

struct Relation {
    1: i64 MyId
    2: i64 OtherUserId
    3: i32 State
}

struct Message {
    1: i32 Id
    2: i64 ToUserId
    3: i64 MyId
    4: string Message
    5: i64 CreateAt
}

service Echo {
    Response Video(1: Video req)
    Response User(1: User req)
    Response FriendUser(1: FriendUser req)
    Response Comment(1: Comment req)
    Response UserVideoFavorite(1: UserVideoFavorite req)
    Response Relation(1: Relation req)
    Relation Message(1: Message req)
}

