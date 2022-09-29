package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils/gin_context"
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
	latestAnswer, err := db_utils.GetLatestAnswer(&db, newAnswer.Data.Key)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	if latestAnswer != nil {
		if latestAnswer.Value == newAnswer.Data.Value {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "duplicate answer",
			})
			return
		}
	}
	err = db_utils.CreateAnswer(&db, newAnswer.Data)
	if err != nil {
		fmt.Printf("here!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! %+v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	fmt.Printf("\ncreated answer\n")
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
