## 业务逻辑层

- user
  - Register(c *gin.Context)
    - 注册用户

  - Login(c *gin.Context)
    - 登录用户,检查用户是否存在

  - UserInfo(c *gin.Context)
    - 查询用户信息

- feed
  - Feed(c *gin.Context)
    - 查询最近时间至多30个视频信息返回

- favorite
  - FavoriteList(c *gin.Context)
    - 如果token验证通过,就返回我所有喜爱的视频列表

  - FavoriteAction(c *gin.Context)
    - 如果token验证通过,就将视频添加到我喜欢列表

- publish
  - Publish(c *gin.Context)
    - 如果token验证通过,判断是否是视频,保存封面,将视频提交到服务器,视频文件信息提交到服务器数据库

  - GetSnapshot(videoPath, snapshotPath string, frameNum int)
    - 根据视频截取视频封面

  - PublishList(c *gin.Context)
    - 如果token验证通过,就返回我所有发布的视频列表

- relation
  - RelationAction(c *gin.Context)
    - 如果token验证通过,就将to_user_id用户加入自己的关注列表

  - FollowList(c *gin.Context)
    - 如果token验证通过,就返回我的关注列表

  - FollowerList(c *gin.Context)
    - 如果token验证通过,就返回我的被关注列表

  - FriendList(c *gin.Context)
    - 如果token验证通过,就返回我的好友列表

- comment
  - CommentAction(c *gin.Context)
    - 先进行token验证,不通过直接返回错误
    - 根据action_type判断是发布还是删除
    - 如果是发布,判断comment_text是空字符串,则直接返回错误,非空则将数据添加,
    - 删除就根据comment_id将comment删除

  - CommentList(c *gin.Context)
    - 如果token和videoId验证通过,返回根据videoId获取的评论信息

- message
  - MessageAction(c *gin.Context)
    - 如果content是空字符串,则直接返回错误.
    - 如果content是非空字符串且token验证通过,就添加信息到数据库

  - MessageChat(c *gin.Context)
    - 对token进行验证,如果不通过就返回错误
    - 读取该用户与toUserId对应用户最新的聊天记录

  - timeToken(myId int64, ToUserId string) 
    - 构建myId与ToUserId形成timeToken,键值对存储最后一次读的时间,以读取最新消息

- token
  - GenerateToken(username, password string) (string, error)
    - GenerateToken 根据用户的用户名和密码产生token

  - ParseToken(token string) (*Claims, error)
    - ParseToken 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）

  - Substring(token string, size int)
    - 从开头截取token的长为size内容

  - FindUserToken(token string) (class.User, bool)
    - FindUserToken 查询用户的token,先在tokenList中查询,如果不存在,就在数据库中查询.
    - 缺少实时性,如果使用user,则尽可能使用不能被修改的数据

<hr>

## 缓存层(redis)
> 预计信息先存到redis里,等待一定时间或者数量再发送到数据库
- 

<hr>

## 持久化层(mysql)

- ReadVideos(latestTime int64, token string) ([]class.Video, int64)
  - 查询数据库中的最迟为lastestTime的30个视频信息,如果token不为空,则查询自身与视频及其视频作者的关系,
  - 如果视频数不够30就查询全部视频
  - 返回视频信息列表与最早的一个视频上传时间

- 
