# douSheng

## 抖声项目服务端

使用前先添加数据库sqlStatement/create.sql  
在Const/theConst中修改参数  

使用gin开发,使用了gorm,ffmpeg,kitex(待定)等外部库,需要先进行go mod管理,再编译运行

> 使用go版本为19.5(不要使用大于等于1.20的版本)   
> 开发环境为linux/amd64

```shell
go build && ./douSheng
```

### 一些注意

- 使用gorm预编译防止sql注入
- 使用ffmpeg "github.com/u2takey/ffmpeg-go",进行视频判断并且截图作为视频封面
- 使用jwt-go用账户和密码生成token.

### 项目功能说明
- Class : 项目使用的数据结构
- Const : 项目启动前,需要配置的一些参数
- Controller : 项目的业务逻辑层
- KClient/kitexService/kitex : 预计未来使用的kitex设计
- public_cover : 存储视频封面或者用户封面的目录,预计未来将分离.
- public_videos : 存储视频的目录,预计未来将分离.
- Service : 包含使用kitex和一些检测修改服务器的设计.
- Setting : 目前使用中的参数设置,预计未来合并入Const.
- Sql : 项目的持久化层,将数据读取,存储,修改入数据库.
- sqlStatement : 一些sql代码,包含mysql的初始化文件(DDL).

### 接口功能说明

- /douyin/feed/ - 视频流接口
  - 随机抽取视频,以传入的时间为最迟(传入的时间一般为now())
  - 返回视频列表和找到的最早的视频更新时间

- /douyin/user/register/ - 用户注册接口
  - 对用户账户和密码拼接后的token进行简单的加密
  - 将符合条件的用户添加入数据库

- /douyin/user/login/ - 用户登录接口
    - 使用用户账户和密码拼接后的token加密后的数据进行验证
    - 正确则返回部分用户信息

- /douyin/user/ - 用户信息
  - 使用用户账户和密码拼接后的token加密后的数据进行验证
  - 返回验证成功的用户信息

- /douyin/publish/action/ - 视频投稿
  - 登录用户选择视频上传。

- /douyin/publish/list/ - 发布列表
  - 登录用户的视频发布列表，直接列出用户所有投稿过的视频

- /douyin/favorite/action/ - 赞操作
  - 登录用户对视频的点赞和取消点赞操作。

- /douyin/favorite/list/ - 喜欢列表
  - 登录用户的所有点赞视频。

- /douyin/comment/action/ - 评论操作
  - 登录用户对视频进行评论。

- /douyin/comment/list/ - 视频评论列表
  - 查看视频的所有评论，按发布时间倒序。

- /douyin/relation/action/ - 关系操作
  - 登录用户对其他用户进行关注或取消关注。

- /douyin/relatioin/follow/list/ - 用户关注列表
  - 登录用户关注的所有用户列表。

- /douyin/relation/friend/list/ - 用户好友列表
  - 所有关注登录用户的粉丝列表。

- /douyin/message/chat/ - 聊天记录
  - 当前登录用户和其他指定用户的聊天消息记录

- /douyin/message/action/ - 消息操作
  - 登录用户对消息的相关操作，目前只支持消息发送