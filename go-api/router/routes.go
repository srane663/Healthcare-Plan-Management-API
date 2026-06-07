package router

import (
    "assignment1/controller"

    "github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {

    v1 := r.Group("/v1")

    {
        v1.POST(
            "/plan",
            controller.CreatePlan,
        )

        v1.GET(
            "/plan/:id",
            controller.GetPlan,
        )

        v1.DELETE(
            "/plan/:id",
            controller.DeletePlan,
        )
    }
}
