package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"learn-go/models"
)

var conn *gorm.DB

func connect() {
	if conn == nil {
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		conn = db
		conn.AutoMigrate(&models.Student{})
		initDb()
	}
}

func GetConn() *gorm.DB {
	connect()
	return conn
}

func initDb() {
	stuList := []models.Student{
		{
			Name:   "jack",
			Age:    33,
			Addr:   "beijing",
			Height: 199,
		},
		{
			Name:   "rose",
			Age:    33,
			Addr:   "beijing",
			Height: 188,
		},
		{
			Name:   "pony",
			Age:    12,
			Addr:   "shenzhen",
			Height: 178,
		},
		{
			Name:   "lily",
			Age:    25,
			Addr:   "shanghai",
			Height: 188,
		},
		{
			Name:   "mary",
			Age:    23,
			Addr:   "shanghai",
			Height: 188,
		},
	}
	conn.CreateInBatches(&stuList, 100)
}
