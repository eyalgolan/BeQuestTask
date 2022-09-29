package gin_context

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils"
	"github.com/gin-gonic/gin"
)

func SetDBInContext(router *gin.Engine, db postgress_utils.Client) {
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
}

func GetDBFromContext(c *gin.Context) postgress_utils.Client {
	return c.MustGet("db").(postgress_utils.Client)
}
