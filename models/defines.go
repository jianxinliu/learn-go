package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Addr   string  `json:"addr"`
	Height float32 `json:"height"`
}
