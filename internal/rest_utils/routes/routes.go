package routes

import (
	"github.com/gin-gonic/gin"
)

func InitialiseRoutes(router *gin.Engine) {
	router.GET("/answers/:key", GetAnswer)
	router.POST("/answers", PostAnswer)
	router.PUT("/answers", PutAnswer)
	router.DELETE("/answers/:key", DeleteAnswer)
	router.GET("/answers/:key/history", GetHistory)
}
