package convert

import (
	"douSheng/cmd/class"
	"douSheng/cmd/comment/kitex_gen/api"
	"douSheng/setting"
)

func CommentListConvert(commentList []*class.Comment) []*api.Comment {
	newCommentList := make([]*api.Comment, 0)
	for _, comment := range commentList {
		comment := CommentConvert(comment)
		if comment == nil {
			continue
		}
		newCommentList = append(newCommentList, comment)
	}

	return newCommentList
}

func CommentConvert(comment *class.Comment) *api.Comment {
	return &api.Comment{
		Id:             comment.Id,
		UserId:         comment.UserId,
		Author:         UserConvert(comment.Author),
		Content:        comment.Content,
		VideoId:        comment.VideoId,
		Type:           int32(comment.Type),
		JSONCreateDate: setting.CommentTimeString(comment.CreateDate),
	}
}
