package main

import (
	"fmt"
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils/gin_context"
	"github.com/gin-gonic/gin"
)
import (
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils/routes"
	"github.com/joeshaw/envdecode"
)

var cfg struct {
	postgress_utils.PostgresConfig
	rest_utils.RestConfig
}

func main() {
	envdecode.MustDecode(&cfg)
	db, err := postgress_utils.ConnectToDB(cfg.PostgresConfig)
	if err != nil {
		return
	}
	defer func() {
		sqlDB, err := db.DB.DB()
		if err != nil {
			panic(err)
		}
		err = sqlDB.Close()
		if err != nil {
			panic(err)
		}
	}()
	router := gin.Default()
	gin_context.SetDBInContext(router, *db)
	routes.InitialiseRoutes(router)
	router.Run(fmt.Sprintf("%s:%s", cfg.RestConfig.Address, cfg.RestConfig.Port))
}
