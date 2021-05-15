package main

import (
	"context"
	"template/internal/myapp/biz"
	proto "template/proto/myapp"
)

func (u *UserService) GetUser(ctx context.Context, r *proto.UserRequest) (*proto.UserResponse, error) {
	return biz.GetUser(r)
}

func (u *UserService) GetLevelInfos(ctx context.Context, r *proto.NoneRequest) (*proto.LevelInfosResponse, error) {
	return biz.GetLevelInfos()
}