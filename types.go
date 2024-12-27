package logportalapi

type LogMessage struct {
	Type         string `json:"type"`
	Status       int    `json:"status"`
	Method       string `json:"method"`
	Path         string `json:"path"`
	TimeStamp    string `json:"timestamp"`
	Location     string `json:"location"`
	TimeForQuery int64  `json:"timeForQuery"`
	Level        string `json:"level"`
	Message      string `json:"message"`
	StrackTrace  string `json:"stackTrace"`
	Query        string `json:"query"`
	Rows         int64  `json:"rows"`
	ResponseTime string `json:"responseTime"`
}
