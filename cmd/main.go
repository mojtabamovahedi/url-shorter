package main

import (
	"flag"
	"github.com/mojtabamovahedi/url-shorter/api/handler/http"
	"github.com/mojtabamovahedi/url-shorter/app"
	"github.com/mojtabamovahedi/url-shorter/config"
	"github.com/mojtabamovahedi/url-shorter/internal/repository"
	"github.com/mojtabamovahedi/url-shorter/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var cfgPath = flag.String("config", "config.yaml", "Path to config file")

func main() {
	flag.Parse()

	cfg := config.MustReadConfig(*cfgPath)
	appContainer := app.MustNewApp(cfg)

	linkRepo := repo.NewLinkRepo(appContainer.DB())

	linkService := service.NewLinkService(linkRepo)

	appContainer.SetLinkService(linkService)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	server := make(chan error, 1)

	go func() {
		server <- http.Run(appContainer, cfg)
	}()

	select {
	case err := <-server:
		if err != nil {
			log.Fatal(err)
		}
	case <-quit:
		log.Println("Shutting down server...")
	}

	appContainer.Close()
	log.Println("Shut down successfully!")

}
