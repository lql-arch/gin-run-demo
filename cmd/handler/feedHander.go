package handler

import (
	"context"
	"douSheng/cmd/class"
	"douSheng/cmd/feed/convert"
	api "douSheng/cmd/feed/kitex_gen/api"
	"douSheng/sql"
	"log"
)

// FeedImpl implements the last service interface defined in the IDL.
type FeedImpl struct{}

// ReadVideos implements the FeedImpl interface.
func (s *FeedImpl) ReadVideos(ctx context.Context, latestTime int64, token string) (resp *api.FeedList, err error) {
	go log.Println("readVideos.")
	var list []*class.Video
	var nextTime int64
	list, nextTime, err = sql.ReadVideos(latestTime, token)
	resp = new(api.FeedList)
	if err != nil {
		resp.StatusCode = 1
		return resp, err
	}

	videoList, err := convert.VideoListConvert(ctx, list)
	resp.VideoList = videoList
	if err != nil {
		return nil, err
	}

	if nextTime == 0 {
		resp.NextTime = latestTime
	}

	resp.StatusCode = 0

	return
}
