package biz

import (
	"template/internal/myapp/service"
	proto "template/proto/myapp"
)

func GetUser(r *proto.UserRequest) (*proto.UserResponse, error) {
	user, err := service.FindUser(r.Id)
	if err != nil {
		return nil, err
	}

	levelInfo, err := service.GetLevelInfo(user.Level)
	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Class:     user.Class,
		Level:     user.Level,
		LevelName: levelInfo.Name,
	}, nil
}

func GetLevelInfos() (*proto.LevelInfosResponse, error) {
	var list []*proto.LevelInfo
	
	tempList, err := service.GetLevelInfoList()
	if err != nil {
		return nil, err
	}

	for _, item := range tempList  {
		list = append(list, &proto.LevelInfo{
			Level:  item.Level,
			Name:	item.Name,
		})
	}

	return &proto.LevelInfosResponse{
		List: list,
	}, nil
}