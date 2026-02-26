/**
 *@Description TODO
 *@Author fsy
 *@Date 2023/5/18 16:19
 */

package main

import (
	"go-gin/app/router"
	"go-gin/app/router/blog"
	"go-gin/app/router/interceptor"
	"go-gin/app/router/shop"
	"log"
	"net/http"
	"time"
)
/*
var sqlDb *sql.DB           //数据库连接
var sqlResponse SqlResponse //响应结构体

// SqlUser 请求结构体

	type SqlUser struct {
		Id      int    `json:"id"`
		Account string `json:"account"`
		Age     int    `json:"age"`
		Address string `json:"address"`
	}

// SqlResponse 响应结构体

	type SqlResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
*/
func init() {
	/*
		连接数据库
		parseTime:时间格式转换(查询结果为时间时，是否自动解析为时间);
		loc=Local：MySQL的时区设置
	*/
	/*	databaseConStr := "root:root1234@tcp(127.0.0.1:3306)/db2019?charset=utf8&parseTime=true&loc=Local"
				var err error
				sqlDb, err = sql.Open("mysql", databaseConStr)
				if err != nil {
					fmt.Println("连接数据库失败:", err)
					return
				}
		// 测试与数据库建立的连接 Ping()校验连接是否正确
				err = sqlDb.Ping()
				if err != nil {
					fmt.Println("连接数据库失败:", err)
					return
				}*/
}

func main() {

	// 加载多个APP的路由配置
	router.Include(shop.Routers, blog.Routers, interceptor.Routers)
	// 初始化路由
	r1 := router.Init()
	// 设置静态资源路由
	r1.Static("/static", "./static")

	r2 := router.Init()

	/* 	// 创建第一个 Gin 实例（监听 8080 端口）
	   	r1 := gin.Default()
	   	r1.GET("/ping", func(c *gin.Context) {
	   		c.JSON(http.StatusOK, gin
	   			"message": "pong from server 1",
	   		})
	   	})

	   	// 创建第二个 Gin 实例（监听 8081 端口）
	   	r2 := gin.Default()
	   	r2.GET("/ping", func(c *gin.Context) {
	   		c.JSON(http.StatusOK, gin.H{
	   			"message": "pong from server 2",
	   		})
	   	})
	*/
	// 定义第一个 HTTP 服务器
	server1 := &http.Server{
		Addr:         ":8080",
		Handler:      r1,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// 定义第二个 HTTP 服务器
	server2 := &http.Server{
		Addr:         ":8081",
		Handler:      r2,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// 使用 Goroutine 启动第一个服务
	go func() {
		if err := server1.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server 1 failed to listen: %v\n", err)
		}
	}()

	// 使用 Goroutine 启动第二个服务
	go func() {
		if err := server2.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server 2 failed to listen: %v\n", err)
		}
	}()

	// 阻塞主协程，防止程序退出
	select {}
}
