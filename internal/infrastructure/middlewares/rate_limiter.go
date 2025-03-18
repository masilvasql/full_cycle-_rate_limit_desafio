package middlewares

import (
	"errors"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/internal/adapters/middleware"
)

var mu sync.Mutex // Mutex global para controlar o acesso concorrente

func RateLimiter(r middleware.RateLimiterInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()

		ip := c.ClientIP()
		token := c.GetHeader("API_KEY")
		err := r.CheckRateLimit(ip, token)
		if err != nil {
			if errors.Is(err, middleware.ErrNotAuthorized) {
				c.JSON(401, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			// Verifica se o erro é de limite atingido
			if errors.Is(err, middleware.LimitReached) {
				c.JSON(429, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			// Caso contrário, retorna um erro genérico
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}

		// Se passou em todas as verificações, a requisição é permitida
		c.Next()
	}
}
