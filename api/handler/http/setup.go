package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojtabamovahedi/url-shorter/app"
	"github.com/mojtabamovahedi/url-shorter/config"
	"github.com/mojtabamovahedi/url-shorter/internal/service"
	"log"
)

func Run(appContainer *app.App, cfg config.Config) error {
	router := gin.Default()

	router.Use(Logger(), Recovery(), Limiter())

	log.Println("Successfully server started.")

	registerRoutes(router, appContainer.LinkService())

	return router.Run(fmt.Sprintf(":%d", cfg.Server.HttpPort))
}

func registerRoutes(router *gin.Engine, service service.LinkService) {
	router.POST("/new", Save(service))
	router.GET("/:short", Redirect(service))
}
