package main

import (
	"blog.zozoo.net/route"
	"github.com/gin-gonic/gin"
	"log"
)

/**
	gin 博客项目
	20201022 19:37
	ganganlee
 */


func main() {
	var err error
	r := gin.Default()

	//注册路由
	route.CreateRoute(r)

	//连接数据库

	err = r.Run(":80") // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatalf("项目启动失败 err:%v",err)
	}
}
