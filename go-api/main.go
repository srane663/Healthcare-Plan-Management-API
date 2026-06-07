package main

import (
    "assignment1/router"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
        c.Next()
    })

    router.InitializeRoutes(r)

    r.Run(":3000")
}
