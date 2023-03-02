namespace go api

include "../feed/feed.thrift"
include "../user/user.thrift"

struct FavoriteResponse {
    1: i32    StatusCode        (go.tag = 'json:"status_code"')
	2: string StatusMsg         (go.tag = 'json:"status_msg,omitempty"')
	4: list<feed.Video> videos  (go.tag = 'json:"video_list"')
}

service favorite {
    FavoriteResponse FavoriteAction(1:string token ,2:i64 videoID,3:i32 actionType)
    FavoriteResponse FavoriteList(1:string token)
}