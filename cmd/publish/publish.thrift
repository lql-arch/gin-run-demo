namespace go api

include "../feed/feed.thrift"

struct PublishResponse {
    1: i32    StatusCode    (go.tag = 'json:"status_code"')
	2: string StatusMsg     (go.tag = 'json:"status_msg,omitempty"')
    3: list<feed.Video> videos  (go.tag = 'json:"video_list"')
}

struct UserVideoFavorite {
    1 : string Token        (go.tag = 'json:"token"' )
    2 : i64 VideoId         (go.tag = 'json:"video_id"' )
    3 : i32 FavoriteState   (go.tag = 'json:"favorite_state"' )
    4 : i32 PublicState     (go.tag = 'json:"public_state"' )
}

struct videoData {
    1: string title
    2: binary data
    3: string fileName
}

service publish {
    PublishResponse Publish(1:string token ,2: videoData video )
    PublishResponse PublishList(1:string token)
}