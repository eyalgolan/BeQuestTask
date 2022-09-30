package routes

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils/business_logic"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils/gin_context"
	"github.com/gin-gonic/gin"
)

func GetAnswer(c *gin.Context) {
	key := c.Param("key")
	db := gin_context.GetDBFromContext(c)
	statusCode, answerData, err := business_logic.Get(key, db)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{
			"message": err,
		})
		return
	}
	c.IndentedJSON(statusCode, gin.H{
		"message": answerData,
	})
	return
}
