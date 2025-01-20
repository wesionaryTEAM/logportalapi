package main

import (
	"fmt"
	"gogin/framework"
	"gogin/mods/logw"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	logportal := logw.New(&logw.Config{
		ID:         "test",
		HistoryLen: 5,
	})

	logger := framework.GetLogger(logportal)

	r := gin.Default()
	i := 0

	r.Use(corsMiddleware())

	r.GET("/sse", func(c *gin.Context) {
		logportal.Subscribe(&logw.SubSSE{
			Request: c.Request,
			Writer:  c.Writer,
		})
	})

	r.GET("/err", func(c *gin.Context) {
		i = i + 1

		logger.Error("Main one ERROR ............ " + fmt.Sprint(i))
		logger.Info("Main one INFO ............ " + fmt.Sprint(i))

		c.String(200, fmt.Sprintf("ID is: %d", i))
	})

	r.Run("localhost:8484")

	time.Sleep(100 * time.Millisecond)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
