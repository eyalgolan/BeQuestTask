package routes

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils/business_logic"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils/gin_context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteAnswer(c *gin.Context) {
	key := c.Param("key")
	db := gin_context.GetDBFromContext(c)
	var answerToDelete *rest_utils.Answer
	// Call BindJSON to bind the received JSON to updateAnswer.
	if err := c.BindJSON(&answerToDelete); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	statusCode, message, err := business_logic.Delete(key, *answerToDelete, db)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{
			"message": err,
		})
		return
	}
	c.IndentedJSON(statusCode, gin.H{
		"message": message,
	})
}
