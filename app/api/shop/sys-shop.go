package shop

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SysShop struct {
}

// 结构体声明并做约束
type Persion struct {
	ID   int    `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func (s *SysShop) GoodsHandler(c *gin.Context) {
	// 从url路径上获取参数
	/*	id := c.Param("id")
		name := c.Param("name")*/

	var persion Persion

	if err := c.ShouldBindUri(&persion); err != nil {
		c.Status(404)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   persion.ID,
		"name": persion.Name,
	})
}
func (s *SysShop) CheckoutHandler(c *gin.Context) {
	// 从url上获取参数 并且设置默认值
	value := c.DefaultQuery("value", "0")
	v, _ := strconv.Atoi(value)

	c.JSON(http.StatusOK, gin.H{
		"value": v,
	})
}

func (s *SysShop) PostHandler(c *gin.Context) {

	username := c.DefaultPostForm("username", "unkown")
	password := c.DefaultPostForm("password", "unkown")

	if username == "admin" && password == "admin" {
		c.JSON(http.StatusOK, gin.H{
			"tokenstatus": "认证通过",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"tokenstatus": "认认证失败",
		})
	}
}
