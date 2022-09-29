package routes

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils/business_logic"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils/gin_context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostAnswer(c *gin.Context) {
	var newAnswer rest_utils.Answer
	db := gin_context.GetDBFromContext(c)

	// Call BindJSON to bind the received JSON to newAnswer.
	if err := c.BindJSON(&newAnswer); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	statusCode, message := business_logic.Create(newAnswer, db)
	c.IndentedJSON(statusCode, gin.H{
		"message": message,
	})
	return
}
