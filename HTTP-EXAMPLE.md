## Built-in HTTP

```go
http.HandleFunc("/is", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Built-in HTTP"))
})

http.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
    logger.Subscribe(&logw.SubSSE{
        Request: r,
        Writer:  w,
    })
})

http.HandleFunc("/err", func(w http.ResponseWriter, r *http.Request) {
    i = i + 1

    portal.Log(&logw.Entry{
        Type:   "GET",
        Status: i,
    })

    w.Write([]byte("ID is:" + strconv.Itoa(i)))
})

if err := http.ListenAndServe(":8484", nil); err != nil {
    fmt.Println("Error starting server:", err)
}
```

## GIN

```go
r := gin.Default()

r.GET("/is", func(c *gin.Context) {
    c.String(200, "GIN")
})

r.GET("/sse", func(c *gin.Context) {
    logger.Subscribe(&logw.SubSSE{
        Request: c.Request,
        Writer:  c.Writer,
    })
})

r.GET("/err", func(c *gin.Context) {
    i = i + 1

    logger.Log(&logw.Entry{
        Type:   "GET",
        Status: i,
    })

    c.String(200, fmt.Sprintf("ID is: %d", i))
})

r.Run(":8484")
```

# TODO
1. Make it more dynamic
2. Catch panics and recover from them (as a middleware)
3. Support dyanmic clients, eg. email, websocket, write to file, etc.## Built-in HTTP

```go
http.HandleFunc("/is", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Built-in HTTP"))
})

http.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
    logger.Subscribe(&logw.SubSSE{
        Request: r,
        Writer:  w,
    })
})

http.HandleFunc("/err", func(w http.ResponseWriter, r *http.Request) {
    i = i + 1

    portal.Log(&logw.Entry{
        Type:   "GET",
        Status: i,
    })

    w.Write([]byte("ID is:" + strconv.Itoa(i)))
})

if err := http.ListenAndServe(":8484", nil); err != nil {
    fmt.Println("Error starting server:", err)
}
```

## GIN

```go
r := gin.Default()

r.GET("/is", func(c *gin.Context) {
    c.String(200, "GIN")
})

r.GET("/sse", func(c *gin.Context) {
    logger.Subscribe(&logw.SubSSE{
        Request: c.Request,
        Writer:  c.Writer,
    })
})

r.GET("/err", func(c *gin.Context) {
    i = i + 1

    logger.Log(&logw.Entry{
        Type:   "GET",
        Status: i,
    })

    c.String(200, fmt.Sprintf("ID is: %d", i))
})

r.Run(":8484")
```

# TODO
1. Make it more dynamic
2. Catch panics and recover from them (as a middleware)
3. Support dyanmic clients, eg. email, websocket, write to file, etc.