namespace go api

struct MessageResponse {
    1: i32    StatusCode    (go.tag = 'json:"status_code"')
	2: string StatusMsg     (go.tag = 'json:"status_msg,omitempty"')
	3: list<Message> message      (go.tag = 'json:"message_list"')
}

struct Message {
    1: i64 Id           (go.tag = 'json:"id"' )
    2: i64 ToUserId     (go.tag = 'json:"to_user_id"' )
    3: i64 MyId         (go.tag = 'json:"from_user_id" gorm:"column:my_id"' )
    4: string Message   (go.tag = 'json:"content" gorm:"message"' )
    5: i64 CreateAt     (go.tag = 'json:"create_time" gorm:"column:create_at"' )
}

service MessageFunc {
    MessageResponse MessageAction(1: string token,2: string content,3:i64 toUserId,4: i32 actionType)
    MessageResponse MessageChat(1: string token , 2: i64 preMsgTime,3: i64 toUserId)
}