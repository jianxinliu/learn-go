package main

import (
	"learn-go/base"
	"learn-go/handler"
	"net/http"
)

func main() {
	r := base.GinEngine()

	// 自定义中间件
	r.Use(base.UidMiddleware)

	// 路由组
	v1 := r.Group("/v1")

	{
		v1.GET("/hello", handler.Hello)
		v1.POST("/login", handler.Login)
		v1.GET("/:name/:id", handler.DoRegister)
	}

	stuG := r.Group("/stu")

	{
		stuG.POST("/add", handler.AddStu)
		stuG.GET("/findByName", handler.FindByName)
		stuG.GET("/list", handler.ListAll)
	}

	// 静态服务
	r.StaticFS("/img", http.Dir("./img"))
	//r.Static("/img", "./img")
	r.StaticFile("/img2", "./img/img2.png")

	r.Run()
}
