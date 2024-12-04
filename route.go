package logportalapi

import (
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(handler *gin.Engine) *SSEEvent {

	stream := getNewSSEServer()

	handler.GET(
		"/log-portal/stream",
		addHeaders(),
		stream.serveHTTP(),
		getStream,
	)

	handler.Use(loggingMiddleware(stream))

	return stream
}

func getStream(c *gin.Context) {
	v, ok := c.Get("clientChan")
	if !ok {
		return
	}
	clientChan, ok := v.(clientChan)
	if !ok {
		return
	}
	c.Stream(func(w io.Writer) bool {
		// Stream message to client from message channel
		if msg, ok := <-clientChan; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}

func addHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}

func loggingMiddleware(stream *SSEEvent) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		responseTime := time.Since(startTime)

		logMessage := LogMessage{
			TimeStamp:    time.Now().Format("2006-01-02 15:04:05"),
			Status:       c.Writer.Status(),
			Method:       c.Request.Method,
			Path:         c.Request.URL.Path,
			Type:         "gin",
			ResponseTime: responseTime.String(),
		}

		// Convert to JSON
		jsonData, err := json.Marshal(logMessage)
		if err != nil {
			log.Printf("JSON encoding error: %s", err)
			return
		}
		stream.Message <- jsonData
	}
}
