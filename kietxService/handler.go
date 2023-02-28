package main

import (
	"context"
	api "douSheng/kitex/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// FindComments implements the EchoImpl interface.
func (s *EchoImpl) FindComments(ctx context.Context, videoId int32, token string) (resp *api.Comment, err error) {
	// TODO: Your code here...

	return
}

// ReviseComment implements the EchoImpl interface.
func (s *EchoImpl) ReviseComment(ctx context.Context, comment *api.Comment) (resp *api.Response, err error) {
	// TODO: Your code here...

	return
}
