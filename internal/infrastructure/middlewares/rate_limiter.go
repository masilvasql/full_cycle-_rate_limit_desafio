package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/internal/adapters/middleware"
)

func RateLimiter(r middleware.RateLimiterInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		token := c.GetHeader("API_KEY")
		r.CheckRateLimit(ip, token)
		c.Next()
	}
}
