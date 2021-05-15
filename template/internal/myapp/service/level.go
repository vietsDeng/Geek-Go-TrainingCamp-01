package service

import (
	"errors"
	model "template/internal/myapp/model/biz"
)

var levelInfos = []model.LevelInfo{
	{
		Level:   1,
		Name: "青铜",
	},
	{
		Level:   2,
		Name: "白银",
	},
	{
		Level:   3,
		Name: "黄金",
	},
}

func GetLevelInfoList() ([]model.LevelInfo, error) {
	return levelInfos, nil
}

func GetLevelInfo(level int64) (model.LevelInfo, error) {
	var info model.LevelInfo

	for _, item := range levelInfos {
		if item.Level == level {
			info := item
			return info, nil
		}
	}

	return info, errors.New("not Found Level Info")
}
