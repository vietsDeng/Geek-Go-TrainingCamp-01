package model

type User struct{
	Id int64
	Name string
	Class int64
	Level int64
}

// 自定义表名
func (User) TableName() string {
	return "user"
}