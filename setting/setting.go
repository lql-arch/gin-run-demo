package setting

import (
	"douSheng/Const"
	"time"
)

var VideoIds int64

func NowString() string {
	return time.Now().Format(Const.TimeTemplate)
}

func TimeString(times int64) string {
	return time.Unix(times, 0).Format(Const.TimeTemplate)
}
