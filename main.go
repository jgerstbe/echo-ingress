package main

import (
    "github.com/labstack/echo/v4"
)

func main() {
    // Load configuration
    config := getAndPrintConfig()
    
    // Create cache instance and add routes
    c := initCache();
    addRoutesToCache(c, config.Routes)

    // Create a new Echo instance
    e := echo.New()

    // Middleware that validates the request with an external service
    e.Use(validateRequestMiddleware(c))

    // Define a route
    e.GET("/*", handleRequest)
    e.PUT("/*", handleRequest)
    e.POST("/*", handleRequest)
    e.PATCH("/*", handleRequest)
    e.DELETE("/*", handleRequest)

    // Start the server
    e.Start(":8080")
}
