package routes

import (
	"BeQuestPrep/internal/db_utils"
	"BeQuestPrep/internal/rest_utils"
	"BeQuestPrep/internal/rest_utils/gin_context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
)

func PutAnswer(c *gin.Context) {
	var updateAnswer rest_utils.Answer
	// Call BindJSON to bind the received JSON to updateAnswer.
	if err := c.BindJSON(&updateAnswer); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	if updateAnswer.Event != "update" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "PUT request can only have an update event",
		})
		return
	}
	db := gin_context.GetDBFromContext(c)
	err := db_utils.UpdateAnswer(&db, updateAnswer.Data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "key doesn't exist",
			})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		return
	}
	err = db_utils.UpdateHistory(&db, updateAnswer.Data)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.IndentedJSON(http.StatusOK, updateAnswer)
}
