package handler

import (
	"bytes"
	"context"
	"douSheng/Const"
	"douSheng/cmd/publish/convert"
	api "douSheng/cmd/publish/kitex_gen/api"
	"douSheng/sql"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// PublishImpl implements the last service interface defined in the IDL.
type PublishImpl struct{}

// Publish implements the PublishImpl interface.
func (s *PublishImpl) Publish(ctx context.Context, token string, video *api.VideoData) (resp *api.PublishResponse, err error) {
	// TODO: Your code here...
	log.Println("Publish")
	log.Println("Publish 需要事务,1未处理")

	user, exist := sql.FindUser(token)
	if !exist {
		return &api.PublishResponse{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		}, fmt.Errorf("用户不存在")
	}

	savePath, fileName := CreateFileName(video, user.Id)

	// 保存视频
	_ = SaveUploadedFile(video, savePath)
	if err != nil {
		return &api.PublishResponse{
			StatusCode: 1,
			StatusMsg:  "视频上传失败",
		}, fmt.Errorf("视频上传失败")
	}
	filePath := Const.ServiceUrl + "/static/" + fileName

	//保存封面并且判断是否是视频
	cover, err := GetSnapshot(savePath, fileName, 1)
	if err != nil { // 如果无法截图就说明文件有误
		log.Println(err)
		if err = os.Remove(savePath); err != nil {
			log.Println(err)
		}
		return &api.PublishResponse{
			StatusCode: 1,
			StatusMsg:  video.Title + "不正确,请确认文件类型",
		}, fmt.Errorf(video.Title + "不正确,请确认文件类型")
	}
	CoverPath := Const.ServiceUrl + "/jpg/" + cover

	// 检查文件是否存在,存在更新文件退出(应成替换,而非更新)
	if flag := sql.FindVideoByFile(filePath, user); flag {
		return &api.PublishResponse{
			StatusCode: 0,
			StatusMsg:  "视频重传成功",
		}, fmt.Errorf("视频重传成功")
	}

	// 存储文件地址到video数据库
	videoId, ok := sql.InsertVideo(filePath, user, video.Title, time.Now().Unix(), CoverPath)
	if ok != nil {
		return &api.PublishResponse{
			StatusCode: 0,
			StatusMsg:  ok.Error(),
		}, ok
	}

	// 存储文件信息到public数据库
	if err = sql.InsertPublic(token, videoId); err != nil {
		return &api.PublishResponse{
			StatusCode: 0,
			StatusMsg:  err.Error(),
		}, err
	}

	// 修改作品数目(+1)
	err = sql.InsertUserPublicCount(user.Id, true)
	if err != nil {
		return &api.PublishResponse{
			StatusCode: 0,
			StatusMsg:  err.Error(),
		}, err
	}

	return &api.PublishResponse{
		StatusCode: 0,
		StatusMsg:  "上传成功",
	}, nil
}

// PublishList implements the PublishImpl interface.
func (s *PublishImpl) PublishList(ctx context.Context, token string) (resp *api.PublishResponse, err error) {
	// TODO: Your code here...
	log.Println("PublishList")

	videoList, err := sql.ReadPublishVideos(token)

	if err != nil {
		log.Println(err)
		return &api.PublishResponse{
			StatusCode: 1,
			StatusMsg:  "查询视频出错",
			Videos:     nil,
		}, err
	}

	videos, _ := convert.VideoListConvert(ctx, videoList)

	return &api.PublishResponse{
		StatusCode: 0,
		StatusMsg:  "",
		Videos:     videos,
	}, nil
}

// GetSnapshot 根据视频截取视频封面
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

func SaveUploadedFile(video *api.VideoData, saveFile string) error {
	reader := bytes.NewReader(video.Data)
	out, err := os.Create(saveFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, reader)
	return err
}

func CreateFileName(video *api.VideoData, userId int64) (savePath string, fileName string) {
	//构建新文件名
	finalName := fmt.Sprintf("%d_%s_%s", userId, video.Title, time.Now().Format("200601021150405"))

	//构建文件保持地址
	saveFile := filepath.Join("./public_videos/", finalName)

	return saveFile, finalName
}
