package base

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"os"
)

func UidMiddleware(c *gin.Context) {
	c.AddParam("uid", uuid.New().String())
}

func GinEngine() *gin.Engine {
	//gin.DisableConsoleColor()

	log, err := os.Create("info.log")
	if err != nil {
		println(err)
	}
	gin.DefaultWriter = io.MultiWriter(log, os.Stdout)
	return gin.Default()
}
