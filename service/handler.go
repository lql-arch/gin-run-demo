package service

import (
	"context"
	api "douSheng/kitex/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct {
}

// Video implements the EchoImpl interface.
func (s *EchoImpl) Video(ctx context.Context, req *api.Video) (resp *api.Response, err error) {

	return
}

// User implements the EchoImpl interface.
func (s *EchoImpl) User(ctx context.Context, req *api.User) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}

// FriendUser implements the EchoImpl interface.
func (s *EchoImpl) FriendUser(ctx context.Context, req *api.FriendUser) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}

// Comment implements the EchoImpl interface.
func (s *EchoImpl) Comment(ctx context.Context, req *api.Comment) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}

// UserVideoFavorite implements the EchoImpl interface.
func (s *EchoImpl) UserVideoFavorite(ctx context.Context, req *api.UserVideoFavorite) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}

// Relation implements the EchoImpl interface.
func (s *EchoImpl) Relation(ctx context.Context, req *api.Relation) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}

// Message implements the EchoImpl interface.
func (s *EchoImpl) Message(ctx context.Context, req *api.Message) (resp *api.Relation, err error) {
	// TODO: Your code here...
	return
}
