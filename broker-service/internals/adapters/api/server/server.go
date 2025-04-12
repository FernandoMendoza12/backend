package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"broker-service/config"

)

func StarServer(cfg *config.Config) error {

	router := gin.Default()

	router.GET("/alive", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok broker alive"})
	})

	return router.Run(":" + cfg.Port)

}
