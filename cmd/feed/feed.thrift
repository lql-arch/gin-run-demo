namespace go api

include "../user/user.thrift"

struct FeedList {
    1: list<Video> VideoList        (go.tag ='json:"video_list,omitempty"')
    2: i64 NextTime                 (go.tag ='son:"next_time,omitempty"')
    3: string token
    4: i32    StatusCode    (go.tag = 'json:"status_code"')
    5: string StatusMsg     (go.tag = 'json:"status_msg,omitempty"')
}

struct Video {
	1:  i64             Id (go.tag = 'json:"id,omitempty"' )
    2:  string          Title (go.tag = 'json:"title"' )
    3:  i64             AuthorId (go.tag = 'json:"author_id"' )
    4:  user.User            Author (go.tag = 'json:"author" gorm:"foreignKey:id;references:author_id"' )
    5:  string 	        PlayUrl (go.tag = 'json:"play_url" json:"play_url,omitempty"' )
    6:  string	        CoverUrl (go.tag = 'json:"cover_url,omitempty"' )
    7:  i64	            FavoriteCount (go.tag = 'json:"favorite_count"' )
    8:  i64	            CommentCount (go.tag = 'json:"comment_count"' )
    9:  bool	        IsFavorite (go.tag = 'json:"is_favorite,omitempty" gorm:"-:all"' )
    10: i64	            CreateAt (go.tag = 'json:"create_at" gorm:"column:create_at"' )
    11: i64	            UpdateAt (go.tag = 'json:"update_at" gorm:"column:update_at"' )
}

service Feed {
     FeedList ReadVideos( 1:i64 latestTime , 2:string token )
}