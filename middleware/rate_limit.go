package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"jsonhutapi/models"
	"net/http"
	"time"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) gin.HandlerFunc {
	// fillInterval 每多少秒填充一个令牌  cap 令牌桶的最大容量
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		i := bucket.TakeAvailable(1)
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if i < 1 {
			fmt.Println("rate limit...")
			c.JSON(http.StatusServiceUnavailable, models.ReturnJsonWithoutData{
				Code: 503,
				Msg:  "Rate limit.. | Please try again later",
			})
			c.Abort()
			return
		}
		//fmt.Println(i)
		c.Next()
	}
}
