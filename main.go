package main

import (
	"code.lssj.me/bloger/controller"
	"code.lssj.me/bloger/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
		return
	}

	r := gin.Default()
	//加载静态资源
	r.Static("static/", "./static")
	//加载模板
	r.LoadHTMLGlob("views/*")

	r.GET("/", controller.IndexHandler)
	r.GET("/article/detail/", controller.ArticleDetailHandler)
	r.GET("/category/", controller.ArticleListByCategory)
	r.GET("/about/me/", controller.AboutMe)
	r.GET("/leave/new/", controller.LeavePage)
	r.POST("/leave/submit/", controller.WriteLeaves)
	r.GET("/article/new/", controller.PostPage)
	r.POST("/article/submit/", controller.PostArticle)
	r.POST("/comment/submit/", controller.PostComment)
	err = r.Run(":8099")
	if err != nil {
		panic(err)
		return
	}

}
