package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jsonhut-server/config"
	"jsonhut-server/dao"
	"jsonhut-server/models"
	"net/http"
	"strconv"
)

func IPCurrentLimiting(methodType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		redisIPName := methodType + "-" + c.ClientIP()
		requestNum, err := dao.RedisGetData(redisIPName)
		if err != nil {
			fmt.Println(err.Error())
		}
		usedRateLimit, _ := strconv.Atoi(requestNum)

		var maxRateLimit int
		switch methodType {
		case "GET":
			maxRateLimit = config.App.GetIPCurrentLimit
			break
		case "POST":
			maxRateLimit = config.App.PostIPCurrentLimit
			break
		case "DETAILS":
			maxRateLimit = config.App.DetailsIPCurrentLimit
			break
		}

		// 判断是否超过指定调用次数限制
		if usedRateLimit >= maxRateLimit {
			c.JSON(http.StatusServiceUnavailable, models.ReturnJsonWithoutData{
				Code: 503,
				Msg:  "Rate limit.. | You can try again in " + strconv.Itoa(dao.RedisGetTTL(redisIPName)) + "s",
			})
			c.Abort()
			return
		}

		// 增加对应ip的调用次数，过期时长1小时
		usedRateLimit++
		if usedRateLimit == 1 {
			// 首次计数
			dao.RedisSetDataWithExpireTime(redisIPName, strconv.Itoa(usedRateLimit), 60*60)
		} else {
			dao.RedisSetDataWithExpireTime(redisIPName, strconv.Itoa(usedRateLimit), dao.RedisGetTTL(redisIPName))
		}

		// 设置头部信息
		c.Header("X-Rate-Limit-Limit", strconv.Itoa(maxRateLimit))                   // 允许的最大调用次数
		c.Header("X-Rate-Limit-Reset", strconv.Itoa(dao.RedisGetTTL(redisIPName)))   // 计数在多少秒后重置
		c.Header("X-Rate-Limit-Remaining", strconv.Itoa(maxRateLimit-usedRateLimit)) // 剩余的调用次数

		c.Next()
	}
}
