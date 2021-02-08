package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"jsonhutapi/controller/GetJson"
	"jsonhutapi/controller/GetJsonDetails"
	"jsonhutapi/controller/PostJson"
	"log"
	"net/http"
	"time"
)

func main() {
	engine := gin.Default()

	//engine.Use(Cors())
	engine.Use(cors.Default(), RateLimitMiddleware(500, 7))
	{
		engine.GET("/", func(context *gin.Context) {
			baseurl := "https://api.jsonhut.com"
			context.JSON(http.StatusOK, gin.H{
				"Get":     baseurl + "/bins/{id}",
				"Post":    baseurl + "/bins",
				"Details": baseurl + "/details/{id}",
			})
		})
		engine.GET("/bins/:id", GetJson.GetJson)
		engine.POST("/bins", PostJson.PostJson)
		engine.GET("/details/:id", GetJsonDetails.GetDetails)
	}

	//engine.Use(cors.New(cors.Config{
	//	AllowOrigins:  []string{"*"},
	//	AllowMethods:  []string{"POST, GET, OPTIONS, DELETE"},
	//	AllowHeaders:  []string{"Origin", "Authorization", "Content-Length", "X-CSRF-Token", "Token,session"},
	//	ExposeHeaders: []string{"Content-Type, x-requested-with, X-Custom-Header, Authorization"},
	//	//AllowAllOrigins:  true,
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return true
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))

	err := engine.Run(":8081")
	if err != nil {
		fmt.Println(err)
	}
}

func RateLimitMiddleware(fillInterval time.Duration, cap int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		i := bucket.TakeAvailable(1)
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if i < 1 {
			fmt.Println("rate limit...")
			c.String(503, "rate limit...")
			c.Abort()
			return
		}
		fmt.Println(i)
		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
