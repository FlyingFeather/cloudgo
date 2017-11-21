package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func newServer() *gin.Engine {
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    mountRoutes(router)
    return router
}

func mountRoutes(router *gin.Engine) {
	  router.GET("/hello/:name", replyhello)
}

func replyhello(ctx *gin.Context) {
	  name := ctx.Param("name")
	  ctx.IndentedJSON(http.StatusOK, gin.H{
        "Test": "Hello " + name,
    })
}
