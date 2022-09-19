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

func GetAnswer(c *gin.Context) {
	db := gin_context.GetDBFromContext(c)
	key := c.Param("key")
	if key == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "GET request must include answer key",
		})
		return
	}
	answer, err := db_utils.GetAnswer(&db, key)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "Not found",
			})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "unable to perform request",
			})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": rest_utils.AnswerData{Key: answer.Key, Value: answer.Value},
	})
}
