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
    - 如果token验证通过,

  - CommentList(c *gin.Context)
    - 如果


<hr>

## 缓存层(redis)

- 

<hr>

## 持久化层(mysql)

- ReadVideos(latestTime int64, token string) ([]class.Video, int64)
  - 查询数据库中的最迟为lastestTime的30个视频信息,如果token不为空,则查询自身与视频及其视频作者的关系,
  - 如果视频数不够30就查询全部视频
  - 返回视频信息列表与最早的一个视频上传时间

- 
