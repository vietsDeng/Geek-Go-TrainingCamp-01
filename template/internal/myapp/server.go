package myapp

import (
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
)

type app struct{
	DBConfig dbConfig
	DB *gorm.DB
}

//Person
type dbConfig struct {
	Driver 	string
	DSN    	string
}

var App app

func newDBConfig() dbConfig {
	var dbConfig dbConfig

	path := "./config/db.toml"
	if _, err := toml.DecodeFile(path, &dbConfig); err != nil {
		panic(err)
	}

	return dbConfig
}

func newDB(dbConfig dbConfig) *gorm.DB {
	db, err := gorm.Open(dbConfig.Driver, dbConfig.DSN)
	if err != nil {
		panic(err)
	}

	return db
}

func NewApp(dbConfig dbConfig, db *gorm.DB) app {
	return app{dbConfig, db}
}

func InitApp()  {
	App = initApp()
}
