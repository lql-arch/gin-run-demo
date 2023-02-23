package controller

import (
	"bytes"
	"douSheng/Const"
	"douSheng/class"
	"douSheng/setting"
	"douSheng/sql"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type VideoListResponse struct {
	class.Response
	VideoList []class.Video `json:"video_list"`
	NextTime  int64         `json:"next_time"`
}

// Publish 需要事务,1未处理
// Publish check token then save upload file to public_videos directory
func Publish(c *gin.Context) {
	log.Println("Publish 需要事务,1未处理")
	token := c.PostForm("token")
	title := c.PostForm("title")

	user, exist := sql.FindUser(token)
	if !exist {
		c.JSON(http.StatusOK, class.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	//得到文件信息
	data, err := c.FormFile("data")

	if err != nil {
		c.JSON(http.StatusOK, class.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//得到文件名
	filename := filepath.Base(data.Filename)

	//构建新文件名
	finalName := fmt.Sprintf("%d_%d_%s", user.Id, setting.VideoIds, filename)
	setting.VideoIds++

	//构建文件保持地址
	saveFile := filepath.Join("./public_videos/", finalName)

	// 保存文件data到saverFile文件里
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, class.Response{
			StatusCode: 1,
			StatusMsg:  "saveVideo:" + err.Error(),
		})
		return
	}

	//保存封面并且判断是否是视频
	tmpCover, err := GetSnapshot(saveFile, filename, 1)
	if err != nil { // 如果无法截图就说明文件有误
		log.Println(err)
		if err = os.Remove(saveFile); err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, class.Response{
			StatusCode: 0,
			StatusMsg:  finalName + "不正确,请确认文件类型",
		})
		return
	}
	CoverPath := Const.ServiceUrl + "/jpg/" + tmpCover

	filePath := Const.ServiceUrl + "/static/" + finalName
	// 检查文件是否存在,存在更新文件退出(应成替换,而非更新)
	if flag := sql.FindVideoByFile(filePath, user); flag {
		c.JSON(http.StatusOK, class.Response{
			StatusCode: 0,
			StatusMsg:  finalName + " renew successfully",
		})
		return
	}
	// 存储文件地址到video数据库
	videoId, ok := sql.InsertVideo(filePath, user, title, time.Now().Unix(), CoverPath)
	if ok != nil {
		c.JSON(http.StatusOK, class.Response{
			StatusCode: 1,
			StatusMsg:  ok.Error(),
		})
		return
	}

	// 存储文件信息到public数据库
	if err = sql.InsertPublic(token, videoId); err != nil {
		c.JSON(http.StatusOK, class.Response{
			StatusCode: 1,
			StatusMsg:  ok.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, class.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)

	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{
			fmt.Sprintf("gte(n,%d)", frameNum),
		}).Output("pipe:", ffmpeg.KwArgs{
		"vframes": 1, "format": "image2", "vcodec": "mjpeg",
	}).WithOutput(buf, os.Stdout).Run()
	if err != nil {
		return "", fmt.Errorf("生成图失败：%v", err)
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		return "", fmt.Errorf("生成图失败：%v", err)
	}

	err = imaging.Save(img, "./public_cover/"+snapshotPath+".png")
	if err != nil {
		return "", fmt.Errorf("生成图失败：%v", err)
	}

	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}

func PublishList(c *gin.Context) {
	token := c.Query("token")

	c.JSON(http.StatusOK, VideoListResponse{
		Response: class.Response{
			StatusCode: 0,
		},
		VideoList: sql.ReadPublishVideos(token),
	})

}
