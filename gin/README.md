# Log Portal

### Setup

Setup log portal in wesionary clean architecture project

`pkg/infrastructure/logportal.go`

```go
package infrastructure

import (
	"github.com/wesionaryTEAM/logportalapi/gin"
)

func RegisterLogPortal(router Router) *logportalapi.SSEEvent {
	return logportalapi.RegisterRoute(router.Engine)
}
```

- Add to dependency

```go
  // ...
  fx.Provide(RegisterLogPortal),
  // ...
```

> `stream.Message <- jsonData` can be use to send any message to stream logger. `jsonData` should be of `logportalapi/LogMessage` type

### Add loggin to GORM

- Create gorm logger module

```go
package framework

import (
	"context"
	"encoding/json"
	"fmt"

	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"github.com/wesionaryTEAM/logportalapi/gin"
)

type Writer interface {
	Printf(string, ...interface{})
}

type SSELogger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
	stream        *logportalapi.SSEEvent
}

func NewSSELogger(
	threshold time.Duration,
	level logger.LogLevel,
	stream *logportalapi.SSEEvent,
) logger.Interface {
	return &SSELogger{
		LogLevel:      level,
		SlowThreshold: threshold,
		stream:        stream,
	}
}

func (l *SSELogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}
func (l *SSELogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.log("info", msg, data...)
}

func (l *SSELogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.log("warn", msg, data...)
}

func (l *SSELogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.log("error", msg, data...)
}

func (l *SSELogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()

	logEntry := logportalapi.LogMessage{
		TimeForQuery: elapsed.Milliseconds(),
		Query:        sql,
		Rows:         rows,
		Type:         "sql",
	}

	if err != nil {
		logEntry.Level = "error"
		logEntry.Message = err.Error()
	} else if elapsed > l.SlowThreshold && l.SlowThreshold != 0 {
		logEntry.Level = "warn"
		logEntry.Message = "Slow SQL detected"
	} else if l.LogLevel == logger.Info {
		logEntry.Level = "info"
	}

	logEntry.Location = utils.FileWithLineNum()

	jsonData, _ := json.Marshal(logEntry)
	l.stream.Message <- jsonData
}
func (l *SSELogger) log(level, msg string, data ...interface{}) {
	logMsg := fmt.Sprintf(msg, data...)
	logEntry := logportalapi.LogMessage{
		Level:    level,
		Message:  logMsg,
		Location: utils.FileWithLineNum(),
	}
	jsonData, _ := json.Marshal(logEntry)
	l.stream.Message <- jsonData
}
```

- Add logger to GORM

```go
func NewDatabase(stream *logportalapi.SSEEvent) Database {
  // ...
  newLogger := framework.NewSSELogger(time.Second, gormlogger.Info, stream)
  database, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{Logger: newLogger})
  // .... other code

  return Database{DB: database}
}

```
