/**
 *@Description TODO
 *@Author fsy
 *@Date 2023/6/20 17:30
 */
package interceptor

import (
	"github.com/gin-gonic/gin"
	"go-gin/app/api/blog"
)

func Routers(e *gin.Engine) {

	api := blog.SysBlog{}

	router := e.Group("")
	{
		router.GET("/add-xxx", api.PostHandler)
	}
}
