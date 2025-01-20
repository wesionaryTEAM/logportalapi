package logw

import (
	"fmt"
	"net/http"
)

func (sse *SubSSE) GetPortalID() string {
	return sse.portal.ID
}

func (sse *SubSSE) IsActive() bool {
	return sse.Active
}

func (sse *SubSSE) SetActive(status bool) {
	sse.Active = status
}

func (sse *SubSSE) Write(p *Payload) {
	if !sse.IsActive() {
		return
	}

	fmt.Fprintf(sse.Writer, "data: %s\n\n", string(p.Json()))
	sse.Writer.(http.Flusher).Flush()
}

func (sse *SubSSE) run() {
	close := sse.Request.Context().Done()

	// Flush headers immediately.
	sse.Writer.Header().Set("Content-Type", "text/event-stream")
	sse.Writer.Header().Set("Cache-Control", "no-cache")
	sse.Writer.Header().Set("Connection", "keep-alive")
	sse.Writer.(http.Flusher).Flush()

	// Send history
	sse.Write(&Payload{
		PortalID: sse.GetPortalID(),
		Type:     "history",
		Entries:  sse.portal.history,
	})

	<-close

	sse.SetActive(false)
}
