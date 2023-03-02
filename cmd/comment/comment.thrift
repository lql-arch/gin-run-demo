namespace go api

include "../user/user.thrift"

struct CommentResponse {
    1: i32              StatusCode      (go.tag = 'json:"status_code"')
	2: string           StatusMsg       (go.tag = 'json:"status_msg,omitempty"')
    3: Comment          comment         (go.tag = 'json:"comment,omitempty"')
    4: list<Comment>    commentList     (go.tag = 'json:"comment_list,omitempty"')
}

struct Comment {
    1:i64               Id (go.tag = 'json:"id,omitempty" gorm:"column:id"' )
    2:i64               UserId (go.tag = 'json:"user_id"' )
    3:user.User         Author (go.tag = 'json:"user" gorm:"foreignKey:id;references:user_id"' )
    4:string            Content (go.tag = 'json:"content,omitempty"' )
    5:i64               VideoId (go.tag = 'json:"video_id"' )
    6:i32               Type (go.tag = 'json:"type"' )
    7:i64               CId (go.tag = 'gorm:"<-:false" json:"c_id"' )
    8:i64               CreateDate (go.tag = 'json:"abandon,omitempty" gorm:"column:create_date"' )
    9:string            JSONCreateDate (go.tag = 'json:"create_date,omitempty" gorm:"column:abandon"')
}

service commentFunc {
    CommentResponse CommentList(1: i64 videoId , 2: string token)
    CommentResponse AddCommentAction(1: string token,2: i32 actionType ,3: string text,4: i64 videoId )
    CommentResponse DeleteCommentAction(1: string token,2: i32 actionType ,3: i64 commentId,4: i64 videoId)
}

