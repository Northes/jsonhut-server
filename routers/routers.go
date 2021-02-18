package routers

import (
	"jsonhutapi/config"
	"jsonhutapi/controller"
	"jsonhutapi/middleware"
	"jsonhutapi/models"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	engine := gin.Default()

	//engine.Use(Cors())
	engine.Use(cors.Default(), middleware.RateLimitMiddleware(200, 8))
	{
		engine.GET("/", func(context *gin.Context) {
			baseurl := config.App.BaseUrl
			context.JSON(http.StatusOK, gin.H{
				"[GET] Request Json":       baseurl + "/bins/{id}",
				"[POST] Create a new Json": baseurl + "/bins",
				"[GET] Get Json details":   baseurl + "/details/{id}",
			})
		})
		engine.GET("/bins/:id", middleware.IPCurrentLimiting("GET"), controller.GetJson)
		engine.POST("/bins", middleware.IPCurrentLimiting("POST"), controller.PostJson)
		engine.GET("/details/:id", middleware.IPCurrentLimiting("DETAILS"), controller.GetJsonDetails)
		engine.NoRoute(func(context *gin.Context) {
			context.JSON(http.StatusNotFound, models.ReturnJsonWithoutData{
				Code: 404,
				Msg:  "You can go to the documentationUrl [https://jsonhut.com/docs]",
			})
		})
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
