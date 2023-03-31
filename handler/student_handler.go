package handler

import (
	"github.com/gin-gonic/gin"
	"learn-go/base"
	"learn-go/dao"
	"learn-go/models"
	"learn-go/utils"
)

func AddStu(c *gin.Context) {
	var stu models.Student
	err := c.ShouldBindJSON(&stu)
	if err != nil {
		base.Fail(c, 400, err.Error())
		return
	}
	tx := dao.StuRepo().Create(&stu)
	base.Success(c, tx.RowsAffected)
}

func FindByName(c *gin.Context) {
	name := c.Query("name")
	if utils.IsBlank(name) {
		base.Fail(c, 400, "empty query key")
		return
	}
	var stu models.Student
	tx := dao.StuRepo().First(&stu, "name = ?", name)
	if tx.RowsAffected < 1 {
		base.Fail(c, 200, "[]")
		return
	}
	base.Success(c, stu)
}

func ListAll(c *gin.Context) {
	var stuList []models.Student
	dao.StuRepo().Select("name", "age", "addr", "height").Order("age desc").Find(&stuList)
	base.Success(c, stuList)
}
