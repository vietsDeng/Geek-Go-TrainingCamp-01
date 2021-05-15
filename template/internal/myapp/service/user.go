package service

import (
	"template/internal/myapp/data"
	model "template/internal/myapp/model/data"
)

func FindUser(id int64) (*model.User, error) {
	return data.FindUser(id)
}