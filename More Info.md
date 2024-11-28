Middleware
One of Gin's strengths lies in its middleware support. Middleware in Gin allows you to define custom handlers that can process requests before they reach your main handler.
This is useful for tasks like logging, authentication, and request modification

```
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()
        // Set example variable
        c.Set("example", "12345")
        // Request before processing
        c.Next()
        // Request after processing
        latency := time.Since(t)
        log.Print(latency)
        // Access status and error from the request
        status := c.Writer.Status()
        log.Println(status)
    }
}
```
Grouping Routes
Gin allows for route grouping, which is a convenient way to organize routes with common middleware or URL prefixes.

```
v1 := r.Group("/v1")
{
    v1.GET("/login", loginEndpoint)
    v1.GET("/submit", submitEndpoint)
    v1.GET("/read", readEndpoint)
}
v2 := r.Group("/v2")
{
    v2.POST("/login", loginEndpoint)
    v2.POST("/submit", submitEndpoint)
    v2.POST("/read", readEndpoint)
}
```

Binding and Validation
Gin provides built-in support for binding and validating JSON, XML, and form data. This simplifies the process of handling request payloads and ensuring they meet expected formats.
```
type Login struct {
    User     string `form:"user" json:"user" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

func main() {
    r := gin.Default()
    r.POST("/loginJSON", func(c *gin.Context) {
        var json Login
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"status": "you are logged in"})
    })
}
```
Error Handling
Error handling in Gin is straightforward and provides multiple ways to capture and respond to errors.


