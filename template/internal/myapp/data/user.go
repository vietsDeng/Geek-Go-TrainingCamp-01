package data

import (
	model "template/internal/myapp/model/data"
)
import "template/internal/myapp"

func FindUser(id int64) (*model.User, error) {
	user := new(model.User)
	db := myapp.App.DB.Where("id = ?", id).First(&user)
	if len(db.GetErrors()) > 0 {
		return nil, db.GetErrors()[0]
	}

	return user, nil
}