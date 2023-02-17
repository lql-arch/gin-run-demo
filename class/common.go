package class

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type FriendUser struct {
	User
	Message string `json:"message" gorm:"column:message"`
	MsgType int    `json:"msgType"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Title         string `json:"title"`
	AuthorId      int64  `json:"author_id"` //sql
	Author        User   `json:"author" gorm:"foreignKey:id;references:author_id"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty" gorm:"-:all"`
	CreateAt      int64  `json:"create_at" gorm:"column:create_at"`
	UpdateAt      int64  `json:"update_at" gorm:"column:update_at"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"column:id"`
	UserToken  string `json:"user_token"` //sql
	Author     User   `json:"user" gorm:"foreignKey:token;references:user_token"`
	Content    string `json:"content,omitempty"`
	CreateDate int64  `json:"create_date,omitempty"`
	VideoId    int    `json:"video_id"`
	Type       int    `json:"type"`
	CId        int64  `gorm:"<-:false" json:"c_id"` //只读,禁止写
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"-:all"`
	Token         string `json:"token"`
	Uid           int64  `gorm:"<-:false" json:"uid"` //只读,禁止写
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type UserVideoFavorite struct {
	Token         string `json:"token"`
	VideoId       int64  `json:"video_id"`
	FavoriteState int    `json:"favorite_state"`
	PublicState   int    `json:"public_state"`
}

type Relation struct {
	MyId        int64 `json:"my_id"`
	OtherUserId int64 `json:"other_user_id"`
	State       int   `json:"state"`
}

type Message struct {
	Id       int    `json:"id"`
	ToUserId int    `json:"to_user_id"`
	MyId     int    `json:"from_user_id" gorm:"column:my_id"`
	Message  string `json:"content" gorm:"message"`
	CreateAt int64  `json:"create_time" gorm:"column:create_at"`
}

type SendMessage struct {
	Token      string `json:"token"`
	ToUserId   int    `json:"to_user_id"`
	ActionType int    `json:"action_type"`
	Message    string `json:"content"`
}

func (n Message) TableName() string {
	return "message"
}

func (r Relation) TableName() string {
	return "relation"
}

func (u UserVideoFavorite) TableName() string {
	return "user_video"
}

func (v Video) TableName() string { //为Video定义表名
	return "videos"
}

func (c Comment) TableName() string { //为comment定义表名
	return "comment"
}

func (user User) TableName() string { //为User定义表名
	return "user"
}

func UserSetByUserSlice(userId int, token string, p func(userId int, token string) (users []User)) (result map[int64]struct{}, users []User) {
	users = p(userId, token)
	result = make(map[int64]struct{})

	for _, user := range users {
		result[user.Id] = struct{}{}
	}

	return result, users
}
