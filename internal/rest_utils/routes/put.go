package routes

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils/business_logic"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils/gin_context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutAnswer(c *gin.Context) {
	var updateAnswer rest_utils.Answer
	db := gin_context.GetDBFromContext(c)

	// Call BindJSON to bind the received JSON to updateAnswer.
	if err := c.BindJSON(&updateAnswer); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	statusCode, message, err := business_logic.Update(updateAnswer, db)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{
			"message": err,
		})
		return
	}
	c.IndentedJSON(statusCode, gin.H{
		"message": message,
	})
	return
}
