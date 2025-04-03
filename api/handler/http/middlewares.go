package http

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"golang.org/x/time/rate"
)

func Logger() gin.HandlerFunc {
	return gin.Logger()
}

func Recovery() gin.HandlerFunc {
	return gin.Recovery()
}

var limiter = rate.NewLimiter(rate.Limit(1), 10)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
