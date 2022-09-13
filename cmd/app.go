package main

import (
	"BeQuestPrep/internal/db_utils/postgress_utils"
	"BeQuestPrep/internal/rest_utils/gin_context"
	"github.com/gin-gonic/gin"
)
import (
	"BeQuestPrep/internal/rest_utils/routes"
	"github.com/joeshaw/envdecode"
)

var cfg struct {
	postgress_utils.PostgresConfig
}

func main() {
	envdecode.MustDecode(&cfg)
	db, err := postgress_utils.ConnectToDB(cfg.PostgresConfig)
	if err != nil {
		return
	}
	router := gin.Default()
	gin_context.SetDBInContext(router, *db)
	routes.InitialiseRoutes(router)
	router.Run("0.0.0.0:8080")
}
