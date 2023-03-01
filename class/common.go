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
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite,omitempty" gorm:"-:all"`
	CreateAt      int64  `json:"create_at" gorm:"column:create_at"`
	UpdateAt      int64  `json:"update_at" gorm:"column:update_at"`
}

type Comment struct {
	Id             int64  `json:"id,omitempty" gorm:"column:id"`
	UserId         int64  `json:"user_id"`                                      //sql
	Author         User   `json:"user" gorm:"foreignKey:id;references:user_id"` // 评论用户信息
	Content        string `json:"content,omitempty"`                            // 评论内容
	VideoId        int64  `json:"video_id"`                                     // 视频评论id
	Type           int    `json:"type"`                                         // 2删除或1创建
	CId            int64  `gorm:"<-:false" json:"c_id"`                         //只读,禁止写
	JSONCreateDate string `json:"create_date,omitempty" gorm:"-:all"`  // 评论发布日期，格式 mm-dd
	CreateDate     int64  `json:"abandon,omitempty" gorm:"column:create_date"`  // 评论发布日期，格式 mm-dd
}

type User struct {
	Id              int64  `json:"id,omitempty"`
	Token           string `json:"token"`
	Name            string `json:"name,omitempty"`                                // 用户名称
	FollowCount     int64  `json:"follow_count"`                                  // 关注总数
	FollowerCount   int64  `json:"follower_count"`                                // 粉丝总数
	IsFollow        bool   `json:"is_follow" gorm:"-:all"`                        // true-已关注，false-未关注
	BackgroundImage string `json:"background_image"`                              //用户个人页顶部大图
	Signature       string `json:"signature"`                                     //个人简介
	Avatar          string `json:"avatar"`                                        //用户头像
	TotalFavorite   int64  `json:"total_favorited" gorm:"column:total_favorited"` //获赞数量
	WorkCount       int64  `json:"work_count"`                                    //作品数量
	FavoriteCount   int64  `json:"favorite_count" gorm:"column:favorited_count"`  //点赞数量
	Uid             int64  `gorm:"<-:false" json:"uid"`                           //只读,禁止写
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
	ToUserId int64  `json:"to_user_id"`
	MyId     int64  `json:"from_user_id" gorm:"column:my_id"`
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

func UserSetByUserSlice(userId int64, token string, p func(userId int64, token string) (users []User)) (result map[int64]struct{}, users []User) {
	users = p(userId, token)
	result = make(map[int64]struct{})

	for _, user := range users {
		result[user.Id] = struct{}{}
	}

	return result, users
}
