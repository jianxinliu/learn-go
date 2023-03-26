package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-go/base"
	"learn-go/utils"
	"os"
)

type LoginReq struct {
	UserId string `json:"userId" binding:"required"`
	Pwd    string `json:"pwd"`
}

func Login(c *gin.Context) {
	var loginReq LoginReq
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		fmt.Fprint(os.Stdout, err)
		base.Fail(c, 500, err.Error())
		return
	}
	//uid := c.Param("uid")

	base.Success(c, loginReq)

	//c.JSON(200, ret)
}

func Hello(c *gin.Context) {
	echo := c.Query("echo")
	retMap := make(map[string]string)
	if utils.IsNotBlank(echo) {
		retMap["echo"] = echo
	}
	query := c.DefaultQuery("name", "jak")
	retMap["query"] = query
	base.Success(c, retMap)
}

type Params struct {
	Id   string `uri:"id" json:"id" binding:"required"`
	Name string `uri:"name" json:"name" binding:"required"`
}

func DoRegister(c *gin.Context) {
	var p Params
	// path variable
	err := c.ShouldBindUri(&p)
	if err != nil {
		base.Fail(c, 400, err.Error())
		return
	}
	base.Success(c, p)
}
