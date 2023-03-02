package main

import (
	rpc2 "douSheng/cmd/rpc"
	"douSheng/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	initRPC()
}

func initRPC() {
	rpc2.InitFeed()
	rpc2.InitUser()
	rpc2.InitFavorite()
	rpc2.InitPublish()
	rpc2.InitComment()
	rpc2.InitRelation()
	rpc2.InitMessage()
}

func main() {
	r := gin.Default()

	initRouter(r)
	go testPprof()

	err := r.Run()
	if err != nil {
		panic("run failed.")
	}
}

func testPprof() {
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}

func initRouter(r *gin.Engine) {
	// public_videos and public_cover directory is used to serve static resources
	r.Static("/static", "./public_videos") // 视频
	r.Static("/jpg", "./public_cover")     // 封面

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)

}
