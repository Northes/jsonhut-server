package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		i := bucket.TakeAvailable(1)
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if i < 1 {
			fmt.Println("rate limit...")
			c.String(http.StatusServiceUnavailable, "rate limit...")
			c.Abort()
			return
		}
		fmt.Println(i)
		c.Next()
	}
}
