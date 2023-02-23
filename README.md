# douSheng

## 抖声项目服务端

使用前先添加数据库sqlStatement/create.sql  
在Const/theConst中修改参数  

使用gin开发,使用了gorm,ffmpeg,kitex(待定)等外部库,需要先进行go mod管理,再编译运行

```shell
go build && ./douSheng
```

### 一些注意

- 使用gorm预编译防止sql注入
- 使用ffmpeg "github.com/u2takey/ffmpeg-go",进行视频判断并且截图作为视频封面

### 功能说明

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
  - 