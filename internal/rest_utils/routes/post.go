package routes

import (
	"KeyValuePermStore/internal/db_utils"
	"KeyValuePermStore/internal/rest_utils"
	"KeyValuePermStore/internal/rest_utils/gin_context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func PostAnswer(c *gin.Context) {
	var newAnswer rest_utils.Answer

	// Call BindJSON to bind the received JSON to newAnswer.
	if err := c.BindJSON(&newAnswer); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	if newAnswer.Event != "create" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "POST request can only have a create event",
		})
		return
	}
	db := gin_context.GetDBFromContext(c)
	err := db_utils.CreateAnswer(&db, newAnswer.Data)
	if err != nil {
		var duplicateAnswerErr *db_utils.DuplicateAnswerErr
		if errors.As(err, &duplicateAnswerErr) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "duplicate answer",
			})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		}
		return
	}
	err = db_utils.CreateEvent(&db, newAnswer.Data)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": newAnswer,
	})
}
