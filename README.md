## How to use?

1. Create a logger

* Regular
```go
logger := logw.New(&logw.Config{
    ID:         "test",
    HistoryLen: 5,
})
```

* Zap sugared
```go
logger := logwzap.New(&logwzap.LogwZap{
    ID:   "test1",
    Logw: logw.New(&logw.Config{HistoryLen: 5}), // ID is not needed here.

    // Zap related config. These are all optional.
    Encoder       zapcore.Encoder
	EncoderConfig *zapcore.EncoderConfig
	MinLogLevel   zapcore.Level // Default: zap.InfoLevel
})
```

2. Add subscriptions

* Function
```go
logger.Subscribe(&logw.SubFunc{
    Call: func(p *logw.Payload) {
        fmt.Println(string(p.Json()))
    },
})
```

* SSE (Default HTTP Module)
```go
http.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
    logger.Subscribe(&logw.SubSSE{
        Request: r,
        Writer:  w,
    })
})
```

* SSE (GIN)
```go
r.GET("/sse", func(c *gin.Context) {
    logger.Subscribe(&logw.SubSSE{
        Request: c.Request,
        Writer:  c.Writer,
    })
})
```

3. Log

```go
// Regular logger.
logger.Log(&logw.Entry{
    Type:   "info",
    Status: 124,
})

// Sugared logger.
logger.Sugar().Info("any", "thing")
logger.Sugar().Infow("Main", "is", "here")

// Don't worry, all the sugar methods are available. 😄
```