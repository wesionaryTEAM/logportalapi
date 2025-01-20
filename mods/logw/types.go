package logw

import (
	"net/http"
	"sync"
)

type Config struct {
	ID         string
	HistoryLen int

	// Automatically populated.
	mutex       sync.RWMutex
	history     []interface{}
	subscribers []Subscriber
}

type Entry struct {
	Type   string `json:"type"`
	Status int    `json:"status"`
	// Method       string `json:"method"`
	// Path         string `json:"path"`
	// TimeStamp    string `json:"timestamp"`
	// Location     string `json:"location"`
	// TimeForQuery int64  `json:"timeForQuery"`
	// Level        string `json:"level"`
	// Message      string `json:"message"`
	// StrackTrace  string `json:"stackTrace"`
	// Query        string `json:"query"`
	// Rows         int64  `json:"rows"`
	// ResponseTime string `json:"responseTime"`
}

type Subscriber interface {
	Write(p *Payload)
	IsActive() bool
	SetActive(status bool)
}

type Payload struct {
	PortalID string        `json:"portal_id"`
	Type     string        `json:"type"`
	Entries  []interface{} `json:"entries,omitempty"`
	Entry    interface{}   `json:"entry,omitempty"`
}

// Subscription types
type SubSSE struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Active  bool

	// Auto
	portal *Config
}

type SubFunc struct {
	Call   func(p *Payload)
	Active bool

	// Auto
	portal *Config
}
