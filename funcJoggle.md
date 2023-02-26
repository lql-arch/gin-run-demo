## 持久化层

- Feed(c *gin.Context)
  - 查询最近时间至多30个视频信息返回


## 缓存层(redis)

- 

## 持久化层(mysql)

- ReadVideos(latestTime int64, token string) ([]class.Video, int64)
  - 查询数据库中的最迟为lastestTime的30个视频信息,如果token不为空,则查询自身与视频及其视频作者的关系,
  - 如果视频数不够30就查询全部视频
  - 返回视频信息列表与最早的一个视频上传时间

- 
