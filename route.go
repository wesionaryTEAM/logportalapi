package logportalapi

import (
	"io"

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
