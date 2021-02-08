package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"jsonhutapi/controller"
	"jsonhutapi/middleware"
	"net/http"
)

func SetupRouter() *gin.Engine {
	engine := gin.Default()

	//engine.Use(Cors())
	engine.Use(cors.Default(), middleware.RateLimitMiddleware(500, 7))
	{
		engine.GET("/", func(context *gin.Context) {
			baseurl := "https://api.jsonhut.com"
			context.JSON(http.StatusOK, gin.H{
				"Get":     baseurl + "/bins/{id}",
				"Post":    baseurl + "/bins",
				"Details": baseurl + "/details/{id}",
			})
		})
		engine.GET("/bins/:id", controller.GetJson)
		engine.POST("/bins", controller.PostJson)
		engine.GET("/details/:id", controller.GetJsonDetails)
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
	return engine
}
