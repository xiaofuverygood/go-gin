package blog

import (
	"github.com/gin-gonic/gin"
)

type SysBlog struct {
}

func (s SysBlog) PostHandler(c *gin.Context) {
	// fmt.Println("123")
	c.JSON(250, "name")
}
