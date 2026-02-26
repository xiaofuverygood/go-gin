package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Option func(*gin.Engine)

var options []Option

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Init 初始化
func Init() *gin.Engine {
	r := gin.Default()

	// 设置CORS头 注意要在注册路由前配置
	r.Use(cors.Default())
	// 拦截请求
	// r.Use(Middleware)

	for _, opt := range options {
		opt(r)
	}

	return r
}

// 中间件拦截器
func Middleware(c *gin.Context) {

	adminReqs := []string{"add", "edit", "del"}

	// 截取路径和参数
	path := c.Request.URL.Path
	params := c.Request.URL.Query()
	param := params.Get("auth")
	request := path[1:strings.Index(path, "-")]

	// 校验权限
	if !checkReq(adminReqs, request) && param != "admin" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"tokenstatus": "没有权限访问",
		})
	}

	// 继续处理下一个中间件或处理函数
	c.Next()

	statusCode := c.Writer.Status()
	fmt.Println(statusCode)

	responseBody := c.Writer
	fmt.Println("响应体：", responseBody)

}

func checkReq(slice []string, req string) bool {
	for _, str := range slice {
		if str == req {
			return false
		}
	}

	return true
}
