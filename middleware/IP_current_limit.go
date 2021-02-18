package middleware

import (
	"fmt"
	"jsonhut-server/config"
	"jsonhut-server/dao"
	"jsonhut-server/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IPCurrentLimiting(mode string) gin.HandlerFunc {
	return func(c *gin.Context) {
		dataName := mode + "-" + c.ClientIP()

		requestNum, err := dao.RedisGetData(dataName)
		if err != nil {
			fmt.Println(err.Error())
		}
		requestInt, _ := strconv.Atoi(requestNum)
		// 判断是否超过指定调用次数限制
		if requestInt >= config.App.GetIPCurrentLimit && mode == "GET" || requestInt >= config.App.PostIPCurrentLimit && mode == "POST" || requestInt >= config.App.DetailsIPCurrentLimit && mode == "DETAILS" {
			c.JSON(http.StatusServiceUnavailable, models.ReturnJsonWithoutData{
				Code: 503,
				Msg:  "Rate limit.. | You can try again in " + strconv.Itoa(dao.RedisGetTTL(dataName)) + "s",
			})
			c.Abort()
			return
		}
		// 增加对应ip的调用次数，时长1小时
		dao.RedisSetData(dataName, strconv.Itoa(requestInt+1), 60*60)

		c.Next()
	}
}
