package shop

import (
	"github.com/gin-gonic/gin"
	"go-gin/app/api/shop"
)

func Routers(e *gin.Engine) {

	api := shop.SysShop{}

	router := e.Group("/shop")
	{
		router.GET("/goods/:id/:name", api.GoodsHandler)
		router.GET("/checkout", api.CheckoutHandler)
		router.GET("/post", api.PostHandler)
	}
}
