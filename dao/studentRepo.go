package dao

import (
	"gorm.io/gorm"
	"learn-go/base"
	"learn-go/models"
)

var dbInited = true

func StuRepo() *gorm.DB {
	if !dbInited {
		initDb()
		dbInited = true
	}
	return base.GetConn()
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

	conn := base.GetConn()
	conn.CreateInBatches(&stuList, 100)
}
