package controller

import (
    "assignment1/service"
    "github.com/gin-gonic/gin"
)

func CreatePlan(c *gin.Context) {

    var plan map[string]interface{}

    if err := c.BindJSON(&plan); err != nil {

        c.JSON(400, gin.H{
            "error": "Invalid JSON",
        })

        return
    }

    status, etag, response :=
        service.CreatePlan(plan)

    c.Header("ETag", etag)

    c.JSON(status, response)
}

func GetPlan(c *gin.Context) {

    id := c.Param("id")

    clientETag :=
        c.GetHeader("If-None-Match")

    status, etag, response :=
        service.GetPlan(id, clientETag)

    c.Header("ETag", etag)

    if status == 304 {

        c.Status(304)
        return
    }

    c.JSON(status, response)
}

func DeletePlan(c *gin.Context) {

    id := c.Param("id")

    status :=
        service.DeletePlan(id)

    c.Status(status)
}
