package blog

import (
	"github.com/gin-gonic/gin"
	"go-gin/app/api/blog"
)

func Routers(e *gin.Engine) {

	api := blog.SysBlog{}

	router := e.Group("blog")
	{
		router.GET("/post", api.PostHandler)
		// router.GET("/comment", api.CommentHandler)
	}
}
