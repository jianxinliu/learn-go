package base

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"learn-go/models"
)

var conn *gorm.DB

func connect() {
	if conn == nil {
		println("new db connection....")
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		conn = db
		conn.AutoMigrate(&models.Student{})
	}
}

func GetConn() *gorm.DB {
	connect()
	return conn
}
