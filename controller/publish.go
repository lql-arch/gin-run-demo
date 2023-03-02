package controller

import (
	"bytes"
	"douSheng/cmd/class"
	"douSheng/cmd/publish/kitex_gen/api"
	"douSheng/cmd/rpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	class.Response
	VideoList []class.Video `json:"video_list"`
	//NextTime  int64         `json:"next_time"`
}

// Publish 需要事务,1未处理
// Publish check token then save upload file to public_videos directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	data, err := c.FormFile("data")

	if err != nil {
		c.JSON(http.StatusOK, Errorf(fmt.Errorf("上传的文件异常")))
		return
	}

	file, err := data.Open()
	if err != nil {
		c.JSON(http.StatusOK, Errorf(fmt.Errorf("上传的文件异常")))
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		c.JSON(http.StatusOK, Errorf(fmt.Errorf("保存文件出错")))
		return
	}

	videoData := &api.VideoData{
		Title:    title,
		Data:     buf.Bytes(),
		FileName: filepath.Base(data.Filename),
	}

	resp, err := rpc.PublishAction(c, token, videoData)

	if err != nil {
		c.JSON(http.StatusOK, Errorf(err))
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}

func PublishList(c *gin.Context) {
	token := c.Query("token")

	resp, err := rpc.PublishList(c, token)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, *Errorf(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
