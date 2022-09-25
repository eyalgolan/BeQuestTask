package routes

import (
	"KeyValuePermStore/internal/db_utils"
	"KeyValuePermStore/internal/rest_utils"
	"KeyValuePermStore/internal/rest_utils/gin_context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
)

func DeleteAnswer(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "DELETE request must include answer key",
		})
		return
	}

	var answerToDelete rest_utils.Answer
	// Call BindJSON to bind the received JSON to answerToDelete.
	if err := c.BindJSON(&answerToDelete); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	if answerToDelete.Data.Key != key {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "key param must equal key in body",
		})
		return
	}
	if answerToDelete.Event != "delete" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "DELETE request can only have a delete event",
		})
		return
	}
	db := gin_context.GetDBFromContext(c)
	err := db_utils.DeleteAnswer(&db, answerToDelete.Data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "delete answer",
			})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		}
		return
	}
	err = db_utils.DeleteEvent(&db, answerToDelete.Data)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": answerToDelete,
	})
}
